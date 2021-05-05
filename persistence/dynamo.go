package persistence

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/vitorhrmiranda/go-serverless/entity"
	"github.com/vitorhrmiranda/go-serverless/values"
)

type DynamoDB struct {
	c *dynamodb.DynamoDB
}

func NewDynamoDB() (DynamoDB, error) {
	var sess *session.Session
	var err error

	impl := DynamoDB{}

	sess, err = session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials("foo", "var", ""),
		Region:      aws.String(values.DEFAULT_REGION),
		Endpoint:    aws.String(values.DYNAMODB_AWS_ENDPOINT),
	})

	if err != nil {
		return impl, err
	}

	impl.c = dynamodb.New(sess)

	return impl, nil
}

func (db *DynamoDB) Create(usr *entity.User) (err error) {
	item, err := dynamodbattribute.MarshalMap(usr)

	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		TableName: aws.String(values.TABLE_NAME),
		Item:      item,
	}

	_, err = db.c.PutItem(input)

	return err
}
