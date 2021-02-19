package param

type ChangePasswordParam struct {
	Token       string `json:"token"`
	Code        string `json:"code"`
	Account     string `json:"account"`
	NewPassword string `json:"new_password"`
}
