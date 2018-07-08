package swagger

type HalLinked interface {
	Clean() interface{}
}

func (rlr *ResponseListResources) Clean() interface{} {
	return rlr.Embedded.Clean()
}

func (rltm *ResponseListTfModules) Clean() interface{} {
	modules := []string{}
	for _, rsc := range rltm.Embedded {
		modules = append(modules, rsc.Name)
	}
	return modules
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

func (r *ResponseTfModule) Clean() interface{} {
	return map[string]string{
		"name": r.Name,
		"type": r.Type_,
	}
}
