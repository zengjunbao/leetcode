package markdown

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestName(t *testing.T) {
	html, err := MdToHtml("111")
	assert.NoError(t, err)
	fmt.Println(string(html))

}
