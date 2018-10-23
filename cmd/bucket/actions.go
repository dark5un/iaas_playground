package main

import (
	"errors"
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/google/go-cloud/blob"
	"github.com/google/go-cloud/blob/gcsblob"
	"github.com/google/go-cloud/blob/s3blob"
	"github.com/google/go-cloud/gcp"
	"github.com/urfave/cli"
)

// Create is a cli action handle
func Create(c *cli.Context) error {
	bucket.WriteAll(context.Background(), "any_key_would_do_15", []byte("Some data"), nil)
	return nil
}

// Delete is a cli action handle
func Delete(c *cli.Context) error {
	return nil
}

// List is a cli action handle
func List(c *cli.Context) error {
	return nil
}

var bucket *blob.Bucket

// Before makes sure we use the correct IAAS
func Before(c *cli.Context) error {
	var (
		err error
	)
	if BucketName == "" {
		return errors.New("name should be set")		
	}
	switch IAAS {
	case "aws":
		bucket, err = setupAWS(context.Background(), BucketName)
		if err != nil {
			return err
		}
	case "gcp":
		bucket, err = setupGCP(context.Background(), BucketName)
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupported IAAS")		
	}
	return nil
}

// setupGCP creates a connection to Google Cloud Storage (GCS).
func setupGCP(ctx context.Context, bucket string) (*blob.Bucket, error) {
	// DefaultCredentials assumes a user has logged in with gcloud.
	// See here for more information:
	// https://cloud.google.com/docs/authentication/getting-started
	creds, err := gcp.DefaultCredentials(ctx)
	if err != nil {
		return nil, err
	}
	c, err := gcp.NewHTTPClient(gcp.DefaultTransport(), gcp.CredentialsTokenSource(creds))
	if err != nil {
		return nil, err
	}
	return gcsblob.OpenBucket(ctx, bucket, c)
}

// setupAWS creates a connection to Simple Cloud Storage Service (S3).
func setupAWS(ctx context.Context, bucket string) (*blob.Bucket, error) {
	creds := credentials.NewEnvCredentials()
	c := &aws.Config{
		// Either hard-code the region or use AWS_REGION.
		Region: aws.String("eu-east-2"),
		// credentials.NewEnvCredentials assumes two environment variables are
		// present:
		// 1. AWS_ACCESS_KEY_ID, and
		// 2. AWS_SECRET_ACCESS_KEY.
		Credentials: creds,
	}
	s := session.Must(session.NewSession(c))
	return s3blob.OpenBucket(ctx, s, bucket)
}
