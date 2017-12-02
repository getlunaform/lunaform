package backend

// IdentityProvider is a specialised database concerned exclusively
// with managing Users and a source of authority for their
// authentication.
type IdentityProvider interface {
	IsEditable() bool
	IsFederated() bool

	CreateUser(username string, password string) (User, error)
	ReadUser(username string) (User, error)
	UpdateUser(user User) error

	LoginUser(user User, password string) bool
	ChangePassword(user User, password string) error

	ConsumeEndpoint(payload []byte) error
}
