package util

import (
	"os"
	"reflect"
	"simple_add_page/src/util"
	"testing"
)

func TestSplitLines(t *testing.T) {
	f, _ := os.Open("../input_file/photo_list.txt")
	defer f.Close()

	type args struct {
		content os.File
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "normal",
			args: args{content: *f},
			want: []string{"AYtest0000011"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.SplitLines(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SplitLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
