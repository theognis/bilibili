package param

type ChangePhoneInfo struct {
	NewPhone   string `form:"new_phone"`
	NewCode    string `form:"new_code"`
	OldCode    string `form:"old_code"`
	OldAccount string `form:"old_account"`
	Token      string `form:"token"`
}
