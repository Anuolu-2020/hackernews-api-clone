package users

type WrongUsernameOrPasswordErr struct{}

func (m *WrongUsernameOrPasswordErr) Error() string {
	return "wrong username or password"
}
