package param

type ChangePhoneInfo struct {
	NewPhone   string `json:"new_phone"`
	NewCode    string `json:"new_code"`
	OldCode    string `json:"old_code"`
	OldAccount string `json:"old_account"`
	Token      string `json:"token"`
}
