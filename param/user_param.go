package param

type UserParam struct {
	Username   string `form:"username"`
	Pwd        string `form:"password"`
	Phone      string `form:"phone"`
	VerifyCode string `form:"verify_code"`
}
