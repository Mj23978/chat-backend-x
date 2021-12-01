package utils

import (
	errs "github.com/pkg/errors"
	"log"
)

func ErrorWrap(err error, wrap string) error {
	if err != nil {
		return errs.Wrap(err, wrap)
	}
	return nil
}

func ErrorFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
