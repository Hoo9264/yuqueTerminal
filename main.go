package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/FlashFeiFei/yuque/request"
	"github.com/FlashFeiFei/yuque/response"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main111()  {
	arg := os.Args[1]
	b ,err := ioutil.ReadFile(arg)
	if err != nil {
		fmt.Println("failed to open file")
		os.Exit(1)
	}
	orig := string(b)
	//orig := "试试"
	key := "0123456789012345"
	//加密
	encryptCode := AesEncrypt(orig, key)
	//名字是当天的时间
	name := time.Now().Format("2006-01-02")
	CreateYuqueDoc(name, encryptCode)
}

func AesEncrypt(orig string, key string) string {
	// 转成字节数组
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}
//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func CreateYuqueDoc(title string, body string) {
	log.Println("my---------user")
	user_request := request.UserRequest{
		AuthToken: request.AuthToken{
			Token: "",
		},
	}
	post := "{\"title\":\"" + title +
		"\",\"public\":\"" + "1" +
		"\",\"body\":\"" + body +
		"\"}"
	jsonStr := []byte(post)
	client := user_request.CreateArticleById("")
	client.Body = bytes.NewBuffer(jsonStr)
	res_user := new(response.ResponseUserSerializer)
	client.Request(res_user)
	data, _ := json.Marshal(res_user)
	log.Println(string(data))
}

