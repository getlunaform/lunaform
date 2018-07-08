package swagger

type HalLinked interface {
	Clean() interface{}
}

func (rlr *ResponseListResources) Clean() interface{} {
	return rlr.Embedded.Clean()
}

func (rl *ResourceList) Clean() interface{} {
	names := []map[string]string{}
	for _, rsc := range rl.Resources {
		names = append(names, map[string]string{
			"name": rsc.Name,
		})
	}
	return names
}
