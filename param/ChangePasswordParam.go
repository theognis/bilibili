package param

type ChangePasswordParam struct {
	Code        string
	Account     string
	NewPassword string
}
