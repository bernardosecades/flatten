package http

import "errors"

type Request struct {
	Input []interface{}
}

func (r Request) Validate() error {
	if len(r.Input) == 0 {
		return errors.New("invalid format in input parameter")
	}
	return nil
}
