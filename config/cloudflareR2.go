package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/rs/zerolog/log"
)

func (cfg Config) LoadAWSConfig() aws.Config {
  conf, err := awsConfig.LoadDefaultConfig(context.TODO(), 
      awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
          cfg.R2.ApiKey, cfg.R2.ApiSecret, "",
      )), awsConfig.WithRegion("auto"),)

	if err != nil {
		log.Fatal().Msgf("Error loading AWS credentials,%v", err)
	}

	log.Info().Msg("Succesfully loaded AWS credentials")

	return conf
}