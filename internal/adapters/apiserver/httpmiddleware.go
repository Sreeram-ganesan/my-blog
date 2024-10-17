package apiserver

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"github.com/Sreeram-ganesan/my-blog/internal/core/app"
	"time"
)

const httpLogFormat = `"[END] %s %s %s" from %s`
const xRequestId = "x-request-id"

// zapLoggerMiddleware provides middleware for adding zap logger into the context of request handler
func zapLoggerMiddleware(logger *zap.SugaredLogger) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqID := c.GetHeader(xRequestId)
		if reqID == "" {
			reqID = uuid.NewString()
		}
		req := c.Request
		// Set the id to ensure that the requestid is in the response
		c.Header(xRequestId, reqID)
		l := logger.With(zap.String("requestId", reqID))
		c.Request = c.Request.WithContext(app.ContextWithLogger(req.Context(), l))
		tbegin := time.Now()
		defer func() {
			l.With(zap.String("duration", time.Since(tbegin).String())).Infof(httpLogFormat,
				req.Method,
				req.URL.Path,
				req.Proto,
				req.RemoteAddr,
			)
		}()
		c.Next()
	}
}
