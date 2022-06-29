package captcha

import (
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type Captcha struct {
	c *base64Captcha.Captcha
}

// NewCaptcha 创建字符串验证码实例
func NewCaptcha(height, width, length int, source string) *Captcha {
	driver := base64Captcha.NewDriverString(
		height,
		width,
		0,
		base64Captcha.OptionShowHollowLine,
		length,
		source,
		&color.RGBA{0, 0, 0, 0},
		&base64Captcha.EmbeddedFontsStorage{},
		[]string{},
		)
	cape := base64Captcha.NewCaptcha(driver, base64Captcha.DefaultMemStore)
	return &Captcha{cape}
}

// VerifyCaptcha 验证是否有效
func (r *Captcha) VerifyCaptcha(id, answer string) bool {
	get := base64Captcha.DefaultMemStore.Get(id, false)
	if get == "" {
		return false
	}
	return r.c.Verify(id, answer, true)
}

// GenerateCaptcha 生成base64
func (r *Captcha) GenerateCaptcha() (id, b64s string, err error) {
	return r.c.Generate()
}

// GetLoginCaptcha 登录验证码
func GetLoginCaptcha() *Captcha {
	source := "1234567890qwertyuioplkjhgfdsazxcvbnm"
	return NewCaptcha(40, 120, 4, source)
}
