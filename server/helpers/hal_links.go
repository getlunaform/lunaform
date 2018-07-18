package helpers

import (
	"github.com/go-openapi/strfmt"
	"github.com/drewsonne/lunaform/server/models"
)

func HalRootRscLinks(ch ContextHelper) *models.HalRscLinks {
	lnks := HalSelfLink(ch.FQEndpoint)
	lnks.Doc = HalDocLink(ch).Doc
	return lnks
}

func HalSelfLink(href string) *models.HalRscLinks {
	return &models.HalRscLinks{
		Self: &models.HalHref{Href: strfmt.URI(href)},
	}
}

func HalDocLink(ch ContextHelper) *models.HalRscLinks {
	return &models.HalRscLinks{
		Doc: &models.HalHref{
			Href: strfmt.URI(
				ch.ServerURL + "/docs#operation/" + ch.OperationID,
			),
		},
	}
}

