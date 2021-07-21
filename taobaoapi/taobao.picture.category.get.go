package taobaoapi

import (
	"context"
	"github.com/saodd/taobaogo/constants"
	"github.com/saodd/taobaogo/taobaomodels"
	"net/url"
	"strconv"
)

// TaobaoPictureCategoryGet 获取图片分类信息
// https://open.taobao.com/api.htm?docId=137&docType=2
func (client *Client) TaobaoPictureCategoryGet(ctx context.Context, data *TaobaoPictureCategoryGetRequest, session string) (*TaobaoPictureCategoryGetResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.picture.category.get",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoPictureCategoryGetResponse `json:"picture_category_get_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

const (
	TaobaoPictureCategoryGetType1 = "sys-fixture"
	TaobaoPictureCategoryGetType2 = "sys-auction"
	TaobaoPictureCategoryGetType3 = "user-define"
)

type TaobaoPictureCategoryGetResponse struct {
	PictureCategories struct {
		PictureCategory []struct {
			PictureCategoryId   int    `json:"picture_category_id"`
			PictureCategoryName string `json:"picture_category_name"`
			Type                string `json:"type"` // TaobaoPictureCategoryGetType
			Created             string `json:"created"`
			Modified            string `json:"modified"`
			Position            int    `json:"position"`
			ParentId            int    `json:"parent_id"`
		} `json:"picture_category"`
	} `json:"picture_categories"`
	RequestId string `json:"request_id"`
}

type TaobaoPictureCategoryGetRequest struct {
	RequestFields

	Type                string                      `json:"type,omitempty"` // TaobaoPictureCategoryGetType
	PictureCategoryId   int                         `json:"picture_category_id,omitempty"`
	PictureCategoryName string                      `json:"picture_category_name,omitempty"`
	ParentId            int                         `json:"parent_id,omitempty"`
	ModifiedTime        taobaomodels.TaobaoDatetime `json:"modified_time,omitempty"`
}

func (r *TaobaoPictureCategoryGetRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	if r.RequestFields["type"] {
		m["type"] = r.Type
	}
	if r.RequestFields["picture_category_id"] {
		m["picture_category_id"] = strconv.Itoa(r.PictureCategoryId)
	}
	if r.RequestFields["picture_category_name"] {
		m["picture_category_name"] = r.PictureCategoryName
	}
	if r.RequestFields["parent_id"] {
		m["parent_id"] = strconv.Itoa(r.ParentId)
	}
	if r.RequestFields["modified_time"] {
		m["modified_time"] = r.ModifiedTime.Format(constants.TaobaoDatetimeFormat)
	}
	return m
}

func (r *TaobaoPictureCategoryGetRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}
