package helper

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

func IsMongoDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}
