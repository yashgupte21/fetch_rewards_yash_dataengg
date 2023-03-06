package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	_ "github.com/lib/pq"
)

type UserLogins struct {
	User_id     string
	App_version int
	Device_type string
	IP          string
	Locale      string
	Device_id   string
}

func main() {

	// database connectoin
	db_connection := "user=postgres dbname=postgres password=postgres host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", db_connection)
	if err != nil {
		panic(err)
	}

	// check db connection
	err = db.Ping()
	if err != nil {
		fmt.Println("Couldn't Connect to database")
		panic(err)
	}

	// localstack endpoint set up
	err1 := os.Setenv("LOCALSTACK_ENDPOINT", "http://localhost:4566")
	if err1 != nil {
		panic(err1)
	}

	sess, err := CreateSession("us-east-1")
	if err != nil {
		panic(err)
	}

	sqsSvc := sqs.New(sess)

	for {
		result, err := sqsSvc.ReceiveMessage( //fetch localstack login_queue
			&sqs.ReceiveMessageInput{
				QueueUrl:            aws.String("http://localhost:4566/000000000000/login-queue"),
				MaxNumberOfMessages: aws.Int64(1),
				WaitTimeSeconds:     aws.Int64(3),
			},
		)

		if err != nil {
			fmt.Printf("failed to receive message with error %v", err)
			continue
		}

		if len(result.Messages) == 0 {
			break
		}
		if (*result.Messages[0].Body) == "{\"foo\": \"oops_wrong_msg_type\", \"bar\": \"123\"}" {
			continue
		}

		var user_logins UserLogins
		err1 := json.Unmarshal([]byte(*result.Messages[0].Body), &user_logins)
		if err != nil {
			panic(err1)
		}

		// ip address masking
		octets := strings.Split(user_logins.IP, ".")

		ip_first_part, _ := strconv.Atoi(octets[0])
		ip_second_part, _ := strconv.Atoi(octets[1])
		ip_third_part, _ := strconv.Atoi(octets[2])
		ip_fourth_part, _ := strconv.Atoi(octets[3])
		a := strconv.FormatInt(int64(ip_first_part), 2)
		b := strconv.FormatInt(int64(ip_second_part), 2)
		c := strconv.FormatInt(int64(ip_third_part), 2)
		d := strconv.FormatInt(int64(ip_fourth_part), 2)
		sum := 0
		for index, item := range d {
			sum += int(item) * (int(math.Pow(2, 7-float64(index))))
		}
		for index, item := range c {
			sum += int(item) * (int(math.Pow(2, 8+7-float64(index))))
		}
		for index, item := range b {
			sum += int(item) * (int(math.Pow(2, 16+7-float64(index))))
		}
		for index, item := range a {
			sum += int(item) * (int(math.Pow(2, 24+7-float64(index))))
		}

		// device id masking

		tmp_str := ""
		for i := 0; i < 11; i++ {
			switch user_logins.Device_id[i] {
			case 48:
				tmp_str = tmp_str + "P"
			case 49:
				tmp_str = tmp_str + "O"
			case 50:
				tmp_str = tmp_str + "S"
			case 51:
				tmp_str = tmp_str + "T"
			case 52:
				tmp_str = tmp_str + "G"
			case 53:
				tmp_str = tmp_str + "R"
			case 54:
				tmp_str = tmp_str + "E"
			case 55:
				tmp_str = tmp_str + "Z"
			case 56:
				tmp_str = tmp_str + "Q"
			case 57:
				tmp_str = tmp_str + "L"
			}
		}

		// load data into postgres docker container
		user_id := user_logins.User_id
		device_type := user_logins.Device_type
		masked_ip := sum
		masked_device_id := tmp_str
		locale := user_logins.Locale
		app_version := user_logins.App_version

		create_date := time.Now()

		fmt.Printf("%d %s \n", app_version, locale)

		psql_statement := "INSERT INTO user_logins(user_id, device_type, masked_ip, masked_device_id, locale, app_version, create_date) VALUES($1, $2, $3, $4, $5, $6, $7);"
		_, err = db.Exec(psql_statement, user_id, device_type, masked_ip, masked_device_id, locale, app_version, create_date)

		if err != nil {
			panic(err)
		}

	}

}

// session setup
func CreateSession(region string) (*session.Session, error) {
	awsConfig := &aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials("AKID", "SECRET_KEY", "TOKEN"),
	}

	if localStackEndpoint := os.Getenv("LOCALSTACK_ENDPOINT"); localStackEndpoint != "" {
		awsConfig.Endpoint = aws.String(localStackEndpoint)
	}

	return session.NewSession(awsConfig)
}
