package resp

import (
	"github.com/arienmalec/alexa-go"
)

type ComResponse struct {
	Useful string

}

func SSMLResponse(content string) alexa.Response {
	var text string
	text = "Hello, " + content + ". That's it!"

	r := alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			OutputSpeech: &alexa.Payload{
				Type: "SSML",
				Text: text,
			},
			ShouldEndSession: true,
		},
	}
	return r
}