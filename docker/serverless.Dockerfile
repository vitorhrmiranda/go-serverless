FROM amaysim/serverless:2.40.0

WORKDIR /app/

RUN yarn global add serverless-localstack

RUN pip --no-cache-dir install awscli-local 
