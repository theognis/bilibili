package param

type ChangeEmailInfo struct {
	NewEmail   string `json:"new_email"`
	NewCode    string `json:"new_code"`
	OldCode    string `json:"old_code"`
	OldAccount string `json:"old_account"`
	Token      string `json:"token"`
}
