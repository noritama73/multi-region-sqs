package service

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

const (
	primaryRegion   awsRegion = endpoints.ApNortheast1RegionID // Tokyo
	secondaryRegion awsRegion = endpoints.ApNortheast3RegionID // Osaka
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

func newSqsClient(region awsRegion) *sqsClient {
	sess := session.Must(session.NewSession())

	svc := sqs.New(sess, aws.NewConfig().WithRegion(region.String()))
	return &sqsClient{
		region: region,
		svc:    svc,
	}
}

type multiRegionSQS struct {
	queueName       string
	primaryClient   *sqsClient
	secondaryClient *sqsClient
}

func NewMultiRegionSQS(queueName string) *multiRegionSQS {
	return &multiRegionSQS{
		queueName:       queueName,
		primaryClient:   newSqsClient(primaryRegion),
		secondaryClient: newSqsClient(secondaryRegion),
	}
}
