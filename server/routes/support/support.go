package supportroutes

import (
	"net/http"
	"net/smtp"

	server_utils "github.com/TimTwigg/Manwe/server/utils"
	io "github.com/TimTwigg/Manwe/utils/io"
	logger "github.com/TimTwigg/Manwe/utils/log"
	email "github.com/jordan-wright/email"
)

func SupportHandler(w http.ResponseWriter, r *http.Request, userid string) {
	switch r.Method {
	case http.MethodPost:
		logger.PostRequest("SupportHandler: POST request")

		email_addr, err := server_utils.GetSessionUserEmail(userid)
		if err != nil {
			logger.Error("SupportHandler: Error getting user email: " + err.Error())
			http.Error(w, "Error getting user email", http.StatusInternalServerError)
			return
		}

		support_address, err := io.GetEnvVar("SUPPORT_EMAIL")
		if err != nil {
			logger.Error("SupportHandler: Error reading environment variable SUPPORT_EMAIL: " + err.Error())
			http.Error(w, "Error reading environment variable SUPPORT_EMAIL", http.StatusInternalServerError)
			return
		}
		support_secret, err := io.GetEnvVar("SUPPORT_EMAIL_PASSWORD")
		if err != nil {
			logger.Error("SupportHandler: Error reading environment variable SUPPORT_EMAIL_PASSWORD: " + err.Error())
			http.Error(w, "Error reading environment variable SUPPORT_EMAIL_PASSWORD", http.StatusInternalServerError)
			return
		}

		e := email.NewEmail()
		e.From = email_addr
		e.To = []string{support_address}
		e.Subject = "Awesome Subject"
		e.Text = []byte("Text Body is, of course, supported!")
		e.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
		err = e.Send("smtp.gmail.com:587", smtp.PlainAuth("", support_address, support_secret, "smtp.gmail.com")) // Auth not working
		if err != nil {
			logger.Error("SupportHandler: Error sending email: " + err.Error())
			http.Error(w, "Error sending email", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}
}
