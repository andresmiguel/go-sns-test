package main

import (
	"bufio"
	"fmt"
	"go-sns-test/snsclient"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// usage:
// go run main.go
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("You need to have configured the AWS credentials")
	fmt.Println("Enter aws-cli profile:")
	scanner.Scan()
	profile := scanner.Text()
	fmt.Println("Enter AWS Region:")
	scanner.Scan()
	region := scanner.Text()
	fmt.Println("Enter AWS SNS Topic ARN:")
	scanner.Scan()
	topicArn := scanner.Text()

	// SDK will use to load credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profile,
		Config: aws.Config{
			Region: aws.String(region),
		},
	})

	if err != nil {
		fmt.Println("NewSession error:", err)
		return
	}

	client := snsclient.SNS{
		Client:   sns.New(sess),
		TopicArn: topicArn,
	}

	for {
		fmt.Println("Enter message to send or ':q' to quit")
		scanner.Scan()
		msg := scanner.Text()
		if msg == ":q" {
			break
		}
		result, err := client.Publish(msg)
		if err != nil {
			fmt.Println("Publish error:", err)
		} else {
			fmt.Println(result)
		}
	}
}
