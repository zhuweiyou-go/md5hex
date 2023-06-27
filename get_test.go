package md5hex

import "testing"

func TestGet(t *testing.T) {
	r := Get("123456")
	if r != "e10adc3949ba59abbe56e057f20f883e" {
		t.Fatal(r)
	}

	t.Log(r)
}
