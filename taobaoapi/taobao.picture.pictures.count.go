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
	RequestFields

	PictureCategoryId int    `json:"picture_category_id,omitempty"`
	PictureId         int    `json:"picture_id,omitempty"`
	Title             string `json:"title,omitempty"`
	// ...
	// 2021-07-22: 目前看来，传入的所有参数都无效
}

func (r *TaobaoPicturePicturesCountRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	if r.RequestFields["picture_category_id"] {
		m["picture_category_id"] = strconv.Itoa(r.PictureCategoryId)
	}
	if r.RequestFields["picture_id"] {
		m["picture_id"] = strconv.Itoa(r.PictureId)
	}
	if r.RequestFields["title"] {
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
