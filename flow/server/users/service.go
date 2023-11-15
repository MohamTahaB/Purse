package users

import (
	"fmt"
)

// Fetches the user info from the user info map (temporary until a DB is made), knowing the id, outputs an error if not found.
func (u *User) FetchUserInfoFromUser(userInfoMap map[string]*UserInfo) (*UserInfo, error) {
	userInfoOutput, exists := userInfoMap[u.Id]
	if exists {
		return userInfoOutput, nil
	} else {
		return nil, fmt.Errorf("info of user of id : %s not found", u.Id)
	}
}

// Fetches the user from the user map (temporary until a DB is made), knowing the id, outputs an error if not found.
func (ui *UserInfo) FetchUserFromUserInfo(userMap map[string]*User) (*User, error) {
	userOutput, exists := userMap[ui.Id]
	if exists {
		return userOutput, nil
	} else {
		return nil, fmt.Errorf("user of id : %s not found", ui.Id)
	}
}
