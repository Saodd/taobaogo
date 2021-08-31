package taobaoapi

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_TaobaoPicturePicturesCount(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *TaobaoPicturePicturesCountRequest
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
				data:    &TaobaoPicturePicturesCountRequest{},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
		{
			name: "指定目录",
			args: args{
				ctx: context.Background(),
				data: &TaobaoPicturePicturesCountRequest{
					RequestFields:     TaobaoPicturePicturesCountFields{PictureCategoryId: true},
					PictureCategoryId: 14147127675499170,
				},
				session: secrets.TaobaoShop.Session,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.TaobaoPicturePicturesCount(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaobaoPicturePicturesCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				j, _ := json.Marshal(got)
				fmt.Println(string(j))
			}
		})
	}
}
