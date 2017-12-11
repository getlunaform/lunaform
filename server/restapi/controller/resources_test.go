package controller

import (
	"testing"
	"github.com/zeebox/terraform-server/server/restapi/operations/resources"
	"net/http"
	"net/http/httptest"
	"io"
	"github.com/go-openapi/runtime/middleware"
)

type mockOpHelper struct{}

func (moh mockOpHelper) GetOperationID(*http.Request, *middleware.Context) string {
	return "mock-id"
}

type mockProducer struct{}

func (mp mockProducer) Produce(w io.Writer, i interface{}) (err error) {
	return nil
}

func TestResourcesController(t *testing.T) {
	mock := http.Request{
		Host: "mock.com",
	}

	t.Run("ListIdentity", func(*testing.T) {
		r := ListResourcesController(nil, mockOpHelper{}, &middleware.Context{}).Handle(resources.ListResourcesParams{
			HTTPRequest: &mock,
			Group:       "tf",
		})

		w := httptest.NewRecorder()
		p := mockProducer{}
		r.WriteResponse(w, p)
	})
}
