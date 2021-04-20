package taobaoapi

import (
	"context"
	"encoding/json"
	"github.com/saodd/taobaogo/taobaomodels"
	"net/url"
)

// TaobaoTopAuthTokenCreate 获取Access Token
// https://open.taobao.com/api.htm?docId=25388&docType=2
func (client *Client) TaobaoTopAuthTokenCreate(ctx context.Context, data *TaobaoTopAuthTokenCreateRequest) (*TaobaoTopAuthTokenCreateResponse, error) {
	var sp = SystemParams{
		Method: "taobao.top.auth.token.create",
	}
	var res struct {
		Resp struct {
			TokenResult string `json:"token_result"`
			RequestId   string `json:"request_id"`
		} `json:"top_auth_token_create_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}

	// 因为淘宝很奇葩地返回一个经过转义的字符串，所以不得不多一个反序列化步骤
	var token taobaomodels.TopAuthToken
	if err := json.Unmarshal([]byte(res.Resp.TokenResult), &token); err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return &TaobaoTopAuthTokenCreateResponse{
		TokenResult: &token,
		RequestId:   res.Resp.RequestId,
	}, nil
}

type TaobaoTopAuthTokenCreateResponse struct {
	TokenResult *taobaomodels.TopAuthToken `json:"token_result"`
	RequestId   string                     `json:"request_id"`
}

type TaobaoTopAuthTokenCreateRequest struct {
	Code string `json:"code"`
}

func (r *TaobaoTopAuthTokenCreateRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	m["code"] = r.Code
	return m
}

func (r *TaobaoTopAuthTokenCreateRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}
