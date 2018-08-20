package handler

import (
	"errors"

	"github.com/truebluejason/new-outlook/internal/data"
	"github.com/truebluejason/new-outlook/internal/helper"
	"github.com/truebluejason/new-outlook/internal/resp"

	"github.com/arienmalec/alexa-go"
)

func HandleGeneral(request alexa.Request) (alexa.Response, error) {

	var (
		desc string
		ssml string
		response alexa.Response
		dbResponse *data.DBResponse
	)
	// Format Data to Query: Description, Useful/NotUseful, True Fact, Inspiration, Initiative
	desc = helper.DescribeEmotion("displeasure", []string{
		"hopelessness",
		"lack of energy",
		"critical thoughts",
	})

	// Create DynamoDB Connection
	ses, err := data.NewSession()
	if err != nil {
		return response, errors.New("Error creating AWS session: " + err.Error())
	}
	ddb := data.NewDynamoDB(ses)

	// Query Data
	dbResponse, err = data.GetContents(ddb, "general")
	if err != nil {
		return response, errors.New("Error fetching data from DynamoDB: " + err.Error())
	}
	dbResponse.Description = desc

	// Return Alexa Response
	ssml = resp.FormatSSML(*dbResponse)
	response = resp.SSMLResponse(ssml)

	return response, nil
}