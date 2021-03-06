package gox_aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/devlibx/gox-base"
	errors2 "github.com/devlibx/gox-base/errors"
	"github.com/devlibx/gox-base/util"
)

type awsSessionContext struct {
	session *session.Session
	gox.CrossFunction
}

func (a *awsSessionContext) GetSession() *session.Session {
	return a.session
}

func NewAwsContext(cf gox.CrossFunction, config Config) (ctx AwsContext, err error) {
	_ctx := &awsSessionContext{CrossFunction: cf}

	// Set default region
	if util.IsStringEmpty(config.Region) {
		config.Region = "ap-south-1"
	}

	// Setup AWS session
	if len(config.Endpoint) > 0 {
		_ctx.session, err = session.NewSession(
			&aws.Config{
				Region:   aws.String(config.Region),
				Endpoint: aws.String(config.Endpoint),
			})
	} else {
		_ctx.session, err = session.NewSession(
			&aws.Config{
				Region: aws.String(config.Region),
			})
	}
	if err != nil {
		return nil, errors2.Wrap(err, "failed to create aws session: endpoint=%s, region=%s", config.Endpoint, config.Region)
	}

	return _ctx, err
}
