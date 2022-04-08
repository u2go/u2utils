package u2utils

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"testing"
)

func TestStrEncodeConvert(t *testing.T) {
	fmt.Println(string(StrEncodeConvert([]byte{180, 180, 210, 181, 176, 229}, simplifiedchinese.GB18030)))
}
