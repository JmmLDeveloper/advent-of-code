package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func main() {

	const input string = "iwrupvqb"

	h := md5.New()
	num := 0


	for true {
		test_case := fmt.Sprintf("%s%d",input,num)

	
		io.WriteString(h, test_case)
		digest := fmt.Sprintf("%x", h.Sum(nil))

		if len(digest) > 5 && digest[0:6] == "000000" {
			fmt.Printf("The result is %d\n",num)
			break
		}

		num++
		h.Reset()
	}
}
