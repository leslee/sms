/*
golang发送短信验证码，集成各大平台短信
作者：邹慧刚
github:github.com/zouhuigang
邮箱：952750120@qq.com

短信平台:http://sms-api.luosimao.com/v1/send.json
*/
package sms

import (
	//"bytes"
	//"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	S_URL     = "http://sms-api.luosimao.com/v1/send.json" //api地址
	S_COM     = "【学优教育】"                                   //【公司名称】
	S_API_KEY = "key-9a59f8ecb5201cc2728bb31719bc35e4"     //luosimao官方APIKEY
)

//真实发送短信
//request map[string]string
func SendSms(strmobile string, content string) string {
	content = content + S_COM
	post_data := make(map[string]string) //post_data := map[string]string{}
	post_data["mobile"] = strmobile
	post_data["message"] = content
	gets := HttpPostH(S_URL, post_data)
	return gets
}

func HttpPostH(queryurl string, postdata map[string]string) string {
	//postdata数据
	//form
	postform := url.Values{} //map[string][]string
	for key, val := range postdata {
		postform.Set(key, val)
	}
	req_new := ioutil.NopCloser(strings.NewReader(postform.Encode())) //把form数据编下码

	//模拟请求
	client := &http.Client{}
	reqest, _ := http.NewRequest("POST", queryurl, req_new)
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded") //application/x-www-form-urlencoded application/json;charset=utf-8  application/x-www-form-urlencoded; param=value
	//realm := "api:key-9a59f8ecb5201cc2728bb31719bc35e4"
	//reqest.Header.Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, realm))
	reqest.SetBasicAuth("api", S_API_KEY)

	response, err := client.Do(reqest)
	if err != nil {
		return err.Error()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err.Error()
	}
	defer response.Body.Close()

	return string(body)

}
