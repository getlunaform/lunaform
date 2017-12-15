package controller

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/zeebox/terraform-server/server/restapi/operations/resources"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestResourceGroupsController(t *testing.T) {
	for _, test := range []struct {
		responseCode int
		response     string
	}{
		{
			responseCode: 200,
			response:     "{\"_embedded\":{\"resources\":[{\"_links\":{\"self\":{\"href\":\"/tf\"}},\"name\":\"tf\"},{\"_links\":{\"self\":{\"href\":\"/identity\"}},\"name\":\"identity\"},{\"_links\":{\"self\":{\"href\":\"/vcs\"}},\"name\":\"vcs\"}]},\"_links\":{\"doc\":{\"href\":\"/docs#operation/\"},\"self\":{}}}",
		},
	} {
		r := ListResourceGroupsController(nil, ContextHelper{}).Handle(resources.ListResourceGroupsParams{
			HTTPRequest: &http.Request{
				Host: "example.com",
			},
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

	for _, test := range []struct {
		group        string
		responseCode int
		response     string
	}{
		{
			group:        "tf",
			responseCode: 200,
			response:     "{\"_embedded\":{\"resources\":[{\"_links\":{\"self\":{\"href\":\"/modules\"}},\"name\":\"modules\"},{\"_links\":{\"self\":{\"href\":\"/stacks\"}},\"name\":\"stacks\"},{\"_links\":{\"self\":{\"href\":\"/state-backends\"}},\"name\":\"state-backends\"},{\"_links\":{\"self\":{\"href\":\"/workspaces\"}},\"name\":\"workspaces\"}]},\"_links\":{\"doc\":{\"href\":\"/docs#operation/\"},\"self\":{}}}",
		},
		{
			group:        "identity",
			responseCode: 200,
			response:     "{\"_embedded\":{\"resources\":[{\"_links\":{\"self\":{\"href\":\"/groups\"}},\"name\":\"groups\"},{\"_links\":{\"self\":{\"href\":\"/providers\"}},\"name\":\"providers\"},{\"_links\":{\"self\":{\"href\":\"/users\"}},\"name\":\"users\"}]},\"_links\":{\"doc\":{\"href\":\"/docs#operation/\"},\"self\":{}}}",
		},
		{
			group:        "vcs",
			responseCode: 200,
			response:     "{\"_embedded\":{\"resources\":[{\"_links\":{\"self\":{\"href\":\"/git\"}},\"name\":\"git\"}]},\"_links\":{\"doc\":{\"href\":\"/docs#operation/\"},\"self\":{}}}",
		},
		{
			group:        "404",
			responseCode: 404,
			response:     "",
		},
	} {

		t.Run(test.group, func(t *testing.T) {
			r := ListResourcesController(nil, ContextHelper{}).Handle(resources.ListResourcesParams{
				HTTPRequest: &http.Request{
					Host: "example.com",
				},
				Group: test.group,
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
