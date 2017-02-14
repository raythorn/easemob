package easemob

import (
	"errors"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname,omitempty"`
}

func (user *User) add(context *Context, data interface{}) (resp []byte, err error) {
	uri := context.uri() + "/users"
	var payload interface{} = user

	if data != nil {
		switch data.(type) {
		case *User, []*User:
			payload = data
		default:
			return nil, errors.New("Invalid parameters")
		}
	}

	return context.send("POST", uri, payload)
}

func (user *User) get(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username
	return context.send("GET", uri, nil)
}

func (user *User) getUsers(context *Context, limit int, cursor string) (resp []byte, err error) {
	uri := context.uri() + "/users"
	var param bool = false
	if limit != 0 {
		uri += "?limit=" + fmt.Sprintf("%d", limit)
		param = true
	}

	if cursor != "" {
		if param {
			uri += "&"
		} else {
			uri += "?"
		}

		uri += "cursor=" + cursor
	}

	return context.send("GET", uri, nil)
}

func (user *User) delete(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username
	return context.send("DELETE", uri, nil)
}

func (user *User) deleteUsers(context *Context, limit int) (resp []byte, err error) {

	if limit == 0 {
		return nil, errors.New("Invalid params")
	}

	uri := context.uri() + "/users?limit=" + fmt.Sprintf("%d", limit)
	return context.send("DELETE", uri, nil)
}

func (user *User) setPassword(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/password"
	data := map[string]string{"newpassword": user.Password}
	return context.send("PUT", uri, data)
}

func (user *User) setNickname(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username
	data := map[string]string{"nickname": user.Nickname}
	return context.send("PUT", uri, data)
}

func (user *User) addFriend(context *Context, friend string) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/contacts/users/" + friend
	return context.send("POST", uri, nil)
}

func (user *User) deleteFriend(context *Context, friend string) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/contacts/users/" + friend
	return context.send("DELETE", uri, nil)
}

func (user *User) getFriends(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/contacts/users"
	return context.send("GET", uri, nil)
}

func (user *User) getBlockUsers(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/blocks/users"
	return context.send("GET", uri, nil)
}

func (user *User) blockUsers(context *Context, blockusers []string) (resp []byte, err error) {
	if len(blockusers) == 0 {
		return nil, errors.New("Invalid Params")
	}

	uri := context.uri() + "/users/" + user.Username + "/blocks/users"
	data := map[string][]string{"usernames": blockusers}
	return context.send("POST", uri, data)
}

func (user *User) deleteBlockUser(context *Context, blockuser string) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/blocks/users/" + blockuser
	return context.send("DELETE", uri, nil)
}

func (user *User) getUserStatus(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/status"
	return context.send("GET", uri, nil)
}

func (user *User) activate(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/activate"
	return context.send("POST", uri, nil)
}

func (user *User) deactivate(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/deactivate"
	return context.send("POST", uri, nil)
}

func (user *User) offline(context *Context) (resp []byte, err error) {
	uri := context.uri() + "/users/" + user.Username + "/disconnect"
	return context.send("GET", uri, nil)
}
