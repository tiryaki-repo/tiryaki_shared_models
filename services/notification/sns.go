package notification

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
)

type snsClient interface {
	Publish(ctx context.Context,
		params *sns.PublishInput,
		optFns ...func(*sns.Options)) (*sns.PublishOutput, error)
}

type SNS struct {
	topicArn string
	client   snsClient
}

func (notifier *SNS) Publish(ctx context.Context, subject string, message interface{}) error {

	input, err := json.Marshal(message)

	if err != nil {
		return err
	}

	log.Printf("Publishing message : %s & topicArn : %s", input, notifier.topicArn)

	publishInput := &sns.PublishInput{
		Message:  aws.String(string(input)),
		Subject:  aws.String(subject),
		TopicArn: aws.String(notifier.topicArn),
	}

	output, err := notifier.client.Publish(ctx, publishInput)

	log.Printf("Output is : %s", *output.MessageId)

	return err
}

func NewSNS(conf aws.Config, topicArn string) *SNS {
	return &SNS{
		topicArn: topicArn,
		client:   sns.NewFromConfig(conf),
	}
}
