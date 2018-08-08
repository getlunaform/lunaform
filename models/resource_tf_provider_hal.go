package models

import "github.com/getlunaform/lunaform/helpers"

func (rtp *ResourceTfProvider) BuildRscHalLinks(ch *helpers.ContextHelper) {
	rtp.Links = helpers.HalSelfLink(nil, ch.BasePath)
}

func (rtp *ResourceTfProvider) BuildRootHalLinks(ch *helpers.ContextHelper) {
	rtp.Links = helpers.HalRootRscLinks(ch)
}
