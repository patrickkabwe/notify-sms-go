package notify_sms

import (
	"io"
	"time"
)

type RecipientType string
type SenderID string
type Token string

const (
	NOTIFY_RECIPIENT_TYPE_CUSTOM        RecipientType = "NOTIFY_RECIEPIENT_TYPE_CUSTOM"
	NOTIFY_RECIPIENT_TYPE_CHANNEL       RecipientType = "NOTIFY_RECIEPIENT_TYPE_CHANNEL"
	NOTIFY_RECIPIENT_TYPE_CONTACT_GROUP RecipientType = "NOTIFY_RECIEPIENT_TYPE_CONTACT_GROUP"
)

type MakeRequestOptions struct {
	Headers map[string]string
}

type APIResponse[T any] struct {
	Success bool          `json:"success"`
	Message string        `json:"message"`
	Payload T             `json:"payload"`
	Error   ErrorResponse `json:"error"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Title   string `json:"title"`
	Message string `json:"message"`
}

type AuthAPIResponse struct {
	Token Token `json:"token"`
}

type SenderResponse struct {
}

type SendersAPIResponse struct {
	Data []Sender `json:"data"`
}
type SenderAPIResponse struct {
	Data Sender `json:"data"`
}

type Sender struct {
	Id             string    `json:"_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Tracker        Tracker   `json:"tracker"`
	Status         string    `json:"status"`
	Active         bool      `json:"active"`
	Account        string    `json:"account"`
	User           string    `json:"user"`
	CreatedOn      time.Time `json:"createdOn"`
	LastModifiedOn time.Time `json:"lastModifiedOn"`
}

type Tracker struct {
	Id             string    `json:"_id"`
	Title          string    `json:"title"`
	AutoApprove    bool      `json:"autoApprove"`
	Status         string    `json:"status"`
	Active         bool      `json:"active"`
	CreatedOn      time.Time `json:"createdOn"`
	LastModifiedOn time.Time `json:"lastModifiedOn"`
}

type CreateSenderIDParams struct {
	BusinessName string
	Description  string
}

type NewClientParams struct {
	Username    string `json:"userName"`
	Password    string `json:"password"`
	makeRequest func(method, endpoint string, params io.Reader, opt MakeRequestOptions) ([]byte, error)
}

type Message struct {
	RecipientType RecipientType `json:"reciepientType"`
	// SenderID - sender id collected from GetSenderIDs
	SenderID string `json:"senderId"`
	// Contacts - phone numbers you are trying to send sms to.
	// e.g [+260979600000]
	Contacts []string `json:"reciepients"`
	// Channel - channel you are trying to send sms message to. set RecipientType to NOTIFY_RECIPIENT_TYPE_CHANNEL
	Channel string
	// ContactGroup - The group id you are trying to send a sms message.
	// PS: Only define this field when you are trying to send sms to a specific group. set RecipientType to NOTIFY_RECIPIENT_TYPE_CONTACT_GROUP
	ContactGroup string `json:"contactGroup"`
	// Message - Text message you are sending to contact(s)
	// e.g Hello Notify
	Message string `json:"message"`
}
