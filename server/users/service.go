package users

import (
	"fmt"
)

func (u *User) FetchUserInfo(userInfoMap map[string]*UserInfo) (*UserInfo, error) {
	userInfoOutput, exists := userInfoMap[u.Id]
	if exists {
		return userInfoOutput, nil
	} else {
		return nil, fmt.Errorf("info of user of id : %s not found", u.Id)
	}
}
