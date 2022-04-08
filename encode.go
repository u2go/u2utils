package u2utils

import (
	"bytes"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
)

var (
	GBK  = simplifiedchinese.GBK
	UTF8 = unicode.UTF8
)

func StrEncodeConvert(str []byte, encode encoding.Encoding) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(str), encode.NewDecoder())
	d, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return d, nil
}
