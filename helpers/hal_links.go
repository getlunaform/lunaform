package helpers

import (
	"github.com/getlunaform/lunaform/models"
	"github.com/go-openapi/strfmt"
)

func newHalRscLinks() *models.HalRscLinks {
	return &models.HalRscLinks{
		HalRscLinks: make(map[string]*models.HalHref),
	}
}

func HalRootRscLinks(ch ContextHelper) (links *models.HalRscLinks) {
	links = newHalRscLinks()

	HalAddCuries(ch, links)
	HalSelfLink(links, ch.Endpoint)
	HalDocLink(links, ch.OperationID)

	return links
}

func HalSelfLink(links *models.HalRscLinks, href string) *models.HalRscLinks {
	if links == nil {
		links = newHalRscLinks()
	}

	links.HalRscLinks["lf:self"] = &models.HalHref{Href: href}

	return links
}

func HalDocLink(links *models.HalRscLinks, operationId string) *models.HalRscLinks {
	if links == nil {
		links = newHalRscLinks()
	}

	links.HalRscLinks["doc:"+operationId] = &models.HalHref{Href: "/" + operationId}

	return links
}

func HalAddCuries(ch ContextHelper, links *models.HalRscLinks) *models.HalRscLinks {
	if links == nil {
		links = &models.HalRscLinks{}
	}

	links.Curies = []*models.HalCurie{
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
