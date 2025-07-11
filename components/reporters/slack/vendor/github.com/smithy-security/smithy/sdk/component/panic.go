package component

import (
	"context"
	"runtime/debug"

	sdklogger "github.com/smithy-security/smithy/sdk/logger"
)

type (
	// PanicHandler defines a generic contract for handling panics following the recover semantics.
	PanicHandler interface {
		// HandlePanic handles a panic and returns an optional error with a signal on whether it should be
		// fatal or not.
		HandlePanic(ctx context.Context, err any) (error, bool)
	}

	defaultPanicHandler struct{}
)

// NewDefaultPanicHandler returns a new default panic handler.
func NewDefaultPanicHandler() (*defaultPanicHandler, error) {
	return &defaultPanicHandler{}, nil
}

// HandlePanic logs a panic and tells the runner to exit from the application.
func (dph *defaultPanicHandler) HandlePanic(ctx context.Context, err any) (error, bool) {
	logger := sdklogger.LoggerFromContext(ctx)
	if err != nil {
		e, ok := err.(error)
		if !ok {
			return nil, true
		}
		logger.With(
			sdklogger.LogKeyPanicStackTrace, string(debug.Stack()),
		).Error("received a panic. Check the stacktrace for more information. Exiting.")
		return e, true
	}
	return nil, false
}
