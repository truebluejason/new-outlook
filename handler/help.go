package handler

import (
	"fmt"
	"github.com/arienmalec/alexa-go"
)

func HandleHelp() alexa.Response {
	return alexa.NewSimpleResponse(
		"Help", 
		fmt.Sprint(
			"Tell new outlook of your current emotional state to hear about what you can do in this very moment. ",
			"The states can be anxiety, bitterness, craving, doubt, and laziness. ",
			"For example, you can say 'Tell New Outlook that I feel lazy'. ",
			"If you're not sure, simply say 'Tell New Outlook that I feel down'.",
		),
	)
}
