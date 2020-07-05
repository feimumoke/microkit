package ps

import (
	"errors"
	"fmt"
)

type IUserService interface {
	GetName(userid int) string
	DelUser(userid int) error
}

type UserService struct {
}

func (u UserService) GetName(userid int) string {
	if userid == 101 {
		return "zhangsan"
	}

	return "lisi"
}

func (u UserService) DelUser(userid int) error {
	if userid == 101 {
		return errors.New("no author")
	}
	fmt.Println("Del user")
	return nil
}
