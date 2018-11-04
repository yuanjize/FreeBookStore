package config

import (
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

const CONFIG_FILE = "./config.yaml"
const Host  = "http://localhost:8080/"
var (
	clientId      string
	clientSecret  string
	codeHost      string
	authHost      string
	captchaSwitch bool
)

type ConfigParam struct {
	GithubClientId     string `yaml:"githubClientId"`
	GithubClientSecret string `yaml:"githubClientSecret"`
	GithubCodeHost     string `yaml:"githubCodeHost"`
	GithubAuthHost     string `yaml:"githubAuthHost"`
	CaptchaSwitch      bool   `yaml:"captcha"`
}

// parse config.yaml
func init() {
	//fmt.Println(os.Getwd())
	file, err := os.Open(CONFIG_FILE)
	if err != nil {
		log.Panicf("open config file %s fail:%s\n", CONFIG_FILE, err)
	}
	decoder := yaml.NewDecoder(file)
	param := &ConfigParam{}
	err = decoder.Decode(param)
	log.Println("config :%#v", param)
	if err != nil {
		log.Panicf("decode config file %s fail:%s\n", CONFIG_FILE, err)
	}
	clientId = param.GithubClientId
	clientSecret = param.GithubClientSecret
	codeHost = param.GithubCodeHost
	authHost = param.GithubAuthHost
	captchaSwitch = param.CaptchaSwitch
}

func GetClientId() string {
	return clientId
}

func GetClientSecret() string {
	return clientSecret
}

func GetCodeHost() string {
	return codeHost
}
func GetAuthHost() string {
	return authHost
}

func CaptchaSwitch() string {
	return authHost
}
