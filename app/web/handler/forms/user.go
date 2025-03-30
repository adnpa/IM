package forms

type RegisterForm struct {
	Username string `json:"username,omitempty"`
	Mobile   string `json:"mobile,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Code     string `json:"code,omitempty"`
}

type PwdLoginForm struct {
	Username  string `json:"username,omitempty"`
	Mobile    string `json:"mobile,omitempty"`
	Email     string `json:"email,omitempty"`
	Password  string `json:"password,omitempty"`
	Captcha   string `json:"captcha,omitempty"`
	CaptchaId string `json:"captcha_id,omitempty"`
}

type UpdateSelfProfileForm struct {
	Avatar   string `json:"avatar,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Gender   int32  `json:"gender,omitempty"`
	Birthday string `json:"birthday,omitempty"`
	Memo     string `json:"memo,omitempty"`
}
