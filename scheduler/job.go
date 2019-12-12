package scheduler

import (
	"gopkg.in/gomail.v2"
	"strconv"
	"weekly_report_golang/docx"
	"weekly_report_golang/report"
	"weekly_report_golang/setting"
)

func job() {
	r := report.New().Load()
	rendererDocx(r)
	sendEmail(r)
}

func rendererDocx(r *report.Report) {
	dr := docx.New(setting.TemplateDocPath, setting.ExportDocPath)
	ed := r.EndDay()

	for key, value := range r.Items() {
		dr.Set(key, value)
	}
	dr.Set("name", setting.MailFromName)
	dr.Set("Y", strconv.Itoa(ed.Year()))
	dr.Set("M", strconv.Itoa(int(ed.Month())))
	dr.Set("D", strconv.Itoa(ed.Day()))
	err := dr.Render()
	if err != nil {
		panic(err)
	}
}

func sendEmail(r *report.Report) {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", setting.MailFrom, setting.MailFromName)
	m.SetAddressHeader("To", setting.MailTo, setting.MailToName)
	m.SetHeader("Subject", setting.MailSubject)
	m.SetBody("text/plain", r.Message())
	m.Attach(setting.ExportDocPath)

	d := gomail.NewDialer(setting.MailHost, setting.MailPort, setting.MailUser, setting.MailPass)
	d.SSL = true

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
