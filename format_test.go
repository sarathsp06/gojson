package main

import (
	"reflect"
	"testing"
)

func Test_formatJSON(t *testing.T) {
	inavalidJSON := []byte(`{sssss`)
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "empty",
			args:    args{},
			want:    []byte(`null`),
			wantErr: false,
		},
		{
			name: "valid json body",
			args: args{data: []byte(`{"name":"Sarath"}`)},
			want: []byte(
				`{
  "name": "Sarath"
}`),
			wantErr: false,
		},
		{
			name:    "inavalid json",
			args:    args{data: inavalidJSON},
			want:    inavalidJSON,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := formatJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("formatJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
