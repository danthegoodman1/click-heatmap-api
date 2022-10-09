package http_server

import (
	"context"
	"errors"
	"net/http"

	"github.com/UltimateTournament/TangiaMonoAPI/gologger"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
)

type CustomContext struct {
	echo.Context
	RequestID string
	UserID    string
	logger    zerolog.Logger
}

func CreateReqContext(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqID := uuid.NewString()
		logger := logger.With().Str("reqID", reqID).Logger()
		ctx := context.WithValue(c.Request().Context(), gologger.ReqIDKey, reqID)
		ctx = logger.WithContext(ctx)
		c.SetRequest(c.Request().WithContext(ctx))
		cc := &CustomContext{
			Context:   c,
			RequestID: reqID,
			// TODO remove and replace with zerolog.Ctx()
			logger: logger,
		}
		return next(cc)
	}
}

// Casts to custom context for the handler, so this doesn't have to be done per handler
func ccHandler(h func(*CustomContext) error) echo.HandlerFunc {
	// TODO: Include the path?
	return func(c echo.Context) error {
		return h(c.(*CustomContext))
	}
}

func (c *CustomContext) internalErrorMessage() string {
	return "internal error, request id: " + c.RequestID
}

func (c *CustomContext) InternalError(err error, msg string) error {
	if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
		c.logger.Warn().CallerSkipFrame(1).Msg(err.Error())
	} else {
		c.logger.Error().CallerSkipFrame(1).Err(err).Msg(msg)
	}
	return c.String(http.StatusInternalServerError, c.internalErrorMessage())
}

// Sets the user property, and updates the logger
func (c *CustomContext) SetUser(userID string) {
	c.UserID = userID
	c.logger = c.logger.With().Str("userID", userID).Logger()
}
