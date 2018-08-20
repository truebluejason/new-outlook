package handler

import (
	"github.com/arienmalec/alexa-go"
)

func HandleBitterness(request alexa.Request) (alexa.Response, error) {
	return alexa.NewSimpleResponse("Response for Bitterness", "Bitterness is..."), nil
}