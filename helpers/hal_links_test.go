package helpers

import (
	"reflect"
	"testing"

	"github.com/getlunaform/lunaform/models/hal"
	"github.com/stretchr/testify/assert"
)

func Test_newHalRscLinks(t *testing.T) {
	tests := []struct {
		name string
		want *hal.HalRscLinks
	}{
		{
			name: "basic",
			want: &hal.HalRscLinks{HalRscLinks: map[string]*hal.HalHref{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := newHalRscLinks()
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestHalRootRscLinks(t *testing.T) {
	type args struct {
		ch *ContextHelper
	}
	tests := []struct {
		name      string
		args      args
		wantLinks *hal.HalRscLinks
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLinks := HalRootRscLinks(tt.args.ch); !reflect.DeepEqual(gotLinks, tt.wantLinks) {
				t.Errorf("HalRootRscLinks() = %v, want %v", gotLinks, tt.wantLinks)
			}
		})
	}
}

func TestHalSelfLink(t *testing.T) {
	for _, tt := range []struct {
		name  string
		links *hal.HalRscLinks
		href  string
		want  *hal.HalRscLinks
	}{
		{
			name:  "nil-hal-src-links",
			links: nil,
			href:  "/my-mock",
			want: &hal.HalRscLinks{
				HalRscLinks: map[string]*hal.HalHref{
					"lf:self": {Href: "/my-mock"},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := HalSelfLink(tt.links, tt.href)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_halDocLink(t *testing.T) {
	for _, tt := range []struct {
		name        string
		links       *hal.HalRscLinks
		operationId string
		want        *hal.HalRscLinks
	}{
		{
			name:        "nil-hal-src-links",
			links:       nil,
			operationId: "my-mock",
			want: &hal.HalRscLinks{
				HalRscLinks: map[string]*hal.HalHref{
					"doc:my-mock": {
						Href: "/my-mock",
					},
				},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := HalDocLink(tt.links, tt.operationId)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_halAddCuries(t *testing.T) {

	for _, tt := range []struct {
		name  string
		ch    *ContextHelper
		links *hal.HalRscLinks
		want  *hal.HalRscLinks
	}{
		{
			name:  "nil-hal-src-links",
			links: nil,
			ch:    &ContextHelper{},
			want: &hal.HalRscLinks{
				Curies: []*hal.HalCurie{{
					Href:      "/{rel}",
					Name:      "lf",
					Templated: true,
				}, {
					Href:      "/docs#operation/{rel}",
					Name:      "doc",
					Templated: true,
				}},
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			got := HalAddCuries(tt.ch, tt.links)
			assert.Equal(t, tt.want, got)
		})
	}
}
