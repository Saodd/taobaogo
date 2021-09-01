package taobaoapi

import (
	"context"
	"fmt"
	"net/url"
)

// AlibabaItemEditSchemaGet 商品编辑获取schema信息
// https://open.taobao.com/api.htm?docId=53967&docType=2
func (client *Client) AlibabaItemEditSchemaGet(ctx context.Context, data *AlibabaItemEditSchemaGetRequest, session string) (*AlibabaItemEditSchemaGetResponse, error) {
	var sp = SystemParams{
		Method:  "alibaba.item.edit.schema.get",
		Session: &session,
	}
	var res struct {
		Resp *AlibabaItemEditSchemaGetResponse `json:"alibaba_item_edit_schema_get_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type AlibabaItemEditSchemaGetResponse struct {
	UpdateItemResult string `json:"result"`
}

type AlibabaItemEditSchemaGetRequest struct {
	RequestFields AlibabaItemEditSchemaGetFields `json:"-"`

	ItemId  int64  `json:"item_id"`
	BizType string `json:"biz_type"` // 可选
}

type AlibabaItemEditSchemaGetFields struct {
	BizType bool
}

func (r *AlibabaItemEditSchemaGetRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	m["item_id"] = fmt.Sprint(r.ItemId)
	if r.RequestFields.BizType {
		m["biz_type"] = r.BizType
	}
	return m
}

func (r *AlibabaItemEditSchemaGetRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

func (r *AlibabaItemEditSchemaGetRequest) Valid() error {
	return nil
}
