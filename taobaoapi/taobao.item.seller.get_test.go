package taobaoapi

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_TaobaoItemSellerGet(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *TaobaoItemSellerGetRequest
		session string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "手动用例1",
			args: args{
				ctx: context.Background(),
				data: &TaobaoItemSellerGetRequest{
					Fields: "num_iid,desc,wireless_desc,sell_point",
					NumIid: secrets.TaobaoShop.Iids[0],
				},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.TaobaoItemSellerGet(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaobaoItemSellerGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				fmt.Println(got.Item)
			}
		})
	}
}
