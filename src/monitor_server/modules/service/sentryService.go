package service

import (
    "fmt"
	"net/http"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"monitor_server/components"
	. "monitor_server/tools"
    "github.com/astaxie/beego/logs"
	simplejson "github.com/bitly/go-simplejson"
)

// sentry相关控制器
type SentryService struct {
	BaseService
}

/**
 * 发送微信消息
 */
func (this SentryService) EventCheck(c *gin.Context) {
    // 调用sentry接口，获得事件列表
	client := &http.Client{}
	req, err := http.NewRequest("GET", "http://mis.dev02.aplum-inc.com:8004/api/0/projects/sentry/aplum_www/events/", nil)
    if err != nil {
        ApiReturnError(c, 200, 1002, "get sentry data fail")
    }
	req.Header.Add("Authorization", "Bearer 907dac1e85f1461a816cc2a164db6ecb0c62d22bf4c54f26bc58e82f59b7ce88")
	resp, err := client.Do(req)
    if err != nil {
        ApiReturnError(c, 200, 1002, err.Error())
    }
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
        ApiReturnError(c, 200, 1002, err.Error())
	}

	// 处理结果
	jsonRes, err := simplejson.NewJson(res)
    if err != nil {
        ApiReturnError(c, 200, 1002, err.Error())
    }
	rows, err := jsonRes.Array()
	if err != nil {
        ApiReturnError(c, 200, 1002, err.Error())
	}
    var key string = "SEND_MSG_"
    var tmpKey string = ""
	for _, row := range rows {
		if each_map, ok := row.(map[string]interface{}); ok {
            // 判断是否已经发过报警信息
            tmpKey = fmt.Sprintf("%s%s", key, each_map["eventID"].(string))
            tmpRes := components.RedisHandleObj.GetString(tmpKey)
            if tmpRes == "ok" {
                continue
            }

			if each_map2, ok2 := each_map["entries"].([]interface{}); ok2 {
				for _, v := range each_map2 {
					if each_map3, ok3 := v.(map[string]interface{}); ok3 {
						if each_map4, ok4 := each_map3["data"].(map[string]interface{}); ok4 {
							if tmpMsg, ok5 := each_map4["message"]; ok5 {
                                // 发送告警消息
                                _, err := components.SendWxMsg("@all", tmpMsg.(string))
                                if err != nil {
                                    logs.Error(1024, err.Error())
                                }

                                // 标志已经发送过了
                                components.RedisHandleObj.Set(tmpKey, "ok", 864000)
							}
						}
					}
				}
			}
		}
	}

	ApiReturnSuccess(c, 0, "success")
}
