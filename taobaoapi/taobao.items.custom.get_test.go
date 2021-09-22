package taobaoapi

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_TaobaoItemsCustomGet(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *TaobaoItemsCustomGetRequest
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
				data: &TaobaoItemsCustomGetRequest{
					OuterId: "HJD",
					Fields:  "num_iid,title,price,modified,outer_id,pic_url",
				},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.TaobaoItemsCustomGet(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaobaoItemsCustomGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && len(got.Items.Item) > 0 {
				j, _ := json.Marshal(got.Items.Item[0])
				fmt.Println(string(j))
			}
		})
	}
}
