package apiserver

// api server ...
type APIServer struct{}

// init New
func New() *APIServer {
	return &APIServer{}
}

// start ...
func (s *APIServer) Start() error {
	return nil
}
