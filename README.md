# Localstack management with Serverless Framework

<h1 align="center">
    <img src="https://res.cloudinary.com/yugovtr/image/upload/v1620795932/1_CxoqbfJvD8PDBRcHVmirHQ-removebg-preview_qjjrk4.png" />

</h1>

This repository is a proof of concept (PoC) of using the `Serverless` framework to manage the infrastructure dependencies for `AWS` in a development environment with `Localstack`

### Dependencies
- [Serverless](https://www.serverless.com)
- [Docker](https://www.docker.com/get-started)
- [AWS Local CLI](https://github.com/localstack/awscli-local)

### Model
![model](https://res.cloudinary.com/yugovtr/image/upload/v1620795454/model_nuhc2i.jpg)

### Install
```bash
make setup
```

### Send message to SNS
```bash
make notify
```

## Metrics
### Read Dynamo
```bash
make scan
```

### Local Stack Logs
```bash
docker logs --tail 1000 -f awslocal
```

> Output
![image](https://user-images.githubusercontent.com/76954948/117922710-432b1680-b2c9-11eb-87d1-8935fa0558c6.png)
