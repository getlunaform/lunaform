package restapi

import (
	"encoding/json"
	"io"
	"reflect"
	"testing"

	"github.com/Flaque/filet"
	"github.com/getlunaform/lunaform/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var api *operations.LunaformAPI

var defaultConfigPayload = `
---
identity:
  defaults:
    - user: admin
      password: mock_password
backend:
  database:
    type: json
    path: config/db.yaml
  identity:
    - type: json
      path: config/auth-db.yaml`

func Test_configFile(t *testing.T) {

	c := Configuration{}

	err := yaml.Unmarshal([]byte(defaultConfigPayload), &c)
	assert.Nil(t, err)

	assert.NotNil(t, c.Identity.Defaults)
	assert.Len(t, c.Identity.Defaults, 1)

	u := c.Identity.Defaults[0]
	assert.NotNil(t, u)
	assert.Equal(t, "admin", u.User)
	assert.Equal(t, "mock_password", u.Password) // @TODO THis should be a bcrypt
}

func Test_cliOptions(t *testing.T) {
	api := operations.LunaformAPI{}

	assert.Empty(t, api.CommandLineOptionsGroups)

	configureFlags(&api)

	assert.Len(t, api.CommandLineOptionsGroups, 1)
	opt := api.CommandLineOptionsGroups[0]

	assert.Equal(t, "Terraform Server", opt.ShortDescription)
	assert.Equal(t, "Server Configuration", opt.LongDescription)

	assert.NotNil(t, opt.Options)
	assert.IsType(t, &ConfigFileFlags{}, opt.Options)

}

func Test_loadCliConfiguration(t *testing.T) {
	defer filet.CleanUp(t)
	file := filet.TmpFile(t, "", defaultConfigPayload)

	cliconfig = ConfigFileFlags{
		ConfigFile: file.Name(),
	}

	cfg, err := parseCliConfiguration()

	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	assert.IsType(t, &Configuration{}, cfg)

}

type mockProducer struct {
	ProducerHandler func(w io.Writer, i interface{}) (err error)
}

func (mp mockProducer) Produce(w io.Writer, i interface{}) (err error) {
	var b []byte
	b, err = json.Marshal(i)
	w.Write(b)
	return
}

func init() {
	swaggerSpec, err := loads.Analyzed(SwaggerJSON, "")
	if err != nil {
		panic(err)
	}

	api = operations.NewLunaformAPI(swaggerSpec)
	configureAPI(api)

	return
}

func TestConfiguration_loadFromFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		cfg     *Configuration
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.cfg.loadFromFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("Configuration.loadFromFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseCliConfiguration(t *testing.T) {
	tests := []struct {
		name    string
		wantCfg *Configuration
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := parseCliConfiguration()
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCliConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("parseCliConfiguration() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

func Test_newDefaultConfiguration(t *testing.T) {
	tests := []struct {
		name string
		want *Configuration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDefaultConfiguration(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDefaultConfiguration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_configureFlags(t *testing.T) {
	type args struct {
		api *operations.LunaformAPI
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configureFlags(tt.args.api)
		})
	}
}
