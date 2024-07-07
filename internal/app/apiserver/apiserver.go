package apiserver

// api server ...
type APIServer struct {
	config *Config
}

// init New
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// start ...
func (s *APIServer) Start() error {
	return nil
}
