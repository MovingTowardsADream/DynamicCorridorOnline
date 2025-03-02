package middleware

type Middlewares struct {
	AuthMiddleware
	CORSMiddleware
}

func New(host string, auth Authorization) *Middlewares {
	return &Middlewares{
		AuthMiddleware{auth: auth},
		CORSMiddleware{host: host},
	}
}
