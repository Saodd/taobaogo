package taobaoapi

import (
	"context"
	"log"
	"testing"
)

func TestClient_TaobaoUserSellerGet(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *TaobaoUserSellerGetRequest
		session string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "手动测试1",
			args: args{
				ctx: context.Background(),
				data: &TaobaoUserSellerGetRequest{
					Fields: TaobaoUserSellerGetDemo,
				},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.TaobaoUserSellerGet(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaobaoUserSellerGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				log.Println(got.User)
			}
		})
	}
}
