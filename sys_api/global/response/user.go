package response

import (
	"fmt"
	"time"
)

type JsonTime time.Time

func (j JsonTime) MarshalJSON() ([]byte, error) {
	stmp := fmt.Sprintf("\"%s\"", time.Time(j).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Id       int32    `json:"id"`
	NickName string   `json:"nick_name"`
	Mobile   string   `json:"mobile"`
	Birthday JsonTime `json:"birthday"`
	Gender   string   `json:"gender"`
}
