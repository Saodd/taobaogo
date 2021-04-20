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
	RequestFields

	Fields        string                      `json:"fields"`
	Q             string                      `json:"q"`              // 可选
	Banner        string                      `json:"banner"`         // 可选
	Cid           int                         `json:"cid"`            // 可选
	SellerCids    string                      `json:"seller_cids"`    // 可选
	PageNo        int                         `json:"page_no"`        // 可选
	PageSize      int                         `json:"page_size"`      // 可选
	HasDiscount   string                      `json:"has_discount"`   // 可选
	OrderBy       string                      `json:"order_by"`       // 可选
	IsTaobao      bool                        `json:"is_taobao"`      // 可选
	IsEx          bool                        `json:"is_ex"`          // 可选
	StartModified taobaomodels.TaobaoDatetime `json:"start_modified"` // 可选
	EndModified   taobaomodels.TaobaoDatetime `json:"end_modified"`   // 可选
	AuctionType   string                      `json:"auction_type"`   // 可选
}

func (r *TaobaoItemsInventoryGetRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	m["fields"] = fmt.Sprint(r.Fields)
	if r.RequestFields["q"] {
		m["q"] = r.Q
	}
	if r.RequestFields["page_no"] {
		m["page_no"] = strconv.Itoa(r.PageNo)
	}
	if r.RequestFields["page_size"] {
		m["page_size"] = strconv.Itoa(r.PageSize)
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
