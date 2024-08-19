package encrypt

import (
	"testing"
)

func TestBase64(t *testing.T) {
	str := "Lucifer"
	newStr := Base64Encode(str)
	p(newStr)
	oldStr := Base64Decode(newStr)
	p(oldStr)
}
