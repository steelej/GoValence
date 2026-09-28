package valence

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/oauth2"
)

type tokenSourceFunc func() (*oauth2.Token, error)

func (f tokenSourceFunc) Token() (*oauth2.Token, error) { return f() }

func TestOAuthTokenSourceAuthUsesCurrentTokenForEachRequest(t *testing.T) {
	var headers []string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers = append(headers, r.Header.Get("Authorization"))
		_, _ = w.Write([]byte(`{}`))
	}))
	defer server.Close()

	calls := 0
	client := New(Config{
		BaseURL: server.URL,
		Auth: NewOAuthTokenSourceAuth(tokenSourceFunc(func() (*oauth2.Token, error) {
			calls++
			if calls == 1 {
				return &oauth2.Token{AccessToken: "first"}, nil
			}
			return &oauth2.Token{AccessToken: "refreshed"}, nil
		})),
	})

	for i := 0; i < 2; i++ {
		if err := client.get("/test", nil, nil); err != nil {
			t.Fatal(err)
		}
	}
	if calls != 2 {
		t.Fatalf("Token() called %d times, want 2", calls)
	}
	if len(headers) != 2 || headers[0] != "Bearer first" || headers[1] != "Bearer refreshed" {
		t.Fatalf("Authorization headers = %v", headers)
	}
}

func TestOAuthTokenSourceAuthRejectsMissingTokenAndPropagatesErrors(t *testing.T) {
	refreshErr := errors.New("refresh failed")
	tests := []struct {
		name   string
		source oauth2.TokenSource
		want   error
	}{
		{"nil source", nil, nil},
		{"source error", tokenSourceFunc(func() (*oauth2.Token, error) { return nil, refreshErr }), refreshErr},
		{"nil token", tokenSourceFunc(func() (*oauth2.Token, error) { return nil, nil }), nil},
		{"empty token", tokenSourceFunc(func() (*oauth2.Token, error) { return &oauth2.Token{}, nil }), nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/test", nil)
			err := NewOAuthTokenSourceAuth(tt.source).AuthenticateRequest(req)
			if err == nil {
				t.Fatal("expected authentication error")
			}
			if tt.want != nil && !errors.Is(err, tt.want) {
				t.Fatalf("error = %v, want wrapped %v", err, tt.want)
			}
			if got := req.Header.Get("Authorization"); got != "" {
				t.Fatalf("Authorization = %q, want empty", got)
			}
		})
	}
}
