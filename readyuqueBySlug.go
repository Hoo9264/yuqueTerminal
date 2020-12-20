package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/FlashFeiFei/yuque/request"
	"github.com/FlashFeiFei/yuque/response"
	"io/ioutil"
	"os"
)
const key = "0123456789012345" //自定义密钥，限制了输入k的长度必须为16, 24或者32
func main()  {
	fmt.Println("传入yuque文章的slug（最后的路径短称）")
	arg := os.Args[1]
	if arg=="" {
		fmt.Println("无效短称")
		os.Exit(1)
	}
	title, body := ReadYuqueDoc(arg)
	title = title[1:len(title)-1]
	body = body[1:len(body)-1]
	decryptCode := AesDecrypt(body, key)
	if err := ioutil.WriteFile(title+"-clean.md", []byte(decryptCode),0666); err!=nil{
		fmt.Println("failed to remove CBC code")
	}
	//arg := os.Args[1]
	//ReadYuqueDoc(arg)
}


func ReadYuqueDoc(slug string) (string,string){
	user_request := request.UserRequest{
		AuthToken: request.AuthToken{
			Token: "xxxxxxxxxxxxxxxxxxx",  //yuque配置的token
		},
	}
	client := user_request.ReadArticleById("", slug)  //namespace为知识库id
	res_doc := new(response.ResponseDocDetailSerializer)
	client.Request(res_doc)
	title, _ := json.Marshal(res_doc.Data.Title)
	body, _ := json.Marshal(res_doc.Data.Body)
	return string(title),string(body)
	//fmt.Println(string(data))
}

func AesDecrypt(cryted string, key string) string {
	// 转成字节数组
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}
//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
