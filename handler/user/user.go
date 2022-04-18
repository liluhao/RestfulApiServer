package user

import (
	"apiserver/model"
)

/*研发经验，建议：如果消息体有JSON参数需要传递，针对每一个API接口定义独立的go struct来接收，
比如CreateRequest和CreateResponse,并将这些结构体统一放在一个Go文件中，以方便后期维护和修改。
这样做可以使代码结构更加规整和清晰，本例统一放在hand1er/user/user.go中*/
type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []*model.UserInfo `json:"userList"`
}

type SwaggerListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	UserList   []model.UserInfo `json:"userList"`
}
