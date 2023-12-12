package service

import (
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	primaryRegion   awsRegion = "ap-northeast-1" // Tokyo
	secondaryRegion awsRegion = "ap-northeast-3" // Osaka
)

type awsRegion string

func (r awsRegion) String() string {
	return string(r)
}

type sqsClient struct {
	region        awsRegion
	queueURLCache *string
	svc           *sqs.SQS
}

func (c *sqsClient) getQueueURL(region awsRegion, queueName string) (string, error) {
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

type multiRegionSQS struct {
	queueName       string
	primaryClient   sqsClient
	secondaryClient sqsClient
}
