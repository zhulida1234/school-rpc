package httputil

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"sync/atomic"
)

type HttpServer struct {
	listener net.Listener
	srv      *http.Server
	closed   atomic.Bool
}

type HttpOption func(svr *HttpServer) error

func StartHttpServer(addr string, handler http.Handler, opts ...HttpOption) (*HttpServer, error) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("listen error", err)
		return nil, errors.New("Init listener fail")
	}
	srvCtx, srvCancel := context.WithCancel(context.Background())
	srv := &http.Server{
		Handler:           handler,
		ReadTimeout:       timeouts.ReadTimeout,
		WriteTimeout:      timeouts.WriteTimeout,
		IdleTimeout:       timeouts.IdleTimeout,
		ReadHeaderTimeout: timeouts.ReadHeaderTimeout,
		BaseContext: func(listener net.Listener) context.Context {
			return srvCtx
		},
	}
	out := &HttpServer{listener: listener, srv: srv}

	for _, opt := range opts {
		if err := opt(out); err != nil {
			srvCancel()
			fmt.Printf("opt fail", err)
			return nil, errors.New("One of http op fail")
		}
	}
	go func() {
		err := out.srv.Serve(listener)
		srvCancel()
		if errors.Is(err, http.ErrServerClosed) {
			out.closed.Store(true)
		} else {
			fmt.Println("unkonw err:", err)
			panic("unkonw err")
		}
	}()
	return out, nil
}

func (hs *HttpServer) shutdown(ctx context.Context) error { return hs.srv.Shutdown(ctx) }

func (hs *HttpServer) Close() error { return hs.srv.Close() }

func (hs *HttpServer) Closed() bool { return hs.closed.Load() }

func (hs *HttpServer) Stop(ctx context.Context) error {
	if err := hs.shutdown(ctx); err != nil {
		if errors.Is(err, ctx.Err()) {
			return hs.Close()
		}
		return err
	}
	return nil
}

func (hs *HttpServer) Addr() net.Addr { return hs.listener.Addr() }

func WithMaxHeaderBytes(max int) HttpOption {
	return func(srv *HttpServer) error {
		srv.srv.MaxHeaderBytes = max
		return nil
	}
}
