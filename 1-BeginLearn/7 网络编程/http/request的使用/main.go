package main

import (
	"fmt"
	_ "github.com/kirinlabs/HttpRequest"
	"github.com/rongcloud/server-sdk-go/v3/sdk"
	"math/rand"
	"time"
)
func main() {

	appKey:="pvxdm17jpwelr"
	appSecret:= "9xWCD0aoN02o"
	rc:=sdk.NewRongCloud(appKey,appSecret)
	rep,err:=rc.UserRegister("u01",
		"name01",
		"http://rongcloud.cn/portrait.jpg",)
	fmt.Println(rep)
	fmt.Println(err)
	fmt.Println(rep.Status)
	fmt.Println(rep.BlockEndTime)


	fmt.Println(rep.Token)
	fmt.Println(rep.UserID)

	status, err := rc.OnlineStatusCheck(rep.UserID)
	fmt.Println(status)
}

func CreateCaptcha() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
}
