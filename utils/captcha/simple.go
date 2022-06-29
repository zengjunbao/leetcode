package captcha

import (
	"errors"
	"flag"
	"github.com/dchest/captcha"
	"io"
	"os"
)

// 默认配置
var (
	flagImage = flag.Bool("i", true, "output image captcha")
	flagAudio = flag.Bool("a", false, "output audio captcha")
	flagLang  = flag.String("lang", "en", "language for audio captcha")
	flagLen   = flag.Int("len", captcha.DefaultLen, "length of captcha")
	flagImgW  = flag.Int("width", captcha.StdWidth, "image captcha width")
	flagImgH  = flag.Int("height", captcha.StdHeight, "image captcha height")
)

// 验证码存储器
var mp = make(map[string]interface{})

// 生成验证码图片，返回验证码id
func NewSimpleCaptcha(path string) (string, error) {
	// 文件存储名称
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var w io.WriterTo
	// 随机数字验证码
	number := captcha.RandomDigits(*flagLen)
	switch {
	case *flagAudio:
		w = captcha.NewAudio("", number, *flagLang)
	case *flagImage:
		w = captcha.NewImage("", number, *flagImgW, *flagImgH)
	}
	// 保存图片

	if _, err := w.WriteTo(f);err != nil {
		return "", err
	}

	id := captcha.NewLen(10)
	// 保存验证码  用于校验
	mp[id] = number

	return id, nil
}

func CheckSimpleCaptcha(id string, num int) (bool, error) {
	value,ok := mp[id]
	if !ok{
		return false,errors.New("captcha not found")
	}
	if value == num {
		return true,nil
	}
	return false, errors.New("captcha check fail")
}
