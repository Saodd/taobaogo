package utils

import "testing"

func TestSignString(t *testing.T) {
	type args struct {
		data   []byte
		secret []byte
	}
	var tests = []struct {
		name string
		args args
		want string
	}{
		{
			name: "用例1",
			args: args{
				data:   []byte(`app_key00000000fieldstitleformatjsonmethodtaobao.shop.seller.getnicklewinsession203495ct8mu294538720c59384c7n204925876v23487695cb23948756c77sign_methodhmactimestamp2021-04-17 11:18:01v2.0`),
				secret: []byte(`helloworld`),
			},
			want: "EC017B645917D92D28B70F46E6C249AA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TaobaoSign(tt.args.data, tt.args.secret); got != tt.want {
				t.Errorf("TaobaoSign() = %v, want %v", got, tt.want)
			}
		})
	}
}
