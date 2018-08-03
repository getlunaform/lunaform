package models

import "github.com/getlunaform/lunaform/models/hal"

type HalLinkable interface {
	Clean() interface{}
}

func (m ResourceTfModule) Clean() interface{} {
	m.Links = nil
	return m
}

func (s ResourceTfStack) Clean() interface{} {
	s.Links = nil
	return s
}

func (w ResourceTfWorkspace) Clean() interface{} {
	w.Links = nil
	return w
}

func (w ResourceTfStateBackend) Clean() interface{} {
	w.Links = nil
	return w
}

func (d ResourceTfDeployment) Clean() interface{} {
	d.Links = nil
	return d
}

func (r Resource) Clean() interface{} {
	r.Links = nil
	return r
}

func (d *ResourceListTfDeployment) Clean() interface{} {
	if d != nil {
		rscs := make([]interface{}, len(d.Deployments))
		for i, rsc := range d.Deployments {
			rscs[i] = rsc.Clean()
		}
		return map[string]interface{}{
			"deployments": rscs,
			"stack":       d.Stack.Clean(),
		}
	}
	return make([]interface{}, 0)
}

func (m *ResourceListTfStateBackend) Clean() interface{} {
	if m != nil {
		rscs := make([]interface{}, len(m.StateBackends))
		for i, rsc := range m.StateBackends {
			rscs[i] = rsc.Clean()
		}
		return map[string]interface{}{
			"state-backends": rscs,
		}
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

func (m *ResponseListTfDeployments) Clean() interface{} {
	return m.Embedded.Clean()
}

func (m *ResponseListResources) Clean() interface{} {
	return m.Embedded.Clean()
}

func (se *ServerError) Clean() interface{} {
	return se
}

// generate links
func (s *ResourceTfStack) GenerateLinks(endpoint string) {
	s.Links = &hal.HalRscLinks{
		HalRscLinks: map[string]*hal.HalHref{
			"lf:self": {Href: endpoint + "/" + s.ID},
		},
	}
}

func (m *ResourceTfModule) GenerateLinks(endpoint string) {
	m.Links = &hal.HalRscLinks{
		HalRscLinks: map[string]*hal.HalHref{
			"lf:self": {Href: endpoint + "/" + m.ID},
		},
	}
}

func (d *ResourceTfDeployment) GenerateLinks(endpoint string) {
	d.Links = &hal.HalRscLinks{
		HalRscLinks: map[string]*hal.HalHref{
			"lf:self": {Href: endpoint + "/" + d.ID},
		},
	}
}

func (w *ResourceTfWorkspace) GenerateLinks(endpoint string) {
	w.Links = &hal.HalRscLinks{
		HalRscLinks: map[string]*hal.HalHref{
			"lf:self": {
				Href: endpoint + "/" + *w.Name,
			},
		},
	}
}

type StringHalResponse string

func (s StringHalResponse) Clean() interface{} {
	return string(s)
}
