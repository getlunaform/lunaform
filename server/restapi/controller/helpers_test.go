package controller

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

type mockProducer struct {
	ProducerHandler func(w io.Writer, i interface{}) (err error)
}

func (mp mockProducer) Produce(w io.Writer, i interface{}) (err error) {
	var b []byte
	b, err = json.Marshal(i)
	w.Write(b)
	return
}

func TestPointerString(t *testing.T) {

	for _, test := range []string{
		"hello",
	} {
		var s interface{}
		s = str(test)
		_, ok := s.(*string)
		assert.True(t, ok)
	}
}

func TestHALSelfLink(t *testing.T) {

	for _, test := range []struct {
		url string
	}{
		{url: "http://example.com/hello-world"},
	} {
		l := halSelfLink(test.url)
		assert.NotNil(t, l)
		assert.Nil(t, l.Doc)
		assert.NotNil(t, l.Self)

		assert.Equal(t, l.Self.Href.String(), test.url)
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
		l := halRootRscLinks(ContextHelper{
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
