package ezpq

import (
	"github.com/lib/pq"
)

type errWrap struct {
	errorMessage string
	err          error
}

func Err(err error) error {
	return newErrWrap(err, ErrorsPQ[err.(*pq.Error).Code])
}

func newErrWrap(err error, pqErr error) error {
	return &errWrap{
		errorMessage: err.Error(),
		err:          pqErr,
	}

}

func (e *errWrap) Error() string {
	return e.errorMessage
}

func (e *errWrap) Unwrap() error {
	return e.err
}
