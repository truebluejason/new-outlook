package handler

import (
	"github.com/arienmalec/alexa-go"
)

func HandleLaziness(request alexa.Request) (alexa.Response, error) {
	return alexa.NewSimpleResponse("Response for Laziness", "Laziness is..."), nil
}