package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGetObject(t *testing.T) {
	emptyMap := make(map[string]json.RawMessage)
	type args struct {
		jsn []byte
	}
	tests := []struct {
		name  string
		args  args
		want  interface{}
		want1 bool
	}{
		{
			name:  "json map",
			args:  args{jsn: []byte(`{"yo" : "yoyoy"}`)},
			want:  &emptyMap,
			want1: true,
		},
		{
			name:  "json list",
			args:  args{jsn: []byte(`[ 1,2,3,4 ]`)},
			want:  &[]json.RawMessage{},
			want1: true,
		},
		{
			name:  "json string",
			args:  args{jsn: []byte(`"yoyoy"`)},
			want:  new(string),
			want1: false,
		},
		{
			name:  "json number",
			args:  args{jsn: []byte(`123`)},
			want:  new(float64),
			want1: false,
		},
		{
			name:  "empty array",
			args:  args{jsn: nil},
			want:  &json.RawMessage{},
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getObject(tt.args.jsn)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("%s: getObject() got = %v, want %v", tt.name, got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("%s: getObject() got1 = %v, want %v", tt.name, got1, tt.want1)
			}
		})
	}
}
