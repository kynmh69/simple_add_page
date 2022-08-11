package util

import (
	"reflect"
	"simple_add_page/src/util"
	"testing"
)

func TestReadConfigFile(t *testing.T) {
	type args struct {
		config_file_path string
	}
	tests := []struct {
		name string
		args args
		want *util.Config
	}{
		// TODO: Add test cases.
		{
			name: "normal",
			args: args{config_file_path: "config/temp_api_yml"},
			want: &util.Config{ApiSettings: util.ApiSettings{ApiSecret: "", DatabaseID: "", TagId: ""}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.ReadConfigFile(tt.args.config_file_path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
