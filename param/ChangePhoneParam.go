package param

type ChangePhoneInfo struct {
	NewPhone        string `form:"new_phone"`
	NewCode         string `form:"new_code"`
	OriginalCode    string `form:"original_code"`
	OriginalAddress string `form:"original_address"`
	Token           string `form:"token"`
}
