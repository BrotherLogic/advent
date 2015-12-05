package main

import "crypto/md5"
import "encoding/hex"
import "fmt"
import "strconv"
import "strings"

func SolveHash(str string, len int) int {
	value := 0
	found := false
	match := ""
	for i := 0; i < len ; i++ {
		match += "0"
	}
	
	for !found {
		value++
		data := []byte(str + strconv.Itoa(value))
		hash := md5.New()
		hash.Write(data)
		md5v := hex.EncodeToString(hash.Sum(nil))
		if strings.HasPrefix(md5v,match) {
			found = true
		}
	}
	return value;
}

func dayfour() {
	fmt.Printf("Answer = %d\n", SolveHash("bgvyzdsv",5))
	fmt.Printf("Answer = %d\n", SolveHash("bgvyzdsv",6))
}
