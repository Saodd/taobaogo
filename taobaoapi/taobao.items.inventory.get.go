package taobaoapi

import (
	"context"
	"fmt"
	"github.com/saodd/taobaogo/taobaomodels"
	"net/url"
	"strconv"
)

// TaobaoItemsInventoryGet 得到当前会话用户库存中的商品列表
// https://open.taobao.com/api.htm?docId=162&docType=2
func (client *Client) TaobaoItemsInventoryGet(ctx context.Context, data *TaobaoItemsInventoryGetRequest, session string) (*TaobaoItemsInventoryGetResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.items.inventory.get",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoItemsInventoryGetResponse `json:"items_inventory_get_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoItemsInventoryGetResponse struct {
	Items struct {
		Item []*taobaomodels.Item `json:"item"`
	} `json:"items"`
	TotalResults int    `json:"total_results"`
	RequestId    string `json:"request_id"`
}

type TaobaoItemsInventoryGetRequest struct {
	Fields        string                       `json:"fields"`
	Q             *string                      `json:"q"`
	Banner        *string                      `json:"banner"`
	Cid           *int                         `json:"cid"`
	SellerCids    *string                      `json:"seller_cids"`
	PageNo        *int                         `json:"page_no"`
	PageSize      *int                         `json:"page_size"`
	HasDiscount   *string                      `json:"has_discount"`
	OrderBy       *string                      `json:"order_by"`
	IsTaobao      *bool                        `json:"is_taobao"`
	IsEx          *bool                        `json:"is_ex"`
	StartModified *taobaomodels.TaobaoDatetime `json:"start_modified"`
	EndModified   *taobaomodels.TaobaoDatetime `json:"end_modified"`
	AuctionType   *string                      `json:"auction_type"`
}

func (r *TaobaoItemsInventoryGetRequest) ToSignMap() map[string]string {
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
	// ...
	return m
}

func (r *TaobaoItemsInventoryGetRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}
