package requestid

import (
	"context"
	"github.com/savsgio/atreugo/v11"
)

// AtreugoMiddleware puts the requestID into the context value.
func AtreugoMiddleware() atreugo.Middleware {
	return func(ctx *atreugo.RequestCtx) error {
		attachedCtx := ctx.AttachedContext()
		if attachedCtx == nil {
			attachedCtx = context.Background()
		}

		requestID := string(ctx.Request.Header.Peek(headerXRequestID))

		requestIDContext := Set(attachedCtx, requestID)
		ctx.AttachContext(requestIDContext)

		// also set the header on the response (needed if the requesting client didn't generate the id themselves)
		ctx.Response.Header.Set(headerXRequestID, requestID)

		return ctx.Next()
	}
}
