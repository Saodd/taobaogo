package taobaoapi

import (
	"context"
	"fmt"
	"net/url"
)

// TaobaoItemSellerGet 获取单个商品详细信息
// https://open.taobao.com/api.htm?docId=24625&docType=2
func (client *Client) TaobaoItemSellerGet(ctx context.Context, data *TaobaoItemSellerGetRequest, session string) (*TaobaoItemSellerGetResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.item.seller.get",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoItemSellerGetResponse `json:"item_seller_get_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoItemSellerGetResponse struct {
	Item struct {
		Desc         string `json:"desc"`
		Newprepay    string `json:"newprepay"`
		NumIid       int64  `json:"num_iid"`
		WirelessDesc string `json:"wireless_desc"`
		// ...还有很多
	} `json:"item"`
	RequestId string `json:"request_id"`
}

type TaobaoItemSellerGetRequest struct {
	RequestFields

	Fields string `json:"fields"`
	NumIid int64  `json:"num_iid"`
}

func (r *TaobaoItemSellerGetRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	m["fields"] = r.Fields
	m["num_iid"] = fmt.Sprint(r.NumIid)
	return m
}

func (r *TaobaoItemSellerGetRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

type T struct {
	ItemSellerGetResponse TaobaoItemSellerGetResponse `json:"item_seller_get_response"`
}
