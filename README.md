# aws package

```
go get -u github.com/souzagmu/svc-aws-pkg
```

```package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	apkg "github.com/souzagmu/svc-aws-pkg"
)

func main() {

	var (
		ctx     = context.TODO()
		profile = "default"
		region  = "us-east-1"
	)

	c, err := apkg.NewAwsBuilder().MustAWSConfig(profile, region).WithEcr().WithS3().WithIam().Build()
	if err != nil {
		log.Fatal(err)
	}

	s3Input := &s3.ListBucketsInput{}

	s3Otp, err := c.S3Client.ListBuckets(ctx, s3Input)
	if err != nil {
		log.Fatal(err)
	}

	for _, bucket := range s3Otp.Buckets {
		fmt.Printf("bucket %s\n", *bucket.Name)
	}

}```
