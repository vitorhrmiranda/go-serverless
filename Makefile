SERVERLESS_RUN = docker-compose run --rm serverless

setup:
	mkdir -p bin
	go build -o bin/user_filter_nick lambda/user_filter_nick/main.go
	go build -o bin/user_gen_password lambda/user_gen_password/main.go
	$(SERVERLESS_RUN) sls deploy

notify:
	$(SERVERLESS_RUN) awslocal sns publish \
		--topic-arn arn:aws:sns:us-east-1:000000000000:Publisher \
		--message file://input.json

scan:
	$(SERVERLESS_RUN) awslocal dynamodb scan --table-name=users
