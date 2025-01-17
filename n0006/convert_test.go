package n0006

import "testing"

func Test_convert(t *testing.T) {
	type args struct {
		s       string
		numRows int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"PAYPALISHIRING", 3}, "PAHNAPLSIIGYIR"},
		{"2", args{"PAYPALISHIRING", 4}, "PINALSIGYAHRPI"},
		{"3", args{"A", 1}, "A"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.s, tt.args.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
