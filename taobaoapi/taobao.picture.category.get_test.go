package taobaoapi

import (
	"context"
	"fmt"
	"testing"
)

func TestClient_TaobaoPictureCategoryGet(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *TaobaoPictureCategoryGetRequest
		session string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "手动用例1-没有条件",
			args: args{
				ctx:     context.Background(),
				data:    &TaobaoPictureCategoryGetRequest{},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.TaobaoPictureCategoryGet(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaobaoPictureCategoryGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				fmt.Println(got.PictureCategories.PictureCategory)
			}
		})
	}
}
