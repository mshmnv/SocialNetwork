package auth

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
)

var authRequiredMethods = []string{"/friend/set/", "/friend/delete/",
	"/post/create", "/post/update", "/post/delete/", "/post/feed", "/add-posts"}

const loginMethod = "/login"

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
	userIDKey              = "user_id"
)

func getRequestBody(r *http.Request) []byte {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("can't read body: %v", err.Error())
		return nil
	}
	r.Body = io.NopCloser(bytes.NewBuffer(body))
	return body
}

func getUserID(r *http.Request) uint64 {
	body := getRequestBody(r)
	bodyData := struct {
		ID uint64 `json:"id"`
	}{}
	err := json.Unmarshal(body, &bodyData)
	if err != nil {
		logger.Errorf("can't decode body: %v", err.Error())
		return 0
	}

	return bodyData.ID
}

func isAuthRequired(r *http.Request) bool {
	for _, authMethod := range authRequiredMethods {
		if ok := strings.HasPrefix(r.URL.Path, authMethod); ok {
			return true
		}
	}
	return false
}

// get or create session token
func getSessionToken(r *http.Request) (*session, error) {
	if r.URL.Path == loginMethod {
		userID := getUserID(r)
		return &session{userID: userID, expiry: time.Now().Add(expirationPeriod)}, nil
	}

	if !isAuthRequired(r) {
		return &session{}, nil
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
	for token, s := range sessions {
		if s.userID == userSession.userID && !userSession.isExpired() {
			printToken(w, token)
			return
		}
	}

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
	printToken(w, newSessionToken)
}

func printToken(w http.ResponseWriter, token string) {
	_, err := io.WriteString(w, "{\"token\": \""+token+"\"}")
	if err != nil {
		logger.Errorf("can't write token: %v", err.Error())
	}
}
