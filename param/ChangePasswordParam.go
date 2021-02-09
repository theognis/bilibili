package param

type ChangePasswordParam struct {
	Token       string
	Code        string
	Account     string
	NewPassword string
}
