package model

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
Mail is the model of mail used to stored in the database
*/
type Mail struct {
	Message_ID                string
	Date                      string
	From                      string
	To                        string
	Subject                   string
	Cc                        string
	Mime_Version              string
	Content_Type              string
	Content_Transfer_Encoding string
	Bcc                       string
	X_From                    string
	X_To                      string
	X_cc                      string
	X_bcc                     string
	X_Folder                  string
	X_Origin                  string
	X_FileName                string
	Content                   string
}

// String returns an email string in json format.
// If it is not possible to return the json format, it returns an empty string
func (mail Mail) String() string {
	return mail.ToJson()
}

// ToJson returns an email string in json format.
// If it is not possible to return the json format, it returns an empty string
func (mail Mail) ToJson() string {
	bytes, err := mail.ToJsonBytes()
	if err != nil {
		log.Println(err)
		return ""
	}

	return string(bytes)
}

// ToJsonBytes returns an array of bytes that represent the email in json format
func (mail Mail) ToJsonBytes() ([]byte, error) {
	return json.Marshal(mail)
}

// MailFromJson returns an email from an array of bytes that represent the format of the email
func MailFromJson(_json []byte) Mail {
	var mail Mail

	if err := json.Unmarshal(_json, &mail); err != nil {
		fmt.Println(err)
		return mail
	}

	return mail
}
