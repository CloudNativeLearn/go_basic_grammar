package main

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {

	Nonce:=CreateCaptcha2()
	Timestamp:= time.Now().Unix()

	Secret := "9xWCD0aoN02o"
	signature := Secret+string(Nonce)+strconv.FormatInt(Timestamp,10)
	fmt.Println(signature)

	t := sha1.New();
	t.Write([]byte(signature))
	signature =fmt.Sprintf("%x",t.Sum(nil))



	fmt.Println(Nonce)
	fmt.Println(Timestamp)
	fmt.Println(signature)


}

func CreateCaptcha2() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
