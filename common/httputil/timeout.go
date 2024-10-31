package httputil

import "time"

type HttpTimeout struct {
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
}

func WithTimeouts(timeout HttpTimeout) HttpOption {
	return func(s *HttpServer) error {
		s.srv.ReadTimeout = timeout.ReadTimeout
		s.srv.ReadHeaderTimeout = timeout.ReadHeaderTimeout
		s.srv.WriteTimeout = timeout.WriteTimeout
		s.srv.IdleTimeout = timeout.IdleTimeout
		return nil
	}
}

var DefaultTimeout = HttpTimeout{
	ReadTimeout:       15 * time.Second,
	ReadHeaderTimeout: 15 * time.Second,
	WriteTimeout:      15 * time.Second,
	IdleTimeout:       15 * time.Second,
}
