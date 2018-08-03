package helpers

import (
	"net/http"
	"reflect"
	"testing"

	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/assert"
	"github.com/getlunaform/lunaform/models/hal"
)

func TestHALSelfLink(t *testing.T) {

	for _, test := range []struct {
		url string
	}{
		{url: "http://example.com/hello-world"},
	} {
		rootLink := hal.HalRscLinks{}
		l := HalSelfLink(&rootLink, test.url)

		assert.Equal(t, test.url, l.HalRscLinks["lf:self"])

	}
}

func TestHALRootRscLinks(t *testing.T) {

	for _, test := range []struct {
		fqe    string
		server string
		opid   string
		docURL string
	}{
		{
			fqe:    "http://example.com/hello-world",
			server: "http://example.com",
			opid:   "my-operation",
			docURL: "http://example.com/docs#operation/my-operation",
		},
	} {
		l := HalRootRscLinks(ContextHelper{
			FQEndpoint:  test.fqe,
			ServerURL:   test.server,
			OperationID: test.opid,
		})
		assert.NotNil(t, l)
		assert.NotNil(t, l.HalRscLinks[""])
		assert.NotNil(t, l.HalRscLinks[""])

		assert.Equal(t, test.fqe, l.HalRscLinks[""])
		assert.Equal(t, test.docURL, l.HalRscLinks[""])
	}
}

func TestUrlPrefix(t *testing.T) {
	for _, test := range []struct {
		host   string
		uri    string
		https  bool
		prefix string
	}{
		{host: "mock.com", uri: "/mock-uri", https: false, prefix: "http://mock.com/mock-uri"},
		{host: "mock.com", uri: "/mock-uri", https: true, prefix: "https://mock.com/mock-uri"},
	} {
		ch := ContextHelper{}
		assert.Equal(
			t,
			test.prefix,
			ch.urlPrefix(test.host, test.uri, test.https),
		)
	}
}

func TestNewServerError(t *testing.T) {
	type args struct {
		code        int32
		errorString string
	}
	tests := []struct {
		name string
		args args
		want *models.ServerError
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

func TestNewContextHelper(t *testing.T) {
	type args struct {
		ctx *middleware.Context
	}
	tests := []struct {
		name string
		args args
		want ContextHelper
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContextHelper(tt.args.ctx); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContextHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContextHelper_SetRequest(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		ch      *ContextHelper
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ch.SetRequest(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("ContextHelper.SetRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestContextHelper_ParseRequest(t *testing.T) {
	type args struct {
		host       string
		basePath   string
		requestUri string
		hasTls     bool
	}
	for _, tt := range []struct {
		name   string
		args   args
		wantCh *ContextHelper
	}{
		{
			name: "base",
			args: args{
				host:       "www.example.com",
				basePath:   "/mock",
				requestUri: "/mock/my-resources",
				hasTls:     true,
			},
			wantCh: &ContextHelper{
				ServerURL:        "https://www.example.com/mock",
				Endpoint:         "/my-resources",
				EndpointSingular: "/my-resource",
				FQEndpoint:       "https://www.example.com/mock/my-resources",
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			ch := &ContextHelper{}

			ch.ParseRequest(
				tt.args.host, tt.args.basePath,
				tt.args.requestUri, tt.args.hasTls,
			)

			assert.Equal(t, tt.wantCh, ch)
		})
	}
}

func TestContextHelper_urlPrefix(t *testing.T) {
	type args struct {
		host  string
		uri   string
		https bool
	}
	tests := []struct {
		name string
		ch   ContextHelper
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ch.urlPrefix(tt.args.host, tt.args.uri, tt.args.https); got != tt.want {
				t.Errorf("ContextHelper.urlPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	type args struct {
		s []string
		e string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsString(tt.args.s, tt.args.e); got != tt.want {
				t.Errorf("ContainsString() = %v, want %v", got, tt.want)
			}
		})
	}
}
