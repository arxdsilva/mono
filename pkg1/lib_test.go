package pkg1

import (
	"reflect"
	"testing"
)

func TestHelloString(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "ok", want: "Hello"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelloString(); got != tt.want {
				t.Errorf("HelloString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHelloByte(t *testing.T) {
	tests := []struct {
		name string
		want []byte
	}{
		{name: "ok", want: []byte("Hello")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelloByte(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HelloByte() = %v, want %v", got, tt.want)
			}
		})
	}
}
