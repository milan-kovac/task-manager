package types

type JWTUser struct {
	ID       uint
	Email    string
	Password string
}

type CurrentUser struct {
	ID    uint
	Email string
}
