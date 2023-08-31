package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

var ratelimitBucket []string

func main() {
	router := gin.Default()
	router.Static("", "./digigram")
	router.POST("/contact", Email())
	router.Run(":80")
}
func Email() gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.PostForm("email")
		name := c.PostForm("name")
		message := c.PostForm("message")
		ratelimitBucket = append(ratelimitBucket, c.ClientIP())
		myemail := "contact@lekech.com"
		password := "nfrjaiamdbckjdkl"
		smtpServer := "smtp.gmail.com"
		smtpPort := 587
		m := gomail.NewMessage()
		m.SetHeader("From", myemail)
		m.SetHeader("To", "bennyfoefenny@icloud.com")
		m.SetHeader("Subject", "[Auto] From "+name)
		m.SetBody("text/html", `
<!DOCTYPE html>
<html>
<head>
<style>
    body {
        font-family: Arial, sans-serif;
        margin: 0;
        padding: 0;
        background-color: #f4f4f4;
    }
    .container {
        max-width: 600px;
        margin: 0 auto;
        padding: 20px;
        background-color: #ffffff;
        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
    }
    h1 {
        color: #333333;
    }
    p {
        color: #666666;
        line-height: 1.6;
    }
</style>
</head>
<body>
    <div class="container">
		<h3>`+message+`</h3>
        <p>From: `+email+`</p>
    </div>
</body>
</html>

`)

		d := gomail.NewDialer(smtpServer, smtpPort, myemail, password)
		if err := d.DialAndSend(m); err != nil {
			fmt.Println("Error sending email:", err)
		} else {
			fmt.Println("Email sent successfully!")
		}
	}
}
