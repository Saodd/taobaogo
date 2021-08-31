package taobaoapi

import (
	"context"
	"encoding/json"
	"net/url"
)

// TaobaoPicassoPictagCheckimg 检测图片是否违规
// https://open.taobao.com/api.htm?docId=57795&docType=2&source=search
func (client *Client) TaobaoPicassoPictagCheckimg(ctx context.Context, data *TaobaoPicassoPictagCheckimgRequest, session string) (*TaobaoPicassoPictagCheckimgResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.picasso.pictag.checkimg",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoPicassoPictagCheckimgResponse `json:"picasso_pictag_checkimg_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoPicassoPictagCheckimgResponse struct {
	Result struct {
		Success bool `json:"success"`
		Model   struct {
			ReasonMsg   string `json:"reason_msg"`
			ReasonCode  string `json:"reason_code"`
			CheckResult string `json:"check_result"` // 检测结果，有3种值： OK表示检测通过； WARNING表示有风险；FAIL表示检测不通过
		} `json:"model"`
		MsgInfo string `json:"msg_info"`
		MsgCode string `json:"msg_code"`
	} `json:"result"`
	RequestId string `json:"request_id"`
}

type TaobaoPicassoPictagCheckimgRequest struct {
	ItemId  int64  `json:"item_id"`
	Biz     string `json:"biz"` // 目前有2种：pictag表示打标检测底图及修饰图片，upload表示直接上传图片
	PicUrl  string `json:"pic_url,omitempty"`
	PicData []byte `json:"pic_data,omitempty"`
}

func (r *TaobaoPicassoPictagCheckimgRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	j, _ := json.Marshal(r)
	m["param"] = string(j)
	return m
}

func (r *TaobaoPicassoPictagCheckimgRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

func (r *TaobaoPicassoPictagCheckimgRequest) Valid() error {
	return nil
}
