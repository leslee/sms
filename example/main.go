package main

import (
	"fmt"
	"github.com/zouhuigang/sms"
)

func main() {
	gets := sms.SendSms("18516573852", "这是一条验证码:123456")

	fmt.Println(gets)
}
