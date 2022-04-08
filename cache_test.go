package u2utils

import (
	"fmt"
	"testing"
)

func TestTmpCacheSet(t *testing.T) {
	PanicOnError(TmpCacheSet("aa", 11))
}

func TestTmpCacheGet(t *testing.T) {
	var v int
	PanicOnError(TmpCacheGet("aa", &v))
	fmt.Println(v)
}
