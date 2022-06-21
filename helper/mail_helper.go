package helper

import (
	"email/config"
	"email/entity"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"io/ioutil"
	"net/http"
	"time"
)

func Send(param entity.Email) (error error) {
	conf := config.New()
	m := mail.NewV3Mail()

	from := mail.NewEmail(param.FromName, param.FromEmail)
	content := mail.NewContent("text/html", param.EmailContent)

	m.SetFrom(from)
	m.AddContent(content)

	// create new *Personalization
	personalization := mail.NewPersonalization()

	// populate `personalization` with data
	to := mail.NewEmail(param.ToName, param.ToEmail)
	//cc := mail.NewEmail(conf.Get("MAIL_CC_NAME"), conf.Get("MAIL_CC_ADDRESS"))

	personalization.AddTos(to)
	//personalization.AddCCs(cc)
	personalization.Subject = param.Subject

	// add `personalization` to `m`
	m.AddPersonalizations(personalization)

	if param.Attachment != "" {
		jsonAtt := []byte(param.Attachment)
		var arrayAtt []string
		_ = json.Unmarshal(jsonAtt, &arrayAtt)

		// read/attach .pdf file
		for _, item := range arrayAtt {
			attPdf := mail.NewAttachment()
			filePdf, err := ioutil.ReadFile(item)
			if err == nil {
				newFileName := param.Key + "-" + time.Now().Format("20060102150405") + ".pdf"
				encodedPdf := base64.StdEncoding.EncodeToString([]byte(filePdf))
				attPdf.SetContent(encodedPdf)
				attPdf.SetType("application/pdf")
				attPdf.SetFilename(newFileName)
				attPdf.SetDisposition("attachment")

				m.AddAttachment(attPdf)
			}
		}
	}

	request := sendgrid.GetRequest(conf.Get("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)

	if err != nil || response.StatusCode != http.StatusOK {
		err = errors.New(response.Body)
		return err
	} else {
		return nil
	}
}
