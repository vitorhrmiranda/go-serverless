package main

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/viper"
	"github.com/vitorhrmiranda/go-serverless/entity"
	"github.com/vitorhrmiranda/go-serverless/persistence"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bxcodec/faker/v3"
)

func handler(sqsEvent events.SQSEvent) (err error) {

	fmt.Println("RUNNING - user_get_password")

	var dynamo = persistence.DynamoDB{}
	if dynamo, err = persistence.NewDynamoDB(); err != nil {
		return err
	}

	for _, message := range sqsEvent.Records {

		usr := new(entity.User)

		fmt.Printf("MESSAGE BODY: %s\n", message.Body)

		if err := json.Unmarshal([]byte(message.Body), usr); err != nil {
			return fmt.Errorf("Unmarshal: %s", err)
		}

		if len(usr.Password) == 0 {
			usr.ID += "-pwd"
			usr.Password = faker.Password()
		}

		fmt.Printf("USER: %#v\n", usr)

		return dynamo.Create(usr)
	}

	return err
}

func main() {
	viper.AutomaticEnv()
	lambda.Start(handler)
}
