package u2utils

import (
	"fmt"
	"testing"
)

func TestJsonTypeConvert(t *testing.T) {
	var out int
	PanicOnError(JsonTypeConvert(1.2, &out))
	if out != 1 {
		panic(fmt.Sprintf("out is %+v", out))
	}
}
