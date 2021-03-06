package u2utils

import (
	"github.com/pkg/errors"
	"github.com/ztrue/tracerr"
	"go.uber.org/multierr"
)

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func RecoverError(e any, err error) error {
	if e != nil {
		var err1 error
		if err0, ok := e.(error); ok {
			err1 = err0
		} else {
			err1 = errors.Errorf("error: %+v", e)
		}
		if err != nil {
			return multierr.Combine(err, err1)
		} else {
			return err1
		}
	} else {
		return err
	}
}

func PrintError(err error) {
	tracerr.PrintSourceColor(err, 1)
}
