package taobaoapi

import (
	"context"
	"testing"
)

func TestClient_AlibabaItemEditSchemaGet(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *AlibabaItemEditSchemaGetRequest
		session string
	}
	tests := []struct {
		name    string
		args    args
		want    *AlibabaItemEditSchemaGetResponse
		wantErr bool
	}{
		{
			name: "手动用例1",
			args: args{
				ctx: context.Background(),
				data: &AlibabaItemEditSchemaGetRequest{
					ItemId: secrets.TmallShop.Iids[0],
				},
				session: secrets.TmallShop.Session,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.AlibabaItemEditSchemaGet(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("AlibabaItemEditSchemaGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
