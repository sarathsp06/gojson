package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func floatAddr(t float64) *float64 { return &t }

func TestDecode(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    interface{}
		wantErr bool
	}{
		{
			name:    "map json",
			args:    args{data: []byte(`{"dd":234}`)},
			want:    &map[string]json.RawMessage{"dd": json.RawMessage(`234`)},
			wantErr: false,
		},
		{
			name:    "invalid json",
			args:    args{data: []byte(`{"dd":234`)},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "numeric type",
			args:    args{data: []byte(`123`)},
			want:    floatAddr(123.0),
			wantErr: false,
		},
		{
			name:    "null",
			args:    args{data: []byte(`null`)},
			want:    &json.RawMessage{'n', 'u', 'l', 'l'},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decode(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookup(t *testing.T) {
	type args struct {
		key  []string
		data []byte
	}
	testJSON := []byte(`{"users":[{"name":"sarath"}]}`)
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "simple map",
			args:    args{key: []string{"name"}, data: []byte(`{"name":"sarath"}`)},
			want:    []byte(`"sarath"`),
			wantErr: false,
		},
		{
			name:    "map with array ",
			args:    args{key: []string{"users", "0", "name"}, data: testJSON},
			want:    []byte(`"sarath"`),
			wantErr: false,
		},
		{
			name:    "invalid index for array ",
			args:    args{key: []string{"users", "-1"}, data: testJSON},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "index bigger than array",
			args:    args{key: []string{"users", "100", "name"}, data: testJSON},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "invalid json",
			args:    args{key: []string{"users"}, data: []byte(`{"usersame":"sarath"}]}`)},
			want:    []byte(`{"usersame":"sarath"}]}`),
			wantErr: true,
		},
		{
			name:    "unavailable key",
			args:    args{key: []string{"thenga"}, data: []byte(`{"name":"sarath"}`)},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := lookup(tt.args.key, tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("lookup() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
