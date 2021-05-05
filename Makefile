setup:
	# needs:
	# npm i -g serverless
	# npm i -g serverless-localstack

	mkdir -p bin
	go build -o bin/user_filter_nick lambda/user_filter_nick/main.go
	go build -o bin/user_gen_password lambda/user_gen_password/main.go
	sls deploy

notify:
	awslocal sqs send-message \
		--queue-url http://awslocal:4566/000000000000/consumer \
		--message-body file://input.json

scan:
	awslocal dynamodb scan --table-name=users
