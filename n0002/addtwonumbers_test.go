package n0002

import (
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"1", args{NewListNode(2, 4, 3), NewListNode(5, 6, 4)}, NewListNode(7, 0, 8)},
		{"2", args{NewListNode(0), NewListNode(0)}, NewListNode(0)},
		{"3", args{NewListNode(9, 9, 9, 9, 9, 9, 9), NewListNode(9, 9, 9, 9)}, NewListNode(8, 9, 9, 9, 0, 0, 0, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
