package u2utils

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"testing"
)

func TestStrEncodeConvert(t *testing.T) {
	r, err := StrEncodeConvert([]byte{180, 180, 210, 181, 176, 229}, simplifiedchinese.GB18030)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(r))
}
