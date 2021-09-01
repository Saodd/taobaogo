package taobaoapi

import (
	"context"
	"net/url"
)

// TaobaoPicassoPictagIscontrolled 判断商家是否是主图管控商家
// https://open.taobao.com/api.htm?docId=57797&docType=2&source=search
func (client *Client) TaobaoPicassoPictagIscontrolled(ctx context.Context, data *TaobaoPicassoPictagIscontrolledRequest, session string) (*TaobaoPicassoPictagIscontrolledResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.picasso.pictag.iscontrolled",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoPicassoPictagIscontrolledResponse `json:"picasso_pictag_iscontrolled_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoPicassoPictagIscontrolledResponse struct {
	Result struct {
		Model struct {
			CloseControl     bool `json:"close_control"`
			SellerControlled bool `json:"seller_controlled"`
		} `json:"model"`
		Success bool `json:"success"`
	} `json:"result"`
	RequestId string `json:"request_id"`
}

type TaobaoPicassoPictagIscontrolledRequest struct {
	RequestFields TaobaoPicassoPictagIscontrolledFields `json:"-"`
}

type TaobaoPicassoPictagIscontrolledFields struct {
}

func (r *TaobaoPicassoPictagIscontrolledRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	return m
}

func (r *TaobaoPicassoPictagIscontrolledRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

func (r *TaobaoPicassoPictagIscontrolledRequest) Valid() error {
	return nil
}
