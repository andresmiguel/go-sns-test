package snsclient

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type (
	SNS struct {
		Client   snsiface.SNSAPI
		TopicArn string
	}
	PublishOutput struct {
		MessageId string
	}
)

func (s *SNS) Publish(message string) (PublishOutput, error) {
	input := &sns.PublishInput{
		Message:  aws.String(message),
		TopicArn: aws.String(s.TopicArn),
	}
	output, err := s.Client.Publish(input)
	if err != nil {
		return PublishOutput{}, err
	}
	return PublishOutput{MessageId: *output.MessageId}, nil
}
