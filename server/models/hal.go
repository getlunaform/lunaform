package models

type HalLinkable interface {
	Clean() interface{}
}

func (m *ResponseListResources) Clean() interface{} {
	m.Links = nil
	return m
}

func (m *ResourceTfModule) Clean() interface{} {
	m.Links = nil
	return m
}
