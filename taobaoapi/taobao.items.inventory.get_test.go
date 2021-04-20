package taobaoapi

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_TaobaoItemsInventoryGet(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *TaobaoItemsInventoryGetRequest
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
				data: &TaobaoItemsInventoryGetRequest{
					Fields:   "num_iid,title,price,modified",
					PageNo:   2, // 不生效
					PageSize: 1, // 不生效
				},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
		{
			name: "限制page_no,page_size",
			args: args{
				ctx: context.Background(),
				data: &TaobaoItemsInventoryGetRequest{
					RequestFields: map[string]bool{"page_no": true, "page_size": true},
					Fields:        "num_iid,title,price,modified",
					PageNo:        2,
					PageSize:      1,
				},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.TaobaoItemsInventoryGet(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaobaoItemsInventoryGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				fmt.Println(got.Items.Item[0])
			}
		})
	}
}
