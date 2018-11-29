package tools

import(
    "fmt"
	"errors"
	"encoding/json"
)

type AccessToken struct {
	AccessTokenStr string `json:"access_token"`
}

/**
 * 获取微信token
 */
func GetWxApiToken() (result []byte, err error) {
	// 获取token
    tokenUrl := fmt.Sprintf(IniConf.String("WxPushMsg::getTokenUrl"), IniConf.String("WxPushMsg::corpid"), IniConf.String("WxPushMsg::corpsecret"))
    res, err := GetRequest(tokenUrl);
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
func SendWxMsg(msg string) (err error) {
	// 获取token
    accessToken, err := GetWxApiToken();
    if err != nil {
		return err
    }
    if accessToken == nil {
		return errors.New("get access_token fail.")
    }

	fmt.Println(msg)

	return nil
}
