package auth

import (
	"github.com/yuanjize/FreeBookStore/config"
)

var AccessToken chan string

func init() {
	AccessToken = make(chan string)
}

type GithubOAuth2 struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

func GetAccessToken() string {
	params := make(map[string]string)
	params["client_id"] = config.GetClientId()
	params["redirect_uri"] = "http://localhost:8080/token"
	//response := http.RequestGet(config.GetCodeHost(), params, nil)
	return ""
}
