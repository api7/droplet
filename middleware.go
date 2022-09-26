package droplet

type Middleware interface {
	SetNext(next Middleware)
	Handle(ctx Context) error
}

type handlerMiddleware struct {
	handler Handler
	next    Middleware
}

func (m *handlerMiddleware) SetNext(next Middleware) {
	m.next = next
}

func (m *handlerMiddleware) Handle(ctx Context) error {
	// if the request is rewritten, the handler should be called
	if ctx.Rewritten() {
		return nil
	}

	rs, err := m.handler(ctx)
	ctx.SetOutput(rs)
	return err
}

func NewHandlerMiddleware(handler Handler) Middleware {
	return &handlerMiddleware{handler: handler}
}
