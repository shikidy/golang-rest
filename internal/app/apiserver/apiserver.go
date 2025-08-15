package apiserver

type ApiServer struct{}

func New() *ApiServer {
	return &ApiServer{}
}

func (a *ApiServer) Start() error {
	return nil
}
