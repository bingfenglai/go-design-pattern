package builder_dp

import "errors"

type buildError struct {
	error
}

func (buildError *buildError) newError(msg string) *buildError {
	if buildError.error == nil {
		buildError.error = errors.New(msg)
	} else {
		s := buildError.error.Error()
		buildError.error = errors.New(s + ";" + msg)

	}
	return buildError
}

func (buildError *buildError) Error() string {
	if buildError.error == nil {
		return ""
	}
	return buildError.error.Error()
}
