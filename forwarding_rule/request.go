package forwarding_rule

import (
	"github.com/simonsk/opsgenie-go-sdk-v2/client"
	"github.com/pkg/errors"
	"time"
)

type Identifier uint32

type CreateRequest struct {
	client.BaseRequest
	FromUser  User      `json:"fromUser"`
	ToUser    User      `json:"toUser"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Alias     string    `json:"alias,omitempty"`
}

func (r *CreateRequest) Validate() error {

	err := validateUser(&r.ToUser, "ToUser cannot be empty!")
	if err != nil {
		return err
	}

	err = validateUser(&r.FromUser, "FromUser cannot be empty!")
	if err != nil {
		return err
	}

	err = validateDates(&r.StartDate, "Start date cannot be empty.")
	if err != nil {
		return err
	}
	err = validateDates(&r.EndDate, "End date cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (r *CreateRequest) ResourcePath() string {

	return "/v2/forwarding-rules"
}

func (r *CreateRequest) Method() string {
	return "POST"
}

func (r *CreateRequest) RequestParams() map[string]string {

	return make(map[string]string)
}

type GetRequest struct {
	client.BaseRequest
	IdentifierType  Identifier
	IdentifierValue string
}

func (r *GetRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r *GetRequest) ResourcePath() string {

	return "/v2/forwarding-rules/" + r.IdentifierValue
}

func (r *GetRequest) Method() string {
	return "GET"
}

func (r *GetRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == Alias {
		params["identifierType"] = "alias"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

type UpdateRequest struct {
	client.BaseRequest
	IdentifierType  Identifier
	IdentifierValue string
	ToUser          User      `json:"toUser"`
	FromUser        User      `json:"fromUser"`
	StartDate       time.Time `json:"startDate"`
	EndDate         time.Time `json:"endDate"`
}

func (r *UpdateRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	err = validateUser(&r.ToUser, "ToUser cannot be empty!")
	if err != nil {
		return err
	}

	err = validateUser(&r.FromUser, "FromUser cannot be empty!")
	if err != nil {
		return err
	}

	err = validateDates(&r.StartDate, "Start date cannot be empty.")
	if err != nil {
		return err
	}
	err = validateDates(&r.EndDate, "End date cannot be empty.")
	if err != nil {
		return err
	}
	return nil
}

func (r *UpdateRequest) ResourcePath() string {

	return "/v2/forwarding-rules/" + r.IdentifierValue
}

func (r *UpdateRequest) Method() string {
	return "PUT"
}

func (r *UpdateRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == Alias {
		params["identifierType"] = "alias"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

type DeleteRequest struct {
	client.BaseRequest
	IdentifierType  Identifier
	IdentifierValue string
}

func (r *DeleteRequest) Validate() error {
	err := validateIdentifier(r.IdentifierValue)
	if err != nil {
		return err
	}
	return nil
}

func (r *DeleteRequest) ResourcePath() string {

	return "/v2/forwarding-rules/" + r.IdentifierValue
}

func (r *DeleteRequest) Method() string {
	return "DELETE"
}

func (r *DeleteRequest) RequestParams() map[string]string {

	params := make(map[string]string)

	if r.IdentifierType == Alias {
		params["identifierType"] = "alias"
	} else {
		params["identifierType"] = "id"
	}

	return params
}

type ListRequest struct {
	client.BaseRequest
}

func (r *ListRequest) Validate() error {
	return nil
}

func (r *ListRequest) ResourcePath() string {

	return "/v2/forwarding-rules"
}

func (r *ListRequest) Method() string {
	return "GET"
}

func (r *ListRequest) RequestParams() map[string]string {
	return make(map[string]string)
}

type User struct {
	Id       string `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

func validateIdentifier(identifier string) error {
	if identifier == "" {
		return errors.New("Forwarding Rule identifier cannot be empty.")
	}
	return nil
}

func validateUser(user *User, message string) error {
	if *user == (User{}) {
		return errors.New(message)
	}
	if user.Id == "" && user.Username == "" {
		return errors.New(message)
	}
	return nil
}

func validateDates(date *time.Time, message string) error {
	if *date == (time.Time{}) {
		return errors.New(message)
	}
	return nil
}

const (
	Id Identifier = iota
	Alias
)
