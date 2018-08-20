package handler

import (
	"github.com/arienmalec/alexa-go"
)

func HandleDoubt(request alexa.Request) (alexa.Response, error) {
	return alexa.NewSimpleResponse("Response for Doubt", "Doubt is..."), nil
}