package helpers

import (
	"net/http"
	"testing"

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

		assert.Equal(t, &hal.HalHref{Href: test.url}, l.HalRscLinks["lf:self"])
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
		l := HalRootRscLinks(&ContextHelper{
			FQEndpoint:  test.fqe,
			ServerURL:   test.server,
			OperationID: test.opid,
		})
		assert.NotNil(t, l)
		assert.NotNil(t, l.HalRscLinks["doc:my-operation"])
		assert.NotNil(t, l.Curies)

		assert.Equal(t, &hal.HalHref{Href: "/" + test.opid}, l.HalRscLinks["doc:my-operation"])
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
	for _, tt := range []struct {
		name string
		args args
		want bool
	}{
		{
			name: "basic-pass",
			args: args{
				s: []string{"one", "two", "three"},
				e: "two",
			},
			want: true,
		},
		{
			name: "basic-fail",
			args: args{
				s: []string{"one", "two", "three"},
				e: "four",
			},
			want: false,
		},
		{
			name: "empty",
			args: args{
				s: []string{},
				e: "",
			},
			want: false,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := ContainsString(tt.args.s, tt.args.e)
			assert.Equal(t, tt.want, got)
		})
	}
}
