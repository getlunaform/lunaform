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
			response:     `{"_embedded":{"resources":[{"_links":{"lf:self":{"href":"/tf/modules"}},"name":"modules"},{"_links":{"lf:self":{"href":"/tf/stacks"}},"name":"stacks"},{"_links":{"lf:self":{"href":"/tf/state-backends"}},"name":"state-backends"},{"_links":{"lf:self":{"href":"/tf/workspaces"}},"name":"workspaces"}]},"_links":{"curies":[{"href":"http://example.com/api/{rel}","name":"lf","templated":true},{"href":"http://example.com/api/docs#operation/{rel}","name":"doc","templated":true}],"doc:list-resources":{"href":"/list-resources"},"lf:self":{"href":"/tf"}}}`,
		},
		{
			group:        "identity",
			responseCode: 200,
			response:     `{"_embedded":{"resources":[{"_links":{"lf:self":{"href":"/identity/groups"}},"name":"groups"},{"_links":{"lf:self":{"href":"/identity/providers"}},"name":"providers"},{"_links":{"lf:self":{"href":"/identity/users"}},"name":"users"}]},"_links":{"curies":[{"href":"http://example.com/api/{rel}","name":"lf","templated":true},{"href":"http://example.com/api/docs#operation/{rel}","name":"doc","templated":true}],"doc:list-resources":{"href":"/list-resources"},"lf:self":{"href":"/identity"}}}`,
		},
		{
			group:        "vcs",
			responseCode: 200,
			response:     `{"_embedded":{"resources":[{"_links":{"lf:self":{"href":"/vcs/git"}},"name":"git"}]},"_links":{"curies":[{"href":"http://example.com/api/{rel}","name":"lf","templated":true},{"href":"http://example.com/api/docs#operation/{rel}","name":"doc","templated":true}],"doc:list-resources":{"href":"/list-resources"},"lf:self":{"href":"/vcs"}}}`,
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
