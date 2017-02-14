package easemob

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Context struct {
	host   string
	org    string
	app    string
	id     string
	secret string
	auth   bool
	token  string
	uuid   string
	expire int64
}

func (c *Context) init() error {

	data := map[string]string{"grant_type": "client_credentials", "client_id": c.id, "client_secret": c.secret}
	uri := c.uri() + "/token"
	resp, err := c.send("POST", uri, data)
	if err != nil {
		return err
	}

	var result map[string]interface{} = nil
	err = json.Unmarshal([]byte(resp), &result)
	if err != nil {
		return err
	}

	if _, exist := result["access_token"]; exist {
		if val, ok := result["access_token"].(string); ok {
			c.token = val
		}
	}

	if _, exist := result["expires_in"]; exist {

		if val, ok := result["expires_in"].(float64); ok {
			c.expire = int64(val)
		}
	}

	if _, exist := result["application"]; exist {
		if val, ok := result["application"].(string); ok {
			c.uuid = val
		}
	}

	if c.token == "" || c.expire == 0 || c.uuid == "" {
		return errors.New("Initialise easemob fail")
	}

	return nil
}

func (c *Context) uri() string {
	return c.host + "/" + c.org + "/" + c.app
}

func (c *Context) send(method, uri string, data interface{}, header ...map[string]string) (result []byte, err error) {
	var body io.Reader = nil

	if data != nil {
		b, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}

		fmt.Println(string(b))
		body = strings.NewReader(string(b))
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if c.token != "" && c.auth {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode >= 400 {
		errstr := "Unkown Request Error!"
		var result map[string]interface{} = nil
		err = json.Unmarshal(b, &result)
		if err != nil {
			return nil, errors.New(errstr)
		}

		if _, exist := result["error_description"]; exist {
			if val, ok := result["error_description"].(string); ok {
				errstr = val
			}
		}
		return nil, errors.New(errstr)
	}

	return b, nil
}
