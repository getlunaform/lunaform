package models

type HalLinkable interface {
	 Clean() interface{}
}

func (m *ResponseListResources) Clean() {
	m.Links = nil
}