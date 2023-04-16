package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	_ "github.com/lib/pq"
	userService "github.com/mshmnv/SocialNetwork/internal/app/api/user"
	"github.com/mshmnv/SocialNetwork/internal/pkg/postgres"
	desc "github.com/mshmnv/SocialNetwork/pkg/api/user"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()

	db, _, err := postgres.Connect()
	if err != nil {
		log.Fatal(err)
	}
	postgresCtx := postgres.NewContext(ctx, db)

	fs := flag.NewFlagSet("", flag.ExitOnError)
	grpcAddr := fs.String("grpc-addr", ":6565", "grpc address")
	httpAddr := fs.String("http-addr", ":8080", "http address")
	if err := fs.Parse(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// register service
	server := grpc.NewServer()
	userServer := userService.NewUserAPI(postgresCtx)
	desc.RegisterUserAPIServer(server, userServer)

	rmux := runtime.NewServeMux()
	mux := http.NewServeMux()

	mux.Handle("/", rmux)
	{
		err := desc.RegisterUserAPIHandlerServer(ctx, rmux, userServer)
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

	if err := g.Run(); err != nil {
		log.Fatal(err)
	}
}