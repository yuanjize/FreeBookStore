package auth

import (
	_ "github.com/yuanjize/FreeBookStore/config"
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	GetAccessToken()
	t.Log("hello")
}

//func TestGetAccessToken2(t *testing.T) {
//	params := make(map[string]string)
//	params["client_id"] = config.GetClientId()
//	post := http.RequestPost(config.GetCodeHost(), params)
//	fmt.Print(*post)
//}
