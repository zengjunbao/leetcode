package captcha

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleCaptcha(t *testing.T) {
	path := "111.jpg"
	id, err := NewSimpleCaptcha(path)
	assert.NoError(t, err)

	number := 123456
	value, err := CheckSimpleCaptcha(id, number)
	assert.NoError(t, err)
	assert.Equal(t, value,true)
}
