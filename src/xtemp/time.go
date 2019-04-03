package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

func main() {
	h := md5.New()
	io.WriteString(h, "zhyuzh3d")
	a := h.Sum(nil)
	s := hex.EncodeToString(a)
	fmt.Println(s)
}
