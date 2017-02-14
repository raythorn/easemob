package easemob

// import (
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"net/http"
// 	"strings"
// )

// type EMCallback func(result []byte)

// type Context struct {
// 	host   string
// 	org    string
// 	app    string
// 	id     string
// 	secret string
// 	auth   bool
// 	token  string
// 	uuid   string
// 	expire int64
// }

// // func New(host, org, app, id, secret string, auth bool) *Context {
// // 	ctx := &Context{
// // 		host:   host,
// // 		org:    org,
// // 		app:    app,package easemob

// import (
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"mime/multipart"
// 	"net/http"
// 	"os"
// 	"strings"
// )

// type EMCallback func(result []byte)

// type EaseMob struct {
// 	host   string
// 	org    string
// 	app    string
// 	id     string
// 	secret string
// 	auth   bool
// 	token  string
// 	uuid   string
// 	expire int64
// }

// type User struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// 	Nickname string `json:"nickname,omitempty"`
// }

// func New(org, app, id, secret string, auth bool) *EaseMob {
// 	return &EaseMob{
// 		host:   "https://a1.easemob.com",
// 		org:    org,
// 		app:    app,
// 		id:     id,
// 		secret: secret,
// 		auth:   auth,
// 		token:  "",
// 		uuid:   "",
// 		expire: 0,
// 	}
// }

// func (em *EaseMob) Token() error {

// 	data := map[string]string{"grant_type": "client_credentials", "client_id": em.id, "client_secret": em.secret}
// 	uri := em.uri() + "/token"
// 	resp, err := em.send("POST", uri, data)
// 	if err != nil {
// 		return err
// 	}

// 	var result map[string]interface{} = nil
// 	err = json.Unmarshal([]byte(resp), &result)
// 	if err != nil {
// 		return err
// 	}

// 	if _, exist := result["access_token"]; exist {
// 		if val, ok := result["access_token"].(string); ok {
// 			em.token = val
// 		}
// 	}

// 	if _, exist := result["expires_in"]; exist {

// 		if val, ok := result["expires_in"].(float64); ok {
// 			em.expire = int64(val)
// 		}
// 	}

// 	if _, exist := result["application"]; exist {
// 		if val, ok := result["application"].(string); ok {
// 			em.uuid = val
// 		}
// 	}

// 	return nil
// }

// func (em *EaseMob) AddUser(user *User, cb EMCallback) error {

// 	uri := em.uri() + "/users"

// 	resp, err := em.send("POST", uri, user)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) AddUsers(users []*User, cb EMCallback) error {
// 	uri := em.uri() + "/users"
// 	resp, err := em.send("POST", uri, users)
// 	if err == nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) GetUser(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username
// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) GetUsers(limit int, cursor string, cb EMCallback) error {
// 	uri := em.uri() + "/users"
// 	var param bool = false
// 	if limit != 0 {
// 		uri += "?limit=" + fmt.Sprintf("%d", limit)
// 		param = true
// 	}

// 	if cursor != "" {
// 		if param {
// 			uri += "&"
// 		} else {
// 			uri += "?"
// 		}

// 		uri += "cursor=" + cursor
// 	}

// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) DeleteUser(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username
// 	resp, err := em.send("DELETE", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) DeleteUsers(limit int, cb EMCallback) error {

// 	if limit == 0 {
// 		return errors.New("Invalid params")
// 	}

// 	uri := em.uri() + "/users?limit=" + fmt.Sprintf("%d", limit)
// 	resp, err := em.send("DELETE", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) SetPassword(username, password string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/password"
// 	data := map[string]string{"newpassword": password}
// 	resp, err := em.send("PUT", uri, data)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) SetUsername(username, nickname string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username
// 	data := map[string]string{"nickname": nickname}
// 	resp, err := em.send("PUT", uri, data)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) AddFriend(username, friend string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/contacts/users/" + friend
// 	resp, err := em.send("POST", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) DeleteFriend(username, friend string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/contacts/users/" + friend
// 	resp, err := em.send("DELETE", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) GetFriends(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/contacts/users"
// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) GetBlockUsers(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/blocks/users"
// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) BlockUsers(username string, blockusers []string, cb EMCallback) error {
// 	if len(blockusers) == 0 {
// 		return errors.New("Invalid Params")
// 	}

// 	uri := em.uri() + "/users/" + username + "/blocks/users"
// 	data := map[string][]string{"usernames": blockusers}
// 	resp, err := em.send("POST", uri, data)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) DeleteBlockUser(username, blockuser string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/blocks/users/" + blockuser
// 	resp, err := em.send("DELETE", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) GetUserStatus(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/status"
// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) GetOfflineMessageCount(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/offline_msg_count"
// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) GetOfflineMessageStatus(username, msgid string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/offline_msg_status/" + msgid
// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) ActivateUser(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/activate"
// 	resp, err := em.send("POST", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) DeactivateUser(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/deactivate"
// 	resp, err := em.send("POST", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) ForceUserOffline(username string, cb EMCallback) error {
// 	uri := em.uri() + "/users/" + username + "/disconnect"
// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) GetMessages(limit int, cursor string, sql string, cb EMCallback) error {
// 	uri := em.uri() + "/chatmessages"
// 	var params bool = false

// 	if limit > 0 {
// 		uri += "?limit=" + fmt.Sprintf("%d", limit)
// 		params = true
// 	}

// 	if cursor != "" {
// 		if params {
// 			uri += "&"
// 		} else {
// 			uri += "?"
// 		}
// 		uri += "cursor=" + cursor
// 		params = true
// 	}

// 	if sql != "" {
// 		if params {
// 			uri += "&"
// 		} else {
// 			uri += "?"
// 		}
// 		uri += "ql=" + sql
// 	}

// 	resp, err := em.send("GET", uri, nil)
// 	if err != nil {
// 		return err
// 	}

// 	cb(resp)

// 	return nil
// }

// func (em *EaseMob) Upload(file string, cb EMCallback) error {
// 	var b bytes.Buffer
// 	writer := multipart.NewWriter(&b)
// 	f, err := os.Open(file)
// 	if err != nil {
// 		return err
// 	}

// 	defer f.Close()

// 	formfile, err := writer.CreateFormFile("file", file)
// 	if err != nil {
// 		return err
// 	}

// 	if _, err = io.Copy(formfile, f); err != nil {
// 		return err
// 	}

// 	writer.Close()

// 	uri := em.uri() + "/chatfiles"
// 	client := &http.Client{}
// 	req, err := http.NewRequest("POST", uri, &b)
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Content-Type", writer.FormDataContentType())
// 	req.Header.Set("restrict-access", "true")
// 	req.Header.Set("Authorization", "Bearer "+em.token)

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return err
// 	}

// 	defer resp.Body.Close()
// 	data, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return err
// 	}

// 	if resp.StatusCode != http.StatusOK {
// 		errstr := "Unkown Request Error!"
// 		var result map[string]interface{} = nil
// 		err = json.Unmarshal(data, &result)
// 		if err != nil {
// 			return errors.New(errstr)
// 		}

// 		if _, exist := result["error_description"]; exist {
// 			if val, ok := result["error_description"].(string); ok {
// 				errstr = val
// 			}
// 		}
// 		return errors.New(errstr)
// 	}

// 	cb(data)

// 	return nil
// }

// func (em *EaseMob) uri() string {
// 	return em.host + "/" + em.org + "/" + em.app
// }

// func (em *EaseMob) send(method, uri string, data interface{}, header ...map[string]string) (result []byte, err error) {

// 	var body io.Reader = nil

// 	if data != nil {
// 		b, err := json.Marshal(data)
// 		if err != nil {
// 			return nil, err
// 		}

// 		fmt.Println(string(b))
// 		body = strings.NewReader(string(b))
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, uri, body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Content-Type", "application/json;charset=utf-8")
// 	if em.token != "" && em.auth {
// 		req.Header.Set("Authorization", "Bearer "+em.token)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.StatusCode >= 400 {
// 		errstr := "Unkown Request Error!"
// 		var result map[string]interface{} = nil
// 		err = json.Unmarshal(b, &result)
// 		if err != nil {
// 			return nil, errors.New(errstr)
// 		}

// 		if _, exist := result["error_description"]; exist {
// 			if val, ok := result["error_description"].(string); ok {
// 				errstr = val
// 			}
// 		}
// 		return nil, errors.New(errstr)
// 	}

// 	return b, nil
// }

// 		id:     id,
// 		secret: secret,
// 		auth:   auth,
// 		token:  "",
// 		uuid:   "",
// 		expire: 0,
// 	}

// 	return ctx
// }

// func (c *Context) Init() error {

// 	data := map[string]string{"grant_type": "client_credentials", "client_id": c.id, "client_secret": c.secret}
// 	uri := c.Uri() + "/token"
// 	resp, err := c.Send("POST", uri, data)
// 	if err != nil {
// 		return err
// 	}

// 	var result map[string]interface{} = nil
// 	err = json.Unmarshal([]byte(resp), &result)
// 	if err != nil {
// 		return err
// 	}

// 	if _, exist := result["access_token"]; exist {
// 		if val, ok := result["access_token"].(string); ok {
// 			c.token = val
// 		}
// 	}

// 	if _, exist := result["expires_in"]; exist {

// 		if val, ok := result["expires_in"].(float64); ok {
// 			c.expire = int64(val)
// 		}
// 	}

// 	if _, exist := result["application"]; exist {
// 		if val, ok := result["application"].(string); ok {
// 			c.uuid = val
// 		}
// 	}

// 	if c.token == "" || c.expire == 0 || c.uuid == "" {
// 		return errors.New("Initialise easemob fail")
// 	}

// 	return nil
// }

// func (c *Context) Uri() string {
// 	return c.host + "/" + c.org + "/" + c.app
// }

// func (c *Context) Send(method, uri string, data interface{}, header ...map[string]string) (result []byte, err error) {
// 	var body io.Reader = nil

// 	if data != nil {
// 		b, err := json.Marshal(data)
// 		if err != nil {
// 			return nil, err
// 		}

// 		fmt.Println(string(b))
// 		body = strings.NewReader(string(b))
// 	}

// 	client := &http.Client{}
// 	req, err := http.NewRequest(method, uri, body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req.Header.Set("Content-Type", "application/json;charset=utf-8")
// 	if c.token != "" && c.auth {
// 		req.Header.Set("Authorization", "Bearer "+c.token)
// 	}

// 	resp, err := client.Do(req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	defer resp.Body.Close()
// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if resp.StatusCode >= 400 {
// 		errstr := "Unkown Request Error!"
// 		var result map[string]interface{} = nil
// 		err = json.Unmarshal(b, &result)
// 		if err != nil {
// 			return nil, errors.New(errstr)
// 		}

// 		if _, exist := result["error_description"]; exist {
// 			if val, ok := result["error_description"].(string); ok {
// 				errstr = val
// 			}
// 		}
// 		return nil, errors.New(errstr)
// 	}

// 	return b, nil
// }
