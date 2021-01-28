package param

type ChangeEmailInfo struct {
	NewEmail        string `form:"new_email"`
	NewCode         string `form:"new_code"`
	OriginalCode    string `form:"original_code"`
	OriginalAddress string `form:"original_address"`
	Token           string `form:"token"`
}
