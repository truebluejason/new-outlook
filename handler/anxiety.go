package handler

import (
	"github.com/arienmalec/alexa-go"
)

func HandleAnxiety(request alexa.Request) (alexa.Response, error) {
	return alexa.NewSimpleResponse("Response for Anxiety", "Anxiety is..."), nil
}