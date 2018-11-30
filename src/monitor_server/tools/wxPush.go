package tools

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// accessToken
type AccessToken struct {
	AccessTokenStr string `json:"access_token"`
}

// 发送给微信时参数
type CustomToUserParam struct {
	ToUser  string         `json:"touser"`
	MsgType string         `json:"msgtype"`
	AgentId int            `json:"agentid"`
	Safe    int            `json:"safe"`
	Text    TextMsgContent `json:"text"`
}
type TextMsgContent struct {
	Content string `json:"content"`
}

/**
 * 获取微信token
 */
func GetWxApiToken() (result []byte, err error) {
	// 获取token
	tokenUrl := fmt.Sprintf(IniConf.String("WxPushMsg::getTokenUrl"), IniConf.String("WxPushMsg::corpid"), IniConf.String("WxPushMsg::corpsecret"))
	res, err := GetRequest(tokenUrl)
	if err != nil {
		return nil, err
	}
	accToken := &AccessToken{}
	err = json.Unmarshal(res, accToken)
	if err != nil {
		return nil, err
	}

	result = []byte(accToken.AccessTokenStr)
	return result, nil
}

/**
 * 发送消息推送
 */
func SendWxMsg(toUser string, msg string) (result []byte, err error) {
	// 获取token
	accessToken, err := GetWxApiToken()
	if err != nil {
		return
	}
	if accessToken == nil {
		return nil, errors.New("get access_token fail.")
	}

	// 发送消息
	apiURL := fmt.Sprintf(IniConf.String("WxPushMsg::pushUrl"), string(accessToken))
    agentId, _ := IniConf.Int("WxPushMsg::agentId")
	tmpMsg := &CustomToUserParam{
		ToUser:  toUser,
		MsgType: "text",
		AgentId: agentId,
		Text:    TextMsgContent{Content: msg},
		Safe:    0,
	}
	body, err := json.MarshalIndent(tmpMsg, " ", "  ")
	if err != nil {
		return nil, err
	}
	postReq, err := http.NewRequest("POST", apiURL, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	postReq.Header.Set("Content-Type", "application/json; encoding=utf-8")
	client := &http.Client{}
	resp, err := client.Do(postReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
