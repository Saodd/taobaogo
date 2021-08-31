package taobaoapi

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

func TestClient_TaobaoPicassoPictagControltimeQuery(t *testing.T) {
	type args struct {
		ctx     context.Context
		data    *TaobaoPicassoPictagControltimeQueryRequest
		session *string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "手动用例",
			args: args{
				ctx:     context.Background(),
				data:    &TaobaoPicassoPictagControltimeQueryRequest{},
				session: nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.TaobaoPicassoPictagControltimeQuery(tt.args.ctx, tt.args.data, tt.args.session)
			if (err != nil) != tt.wantErr {
				t.Errorf("TaobaoPicassoPictagControltimeQuery() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				j, _ := json.Marshal(got)
				fmt.Println(string(j))
			}
		})
	}
}
