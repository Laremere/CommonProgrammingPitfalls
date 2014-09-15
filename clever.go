package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var diffs []byte = []byte{67, 41, 249, 17, 239, 13, 186, 244, 65, 17, 243, 9, 185, 77, 172, 41, 246}
	for i := 1; i < len(diffs); i++ {
		diffs[i] = diffs[i-1] + diffs[i]
	}
	io.Copy(os.Stdout, bytes.NewReader(diffs))
}
