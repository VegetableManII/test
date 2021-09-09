package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"git.lianjia.com/infrastructure/api-signature-go/signer"
	"gopkg.in/jcmturner/gokrb5.v7/credentials"
)

var (
	httpHost = "http://test3-core.permission.lianjia.com/test_captcha/graphql"

	clientAccessKeyId     = "7GGHIPZX2JN9MH6XH555"
	clientSecretAccessKey = "8/5K62xHH7D8LLfbX+Q/e8PwS/JCvk/jwZaY1Ijc"
)

func credentialProvider(accessKeyId string) (*credentials.Credential, error) {
	return &credentials.Credential{
		AccessKeyId:     clientAccessKeyId,
		SecretAccessKey: clientSecretAccessKey,
	}, nil
}

func main() {

	url := "http://test3-core.permission.lianjia.com/test_captcha/graphql"
	method := "POST"

	payload := strings.NewReader("{\"query\":\"mutation{createRole(input:{id:\\\"哈哈哈\\\",name:\\\"测试\\\"}){id name}}\",\"variables\":{}}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "LJ-HMAC-SHA256 accessKeyId=7GGHIPZX2JN9MH6XH555; nonce=91ba5a2d-b9d1-4045-bf01-139bc5f9c792; timestamp=1630995713; signature=JGX49NCbR7sLT9NqXSGw3mygi43c4dIWg+C6ws+ayQw=")
	req.Header.Add("AK", "7GGHIPZX2JN9MH6XH555")
	req.Header.Add("SK", "8/5K62xHH7D8LLfbX+Q/e8PwS/JCvk/jwZaY1Ijc")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "lianjia_ssid=c6b63edf-38a3-4f2a-98b6-7b746df1faf0; lianjia_uuid=99fbf957-ef23-47ce-92f3-ecb783ad885d")
	r, err := signer.SignRequestWithHeader(req, credentialProvider, nil, clientAccessKeyId, "", "")
	res, err := client.Do(r)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
