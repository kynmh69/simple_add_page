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
		{
			name: "normal",
			args: args{config_file_path: "/Users/Hiroki/Applications/GoProjects/simple_add_page/config/temp_api.yml"},
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

func TestReadPhotoIdFile(t *testing.T) {

	type args struct {
		file_path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "normal",
			args: args{file_path: ""},
			want: "../input_file/photo_list.txt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.ReadPhotoIdFile(tt.args.file_path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadPhotoIdFile() = %v, want %v", got.Name(), tt.want)
			}
		})
	}
}
