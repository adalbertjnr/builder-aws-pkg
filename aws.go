package apkg

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
	ErrNoProfile = errors.New("empty profile")
	ErrNoRegion  = errors.New("empty region")
)

type ProfileRegion struct {
	Profile, Region string
}

type Aws struct {
	AwsCfg    aws.Config
	Pr        ProfileRegion
	R53Client *route53.Client
	S3Client  *s3.Client
	IamClient *iam.Client
	EcrClient *ecr.Client
	SsmClient *ssm.Client
}

type AwsBuilder struct {
	Aws
}

func (b *AwsBuilder) MustAWSConfig(profile, region string) *AwsBuilder {
	b.Pr.Profile = profile
	b.Pr.Region = region
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithSharedConfigProfile(profile), config.WithDefaultRegion(region))
	if err != nil {
		panic(err)
	}
	b.AwsCfg = cfg
	return b
}

func NewAwsBuilder() *AwsBuilder {
	return &AwsBuilder{
		Aws{},
	}
}

func (b *AwsBuilder) WithR53() *AwsBuilder {
	r53 := route53.NewFromConfig(b.AwsCfg)
	b.R53Client = r53
	return b
}

func (b *AwsBuilder) WithS3() *AwsBuilder {
	s3 := s3.NewFromConfig(b.AwsCfg)
	b.S3Client = s3
	return b
}

func (b *AwsBuilder) WithIam() *AwsBuilder {
	iam := iam.NewFromConfig(b.AwsCfg)
	b.IamClient = iam
	return b
}

func (b *AwsBuilder) WithEcr() *AwsBuilder {
	ecr := ecr.NewFromConfig(b.AwsCfg)
	b.EcrClient = ecr
	return b
}

func (b *AwsBuilder) WithSSM() *AwsBuilder {
	ssm := ssm.NewFromConfig(b.AwsCfg)
	b.SsmClient = ssm
	return b
}

func (b *AwsBuilder) Build() (*Aws, error) {

	if b.Pr.Profile == "" {
		return nil, ErrNoProfile
	}

	if b.Pr.Region == "" {
		return nil, ErrNoRegion
	}

	return &b.Aws, nil
}
