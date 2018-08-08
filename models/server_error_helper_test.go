package models

import (
	"reflect"
	"testing"
)

func TestNewServerError(t *testing.T) {
	type args struct {
		code        int32
		errorString string
	}
	tests := []struct {
		name string
		args args
		want *ServerError
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewServerError(tt.args.code, tt.args.errorString); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewServerError() = %v, want %v", got, tt.want)
			}
		})
	}
}