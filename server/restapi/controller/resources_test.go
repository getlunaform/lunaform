package controller

import (
	"testing"
	"github.com/zeebox/terraform-server/server/restapi/operations/resources"
	"net/http"
	"net/http/httptest"
	"io"
	"fmt"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

type mockOpHelper struct {
	OperationID string
	FQEndpoint  string
	ServerURL   string
	Request     *http.Request
}

func (moh mockOpHelper) GetOperationID() string {
	return moh.OperationID
}

func (moh mockOpHelper) GetFQEndpoint() string {
	return moh.FQEndpoint
}

func (moh mockOpHelper) GetServerURL() string {
	return moh.FQEndpoint
}
func (moh mockOpHelper) SetRequest(req *http.Request) {
	moh.Request = req
}

type mockProducer struct {
	ProducerHandler func(w io.Writer, i interface{}) (err error)
}

func (mp mockProducer) Produce(w io.Writer, i interface{}) (err error) {
	return mp.ProducerHandler(w, i)
}

func TestResourcesController(t *testing.T) {

	for _, test := range []struct {
		host string
		responseCode int
	} {
		{host:"mock.com"}
	}

	mock := http.Request{
		Host: "mock.com",
	}

	t.Run("ListIdentity", func(*testing.T) {
		r := ListResourcesController(nil, mockOpHelper{}).Handle(resources.ListResourcesParams{
			HTTPRequest: &mock,
			Group:       "tf",
		})

		w := httptest.NewRecorder()
		p := mockProducer{}
		p.ProducerHandler = func(w io.Writer, i interface{}) (err error) {
			return nil
		}
		r.WriteResponse(w, p)

		assert.Equal(t, 200, w.Result().StatusCode)

		bdy, err := ioutil.ReadAll(w.Body)
		fmt.Println(w, p, bdy, err)
	})
}
