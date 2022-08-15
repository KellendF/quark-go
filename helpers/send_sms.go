package helpers

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"regexp"

	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	dysmsapi20170525 "github.com/alibabacloud-go/dysmsapi-20170525/v2/client"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/parnurzeal/gorequest"
	"github.com/quarkcms/quark-go/pkg/ui/admin/utils"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func AliyunClient(accessKeyId *string, accessKeySecret *string) (_result *dysmsapi20170525.Client, _err error) {
	config := &openapi.Config{
		// 您的 AccessKey ID
		AccessKeyId: accessKeyId,
		// 您的 AccessKey Secret
		AccessKeySecret: accessKeySecret,
	}
	// 访问的域名
	config.Endpoint = tea.String("dysmsapi.aliyuncs.com")
	_result = &dysmsapi20170525.Client{}
	_result, _err = dysmsapi20170525.NewClient(config)
	return _result, _err
}

func AliyunSendSms(phone string, code string) (bool, string) {

	// 匹配规则
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	if !reg.MatchString(phone) {
		return false, "手机号格式错误！"
	}

	accessKeyId := utils.WebConfig("ACCESS_KEY_ID")
	accessKeySecret := utils.WebConfig("ACCESS_KEY_SECRET")
	signName := utils.WebConfig("SIGN_NAME")
	templateCode := utils.WebConfig("TEMPLATE_CODE")

	client, _err := AliyunClient(tea.String(accessKeyId), tea.String(accessKeySecret))
	if _err != nil {
		return false, _err.Error()
	}

	sendSmsRequest := &dysmsapi20170525.SendSmsRequest{
		PhoneNumbers:  tea.String(phone),
		SignName:      tea.String(signName),
		TemplateCode:  tea.String(templateCode),
		TemplateParam: tea.String("{\"code\":\"" + code + "\"}"),
	}
	runtime := &util.RuntimeOptions{}
	tryErr := func() (_e error) {
		defer func() {
			if r := tea.Recover(recover()); r != nil {
				_e = r
			}
		}()
		// 复制代码运行请自行打印 API 的返回值
		_, _err = client.SendSmsWithOptions(sendSmsRequest, runtime)
		if _err != nil {
			return _err
		}

		return nil
	}()

	if tryErr != nil {
		var error = &tea.SDKError{}
		if _t, ok := tryErr.(*tea.SDKError); ok {
			error = _t
		} else {
			error.Message = tea.String(tryErr.Error())
		}
		// 如有需要，请打印 error
		util.AssertAsString(error.Message)
	}

	if _err != nil {
		return false, _err.Error()
	}

	return true, ""
}

// 发送短信
func SiooSendSms(phone string, content string) (bool, string) {
	// 匹配规则
	regRuler := "^1[345789]{1}\\d{9}$"

	// 正则调用规则
	reg := regexp.MustCompile(regRuler)

	// 返回 MatchString 是否匹配
	if !reg.MatchString(phone) {
		return false, "手机号格式错误！"
	}

	uid := utils.WebConfig("SIOO_UID")
	password := utils.WebConfig("SIOO_PASSWORD")

	if uid == "" || password == "" {
		return false, "接口配置错误！"
	}

	md5Byte := md5.Sum([]byte(password))
	md5Password := fmt.Sprintf("%x", md5Byte)

	// 接口url
	url := "https://submit.10690221.com/send/ordinarykv?uid=" + uid + "&password=" + md5Password + "&mobile=" + phone + "&msg=" + content

	request := gorequest.New()
	_, body, _ := request.Get(url).End()

	type Data struct {
		Msg   string
		Code  int
		MsgId string
	}

	var data Data
	json.Unmarshal([]byte(body), &data)

	if data.Code == 0 {
		return true, "发送成功"
	} else {
		return true, data.Msg
	}
}
