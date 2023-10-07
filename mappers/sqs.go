package mapper

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
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

	var message T

	if err := json.Unmarshal([]byte(s.Body), &message); err != nil {
		return message, err
	}

	return message, nil

}
