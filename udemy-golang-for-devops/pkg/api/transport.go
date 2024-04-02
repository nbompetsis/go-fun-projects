package api

import (
	"net/http"
)

type JWTTrasnport struct {
	Transport http.RoundTripper
	Token     string
	Password  string
	Username  string
	LoginURL  string
}

func (m *JWTTrasnport) RoundTrip(req *http.Request) (res *http.Response, err error) {
	if m.Token == "" {
		if m.Password != "" {
			m.Token, err = doLoginRequest(http.Client{}, m.LoginURL,
				LoginRequest{Username: m.Username, Password: m.Password})
			if err != nil {
				return nil, err
			}
		}
	}
	if m.Token != "" {
		req.Header.Add("Authorization", "Bearer "+m.Token)
	}
	return m.Transport.RoundTrip(req)
}
