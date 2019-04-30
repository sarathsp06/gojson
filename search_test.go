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
	testJSON := []byte(`{"users":[{"name":"sarath"},{"name":"syam"}]}`)
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
			name:    "invalid range for array:valid index invalid range",
			args:    args{key: []string{"users", "2:0"}, data: testJSON},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "invalid index",
			args:    args{key: []string{"users", "-1"}, data: testJSON},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "valid slice for array",
			args:    args{key: []string{"users", "0:2"}, data: testJSON},
			want:    []byte(`[{"name":"sarath"},{"name":"syam"}]`),
			wantErr: false,
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
				t.Errorf("name:%s, lookup() = %v, want %v", tt.name, got, tt.want)
			}
		})
	}
}

func equalFunc(f1, f2 interface{}) bool {
	return reflect.ValueOf(f1).Pointer() == reflect.ValueOf(f2).Pointer()
}

func Test_getSliceOperation(t *testing.T) {
	type args struct {
		op string
	}
	tests := []struct {
		name string
		args args
		want sliceOp
	}{
		{
			name: "Number",
			args: args{op: "23"},
			want: sliceIdx,
		},
		{
			name: "slice range",
			args: args{op: "3:23"},
			want: sliceRange,
		},
		{
			name: "map",
			args: args{op: "name"},
			want: sliceMap,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSliceOperation(tt.args.op); !equalFunc(tt.want, got) {
				t.Errorf("getSliceOperation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func rawMessages(args ...string) []json.RawMessage {
	var result []json.RawMessage
	for _, val := range args {
		result = append(result, json.RawMessage(val))
	}
	return result
}

func Test_sliceSerialize(t *testing.T) {
	type args struct {
		obj []json.RawMessage
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "single item slice",
			args: args{obj: []json.RawMessage{[]byte(`123`)}},
			want: []byte(`[123]`),
		},
		{
			name: "empty slice",
			args: args{obj: nil},
			want: []byte(`[]`),
		},
		{
			name: "multi item slice",
			args: args{obj: []json.RawMessage{[]byte(`123`), []byte(`test`), []byte(`"test"`)}},
			want: []byte(`[123,test,"test"]`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceSerialize(tt.args.obj); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceSerialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sliceRange(t *testing.T) {
	type args struct {
		obj []json.RawMessage
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name:    "valid range",
			args:    args{obj: rawMessages(`233`, `123`, `"sara"`, `gama`), key: `1:3`},
			want:    []byte(`[123,"sara"]`),
			wantErr: false,
		},
		{
			name:    "invalid range",
			args:    args{obj: rawMessages(`233`, `123`, `"sara"`, `gama`), key: `3:1`},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "n:",
			args:    args{obj: rawMessages(`233`, `123`, `"test"`, `gama`), key: `1:`},
			want:    []byte(`[123,"test",gama]`),
			wantErr: false,
		},
		{
			name:    ":n",
			args:    args{obj: rawMessages(`233`, `123`, `"test"`, `test`), key: `:2`},
			want:    []byte(`[233,123]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sliceRange(tt.args.obj, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("sliceRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceRange() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func Test_sliceMap(t *testing.T) {
	type args struct {
		obj []json.RawMessage
		key string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "array of objects with name field",
			args: args{obj: rawMessages(
				`{"name":"sarath"}`,
				`{"name":"syam"}`,
				`{"name":"23"}`,
			),
				key: "name"},
			want:    []byte(`["sarath","syam","23"]`),
			wantErr: false,
		},
		{
			name: "array of objects with  missing fields",
			args: args{obj: rawMessages(
				`{"name":"sarath"}`,
				`{"1name":"syam"}`,
				`{"2name":"23"}`,
			),
				key: "name"},
			want:    []byte(`["sarath"]`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sliceMap(tt.args.obj, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("sliceMap() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceMap() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
