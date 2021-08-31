package taobaoapi

import (
	"context"
	"github.com/saodd/taobaogo/taobaomodels"
	"net/url"
	"strconv"
)

// TaobaoPicturePicturesGet (图片空间)图片获取
// https://open.taobao.com/api.htm?docId=25591&docType=2&source=search
func (client *Client) TaobaoPicturePicturesGet(ctx context.Context, data *TaobaoPicturePicturesGetRequest, session string) (*TaobaoPicturePicturesGetResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.picture.pictures.get",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoPicturePicturesGetResponse `json:"picture_pictures_get_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoPicturePicturesGetResponse struct {
	Pictures struct {
		Picture []struct {
			PictureId         int                         `json:"picture_id"`
			PictureCategoryId int                         `json:"picture_category_id"`
			PicturePath       string                      `json:"picture_path"`
			Title             string                      `json:"title"`
			Sizes             int                         `json:"sizes"`
			Pixel             string                      `json:"pixel"`
			Status            string                      `json:"status"`
			Deleted           string                      `json:"deleted"`
			Created           taobaomodels.TaobaoDatetime `json:"created"`
			Modified          taobaomodels.TaobaoDatetime `json:"modified"`
			Referenced        bool                        `json:"referenced"`
			Md5               string                      `json:"md5"`
			ClientType        string                      `json:"client_type"`
		} `json:"picture"`
	} `json:"pictures"`
	RequestId string `json:"request_id"`
}

type TaobaoPicturePicturesGetRequest struct {
	RequestFields TaobaoPicturePicturesGetFields

	PictureCategoryId int    `json:"picture_category_id,omitempty"`
	PictureId         int    `json:"picture_id,omitempty"`
	OrderBy           string `json:"order_by,omitempty"`
	Title             string `json:"title,omitempty"`
	PageSize          int    `json:"page_size,omitempty"`
	CurrentPage       int    `json:"current_page,omitempty"`
	// ...
}
type TaobaoPicturePicturesGetFields struct {
	PictureCategoryId bool
	PictureId         bool
	OrderBy           bool
	Title             bool
	PageSize          bool
	CurrentPage       bool
}

func (r *TaobaoPicturePicturesGetRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	if r.RequestFields.PictureCategoryId {
		m["picture_category_id"] = strconv.Itoa(r.PictureCategoryId)
	}
	if r.RequestFields.PictureId {
		m["picture_id"] = strconv.Itoa(r.PictureId)
	}
	if r.RequestFields.OrderBy {
		m["order_by"] = r.OrderBy
	}
	if r.RequestFields.Title {
		m["title"] = r.Title
	}
	if r.RequestFields.PageSize {
		m["page_size"] = strconv.Itoa(r.PageSize)
	}
	if r.RequestFields.CurrentPage {
		m["current_page"] = strconv.Itoa(r.CurrentPage)
	}
	return m
}

func (r *TaobaoPicturePicturesGetRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

func (r *TaobaoPicturePicturesGetRequest) Valid() error {
	return nil
}
