# fetch_rewards_yash_dataengg
Fetch Rewards Data Engineer Take Home - Yash Pradeep Gupte

## Table of Contents
* [Technologies ](#technologies)
* [Installation Guide ](#installation-guide)
* [Deployment and Production ideas ](#deployment-and-production-ideas)
* [Scaling ](#scaling)
* [Recovering PII ](#recovering-pii)
* [Assumptions ](#assumptions)



---
## Technologies
Project is created with:
* Windows Machine
* Go: go1.20.1 
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

4. Install modules for main.go by using following commands:

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

## Deployment and Production ideas
We can deploy this application on cloud platform provider such as Amazon AWS. 
* Create a docker compose yaml file which will include docker run commands structure for both the Postgres and AWS SQS Localstack image
* Add  a new container in the application.yaml docker compose file. This will be a image repository from AWS ECR (docker registry of AWS)
* Configure ports for containers 
* Run docker-compose command to run and deploy all containers 
```
docker-compose -f  application.yml up
```
* The entire application will be deployed on AWS 

--- 

## Scaling
As this application scale with growing databse, we can perform verticle or horizontal scaling. For realtional database such a Postgres we can increase the processing power of a single server or cluster. In horizontal scaling, we will add more nodes and clusters which can help in read performance balancing taffic between nodes. We can also introduce load balancers in hortizontal scaling. Data sharding is another technique we can implement as and when the application scales. If the database grows drastically and we can not upgrade the hardware then we can use sharding techniques.

--- 

## Assumptions
 * PII such as device id and ip address are masked using strategies.
 * device_id format : follows a format of 2 digits-3digits-4 digits
 * ip : assuming ip v4 addresses
 * masked_device_id : converting integers into characters 
 * masked_ip : converting ip address to its decimal value. We get 4 octates and convert these octates into binary values and convert this whole binary of 32 bits into it's corresponding decimal value 


---

## Recovering PII
In this project I have masked device id and ip address, which can be recovered as follow:
* device_id : reverse the masking technique by converting characters back to their integer values and after insert '-' after the second and fifth character
* ip : convert masked_ip to ip by diving the decimal presentation of ip address by 2 until we get 32 bits






