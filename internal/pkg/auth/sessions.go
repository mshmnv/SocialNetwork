package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

type session struct {
	userID uint64
	expiry time.Time
}

var sessions = map[string]session{}

func (s session) isExpired() bool {
	return s.expiry.Before(time.Now())
}

const (
	sessionTokenCookieName = "session-token"
	expirationPeriod       = 10 * time.Minute
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		for key, val := range sessions {
			logger.Errorf("SESSION %s, userID: %d", key, val.userID)
		}

		var userSession *session
		var err error

		//if r.URL.Path != "/user/register" && r.URL.Path != "/login" {
		if userSession, err = getSessionToken(r); err != nil {
			http.Error(w, "Please login: "+err.Error(), http.StatusUnauthorized)
			return
		}
		//}

		lrw := negroni.NewResponseWriter(w)
		next.ServeHTTP(lrw, r)

		if r.URL.Path == "/login" && lrw.Status() == http.StatusOK { // todo whe recreate
			createSession(w, userSession)
		}
	})
}

func getUserID(r *http.Request) uint64 {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("can't read body: %v", err.Error())
		return 0
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	bodyData := struct {
		ID string `json:"id"`
	}{}
	err = json.Unmarshal(body, &bodyData)
	if err != nil {
		logger.Errorf("can't get user id: %v", err.Error())
		return 0
	}

	userID, err := strconv.ParseUint(bodyData.ID, 10, 64)
	if err != nil {
		logger.Errorf("can't get user id: %v", err.Error())
		return 0
	}
	return userID
}

// get or create session token
func getSessionToken(r *http.Request) (*session, error) {
	if r.URL.Path == "/login" {
		userID := getUserID(r)
		return &session{userID: userID, expiry: time.Now().Add(expirationPeriod)}, nil
	}

	if r.URL.Path == "/user/register" {
		return nil, nil
	}

	c, err := r.Cookie(sessionTokenCookieName)
	if err != nil {
		return nil, err
	}
	sessionToken := c.Value

	userSession, ok := sessions[sessionToken]
	if !ok {
		return nil, errors.New("no active user session")
	}

	if userSession.isExpired() {
		delete(sessions, sessionToken)
		return nil, errors.New("user session is expired")
	}
	return &userSession, nil
}

func createSession(w http.ResponseWriter, userSession *session) {
	newSessionToken := uuid.NewString()

	http.SetCookie(w, &http.Cookie{
		Name:     sessionTokenCookieName,
		Value:    newSessionToken,
		Expires:  userSession.expiry,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})

	sessions[newSessionToken] = session{userID: userSession.userID, expiry: userSession.expiry}

	logger.Infof("Session created: %s", newSessionToken)
}
