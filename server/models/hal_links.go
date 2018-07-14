package models

import (
	"github.com/go-openapi/strfmt"
	"github.com/drewsonne/lunaform/server"
)

func HalRootRscLinks(ch server.ContextHelper) *HalRscLinks {
	lnks := HalSelfLink(ch.FQEndpoint)
	lnks.Doc = HalDocLink(ch).Doc
	return lnks
}

func HalSelfLink(href string) *HalRscLinks {
	return &HalRscLinks{
		Self: &HalHref{Href: strfmt.URI(href)},
	}
}

func HalDocLink(ch server.ContextHelper) *HalRscLinks {
	return &HalRscLinks{
		Doc: &HalHref{
			Href: strfmt.URI(
				ch.ServerURL + "/docs#operation/" + ch.OperationID,
			),
		},
	}
}
