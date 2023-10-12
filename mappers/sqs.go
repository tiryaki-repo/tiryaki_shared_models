package mapper

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/tiryaki-repo/tiryaki_shared_models/services/validator"
)

func ToEventList[T any](s []events.SQSMessage) ([]T, error) {
	list := make([]T, 0)

	for _, msg := range s {
		value, err := ToEvent[T](msg)
		if err != nil {
			return nil, err
		}
		list = append(list, value)
	}

	return list, nil
}

func ToEvent[T any](s events.SQSMessage) (T, error) {

	validator := validator.NewValidator()

	var message T
	var sqsBody SQSBody

	if err := json.Unmarshal([]byte(s.Body), &sqsBody); err != nil {
		return message, err
	}

	log.Printf("message body is :%s", sqsBody)

	if err := json.Unmarshal([]byte(sqsBody.Message), &message); err != nil {
		return message, err
	}

	if err := validator.Validate(message); err != nil {
		return message, err

	}
	return message, nil

}
