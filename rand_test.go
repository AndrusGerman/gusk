package gusk

import (
	"testing"
)

func Test_randStringBytesMaskImprSrcUnsafe(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Len 4",
			args: args{n: 4},
			want: "Error len ouput",
		},
		{
			name: "Len 8",
			args: args{n: 8},
			want: "Error len ouput",
		},
		{
			name: "Len 16",
			args: args{n: 16},
			want: "Error len ouput",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randStringBytesMaskImprSrcUnsafe(tt.args.n); got != tt.want {
				if lg := len(got); lg != tt.args.n {
					t.Errorf("randStringBytesMaskImprSrcUnsafe() = '%v' != len('%v') = error Want %v", tt.args.n, got, tt.want)
				}
			}
		})
	}
}
