package test

import (
	"github.com/DolphinDong/ops-tools/tools"
	"testing"
)

func TestStringInMapKeys(t *testing.T) {
	m := map[string]string{
		"1": "1",
		"2": "1",
		"3": "1",
		"4": "1",
		"5": "1",
	}
	type args struct {
		element string
		m       map[string]string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"1", args{"x", m}, false,
		},
		{
			"2", args{"1", m}, true,
		},
		{
			"3", args{"2", m}, true,
		},
		{
			"4", args{"6", m}, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tools.StringInMapKeys(tt.args.element, tt.args.m); got != tt.want {
				t.Errorf("StringInMapKeys() = %v, want %v", got, tt.want)
			}
		})
	}
}
