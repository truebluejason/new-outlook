package data

import (
	"math/rand"

	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/expression"
    "github.com/aws/aws-sdk-go/aws/session"
)

type DBResponse struct {
	Description string `json:"description"`
	Use			string `json:"use"`
	Initiative	string `json:"initiative"`
}

type DBUseResponse struct {
	Value string `json:"value"`
}

func NewDynamoDB(session *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(session)
}

// Use SCAN and limit returned values to 7
func GetContents(ddb *dynamodb.DynamoDB, intent string) (*DBResponse, error) {
	var (
		dbResponse DBResponse
	)

	filter := expression.Name("category").Equal(expression.Value(intent))
	projection := expression.NamesList(expression.Name("value"))

	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
	if err != nil {
		return &dbResponse, err
	}

	var limit int64 = 7
	resultSlice := make([]DBUseResponse, limit)

	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
        FilterExpression:          expr.Filter(),
        Limit:					   &limit,
        ProjectionExpression:      expr.Projection(),
        TableName:                 aws.String("new-outlook-uses"),
	}

	result, err := ddb.Scan(params)
	if err != nil {
		return &dbResponse, err
	}

	for _, el := range result.Items {
		item := DBUseResponse{}
		err = dynamodbattribute.UnmarshalMap(el, &item)
		if err != nil {
			return &dbResponse, err
		}
		resultSlice = append(resultSlice, item)
	}

	randIndex := rand.Intn(len(resultSlice))
	dbResponse.Use = resultSlice[randIndex].Value

	return &dbResponse, nil
}