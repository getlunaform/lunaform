package helpers

import (
	"testing"

	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
)

type mockBasePathContext struct {
	mockPath     string
	matchedRoute *middleware.MatchedRoute
	foundRoute   bool
}

func (mbpc *mockBasePathContext) BasePath() string {
	return mbpc.mockPath
}

func (mbpc *mockBasePathContext) LookupRoute(*http.Request) (*middleware.MatchedRoute, bool) {
	return mbpc.matchedRoute, mbpc.foundRoute
}

func Test_newContextHelperWithContext(t *testing.T) {
	type args struct {
		ctx BasePathContext
	}
	tests := []struct {
		name    string
		args    args
		want    *ContextHelper
		wantErr string
		fail    bool
	}{
		{
			name: "nil",
			args: args{
				ctx: nil,
			},
			want:    nil,
			wantErr: "context must not be 'nil'",
			fail:    true,
		},
		{
			name: "non-nil",
			args: args{
				ctx: &mockBasePathContext{mockPath: "my-pass"},
			},
			want: &ContextHelper{
				ctx:      &mockBasePathContext{mockPath: "my-pass"},
				BasePath: "my-pass",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewContextHelperWithContext(tt.args.ctx)
			if tt.fail {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_newContextHelper(t *testing.T) {
	type args struct {
		ctx BasePathContext
	}
	tests := []struct {
		name    string
		args    args
		want    *ContextHelper
		wantErr string
		fail    bool
	}{
		{
			name: "basic",
			args: args{
				ctx: &mockBasePathContext{},
			},
			want: &ContextHelper{
				ctx: &mockBasePathContext{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewContextHelper(tt.args.ctx)
			if tt.fail {
				assert.EqualError(t, err, tt.wantErr)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
