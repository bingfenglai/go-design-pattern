package builder_dp

import "errors"

type buildErrorWrapper struct {
	error
}

func (buildError *buildErrorWrapper) newError(msg string) *buildErrorWrapper {
	if buildError.error == nil {
		buildError.error = errors.New(msg)
	} else {
		s := buildError.error.Error()
		buildError.error = errors.New(s + ";" + msg)

	}
	return buildError
}

func (buildError *buildErrorWrapper) Error() string {
	if buildError.error == nil {
		return ""
	}
	return buildError.error.Error()
}
