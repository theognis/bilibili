package param

type ChangePasswordParam struct {
	Code        string `form:"code"`
	Account     string `form:"account"`
	NewPassword string `form:"new_password"`
}
