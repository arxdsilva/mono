package pkg2

import "testing"

func TestWorldString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "ok", want: "World"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WorldString(); got != tt.want {
				t.Errorf("WorldString() = %v, want %v", got, tt.want)
			}
		})
	}
}
