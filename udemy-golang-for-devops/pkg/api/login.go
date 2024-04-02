package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func doLoginRequest(client http.Client, loginURL string, req LoginRequest) (string, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return "", err
	}

	ressponse, err := client.Post(loginURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}

	defer ressponse.Body.Close()

	res, err := io.ReadAll(ressponse.Body)
	if err != nil {
		return "", err
	}

	if ressponse.StatusCode != 200 {
		return "", fmt.Errorf("HTTP login response status %d", ressponse.StatusCode)
	}

	var loginResponse LoginResponse
	err = json.Unmarshal(res, &loginResponse)
	if err != nil {
		return "", err
	}
	fmt.Print("Login successful", loginResponse)
	return loginResponse.Token, nil
}
