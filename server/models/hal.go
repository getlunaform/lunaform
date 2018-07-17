package models

type HalLinkable interface {
	Clean() interface{}
}

func (m *ResponseListResources) Clean() interface{} {
	return m.Embedded.Clean()
}

func (m *ResourceTfModule) Clean() interface{} {
	m.Links = nil
	return m
}

func (s *ResourceTfStack) Clean() interface{} {
	s.Links = nil
	return s
}

func (w *ResourceTfWorkspace) Clean() interface{} {
	w.Links = nil
	return w
}

func (w *ResourceTfStateBackend) Clean() interface{} {
	w.Links = nil
	return w
}

func (r *Resource) Clean() interface{} {
	return &Resource{
		Name: r.Name,
	}
}

func (m *ResourceListTfStateBackend) Clean() interface{} {
	if m != nil {
		rscs := make([]interface{}, len())
	}
	return make([]interface{}, 0)
}

func (m *ResourceListTfModule) Clean() interface{} {
	if m != nil {
		rscs := make([]interface{}, len(m.Modules))
		for i, rsc := range m.Modules {
			rscs[i] = rsc.Clean()
		}
		return map[string]interface{}{
			"modules": rscs,
		}
	}
	return make([]interface{}, 0)
}

func (m *ResourceListTfStack) Clean() interface{} {
	if m != nil {
		rscs := make([]interface{}, len(m.Stacks))
		for i, rsc := range m.Stacks {
			rscs[i] = rsc.Clean()
		}
		return map[string]interface{}{
			"stacks": rscs,
		}
	}
	return make([]interface{}, 0)
}

func (w *ResourceListTfWorkspace) Clean() interface{} {
	if w != nil {
		workspaces := make([]interface{}, len(w.Workspaces))
		for i, workspace := range w.Workspaces {
			workspaces[i] = workspace.Clean()
		}
		return map[string]interface{}{
			"workspaces": workspaces,
		}
	}
	return make([]interface{}, 0)
}

func (list *ResourceList) Clean() interface{} {
	rscs := make([]interface{}, len(list.Resources))
	for i, rsc := range list.Resources {
		rscs[i] = rsc.Clean()
	}
	return rscs
}

func (m *ResponseListTfModules) Clean() interface{} {
	return m.Embedded.Clean()
}

func (m *ResponseListTfStacks) Clean() interface{} {
	return m.Embedded.Clean()
}

func (w *ResponseListTfWorkspaces) Clean() interface{} {
	return w.Embedded.Clean()
}

func (m *ResponseListTfStateBackends) Clean() interface{} {
	return m.Embedded.Clean()
}

func (se *ServerError) Clean() interface{} {
	return se
}
