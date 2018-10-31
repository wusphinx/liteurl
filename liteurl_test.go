package liteurl_test

import (
	"math"
	"testing"

	"github.com/wusphinx/liteurl"
)

func TestConvertToBase(t *testing.T) {
	var i uint64
	for i = 1; i < 10000000; i++ {
		str := liteurl.ConvertToBase(i)
		if liteurl.RestoreBase10(str) != i {
			t.Errorf("str:%s not match i%d", str, i)
		}
	}

	if r := liteurl.ConvertToBase(math.MaxUint64); math.MaxUint64 != liteurl.RestoreBase10(r) {
		t.FailNow()
	}
}

func TestGetMD5(t *testing.T) {
	r := liteurl.GetMD5("https://www.google.com/")
	if r != "d75277cdffef995a46ae59bdaef1db86" {
		t.FailNow()
	}
}
