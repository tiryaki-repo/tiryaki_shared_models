package secretmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type client interface {
	GetSecretValue(ctx context.Context, params *secretsmanager.GetSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.GetSecretValueOutput, error)
}

type SecretManager struct {
	ca client
}

func (c *SecretManager) getSecret(ctx context.Context, sn string) (*string, error) {

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(sn),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := c.ca.GetSecretValue(ctx, input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		return nil, err
	}

	return result.SecretString, nil
}

func NewSecretManager(cfg aws.Config) *SecretManager {
	return &SecretManager{
		ca: secretsmanager.NewFromConfig(cfg),
	}
}
