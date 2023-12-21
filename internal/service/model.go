package service

import "github.com/aws/aws-sdk-go/aws/endpoints"

const (
	primaryRegion   awsRegion = endpoints.ApNortheast1RegionID // Tokyo
	secondaryRegion awsRegion = endpoints.ApNortheast3RegionID // Osaka
)

type awsRegion string

func (r awsRegion) String() string {
	return string(r)
}
