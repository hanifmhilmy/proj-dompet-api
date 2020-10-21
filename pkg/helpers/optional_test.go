package helpers

import (
	"testing"
)

func TestToBool(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 bool
	}{
		{
			name: "error invalid types",
			args: args{
				value: 1234,
			},
			want:  false,
			want1: false,
		}, {
			name: "correct types",
			args: args{
				value: true,
			},
			want:  true,
			want1: true,
		}, {
			name: "correct types",
			args: args{
				value: false,
			},
			want:  false,
			want1: true,
		}, {
			name: "nil args",
			args: args{
				value: nil,
			},
			want:  false,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToBool(tt.args.value)
			if got != tt.want {
				t.Errorf("ToBool() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToBool() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToFloat64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 bool
	}{
		{
			name: "nil args",
			args: args{
				value: nil,
			},
			want:  0,
			want1: false,
		}, {
			name: "invalid types",
			args: args{
				value: true,
			},
			want:  0,
			want1: false,
		}, {
			name: "invalid int types",
			args: args{
				value: 1234,
			},
			want:  0,
			want1: false,
		}, {
			name: "correct types",
			args: args{
				value: 112.41,
			},
			want:  112.41,
			want1: true,
		}, {
			name: "correct types",
			args: args{
				value: 1234.11,
			},
			want:  1234.11,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToFloat64(tt.args.value)
			if got != tt.want {
				t.Errorf("ToFloat64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToFloat64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	const MaxUint = ^uint(0)
	const bigInt = int64(MaxUint >> 1)

	type args struct {
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 bool
	}{
		{
			name: "nil args",
			args: args{
				value: nil,
			},
			want:  0,
			want1: false,
		}, {
			name: "invalid types",
			args: args{
				value: true,
			},
			want:  0,
			want1: false,
		}, {
			name: "invalid float types",
			args: args{
				value: 1234.11,
			},
			want:  0,
			want1: false,
		}, {
			name: "invalid int types - too big",
			args: args{
				value: bigInt,
			},
			want:  0,
			want1: false,
		}, {
			name: "correct types",
			args: args{
				value: 1234,
			},
			want:  1234,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt(tt.args.value)
			if got != tt.want {
				t.Errorf("ToInt() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToInt() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToInt64(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  int64
		want1 bool
	}{
		{
			name: "nil args",
			args: args{
				value: nil,
			},
			want:  0,
			want1: false,
		}, {
			name: "invalid types",
			args: args{
				value: true,
			},
			want:  0,
			want1: false,
		}, {
			name: "invalid float types",
			args: args{
				value: 1234.11,
			},
			want:  0,
			want1: false,
		}, {
			name: "correct types",
			args: args{
				value: int64(1234),
			},
			want:  1234,
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToInt64(tt.args.value)
			if got != tt.want {
				t.Errorf("ToInt64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestToString(t *testing.T) {
	type args struct {
		value interface{}
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 bool
	}{
		{
			name: "nil args",
			args: args{
				value: nil,
			},
			want:  "",
			want1: false,
		}, {
			name: "invalid types",
			args: args{
				value: true,
			},
			want:  "",
			want1: false,
		}, {
			name: "invalid float types",
			args: args{
				value: 1234.11,
			},
			want:  "",
			want1: false,
		}, {
			name: "correct types",
			args: args{
				value: "1234",
			},
			want:  "1234",
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ToString(tt.args.value)
			if got != tt.want {
				t.Errorf("ToString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ToString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
