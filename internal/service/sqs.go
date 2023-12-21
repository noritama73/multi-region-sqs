package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type multiRegionSQS struct {
	queueName       string
	primaryClient   *sqsClient
	secondaryClient *sqsClient
}

func NewMultiRegionSQS(queueName string) MultiRegionSQS {
	return &multiRegionSQS{
		queueName:       queueName,
		primaryClient:   newSqsClient(primaryRegion),
		secondaryClient: newSqsClient(secondaryRegion),
	}
}

type sqsClient struct {
	region        awsRegion
	queueURLCache *string
	svc           *sqs.SQS
}

func newSqsClient(region awsRegion) *sqsClient {
	sess := session.Must(session.NewSession())

	svc := sqs.New(sess, aws.NewConfig().WithRegion(region.String()))
	return &sqsClient{
		region: region,
		svc:    svc,
	}
}

func (c *sqsClient) getQueueURL(queueName string) (string, error) {
	if c.queueURLCache != nil {
		return *c.queueURLCache, nil
	}

	getQueueUrlOutput, err := c.svc.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queueName,
	})
	if err != nil {
		return "", err
	}

	c.queueURLCache = getQueueUrlOutput.QueueUrl
	return *getQueueUrlOutput.QueueUrl, nil
}

func (c *sqsClient) sendMessage(queueName string, message string) error {
	queueURL, err := c.getQueueURL(queueName)
	if err != nil {
		return err
	}

	_, err = c.svc.SendMessage(&sqs.SendMessageInput{
		MessageBody: &message,
		QueueUrl:    &queueURL,
	})
	if err != nil {
		return err
	}

	return nil
}

func (c *sqsClient) receiveMessage(queueName string) ([]*sqs.Message, error) {
	queueURL, err := c.getQueueURL(queueName)
	if err != nil {
		return nil, err
	}

	receiveMessageOutput, err := c.svc.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl: &queueURL,
	})
	if err != nil {
		return nil, err
	}

	return receiveMessageOutput.Messages, nil
}
