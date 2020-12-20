package request

import (
	"fmt"
	"net/http"
)

type UserRequest struct {
	AuthToken
}

func (r *UserRequest) NewUserRequestByLogin(login string) *ClientHandle {
	return NewClientHandle(r.AuthToken.Token, fmt.Sprintf("/users/%s", login), http.MethodGet, nil)
}

//通过id获取用户信息
func (r *UserRequest) NewUserRequestById(id int64) *ClientHandle {
	return NewClientHandle(r.AuthToken.Token, fmt.Sprintf("/users/%d", id), http.MethodGet, nil)
}

func (r *UserRequest) ReadArticleById(namespace string, slug string) *ClientHandle{
	return NewClientHandle(r.AuthToken.Token, fmt.Sprintf("/repos/%s/docs/%s", namespace, slug), http.MethodGet, nil)
}

func (r *UserRequest) CreateArticleById(namespace string) *ClientHandle{
	return NewClientHandle(r.AuthToken.Token, fmt.Sprintf("/repos/%s/docs", namespace, ), http.MethodPost, nil)
}