package taobaoapi

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_TaobaoPicassoPictagCheckimg(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *TaobaoPicassoPictagCheckimgRequest
		session string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "手动用例",
			args: args{
				ctx: context.Background(),
				data: &TaobaoPicassoPictagCheckimgRequest{
					ItemId: secrets.TaobaoShop.Iids[0],
					Biz:    "pictag",
					PicUrl: "https://img.alicdn.com/bao/uploaded/i1/696944147/O1CN01aQ3LOq1gVNBRp5rLi_!!0-item_pic.jpg",
				},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.TaobaoPicassoPictagCheckimg(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaobaoPicassoPictagCheckimg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				j, _ := json.Marshal(got)
				fmt.Println(string(j))
			}
		})
	}
}
