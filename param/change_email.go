package param

type ChangeEmailInfo struct {
	NewEmail   string `form:"new_email"`
	NewCode    string `form:"new_code"`
	OldCode    string `form:"old_code"`
	OldAccount string `form:"old_account"`
	Token      string `form:"token"`
}
