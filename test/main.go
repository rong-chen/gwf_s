package main

import (
	"fmt"
	"gwf/api/email/service"
)

func main() {
	var emailBody service.Mailer
	emailBody.T = "2326229033@qq.com"
	emailBody.F = "1416307833@qq.com"
	emailBody.Account = "1416307833@qq.com"
	emailBody.HtmlBody = fmt.Sprintf(HtmlBody, "q2326229033的野狗出来挨骂")
	emailBody.C = "1416307833@qq.com"
	// yetlmmkncinzhdjj
	emailBody.Password = "yetlmmkncinzhdjj"
	err := emailBody.SendEmail()
	if err != nil {
		panic(err)
	}
}

var HtmlBody = `
	 <html>
        <head>
            <style>
                body {
                    font-family: Arial, sans-serif;
                    background-color: #f0f0f0;
                }
                .container {
                    max-width: 600px;
                    margin: 0 auto;
                    padding: 20px;
                    background-color: #ffffff;
                    border-radius: 10px;
                    box-shadow: 0px 0px 10px 0px rgba(0,0,0,0.1);
                }
                h1 {
                    color: #333333;
                }
                p {
                    color: #666666;
                }
				.button {
					display: inline-block;
					padding: 5px 10px;
					background-color: #007bff;
					color: #ffffff;
					text-decoration: none;
					border-radius: 5px;
				}
            </style>
        </head>
        <body>
            <div class="container">
                <h1>提醒</h1>
                <p>%s</p>
            </div>
        </body>
        </html>
`
