package impl

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/devlibx/gox-base/errors"
	"strings"
	"time"
)

type ParameterInfo struct {
	Key   string
	Value string
}

func ReadSsmFromPath(path string, session *session.Session) ([]ParameterInfo, error) {
	svc := ssm.New(session)
	result := make([]ParameterInfo, 0)

	var out *ssm.GetParametersByPathOutput
	var err error
	count := 0
	for {
		var nextToken *string = nil
		if out != nil {
			nextToken = out.NextToken
			if out.NextToken == nil {
				break
			}
		}
		out, err = svc.GetParametersByPath(&ssm.GetParametersByPathInput{
			MaxResults:     aws.Int64(10),
			Path:           aws.String(path),
			Recursive:      aws.Bool(true),
			WithDecryption: aws.Bool(true),
			NextToken:      nextToken,
		})
		if err != nil {
			return nil, errors.Wrap(err, "failed to read ssm var from path")
		}

		for _, v := range out.Parameters {
			name := strings.ReplaceAll(*v.Name, path, "")
			p := ParameterInfo{
				Key:   name,
				Value: *v.Value,
			}
			result = append(result, p)
		}

		if out.Parameters == nil || len(out.Parameters) == 0 {
			break
		}
		time.Sleep(10 * time.Millisecond)

		count++
		if count > 50 {
			break
		}
	}

	return result, nil
}
