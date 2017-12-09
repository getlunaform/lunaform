package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointerString(t *testing.T) {
	tests := []string{
		"hello",
	}

	for _, test := range tests {
		var s interface{}
		s = str(test)
		_, ok := s.(*string)
		assert.True(t, ok)
	}
}

func TestHALSelfLink(t *testing.T) {
	tests := []struct {
		url string
	}{
		{url: "http://example.com/hello-world"},
	}

	for _, test := range tests {
		l := halSelfLink(test.url)
		assert.NotNil(t, l)
		assert.Nil(t, l.Doc)
		assert.NotNil(t, l.Self)

		assert.Equal(t, l.Self.Href.String(), test.url)
	}
}
func TestHALRootRscLinks(t *testing.T) {

	tests := []struct {
		fqe    string
		server string
		opid   string
		docURL string
	}{
		{fqe: "http://example.com/hello-world", server: "http://example.com", opid: "my-operation", docURL: "http://example.com/docs#operation/my-operation"},
	}

	for _, test := range tests {
		l := halRootRscLinks(&apiHostBase{
			FQEndpoint:  test.fqe,
			ServerURL:   test.server,
			OperationID: test.opid,
		})
		assert.NotNil(t, l)
		assert.NotNil(t, l.Doc)
		assert.NotNil(t, l.Self)

		assert.Equal(t, test.fqe, l.Self.Href.String())
		assert.Equal(t, test.docURL, l.Doc.Href.String())
	}
}
