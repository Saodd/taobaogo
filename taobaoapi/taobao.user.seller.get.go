package taobaoapi

import (
	"context"
	"github.com/saodd/taobaogo/utils"
	"net/url"
)

// TaobaoUserSellerGet 查询卖家用户信息
// https://open.taobao.com/api.htm?docId=21349&docType=2
func (client *Client) TaobaoUserSellerGet(ctx context.Context, data *TaobaoUserSellerGetRequest, session string) (*TaobaoUserSellerGetResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.user.seller.get",
		Session: &session,
	}
	var res struct {
		Resp *TaobaoUserSellerGetResponse `json:"user_seller_get_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}

	// extra: 解码
	if err := utils.QueryUnescape(&res.Resp.User.Nick); err != nil {
		client.HandleError(ctx, err, map[string]interface{}{"nick": res.Resp.User.Nick})
		return nil, err
	}
	if err := utils.QueryUnescape(&res.Resp.User.Avatar); err != nil {
		client.HandleError(ctx, err, map[string]interface{}{"avatar": res.Resp.User.Avatar})
		return nil, err
	}

	return res.Resp, nil
}

type TaobaoUserSellerGetResponse struct {
	User struct {
		UserId       int64  `json:"user_id"`
		Nick         string `json:"nick"`
		Sex          string `json:"sex"`
		SellerCredit struct {
			GoodNum  int `json:"good_num"`
			Level    int `json:"level"`
			Score    int `json:"score"`
			TotalNum int `json:"total_num"`
		} `json:"seller_credit"`
		Type                   string `json:"type"`
		HasMorePic             bool   `json:"has_more_pic"`
		ItemImgNum             int    `json:"item_img_num"`
		ItemImgSize            int    `json:"item_img_size"`
		PropImgNum             int    `json:"prop_img_num"`
		PropImgSize            int    `json:"prop_img_size"`
		AutoRepost             string `json:"auto_repost"`
		PromotedType           string `json:"promoted_type"`
		Status                 string `json:"status"`
		AlipayBind             string `json:"alipay_bind"`
		ConsumerProtection     bool   `json:"consumer_protection"`
		Avatar                 string `json:"avatar"`
		Liangpin               bool   `json:"liangpin"`
		SignFoodSellerPromise  bool   `json:"sign_food_seller_promise"`
		HasShop                bool   `json:"has_shop"`
		IsLightningConsignment bool   `json:"is_lightning_consignment"`
		HasSubStock            bool   `json:"has_sub_stock"`
		IsGoldenSeller         bool   `json:"is_golden_seller"`
		MagazineSubscribe      bool   `json:"magazine_subscribe"`
		VerticalMarket         string `json:"vertical_market"`
		OnlineGaming           bool   `json:"online_gaming"`
		IsTjbSeller            bool   `json:"is_tjb_seller"`
		VipInfo                string `json:"vip_info"`
	} `json:"user"`
	RequestId string `json:"request_id"`
}

type TaobaoUserSellerGetRequest struct {
	Fields string `json:"fields"`
}

func (r *TaobaoUserSellerGetRequest) ToSignMap() map[string]string {
	m := make(map[string]string)
	m["fields"] = r.Fields
	return m
}

func (r *TaobaoUserSellerGetRequest) ToValues() url.Values {
	value := url.Values{}
	for k, v := range r.ToSignMap() {
		value.Set(k, v)
	}
	return value
}

const (
	TaobaoUserSellerGetDemo = "user_id,nick,seller_credit,type,avatar"
)
