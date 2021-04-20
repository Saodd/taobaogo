package taobaoapi

import (
	"context"
	"fmt"
	"github.com/saodd/taobaogo/taobaomodels"
	"net/url"
	"strconv"
)

// TaobaoItemsOnsaleGet 获取当前会话用户出售中的商品列表
// https://open.taobao.com/api.htm?docId=18&docType=2
func (client *Client) TaobaoItemsOnsaleGet(ctx context.Context, data *TaobaoItemsOnsaleGetRequest, session string) (*TaobaoItemsOnsaleGetResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.items.onsale.get",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoItemsOnsaleGetResponse `json:"items_onsale_get_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoItemsOnsaleGetResponse struct {
	Items struct {
		Item []*taobaomodels.Item `json:"item"`
	} `json:"items"`
	TotalResults int    `json:"total_results"`
	RequestId    string `json:"request_id"`
}

type TaobaoItemsOnsaleGetRequest struct {
	Fields   string  `json:"fields"`
	Q        *string `json:"q"`
	PageNo   *int    `json:"page_no"`
	PageSize *int    `json:"page_size"`
	// ...
}

func (r *TaobaoItemsOnsaleGetRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	m["fields"] = fmt.Sprint(r.Fields)
	if r.Q != nil {
		m["q"] = *r.Q
	}
	if r.PageNo != nil {
		m["page_no"] = strconv.Itoa(*r.PageNo)
	}
	if r.PageSize != nil {
		m["page_size"] = strconv.Itoa(*r.PageSize)
	}
	return m
}

func (r *TaobaoItemsOnsaleGetRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}
