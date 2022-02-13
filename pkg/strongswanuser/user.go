package strongswanuser

import (
	"fmt"
	"sync"
)

type user struct {
	Username string
	Password string
}

func (user *user) passwordValid() bool {
	return user.Password != ""
}

type List struct {
	sync.Mutex
	data []*user
}

func NewList() *List {
	return &List{data: make([]*user, 0)}
}

func (list *List) Append(username string, password string) error {
	list.Lock()
	defer list.Unlock()
	user := &user{Username: username, Password: password}
	for _, currentUser := range list.data {
		if currentUser.Username == user.Username {
			return fmt.Errorf("username already exists")
		}
	}
	if user.passwordValid() {
		list.data = append(list.data, user)
	} else {
		return fmt.Errorf("password not valid")
	}
	return nil
}

func (list *List) Remove(username string) error {
	list.Lock()
	defer list.Unlock()
	temp := make([]*user, 0)
	for _, currentUser := range list.data {
		if currentUser.Username != username {
			temp = append(temp, currentUser)
		}
	}
	if len(temp) != len(list.data) {
		list.data = temp
		return nil
	} else {
		return fmt.Errorf("user not found")
	}
}

func (list *List) Users() []*user {
	return list.data
}

func (list *List) Json() string {
	jsonstring := "["
	for _, user := range list.data {
		jsonstring = fmt.Sprintf("%s\"%s\", ", jsonstring, user.Username)
	}
	return jsonstring[:len(jsonstring)-2] + "]\n"
}
