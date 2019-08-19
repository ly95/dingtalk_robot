package robot

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

// Robot class
type Robot struct {
	URI   string
	Token string
}

// NewRobot create Robot instance
func NewRobot(token string) *Robot {
	return &Robot{
		URI:   "https://oapi.dingtalk.com/robot/send",
		Token: token,
	}
}

// Send send message to dingtalk serve
func (r *Robot) Send(msg interface{}) (*Result, error) {
	url := r.URI + "?access_token=" + r.Token

	// b , err := json. Marshal ( group )
	msgBuf, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msgBuf))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		res := new(Result)
		if err := json.Unmarshal(body, res); err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, errors.New(string(body))
}
