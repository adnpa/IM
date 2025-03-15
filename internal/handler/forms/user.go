package forms

type RegisterForm struct {
	Mobile   string `json:"mobile,omitempty"`
	Password string `json:"password,omitempty"`
	Code     string `json:"code,omitempty"`
}

type PwdLoginForm struct {
	Mobile    string `json:"mobile,omitempty"`
	Password  string `json:"password,omitempty"`
	Captcha   string `json:"captcha,omitempty"`
	CaptchaId string `json:"captcha_id,omitempty"`
}
