package yonoma

type Email struct {
	FromEmail    string `json:"from_email"`
	ToEmail      string `json:"to_email"`
	Subject      string `json:"subject"`
	MailTemplate string `json:"mail_template"`
}

func (c *Client) Send(email Email) ([]byte, error) {
	return c.request("POST", "email/send", email)
}
