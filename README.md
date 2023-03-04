# fetch_rewards_yash_dataengg
Fetch Rewards Data Engineer Take Home - Yash Pradeep Gupte

## Table of Contents
* [Technologies ](#technologies)
* [Installation Guide ](#installation-guide)
* [Deployment ideas ](#deployment)
* [Production Ready Additonal Components ](#production)
* [Scaling ](#scaling)
* [Recovering PII ](#recover-pii)
* [Asseumptions ](#assumptions)



---
## Technologies
Project is created with:
* Go: go1.20.1 darwin/arm64 
* Docker: 20.10.23

--- 
## Installation Guide 

I have used GoLang to implement my solution 

To install golang on your system follow the steps provided below:

1. Go to https://go.dev/dl/ and download golang for your system in your root or HOME directory

2. Verify if golang has been installed. Open terminal and type the following command

``` 
go version 
```

3. Clone this repository in your root location 

4. Install dependencies for main.go by using following commands:

```
go mod init examples.com/fetch_rewards_yash_dataengg
```

```
go mod tidy
```

5. Install awscli-local on your system

```
pip install awscli-local
```

5. If you do not have docker install on your system, install docker :

https://docs.docker.com/get-docker/

6. After installing docker and docker-compose , open a new terminal and pull the provided docker images and run these images (container will begin to run on your docker desktop)

Terminal a. 

```
docker pull fetchdocker/data-takehome-postgres
```

```
docker run --rm -it -p 5432:5432 fetchdocker/data-takehome-postgres
```

Terminal b.

```
docker pull fetchdocker/data-takehome-localstack
```

```
docker run --rm -it -p 4566:4566 fetchdocker/data-takehome-localstack
```

7. Next step is to run main.go program in a seperate terminal 

```
go run main.go
```

Executing this program will Extract AWS SQS Queue Messages from Localstack, convert them into JSON format, Transform the data as required and Load it on the Postgres database docker container.

--- 

## Deployment ideas


--- 
## Production Ready Additonal Components

--- 

## Scaling

--- 

## Recovering PII

--- 

## Assumptions

