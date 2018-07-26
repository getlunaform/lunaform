package restapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/getlunaform/lunaform/helpers"
	"github.com/getlunaform/lunaform/restapi/operations/resources"
	"github.com/stretchr/testify/assert"
)

func Test_resourceGroupsController(t *testing.T) {
	oh := helpers.NewContextHelper(api.Context())
	requestHost := "example.com"

	for _, test := range []struct {
		requestURL    string
		requestMethod string
		responseCode  int
		response      string
	}{
		{
			requestURL:    "/api",
			requestMethod: "GET",
			responseCode:  200,
			response:      `{"_embedded":{"resources":[{"_links":{"lf:self":{"href":"/tf"}},"name":"tf"},{"_links":{"lf:self":{"href":"/identity"}},"name":"identity"},{"_links":{"lf:self":{"href":"/vcs"}},"name":"vcs"}]},"_links":{"curies":[{"href":"http://example.com/api/{rel}","name":"lf","templated":true},{"href":"http://example.com/api/docs#operation/{rel}","name":"doc","templated":true}],"doc:list-resource-groups":{"href":"/list-resource-groups"},"lf:self":{"href":"/"}}}`,
		},
	} {

		u, err := url.Parse("http://" + requestHost + test.requestURL)
		assert.NoError(t, err)

		oh.Request = &http.Request{
			Host:       requestHost,
			Method:     test.requestMethod,
			RequestURI: test.requestURL,
			URL:        u,
		}

		r := ListResourceGroupsController(oh).Handle(resources.ListResourceGroupsParams{
			HTTPRequest: oh.Request,
		})

		w := httptest.NewRecorder()
		p := mockProducer{}
		r.WriteResponse(w, p)

		buf := new(bytes.Buffer)
		buf.ReadFrom(w.Result().Body)
		body := buf.String()

		assert.Equal(t, test.responseCode, w.Result().StatusCode)
		assert.Equal(t, test.response, body)

	}
}
