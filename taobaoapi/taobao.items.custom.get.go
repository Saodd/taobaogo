package taobaoapi

import (
	"context"
	"github.com/saodd/taobaogo/taobaomodels"
	"net/url"
)

// TaobaoItemsCustomGet 根据外部ID取商品
// https://open.taobao.com/api.htm?docId=163&docType=2&source=search
func (client *Client) TaobaoItemsCustomGet(ctx context.Context, data *TaobaoItemsCustomGetRequest, session string) (*TaobaoItemsCustomGetResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.items.custom.get",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoItemsCustomGetResponse `json:"items_custom_get_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoItemsCustomGetResponse struct {
	RequestId string `json:"request_id"`
	Items     struct {
		Item []taobaomodels.Item `json:"item"`
	} `json:"items"`
}

type TaobaoItemsCustomGetRequest struct {
	OuterId string `json:"outer_id"`
	Fields  string `json:"fields"`
}

func (r *TaobaoItemsCustomGetRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	m["outer_id"] = r.OuterId
	m["fields"] = r.Fields
	return m
}

func (r *TaobaoItemsCustomGetRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

func (r *TaobaoItemsCustomGetRequest) Valid() error {
	return nil
}
