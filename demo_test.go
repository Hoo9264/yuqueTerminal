package main

import (
	"encoding/json"
	"fmt"
	"github.com/FlashFeiFei/yuque/request"
	"github.com/FlashFeiFei/yuque/request/front"
	"github.com/FlashFeiFei/yuque/response"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	content, _ := ioutil.ReadAll(req.Body)

	doc := response.ResponseDocDetailSerializer{}

	json.Unmarshal(content, &doc)
	log.Println(string(content))
	log.Println("===============================================")
	log.Println(doc)
	log.Println(fmt.Sprintf("v=%v, t=%T", doc, doc))
	log.Println(doc.Data.ID)
	log.Println(doc.Data.Body)
	log.Println(doc.Data.DeletedAt)
}

func UserInfo(w http.ResponseWriter, req *http.Request) {
	client := http.Client{}

	creq, _ := http.NewRequest(http.MethodGet, "https://www.yuque.com/api/v2/users/", nil)
	creq.Header.Add("X-Auth-Token", "xxxxxxxxxxxxxxxxxxx")
	resp, _ := client.Do(creq)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(body)
	log.Println(string(body))
}

func UserInfo2(w http.ResponseWriter, req *http.Request) {
	log.Println("my---------user")
	user_request := request.UserRequest{
		AuthToken: request.AuthToken{
			Token: "xxxxxxxxxxxxxxxxxxxxxxxxx",
		},
	}
	client := user_request.NewUserRequestById(1231375)
	res_user := new(response.ResponseUserSerializer)
	client.Request(res_user)
	data, _ := json.Marshal(res_user)
	log.Println(string(data))
}

//前端调用的api,非文档的，以后可能会被封掉
func DocDetail(w http.ResponseWriter, req *http.Request) {
	client := http.Client{}

	creq, _ := http.NewRequest(http.MethodGet, "https://www.yuque.com/api/docs/tqzqet?book_id=1955564", nil)
	resp, _ := client.Do(creq)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))
}

func DocDetail2(w http.ResponseWriter, req *http.Request) {
	log.Println("DocDetail2")
	response := front.GetDocIntorSerializer("xxxxxxxxx", 111111111)
	data, err := json.Marshal(response)
	log.Println(string(data), err)
}

func TestRun(t *testing.T) {
	fmt.Println("testtest")
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/user", UserInfo)
	http.HandleFunc("/myuser", UserInfo2)
	http.HandleFunc("/DocDetail", DocDetail)
	http.HandleFunc("/DocDetail2", DocDetail2)

	http.ListenAndServe(":12345", nil)
}


