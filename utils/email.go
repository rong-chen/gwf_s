package utils

// EmailBody F: From 发送者
// EmailBody T: To 目标者
// EmailBody C:Cc 抄送
type EmailBody struct {
	Account  string
	Password string
	F, T, C  string
	HtmlBody string
}
