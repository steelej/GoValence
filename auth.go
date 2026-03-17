package valence

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Authenticator modifies an outgoing request to add authentication.
type Authenticator interface {
	AuthenticateRequest(req *http.Request) error
}

// D2LAuth implements D2L ID-keyset (x_a/x_b/x_c/x_d/x_t) authentication.
type D2LAuth struct {
	AppID   string
	AppKey  string
	UserID  string
	UserKey string
}

func (a *D2LAuth) AuthenticateRequest(req *http.Request) error {
	ts := time.Now().Unix()
	method := strings.ToUpper(req.Method)
	path := strings.ToLower(req.URL.Path)

	msg := fmt.Sprintf("%s&%s&%d", method, path, ts)

	appSig, err := d2lSign(a.AppKey, msg)
	if err != nil {
		return err
	}
	userSig, err := d2lSign(a.UserKey, msg)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Set("x_a", a.AppID)
	q.Set("x_b", a.UserID)
	q.Set("x_c", appSig)
	q.Set("x_d", userSig)
	q.Set("x_t", fmt.Sprintf("%d", ts))
	req.URL.RawQuery = q.Encode()
	return nil
}

func d2lSign(key, message string) (string, error) {
	mac := hmac.New(sha256.New, []byte(key))
	if _, err := mac.Write([]byte(message)); err != nil {
		return "", err
	}
	sig := mac.Sum(nil)
	// URL-safe base64, no padding
	encoded := base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(sig)
	return encoded, nil
}

// OAuthAuth implements Bearer token authentication.
type OAuthAuth struct {
	Token string
}

func (a *OAuthAuth) AuthenticateRequest(req *http.Request) error {
	req.Header.Set("Authorization", "Bearer "+a.Token)
	return nil
}

// NewD2LAuth constructs a D2LAuth authenticator.
func NewD2LAuth(appID, appKey, userID, userKey string) *D2LAuth {
	return &D2LAuth{AppID: appID, AppKey: appKey, UserID: userID, UserKey: userKey}
}

// NewOAuthAuth constructs an OAuthAuth authenticator.
func NewOAuthAuth(token string) *OAuthAuth {
	return &OAuthAuth{Token: token}
}

// buildURL constructs a full URL from the client base URL and path segments,
// applying any provided query parameters.
func buildURL(base, path string, params url.Values) string {
	u := base + path
	if len(params) > 0 {
		u += "?" + params.Encode()
	}
	return u
}
