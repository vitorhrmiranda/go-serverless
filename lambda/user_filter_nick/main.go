package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/spf13/viper"
	"github.com/vitorhrmiranda/go-serverless/entity"
	"github.com/vitorhrmiranda/go-serverless/persistence"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(sqsEvent events.SQSEvent) (err error) {

	fmt.Println("RUNNING - user_filter_nick")

	var dynamo = persistence.DynamoDB{}
	if dynamo, err = persistence.NewDynamoDB(); err != nil {
		return err
	}

	for _, message := range sqsEvent.Records {

		usr := new(entity.User)

		if err := json.Unmarshal([]byte(message.Body), usr); err != nil {
			return err
		}

		usr.Nick = strings.ToLower(usr.Nick)

		re := regexp.MustCompile(`\s(\D+)\s`)
		usr.Nick = re.ReplaceAllString(usr.Nick, " ")

		fmt.Printf("USER: %#v\n", usr)

		return dynamo.Create(usr)
	}

	return err
}

func main() {
	viper.AutomaticEnv()
	lambda.Start(handler)
}
