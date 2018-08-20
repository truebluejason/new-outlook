package handler

import (
	"github.com/arienmalec/alexa-go"
)

func HandleCraving(request alexa.Request) (alexa.Response, error) {
	return alexa.NewSimpleResponse("Response for Craving", "Craving is..."), nil
}