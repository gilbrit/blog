package controllers

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"net/smtp"
)

type ContactController struct {
	beego.Controller
}

func sendmail(c *ContactController, subject string, body string) {
	// Connect to the remote SMTP server.
	mail, err := smtp.Dial("localhost:25")
	if err != nil {
		beego.Critical(err)
		return
	}
	defer mail.Close()
	// Set the sender and recipient.
	body = "From: webform@linuxbrit.co.uk\n" + body
	body = "Subject: " + subject + "\n\n" + body
	mail.Mail("tom@linuxbrit.co.uk")
	mail.Rcpt("tom@linuxbrit.co.uk")
	// Send the email body.
	wc, err := mail.Data()
	if err != nil {
		beego.Critical(err)
		return
	}
	defer wc.Close()
	buf := bytes.NewBufferString(body)
	if _, err = buf.WriteTo(wc); err != nil {
		beego.Critical(err)
		return
	}
}

func (c *ContactController) Post() {
	subject := "Email from linuxbrit contact form"
	body := ""
	r := c.Ctx.Request

	body = body + fmt.Sprintln("Method:", r.Method)
	body = body + fmt.Sprintln("URL:", r.URL.String())
	query := r.URL.Query()
	for k := range query {
		body = body + fmt.Sprintln("Query", k+":", query.Get(k))
	}
	r.ParseForm()
	form := r.Form
	for k := range form {
		body = body + fmt.Sprintln("Form", k+":", form.Get(k))
	}
	post := r.PostForm
	for k := range post {
		body = body + fmt.Sprintln("PostForm", k+":", post.Get(k))
	}
	body = body + fmt.Sprintln("RemoteAddr:", r.RemoteAddr)
	if referer := r.Referer(); len(referer) > 0 {
		body = body + fmt.Sprintln("Referer:", referer)
	}
	if ua := r.UserAgent(); len(ua) > 0 {
		body = body + fmt.Sprintln("UserAgent:", ua)
	}
	for _, cookie := range r.Cookies() {
		body = body + fmt.Sprintln("Cookie", cookie.Name+":", cookie.Value, cookie.Path, cookie.Domain, cookie.RawExpires)
	}
	sendmail(c, subject, body)

	c.Redirect("/?contact=true", 302)
}
