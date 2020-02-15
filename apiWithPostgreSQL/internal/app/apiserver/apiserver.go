package apiserver

// APIServer ...
type APIServer struct{
	config *Config
}

// New ...
func New() *APIServer{
	return &APIServer{};
}

// Start ...
func (s *APIServer) Start() error {
	return nil
}