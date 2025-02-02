package heartbeat

import (
	"github.com/simonsk/opsgenie-go-sdk-v2/og"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddRequest_Validate(t *testing.T) {
	request := &AddRequest{Description: "Descriptio2", Interval: 2, IntervalUnit: Minutes, Enabled: true, OwnerTeam: og.OwnerTeam{Name: "Sales"}}
	err := request.Validate()
	assert.Equal(t, err.Error(), errors.New("Invalid request. Name cannot be empty. ").Error())

	request = &AddRequest{Name: "NewSDK", Description: "Descriptio2", Interval: 2, IntervalUnit: Minutes, Enabled: true}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Invalid request. Owner team cannot be empty. ").Error())

	request = &AddRequest{Name: "NewSDK", Description: "Descriptio2", Interval: 0, IntervalUnit: Minutes, Enabled: true, OwnerTeam: og.OwnerTeam{Name: "Sales"}}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Invalid request. Interval cannot be smaller than 1. ").Error())

	request = &AddRequest{Name: "NewSDK", Description: "Descriptio2", Interval: 10, Enabled: true, OwnerTeam: og.OwnerTeam{Name: "Sales"}}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Invalid request. IntervalUnit cannot be empty. ").Error())

	request = &AddRequest{Name: "NewSDK", Description: "Descriptio2", Interval: 10, IntervalUnit: Minutes, Enabled: true, OwnerTeam: og.OwnerTeam{}}
	err = request.Validate()
	assert.Equal(t, err.Error(), errors.New("Invalid request. Owner team cannot be empty. ").Error())
}

func TestGetRequest_Validate(t *testing.T) {
	getRequest := &getRequest{}
	err := getRequest.Validate()

	assert.Equal(t, err.Error(), errors.New("HeartbeatName cannot be empty").Error())
}
