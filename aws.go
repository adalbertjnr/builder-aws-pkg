package main

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ecr"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/route53"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

var (
	ErrNoProfile = errors.New("profile is nil or not valid")
	ErrNoRegion  = errors.New("region is nil or not valid")
)

type ProfileRegion struct {
	profile, region string
}

type Aws struct {
	awsCfg    aws.Config
	pr        ProfileRegion
	r53Client *route53.Client
	s3Client  *s3.Client
	iamClient *iam.Client
	ecrClient *ecr.Client
	ssmClient *ssm.Client
}

type AwsBuilder struct {
	Aws
}

func (b *AwsBuilder) MustAWSConfig(profile, region string) *AwsBuilder {
	b.pr.profile = profile
	b.pr.region = region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile), config.WithDefaultRegion(region))
	if err != nil {
		panic(err)
	}
	b.awsCfg = cfg
	return b
}

func NewAwsBuilder() *AwsBuilder {
	return &AwsBuilder{
		Aws{},
	}
}

func (b *AwsBuilder) WithR53() *AwsBuilder {
	r53 := route53.NewFromConfig(b.awsCfg)
	b.r53Client = r53
	return b
}

func (b *AwsBuilder) WithS3() *AwsBuilder {
	s3 := s3.NewFromConfig(b.awsCfg)
	b.s3Client = s3
	return b
}

func (b *AwsBuilder) WithIam() *AwsBuilder {
	iam := iam.NewFromConfig(b.awsCfg)
	b.iamClient = iam
	return b
}

func (b *AwsBuilder) WithEcr() *AwsBuilder {
	ecr := ecr.NewFromConfig(b.awsCfg)
	b.ecrClient = ecr
	return b
}

func (b *AwsBuilder) WithSSM() *AwsBuilder {
	ssm := ssm.NewFromConfig(b.awsCfg)
	b.ssmClient = ssm
	return b
}

func (b *AwsBuilder) Build() (*Aws, error) {

	if b.pr.profile == "" {
		return nil, ErrNoProfile
	}

	if b.pr.region == "" {
		return nil, ErrNoRegion
	}

	return &b.Aws, nil
}
