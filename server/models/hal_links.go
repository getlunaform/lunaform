package models

import (
	"github.com/drewsonne/lunaform/server/restapi"
	"github.com/go-openapi/strfmt"
)

func HalRootRscLinks(ch restapi.ContextHelper) *HalRscLinks {
	lnks := HalSelfLink(ch.FQEndpoint)
	lnks.Doc = HalDocLink(ch).Doc
	return lnks
}

func HalSelfLink(href string) *HalRscLinks {
	return &HalRscLinks{
		Self: &HalHref{Href: strfmt.URI(href)},
	}
}

func HalDocLink(ch restapi.ContextHelper) *HalRscLinks {
	return &HalRscLinks{
		Doc: &HalHref{
			Href: strfmt.URI(
				ch.ServerURL + "/docs#operation/" + ch.OperationID,
			),
		},
	}
}
