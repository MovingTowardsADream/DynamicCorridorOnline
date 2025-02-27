package middleware

type Middlewares struct {
	AuthMiddleware
}

func New(auth Authorization) *Middlewares {
	return &Middlewares{
		AuthMiddleware{auth: auth},
	}
}
