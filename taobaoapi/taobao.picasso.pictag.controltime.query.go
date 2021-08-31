package taobaoapi

import (
	"context"
	"net/url"
)

// TaobaoPicassoPictagControltimeQuery 查询主图价格管控时段
// https://open.taobao.com/api.htm?docId=58234&docType=2&source=search
func (client *Client) TaobaoPicassoPictagControltimeQuery(ctx context.Context, data *TaobaoPicassoPictagControltimeQueryRequest, session *string) (*TaobaoPicassoPictagControltimeQueryResponse, error) {
	var sp = SystemParams{
		Method:  "taobao.picasso.pictag.controltime.query",
		Session: session,
	}
	var res struct {
		Resp *TaobaoPicassoPictagControltimeQueryResponse `json:"picasso_pictag_controltime_query_response"`
	}
	err := client.Do(ctx, data, sp, &res)
	if err != nil {
		client.HandleError(ctx, err)
		return nil, err
	}
	return res.Resp, nil
}

type TaobaoPicassoPictagControltimeQueryResponse struct {
	Result struct {
		Success              bool `json:"success"`
		ControlTimeRangeList struct {
			ControlTimeRange []struct {
				EndTime      string `json:"end_time"`
				StartTime    string `json:"start_time"`
				ActivityName string `json:"activity_name"`
			} `json:"control_time_range"`
		} `json:"control_time_range_list"`
		MsgInfo string `json:"msg_info"`
		MsgCode string `json:"msg_code"`
	} `json:"result"`
	RequestId string `json:"request_id"`
}

type TaobaoPicassoPictagControltimeQueryRequest struct {
	RequestFields
}

func (r *TaobaoPicassoPictagControltimeQueryRequest) ToSignMap() map[string]string {
	m := make(map[string]string)

	return m
}

func (r *TaobaoPicassoPictagControltimeQueryRequest) ToValues() url.Values {
	value := url.Values{}
	return value
}
