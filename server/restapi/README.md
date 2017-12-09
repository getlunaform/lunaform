

# restapi
`import "github.com/zeebox/terraform-server/server/restapi"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package restapi terraform-server
This is a RESTful server for managing Terraform plan and apply jobs and the auditing of actions to approve those apply jobs.
# Introduction
The inspiration for this project is the AWS CloudFormation API's. The intention is to implement a locking mechanism
not only for the terraform state, but for the plan and apply of terraform modules. Once a `module` plan starts, it
is instantiated as a `stack` within the nomencalture of `terraform-server`.
## Terms


	- `module` - The same definition as used within the [terraform ecosystem](<a href="https://www.terraform.io/docs/modules/index.html">https://www.terraform.io/docs/modules/index.html</a>).
	- `stack` - A _stack_ is a _module_ bound to a specific set of parameters. Taking the analogy of classes and objects, if
	   _modules_ are classes, then _stacks_ are instantiated objects.
	- `identity-provider` - Is a source of users either locally managed, or through a federation.

# Authentication
`terraform-server` offers two forms of authentication:


	- Basic Auth
	- API Key

<!-- ReDoc-Inject: <security-definitions> -->

## Identity Providers
Two types of Identity Providers (IdP) are offered. The first are locally managed IdP's which `terraform-server`
handles all management of. The second are read-only federated IdP's.

### Local


	- memory
	- json file

### Federated


	- [SAMLv2.0](<a href="https://www.oasis-open.org/standards#samlv2.0">https://www.oasis-open.org/standards#samlv2.0</a>)
	
	   Schemes:
	     http
	     https
	   Host: localhost
	   BasePath: /api/
	   Version: 0.1.0
	   License: Apache 2.0 <a href="https://github.com/zeebox/terraform-server/blob/master/LICENSE">https://github.com/zeebox/terraform-server/blob/master/LICENSE</a>
	   Contact: <drew.sonne@gmail.com>
	
	   Consumes:
	   - application/vnd.terraform.server.v1+json
	   - application/x-www-form-urlencoded
	
	   Produces:
	   - application/vnd.terraform.server.v1+json

swagger:meta




## <a name="pkg-index">Index</a>
* [Variables](#pkg-variables)
* [type CfgBackend](#CfgBackend)
* [type CfgIdentity](#CfgIdentity)
* [type CfgIdentityDefault](#CfgIdentityDefault)
* [type ConfigFileFlags](#ConfigFileFlags)
* [type Configuration](#Configuration)
* [type Middleware](#Middleware)
  * [func NewMiddleware(h http.Handler) *Middleware](#NewMiddleware)
  * [func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request)](#Middleware.ServeHTTP)
* [type Server](#Server)
  * [func NewServer(api *operations.TerraformServerAPI) *Server](#NewServer)
  * [func (s *Server) ConfigureAPI()](#Server.ConfigureAPI)
  * [func (s *Server) ConfigureFlags()](#Server.ConfigureFlags)
  * [func (s *Server) Fatalf(f string, args ...interface{})](#Server.Fatalf)
  * [func (s *Server) GetHandler() http.Handler](#Server.GetHandler)
  * [func (s *Server) HTTPListener() (net.Listener, error)](#Server.HTTPListener)
  * [func (s *Server) Listen() error](#Server.Listen)
  * [func (s *Server) Logf(f string, args ...interface{})](#Server.Logf)
  * [func (s *Server) Serve() (err error)](#Server.Serve)
  * [func (s *Server) SetAPI(api *operations.TerraformServerAPI)](#Server.SetAPI)
  * [func (s *Server) SetHandler(handler http.Handler)](#Server.SetHandler)
  * [func (s *Server) Shutdown() error](#Server.Shutdown)
  * [func (s *Server) TLSListener() (net.Listener, error)](#Server.TLSListener)
  * [func (s *Server) UnixListener() (net.Listener, error)](#Server.UnixListener)


#### <a name="pkg-files">Package files</a>
[config.go](/src/github.com/zeebox/terraform-server/server/restapi/config.go) [configure_terraform_server.go](/src/github.com/zeebox/terraform-server/server/restapi/configure_terraform_server.go) [doc.go](/src/github.com/zeebox/terraform-server/server/restapi/doc.go) [embedded_spec.go](/src/github.com/zeebox/terraform-server/server/restapi/embedded_spec.go) [se4_middleware.go](/src/github.com/zeebox/terraform-server/server/restapi/se4_middleware.go) [server.go](/src/github.com/zeebox/terraform-server/server/restapi/server.go) 



## <a name="pkg-variables">Variables</a>
``` go
var SwaggerJSON json.RawMessage
```
SwaggerJSON embedded version of the swagger document used at generation time




## <a name="CfgBackend">type</a> [CfgBackend](/src/target/config.go?s=465:677#L25)
``` go
type CfgBackend struct {
    DatabaseType string      `json:"database_type"`
    Database     interface{} `json:"database"`
    IdentityType string      `json:"identity_type"`
    Identity     interface{} `json:"identity"`
}
```









## <a name="CfgIdentity">type</a> [CfgIdentity](/src/target/config.go?s=285:361#L16)
``` go
type CfgIdentity struct {
    Defaults []CfgIdentityDefault `json:"defaults"`
}
```









## <a name="CfgIdentityDefault">type</a> [CfgIdentityDefault](/src/target/config.go?s=363:463#L20)
``` go
type CfgIdentityDefault struct {
    User     string `json:"user"`
    Password string `json:"password"`
}
```









## <a name="ConfigFileFlags">type</a> [ConfigFileFlags](/src/target/config.go?s=862:982#L43)
``` go
type ConfigFileFlags struct {
    ConfigFile string `short:"c" long:"config" description:"Path to configuration on disk"`
}
```









## <a name="Configuration">type</a> [Configuration](/src/target/config.go?s=175:283#L11)
``` go
type Configuration struct {
    Identity CfgIdentity `json:"identity"`
    Backend  CfgBackend  `json:"backend"`
}
```









## <a name="Middleware">type</a> [Middleware](/src/target/se4_middleware.go?s=154:225#L10)
``` go
type Middleware struct {
    SE4 goose4.Goose4
    // contains filtered or unexported fields
}
```
Middleware handles the "/service" prefix to comply with the SE4 prefix







### <a name="NewMiddleware">func</a> [NewMiddleware](/src/target/se4_middleware.go?s=315:361#L17)
``` go
func NewMiddleware(h http.Handler) *Middleware
```
NewMiddleware takes an http handler
to wrap and returns mutable Middleware object





### <a name="Middleware.ServeHTTP">func</a> (\*Middleware) [ServeHTTP](/src/target/se4_middleware.go?s=476:546#L24)
``` go
func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request)
```
ServeHTTP wraps our requests and handles any calles to `/service*`.




## <a name="Server">type</a> [Server](/src/target/server.go?s=1156:4202#L64)
``` go
type Server struct {
    EnabledListeners []string         `long:"scheme" description:"the listeners to enable, this can be repeated and defaults to the schemes in the swagger spec"`
    CleanupTimeout   time.Duration    `long:"cleanup-timeout" description:"grace period for which to wait before shutting down the server" default:"10s"`
    MaxHeaderSize    flagext.ByteSize `long:"max-header-size" description:"controls the maximum number of bytes the server will read parsing the request header's keys and values, including the request line. It does not limit the size of the request body." default:"1MiB"`

    SocketPath flags.Filename `long:"socket-path" description:"the unix socket to listen on" default:"/var/run/terraform-server.sock"`

    Host         string        `long:"host" description:"the IP to listen on" default:"localhost" env:"HOST"`
    Port         int           `long:"port" description:"the port to listen on for insecure connections, defaults to a random value" env:"PORT"`
    ListenLimit  int           `long:"listen-limit" description:"limit the number of outstanding requests"`
    KeepAlive    time.Duration `long:"keep-alive" description:"sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)" default:"3m"`
    ReadTimeout  time.Duration `long:"read-timeout" description:"maximum duration before timing out read of the request" default:"30s"`
    WriteTimeout time.Duration `long:"write-timeout" description:"maximum duration before timing out write of the response" default:"60s"`

    TLSHost           string         `long:"tls-host" description:"the IP to listen on for tls, when not specified it's the same as --host" env:"TLS_HOST"`
    TLSPort           int            `long:"tls-port" description:"the port to listen on for secure connections, defaults to a random value" env:"TLS_PORT"`
    TLSCertificate    flags.Filename `long:"tls-certificate" description:"the certificate to use for secure connections" env:"TLS_CERTIFICATE"`
    TLSCertificateKey flags.Filename `long:"tls-key" description:"the private key to use for secure conections" env:"TLS_PRIVATE_KEY"`
    TLSCACertificate  flags.Filename `long:"tls-ca" description:"the certificate authority file to be used with mutual tls auth" env:"TLS_CA_CERTIFICATE"`
    TLSListenLimit    int            `long:"tls-listen-limit" description:"limit the number of outstanding requests"`
    TLSKeepAlive      time.Duration  `long:"tls-keep-alive" description:"sets the TCP keep-alive timeouts on accepted connections. It prunes dead TCP connections ( e.g. closing laptop mid-download)"`
    TLSReadTimeout    time.Duration  `long:"tls-read-timeout" description:"maximum duration before timing out read of the request"`
    TLSWriteTimeout   time.Duration  `long:"tls-write-timeout" description:"maximum duration before timing out write of the response"`
    // contains filtered or unexported fields
}
```
Server for the terraform server API







### <a name="NewServer">func</a> [NewServer](/src/target/server.go?s=666:724#L42)
``` go
func NewServer(api *operations.TerraformServerAPI) *Server
```
NewServer creates a new api terraform server server but does not configure it





### <a name="Server.ConfigureAPI">func</a> (\*Server) [ConfigureAPI](/src/target/server.go?s=821:852#L50)
``` go
func (s *Server) ConfigureAPI()
```
ConfigureAPI configures the API and handlers.




### <a name="Server.ConfigureFlags">func</a> (\*Server) [ConfigureFlags](/src/target/server.go?s=1032:1065#L57)
``` go
func (s *Server) ConfigureFlags()
```
ConfigureFlags configures the additional flags defined by the handlers. Needs to be called before the parser.Parse




### <a name="Server.Fatalf">func</a> (\*Server) [Fatalf](/src/target/server.go?s=4616:4670#L107)
``` go
func (s *Server) Fatalf(f string, args ...interface{})
```
Fatalf logs message either via defined user logger or via system one if no user logger is defined.
Exits with non-zero status after printing




### <a name="Server.GetHandler">func</a> (\*Server) [GetHandler](/src/target/server.go?s=12292:12334#L382)
``` go
func (s *Server) GetHandler() http.Handler
```
GetHandler returns a handler useful for testing




### <a name="Server.HTTPListener">func</a> (\*Server) [HTTPListener](/src/target/server.go?s=12763:12816#L402)
``` go
func (s *Server) HTTPListener() (net.Listener, error)
```
HTTPListener returns the http listener




### <a name="Server.Listen">func</a> (\*Server) [Listen](/src/target/server.go?s=10495:10526#L305)
``` go
func (s *Server) Listen() error
```
Listen creates the listeners for the server




### <a name="Server.Logf">func</a> (\*Server) [Logf](/src/target/server.go?s=4304:4356#L97)
``` go
func (s *Server) Logf(f string, args ...interface{})
```
Logf logs message either via defined user logger or via system one if no user logger is defined.




### <a name="Server.Serve">func</a> (\*Server) [Serve](/src/target/server.go?s=5321:5357#L144)
``` go
func (s *Server) Serve() (err error)
```
Serve the api




### <a name="Server.SetAPI">func</a> (\*Server) [SetAPI](/src/target/server.go?s=4884:4943#L117)
``` go
func (s *Server) SetAPI(api *operations.TerraformServerAPI)
```
SetAPI configures the server with the specified API. Needs to be called before Serve




### <a name="Server.SetHandler">func</a> (\*Server) [SetHandler](/src/target/server.go?s=12421:12470#L387)
``` go
func (s *Server) SetHandler(handler http.Handler)
```
SetHandler allows for setting a http handler on this server




### <a name="Server.Shutdown">func</a> (\*Server) [Shutdown](/src/target/server.go?s=12166:12199#L376)
``` go
func (s *Server) Shutdown() error
```
Shutdown server and clean up resources




### <a name="Server.TLSListener">func</a> (\*Server) [TLSListener](/src/target/server.go?s=12976:13028#L412)
``` go
func (s *Server) TLSListener() (net.Listener, error)
```
TLSListener returns the https listener




### <a name="Server.UnixListener">func</a> (\*Server) [UnixListener](/src/target/server.go?s=12548:12601#L392)
``` go
func (s *Server) UnixListener() (net.Listener, error)
```
UnixListener returns the domain socket listener








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
