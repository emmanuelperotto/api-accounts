package infra

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"os"
)

type MessageAttribute struct {
	Key         string
	DataType    string
	StringValue string
}

type Message struct {
	Value      string
	Topic      string
	Attributes []MessageAttribute
}

type publisher interface {
	Publish(Message) error
}

var (
	MessagingService publisher
)

func SetupMessagingService() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile:           os.Getenv("AWS_PROFILE"),
		SharedConfigState: session.SharedConfigEnable,
	}))

	MessagingService = SNS{client: sns.New(sess)}
}
