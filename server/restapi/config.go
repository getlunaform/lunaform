package restapi

type cfg struct {
	Identity struct {
		Defaults []struct {
			User     string `json:"user"`
			Password string `json:"password"`
		} `json:"defaults"`
	} `json:"identity"`
	Backend struct {
		Database cfgBackendDb  `json:"database"`
		Identity cfgBackendIdp `json:"identity"`
	}
}

type cfgBackendDb struct{}
type cfgBackendIdp struct{}

func (cbd *cfgBackendDb) UnmarshallJSON() {

}
