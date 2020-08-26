package favicon

import (
	"reflect"
	"testing"
)

func Test_getFavicon(t *testing.T) {
	type args struct {
		targetUrl string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "12bit",
			args: args{
				targetUrl: "https://12bit.vn",
			},
			want:    "https://12bit.vn/img/icon/favicon.png",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getFavicon(tt.args.targetUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFavicon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFavicon() got = %v, want %v", got, tt.want)
			}
		})
	}
}
