package alert

import (
	"github.com/simonsk/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
)

type UpdateMessageRequest struct {
	client.BaseRequest
	IdentifierType  AlertIdentifier
	IdentifierValue string
	Message         string `json:"message,omitempty"`
}

func (r *UpdateMessageRequest) Validate() error {
	if r.Message == "" {
		return errors.New("Message can not be empty")
	}
	if r.IdentifierValue == "" {
		return errors.New("Identifier can not be empty")
	}
	return nil
}

func (r *UpdateMessageRequest) ResourcePath() string {

	return "/v2/alerts/" + r.IdentifierValue + "/message"
}

func (r *UpdateMessageRequest) Method() string {
	return "PUT"
}

func (r *UpdateMessageRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == ALIAS {
		params["identifierType"] = "alias"

	} else if r.IdentifierType == TINYID {
		params["identifierType"] = "tiny"

	} else {
		params["identifierType"] = "id"

	}
	return params
}
