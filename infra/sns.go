package infra

import (
	"accounts/entities"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
	"os"
	"strconv"
)

type SNS struct {
	client *sns.SNS
}

func (s SNS) Publish(message Message) error {
	snsMessageAttributes := map[string]*sns.MessageAttributeValue{}
	for _, attribute := range message.Attributes {
		snsMessageAttributes[attribute.Key] = &sns.MessageAttributeValue{
			BinaryValue: nil,
			DataType:    aws.String(attribute.DataType),
			StringValue: aws.String(attribute.StringValue),
		}
	}

	input := sns.PublishInput{
		Message:           aws.String(message.Value),
		MessageAttributes: snsMessageAttributes,
		TopicArn:          aws.String(message.Topic),
	}

	output, err := s.client.Publish(&input)

	if err != nil {
		log.Println("[SNS publish error]", err)
		return err
	}

	log.Println("Published event", output)

	return nil
}

func PublishAccountCreatedEvent(account entities.Account) error {
	err := MessagingService.Publish(Message{
		Value: "AccountCreated",
		Topic: os.Getenv("TOPIC_ARN"),
		Attributes: []MessageAttribute{
			{
				Key:         "ID",
				DataType:    "Number",
				StringValue: strconv.FormatInt(account.ID, 10),
			},
			{
				Key:         "Code",
				DataType:    "String",
				StringValue: account.Code,
			},
			{
				Key:         "Agency",
				DataType:    "String",
				StringValue: account.Agency,
			},
		},
	})

	return err
}