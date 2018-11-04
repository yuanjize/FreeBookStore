package util

import (
	"github.com/mojocn/base64Captcha"
	"sync"
)

type CaptchaConfig struct {
	Id              string
	CaptchaType     string
	VerifyValue     string
	ConfigAudio     base64Captcha.ConfigAudio
	ConfigCharacter base64Captcha.ConfigCharacter
	ConfigDigit     base64Captcha.ConfigDigit
}

var (
	captchaConfig     *CaptchaConfig
	captchaConfigOnce sync.Once
)

// 获取base64验证码基本配置
func GetCaptchaConfig() *CaptchaConfig {
	captchaConfigOnce.Do(func() {
		captchaConfig = &CaptchaConfig{
			Id:          "",
			CaptchaType: "character",
			VerifyValue: "",
			ConfigAudio: base64Captcha.ConfigAudio{},
			ConfigCharacter: base64Captcha.ConfigCharacter{
				Height:             60,
				Width:              240,
				Mode:               2,
				IsUseSimpleFont:    false,
				ComplexOfNoiseText: 0,
				ComplexOfNoiseDot:  0,
				IsShowHollowLine:   false,
				IsShowNoiseDot:     false,
				IsShowNoiseText:    false,
				IsShowSlimeLine:    false,
				IsShowSineLine:     false,
				CaptchaLen:         0,
			},
			ConfigDigit: base64Captcha.ConfigDigit{},
		}
	})
	return captchaConfig
}

const (
	CAPTCHA_IS_RIGHT = 0
	CAPTCHA_IS_ERROR = -7
)

//  验证 验证码是否正确
// captchaId: 存于session中
// verifyValue: 客户端发来的验证码
func VerfiyCaptcha(captchaId, verifyValue string) bool {
	verifyResult := base64Captcha.VerifyCaptcha(captchaId, verifyValue)
	return verifyResult
}
