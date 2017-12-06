package controller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointerString(t *testing.T) {
	var s interface{}
	s = str("hallo")
	_, ok := s.(*string)
	assert.True(t, ok)
}

func TestHALSelfLink(t *testing.T) {
	l := HALSelfLink("http://example.com/hello-world")
	assert.NotNil(t, l)
	assert.Nil(t, l.Doc)
	assert.NotNil(t, l.Self)

	assert.Equal(t, l.Self.Href.String(), "http://example.com/hello-world")
}

func TestHALRootRscLinks(t *testing.T) {
	l := HALRootRscLinks(&APIHostBase{
		FQEndpoint:  "http://example.com/hello-world",
		ServerURL:   "http://example.com",
		OperationId: "my-operation",
	})
	assert.NotNil(t, l)
	assert.NotNil(t, l.Doc)
	assert.NotNil(t, l.Self)

	assert.Equal(t, "http://example.com/hello-world", l.Self.Href.String())
	assert.Equal(t, "http://example.com/docs#operation/my-operation", l.Doc.Href.String())
}
