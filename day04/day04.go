package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	h := md5.New()
	count := 1
	fiveFound := false
	for {
		h.Reset()
		io.WriteString(h, fmt.Sprintf("%s%d", os.Args[1], count))
		hash := hex.EncodeToString(h.Sum(nil))
		if hash[:5] == "00000" && !fiveFound {
			fmt.Println(count)
			fiveFound = true
		}
		if hash[:6] == "000000" {
			fmt.Println(count)
			break
		}
		count++
	}

}
