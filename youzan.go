package yzapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const URLOauthToken string = "https://open.youzanyun.com/auth/token"

// URLAPIBase API URL
const URLAPIBase string = "https://open.youzanyun.com/api/%s/%s?access_token=%s"

func (this *Server) GetAccessToken() (string, error) {
	token := this.AccessToken.AccessToken
	expires := this.AccessToken.Expires
	fmt.Println(token, expires, time.Now().UnixNano()/1e6)
	if token == "" || expires < time.Now().Unix()/1e6 {
		fmt.Print("重新获取Token\n")
		err := this.InitAccessToken()
		if err == nil {
			return this.AccessToken.AccessToken, nil
		} else {
			return "", err
		}
	} else {
		fmt.Print("缓存获取Token\n")
		return token, nil
	}
}

func (this *Server) InitAccessToken() error {
	params := map[string]interface{}{
		"client_id":      this.ClientID,
		"client_secret":  this.ClientSecret,
		"authorize_type": this.AuthorizeType,
		"grant_id":       this.GrantID,
	}
	var msg AccessTokenMsg
	json_file, err := ioutil.ReadFile("accesstoken.json")
	if err == nil {
		json.Unmarshal(json_file, &msg)
		this.AccessToken = msg.AccessToken
	}
	if this.AccessToken.Expires >= time.Now().Unix() {
		return nil
	}
	postData := BuildPostParams(params)
	rs, err := PostJSON(URLOauthToken, postData)
	if err != nil {
		return err
	} else {
		json.Unmarshal(rs, &msg)
		if msg.Code == 200 {
			this.AccessToken = msg.AccessToken
			this.AccessTokenWriteToFile("accesstoken.json")
			return nil
		} else {
			this.AccessToken.AccessToken = ""
			return errors.New(msg.Message)

		}
	}

}

//获取的Token信息存入到文件
func (this *Server) AccessTokenWriteToFile(filename string) error {
	token := this.AccessToken
	token_byte, err := json.Marshal(token)
	err = ioutil.WriteFile(filename, token_byte, os.ModePerm)
	return err
}

//初始换一个有赞Serve
func NewServer(ClientID, ClientSecret string, GrantID int64, Refresh bool) *Server {
	s := new(Server)
	s.AuthorizeType = "silent"
	s.ClientID = ClientID
	s.ClientSecret = ClientSecret
	s.GrantID = GrantID
	s.Refresh = Refresh
	s.InitAccessToken()
	return s
}

func (this *Server) Call(apiName, apiVersion string, apiParams map[string]interface{}) ([]byte, error) {
	tk, _ := this.GetAccessToken()
	url := BuildURL(apiName, apiVersion, tk)
	params := BuildPostParams(apiParams)
	body, err := PostJSON(url, params)
	if err != nil {
		return nil, err
	} else {
		return body, nil
	}
}
