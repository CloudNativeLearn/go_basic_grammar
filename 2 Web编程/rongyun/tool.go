package main

import (
	"github.com/rongcloud/server-sdk-go/v3/sdk"
	"sync"
)
var rc *sdk.RongCloud
var mu sync.Mutex

func RongYunUtil() *sdk.RongCloud{
	if rc == nil{
		mu.Lock()
		if rc == nil {
			appKey:="pvxdm17jpwelr"
			appSecret:="9xWCD0aoN02o"
			rc = sdk.NewRongCloud(appKey,appSecret)
			mu.Unlock()
		}
	}
	return rc
}