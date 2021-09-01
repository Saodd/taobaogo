package taobaoapi

import (
	"context"
	"encoding/json"
	"net/url"
)

// TaobaoPicassoPictagReport 上报商家强行使用违规图片记录
// https://open.taobao.com/api.htm?docId=57796&docType=2&source=search
func (client *Client) TaobaoPicassoPictagReport(ctx context.Context, data *TaobaoPicassoPictagReportRequest, session string) (*TaobaoPicassoPictagReportResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.picasso.pictag.report",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoPicassoPictagReportResponse `json:"???"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoPicassoPictagReportResponse struct {
	RequestId string `json:"request_id"`
}

type TaobaoPicassoPictagReportRequest struct {
	RequestFields TaobaoPicassoPictagReportFields `json:"-"`
	ItemId        int64                           `json:"item_id"`
	PicUrl        string                          `json:"pic_url"`
}

type TaobaoPicassoPictagReportFields struct {
}

func (r *TaobaoPicassoPictagReportRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	j, _ := json.Marshal(r)
	m["param"] = string(j)
	return m
}

func (r *TaobaoPicassoPictagReportRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

func (r *TaobaoPicassoPictagReportRequest) Valid() error {
	return nil
}
