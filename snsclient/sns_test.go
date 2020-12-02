package snsclient

import (
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sns/snsiface"
)

type mockedPublish struct {
	snsiface.SNSAPI
	Resp sns.PublishOutput
	Err  error
}

func (m mockedPublish) Publish(*sns.PublishInput) (*sns.PublishOutput, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return &m.Resp, nil
}

func TestSNS_Publish(t *testing.T) {
	cases := []struct {
		Resp     sns.PublishOutput
		Expected PublishOutput
		Err      error
	}{
		// case 1, success
		{
			Resp:     sns.PublishOutput{MessageId: aws.String("mid-01")},
			Expected: PublishOutput{MessageId: "mid-01"},
		},
		// case 2, error
		{
			Resp:     sns.PublishOutput{},
			Expected: PublishOutput{MessageId: ""},
			Err:      errors.New("ups... something went wrong"),
		},
	}
	for i, c := range cases {
		client := SNS{
			Client:   mockedPublish{Resp: c.Resp, Err: c.Err},
			TopicArn: fmt.Sprintf("mockTopic_%d", i),
		}
		resp, err := client.Publish("Some msg")
		if err != c.Err {
			t.Fatalf("%d, expected error [%v], got [%v]", i, err, c.Err)
		}
		if err == nil && resp.MessageId != c.Expected.MessageId {
			t.Fatalf("%d, expected %s, got %s", i, c.Expected.MessageId, resp.MessageId)
		}
	}
}
