package identity

type Provider interface {
	IsEditable() bool
	IsFederated() bool

	CreateUser(username string, password string) (User, error)
	ReadUser(username string) (User, error)

	LoginUser(user User, password string) bool
	ChangePassword(user User, password string) error

	ConsumeEndpoint(payload []byte) error
}
