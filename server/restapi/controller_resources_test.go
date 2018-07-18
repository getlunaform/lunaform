package restapi

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/drewsonne/lunaform/server/restapi/operations/resources"
	"github.com/drewsonne/lunaform/server/helpers"
)

func TestResourceGroupsController(t *testing.T) {
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
			response:      `{"_embedded":{"resources":[{"_links":{"self":{"href":"http://example.com/api/tf"}},"name":"tf"},{"_links":{"self":{"href":"http://example.com/api/identity"}},"name":"identity"},{"_links":{"self":{"href":"http://example.com/api/vcs"}},"name":"vcs"}]},"_links":{"doc":{"href":"http://example.com/api/docs#operation/list-resource-groups"},"self":{"href":"http://example.com/api"}}}`,
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

		r := ListResourceGroupsController(nil, oh).Handle(resources.ListResourceGroupsParams{
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

func TestResourcesController(t *testing.T) {

	oh := helpers.NewContextHelper(api.Context())
	requestHost := "example.com"

	for _, test := range []struct {
		group        string
		requestURL   string
		responseCode int
		response     string
	}{
		{
			group:        "tf",
			responseCode: 200,
			response:     "{\"_embedded\":{\"resources\":[{\"_links\":{\"self\":{\"href\":\"http://example.com/api/tf/modules\"}},\"name\":\"modules\"},{\"_links\":{\"self\":{\"href\":\"http://example.com/api/tf/stacks\"}},\"name\":\"stacks\"},{\"_links\":{\"self\":{\"href\":\"http://example.com/api/tf/state-backends\"}},\"name\":\"state-backends\"},{\"_links\":{\"self\":{\"href\":\"http://example.com/api/tf/workspaces\"}},\"name\":\"workspaces\"}]},\"_links\":{\"doc\":{\"href\":\"http://example.com/api/docs#operation/list-resources\"},\"self\":{\"href\":\"http://example.com/api/tf\"}}}",
		},
		{
			group:        "identity",
			responseCode: 200,
			response:     "{\"_embedded\":{\"resources\":[{\"_links\":{\"self\":{\"href\":\"http://example.com/api/identity/groups\"}},\"name\":\"groups\"},{\"_links\":{\"self\":{\"href\":\"http://example.com/api/identity/providers\"}},\"name\":\"providers\"},{\"_links\":{\"self\":{\"href\":\"http://example.com/api/identity/users\"}},\"name\":\"users\"}]},\"_links\":{\"doc\":{\"href\":\"http://example.com/api/docs#operation/list-resources\"},\"self\":{\"href\":\"http://example.com/api/identity\"}}}",
		},
		{
			group:        "vcs",
			responseCode: 200,
			response:     "{\"_embedded\":{\"resources\":[{\"_links\":{\"self\":{\"href\":\"http://example.com/api/vcs/git\"}},\"name\":\"git\"}]},\"_links\":{\"doc\":{\"href\":\"http://example.com/api/docs#operation/list-resources\"},\"self\":{\"href\":\"http://example.com/api/vcs\"}}}",
		},
		{
			group:        "404",
			responseCode: 404,
			response:     "",
		},
	} {

		t.Run(test.group, func(t *testing.T) {

			requestURI := "/api/" + test.group
			u, err := url.Parse("http://" + requestHost + requestURI)
			assert.NoError(t, err)

			oh.Request = &http.Request{
				Method:     "GET",
				Host:       requestHost,
				RequestURI: requestURI,
				URL:        u,
			}

			r := ListResourcesController(oh).Handle(resources.ListResourcesParams{
				HTTPRequest: oh.Request,
				Group:       test.group,
			})

			w := httptest.NewRecorder()
			p := mockProducer{}
			r.WriteResponse(w, p)

			buf := new(bytes.Buffer)
			buf.ReadFrom(w.Result().Body)
			body := buf.String()

			assert.Equal(t, test.responseCode, w.Result().StatusCode)
			assert.Equal(t, test.response, body)

		})
	}
}
