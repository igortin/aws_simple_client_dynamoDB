package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// New initialize client
func New(profileaName string, awsRegion string) (*Client, error) {
	sess, err := session.NewSessionWithOptions(session.Options{
		Config:  aws.Config{Region: aws.String(awsRegion)},
		Profile: profileaName,
	})

	if err != nil {
		return nil, err
	}

	return &Client{
		dbsvc: dynamodb.New(sess),
	}, nil
}

// Client struct
type Client struct {
	dbsvc *dynamodb.DynamoDB
}

// PutItem Method
func (client *Client) PutItem(e *Event) (*dynamodb.PutItemOutput, error) {

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"EventID": {
				S: aws.String(e.EventID),
			},
			"EventInfo": {
				S: aws.String(e.EventInfo),
			},
			"Body": {
				S: aws.String(string(e.Body)),
			},
		},
		ReturnConsumedCapacity: aws.String("TOTAL"),
		TableName:              aws.String("ApiCall"),
	}
	res, err := client.dbsvc.PutItem(input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
