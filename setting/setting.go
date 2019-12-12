package setting

import (
	"os"
	"strconv"
)

var (
	TemplateDocPath = getEnv("TEMPLATE_DOC_PATH", "template.docx")
	ExportDocPath   = getEnv("EXPORT_DOC_PATH", "项目部周报（邹昆峰）.docx")
	JsonFilePath    = getEnv("JSON_FILE_PATH", "data/report.json")
	MailSubject     = getEnv("MAIL_SUBJECT", "项目部周报（邹昆峰）.docx")
	MailHost        = getEnv("MAIL_HOST", "smtp.qiye.aliyun.com")
	MailPort        = getEnvInt("MAIL_PORT", 465)
	MailFrom        = getEnv("MAIL_FROM", "zoukf@zparkhr.com.cn")
	MailFromName    = getEnv("MAIL_FROM_NAME", "邹昆峰")
	MailTo          = getEnv("MAIL_TO", "lins@zparkhr.com.cn")
	MailToName      = getEnv("MAIL_TO_NAME", "陈艳男")
	MailUser        = getEnv("MAIL_USER", "zoukf@zparkhr.com.cn")
	MailPass        = getEnv("MAIL_PASS", "")
)

func getEnv(key, defaultValue string) (value string) {
	value = os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) (value int) {
	valueString := os.Getenv(key)
	value, err := strconv.Atoi(valueString)
	if err != nil {
		return defaultValue
	}
	return value
}
