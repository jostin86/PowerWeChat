package security

import (
	"github.com/ArtisanCloud/PowerLibs/v2/object"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/power"
	response2 "github.com/ArtisanCloud/PowerWeChat/v2/src/kernel/response"
	"github.com/ArtisanCloud/PowerWeChat/v2/src/miniProgram/security/response"
)

type Client struct {
	*kernel.BaseClient
}

// 向插件开发者发起使用插件的申请
// https://developers.weixin.qq.com/miniprogram/dev/framework/security.imgSecCheck.html
func (comp *Client) ImgSecCheck(path string, form *power.HashMap) (*response2.ResponseMiniProgram, error) {

	result := &response2.ResponseMiniProgram{}

	var files *object.HashMap
	if path != "" {
		files = &object.HashMap{
			"media": path,
		}
	}

	var formData *kernel.UploadForm
	if form != nil {
		formData = &kernel.UploadForm{
			FileName: (*form)["name"].(string),
			Contents: []*kernel.UploadContent{
				&kernel.UploadContent{
					Name:  "media",
					Value: (*form)["value"],
				},
			},
		}
	}

	_, err := comp.HttpUpload("wxa/img_sec_check", files, formData, nil, nil, result)

	return result, err
}

// 异步校验图片/音频是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/framework/security.mediaCheckAsync-v1.html#请求地址
func (comp *Client) MediaCheckAsync(mediaURL string, mediaType int, version int, openID string, scene int) (*response.ResponseSecurityMediaCheckASync, error) {

	result := &response.ResponseSecurityMediaCheckASync{}

	data := &object.HashMap{
		"media_url":  mediaURL,
		"media_type": mediaType,
		"version":    version,
		"openid":     openID,
		"scene":      scene,
	}

	_, err := comp.HttpPostJson("wxa/media_check_async", data, nil, nil, result)

	return result, err
}

// 检查一段文本是否含有违法违规内容
// https://developers.weixin.qq.com/miniprogram/dev/framework/security.msgSecCheck-v1.html#HTTPS%20调用
func (comp *Client) MsgSecCheck(
	openID string, scene int, version int, content string,
	nickname string, title string, signature string) (*response.ResponseSecurityMsgCheckASync, error) {

	result := &response.ResponseSecurityMsgCheckASync{}

	data := &object.HashMap{
		"openid":    openID,
		"scene":     scene,
		"version":   version,
		"content":   content,
		"nickname":  nickname,
		"title":     title,
		"signature": signature,
	}

	_, err := comp.HttpPostJson("wxa/msg_sec_check", data, nil, nil, result)

	return result, err
}
