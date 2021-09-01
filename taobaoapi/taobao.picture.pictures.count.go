package taobaoapi

import (
	"context"
	"net/url"
	"strconv"
)

// TaobaoPicturePicturesCount (图片空间)图片总数查询
// https://open.taobao.com/api.htm?docId=25592&docType=2&source=search
func (client *Client) TaobaoPicturePicturesCount(ctx context.Context, data *TaobaoPicturePicturesCountRequest, session string) (*TaobaoPicturePicturesCountResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.picture.pictures.count",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoPicturePicturesCountResponse `json:"picture_pictures_count_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoPicturePicturesCountResponse struct {
	Totals    int    `json:"totals"`
	RequestId string `json:"request_id"`
}

type TaobaoPicturePicturesCountRequest struct {
	RequestFields TaobaoPicturePicturesCountFields `json:"-"`

	PictureCategoryId int    `json:"picture_category_id,omitempty"`
	PictureId         int    `json:"picture_id,omitempty"`
	Title             string `json:"title,omitempty"`
	// ...
	// 2021-07-22: 目前看来，传入的所有参数都无效
}

type TaobaoPicturePicturesCountFields struct {
	PictureCategoryId bool
	PictureId         bool
	Title             bool
}

func (r *TaobaoPicturePicturesCountRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	if r.RequestFields.PictureCategoryId {
		m["picture_category_id"] = strconv.Itoa(r.PictureCategoryId)
	}
	if r.RequestFields.PictureId {
		m["picture_id"] = strconv.Itoa(r.PictureId)
	}
	if r.RequestFields.Title {
		m["title"] = r.Title
	}
	return m
}

func (r *TaobaoPicturePicturesCountRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

func (r *TaobaoPicturePicturesCountRequest) Valid() error {
	return nil
}
