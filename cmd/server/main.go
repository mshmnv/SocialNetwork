package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	dialogService "github.com/mshmnv/SocialNetwork/internal/app/api/dialog"
	friendService "github.com/mshmnv/SocialNetwork/internal/app/api/friend"
	postService "github.com/mshmnv/SocialNetwork/internal/app/api/post"
	userService "github.com/mshmnv/SocialNetwork/internal/app/api/user"
	"github.com/mshmnv/SocialNetwork/internal/pkg/auth"
	"github.com/mshmnv/SocialNetwork/internal/pkg/metrics"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	"github.com/mshmnv/SocialNetwork/internal/pkg/redis"
	dialogDesc "github.com/mshmnv/SocialNetwork/pkg/api/dialog"
	friendDesc "github.com/mshmnv/SocialNetwork/pkg/api/friend"
	postDesc "github.com/mshmnv/SocialNetwork/pkg/api/post"
	userDesc "github.com/mshmnv/SocialNetwork/pkg/api/user"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}

	shardedDB, err := postgres.ConnectSharded()
	if err != nil {
		log.Fatal(err)
	}

	redisCl, err := redis.Connect()
	if err != nil {
		log.Fatal(err)
	}

	app := startServer(ctx, db, shardedDB, redisCl)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}

func startServer(ctx context.Context, db *postgres.DB, shardedDB *postgres.ShardedDB, redisClient *redis.Client) run.Group {
	fs := flag.NewFlagSet("", flag.ExitOnError)
	grpcAddr := fs.String("grpc-addr", ":"+os.Getenv("SERVER_GRPC_PORT"), "grpc address")
	httpAddr := fs.String("http-addr", ":"+os.Getenv("SERVER_HTTP_PORT"), "http address")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// register service
	server := grpc.NewServer()

	userServer := userService.NewUserAPI(db)
	userDesc.RegisterUserAPIServer(server, userServer)
	friendServer := friendService.NewFriendAPI(db)
	friendDesc.RegisterFriendAPIServer(server, friendServer)
	postServer := postService.NewPostAPI(db, redisClient)
	postDesc.RegisterPostAPIServer(server, postServer)
	dialogServer := dialogService.NewDialogAPI(shardedDB, db)
	dialogDesc.RegisterDialogAPIServer(server, dialogServer)

	rmux := runtime.NewServeMux()
	mux := http.NewServeMux()

	mux.Handle("/", metrics.PrometheusMiddleware(auth.AuthenticationMiddleware(rmux)))
	{
		err := userDesc.RegisterUserAPIHandlerServer(ctx, rmux, userServer)
		if err != nil {
			log.Fatal(err)
		}
		err = friendDesc.RegisterFriendAPIHandlerServer(ctx, rmux, friendServer)
		if err != nil {
			log.Fatal(err)
		}
		err = postDesc.RegisterPostAPIHandlerServer(ctx, rmux, postServer)
		if err != nil {
			log.Fatal(err)
		}
		err = dialogDesc.RegisterDialogAPIHandlerServer(ctx, rmux, dialogServer)
		if err != nil {
			log.Fatal(err)
		}
	}

	// metrics

	mux.Handle("/metrics", promhttp.Handler())

	// serve

	var g run.Group
	{
		grpcListener, err := net.Listen("tcp", *grpcAddr)
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() error {
			log.Printf("Serving grpc address %s", *grpcAddr)
			return server.Serve(grpcListener)
		}, func(error) {
			grpcListener.Close()
		})
	}
	{
		httpListener, err := net.Listen("tcp", *httpAddr)
		if err != nil {
			log.Fatal(err)
		}
		g.Add(func() error {
			log.Printf("Serving http address %s", *httpAddr)
			return http.Serve(httpListener, mux)
		}, func(err error) {
			httpListener.Close()
		})
	}
	return g
}
