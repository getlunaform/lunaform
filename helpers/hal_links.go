package helpers

import (
	"github.com/getlunaform/lunaform/models/hal"
	"github.com/go-openapi/strfmt"
)

func newHalRscLinks() *hal.HalRscLinks {
	return &hal.HalRscLinks{
		HalRscLinks: map[string]*hal.HalHref{},
	}
}

func HalRootRscLinks(ch *ContextHelper) (links *hal.HalRscLinks) {
	links = newHalRscLinks()

	HalAddCuries(ch, links)
	HalSelfLink(links, ch.Endpoint)
	HalDocLink(links, ch.OperationID)

	return links
}

func HalSelfLink(links *hal.HalRscLinks, href string) *hal.HalRscLinks {
	if links == nil {
		links = newHalRscLinks()
	}

	noRscLinks := (links.HalRscLinks == nil) || (len(links.HalRscLinks) == 0)
	if noRscLinks {
		links.HalRscLinks = map[string]*hal.HalHref{}
	}

	links.HalRscLinks["lf:self"] = &hal.HalHref{Href: href}

	return links
}

func HalDocLink(links *hal.HalRscLinks, operationId string) *hal.HalRscLinks {
	if links == nil {
		links = newHalRscLinks()
	}

	noRscLinks := (links.HalRscLinks == nil) || (len(links.HalRscLinks) == 0)
	if noRscLinks {
		links.HalRscLinks = map[string]*hal.HalHref{}
	}

	links.HalRscLinks["doc:"+operationId] = &hal.HalHref{Href: "/" + operationId}

	return links
}

func HalAddCuries(ch *ContextHelper, links *hal.HalRscLinks) *hal.HalRscLinks {
	if links == nil {
		links = &hal.HalRscLinks{}
	}

	links.Curies = []*hal.HalCurie{
		{
			Name:      "lf",
			Href:      strfmt.URI(ch.ServerURL + "/{rel}"),
			Templated: true,
		},
		{
			Name:      "doc",
			Href:      strfmt.URI(ch.ServerURL + "/docs#operation/{rel}"),
			Templated: true,
		},
	}
	return links
}
