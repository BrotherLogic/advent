package main

import "crypto/md5"
import "encoding/hex"
import "fmt"
import "strconv"
import "strings"

func SolveHash(str string) int {
	value := 0
	found := false

	for !found {
		value++
		data := []byte(str + strconv.Itoa(value))
		hash := md5.New()
		hash.Write(data)
		md5v := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(md5v,"00000") {
			found = true
		}
	}
	return value;
}

func dayfour() {
	fmt.Printf("Answer = %d\n", SolveHash("bgvyzdsv"))
}
