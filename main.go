package main

import (
	"fmt"

	"github.com/arienmalec/alexa-go"
	"github.com/aws/aws-lambda-go/lambda"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
/*
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	body, err := json.Marshal(map[string]interface{}{
		"message": "Go Serverless v1.0! Your function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}
*/

func DispatchIntents(request alexa.Request) alexa.Response {
	var (
		response alexa.Response
		//err error
	)
	/*
	switch request.Body.Intent.Name {
	case "AnxietyIntent":
		response, err = handler.HandleAnxiety(request)
	case "BitternessIntent":
		response, err = handler.HandleBitterness(request)
	case "CravingIntent":
		response, err = handler.HandleCraving(request)
	case "DoubtIntent":
		response, err = handler.HandleDoubt(request)
	case "GeneralIntent":
		response, err = handler.HandleGeneral(request)
	case "LazinessIntent":
		response, err = handler.HandleLaziness(request)
	case alexa.HelpIntent:
		response = handler.HandleHelp()
	}

	if err != nil {
		fmt.Errorf(err.Error())
	}
	*/
	response = HandleHelp()

	return response
}

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

func Handler(request alexa.Request) (alexa.Response, error) {
	return DispatchIntents(request), nil
}

func main() {
	lambda.Start(Handler)
}
