package app

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type App struct {
	ctx      context.Context
	cancel   func()
	opts     options
	instance *registry.ServiceInstance
}

// Info is application context value.
type Info interface {
	ID() string
	Name() string
	Version() string
	Metadata() map[string]string
	Endpoint() []string
}

// Option is an application option.
type Option func(o *options)

// options is an application options.
type options struct {
	id        string
	name      string
	version   string
	metadata  map[string]string
	endpoints []*url.URL

	ctx   context.Context
	signs []os.Signal

	logger           *log.Helper
	registrar        registry.Registrar
	registrarTimeout time.Duration
	servers          []transport.Server
}

// New create an application lifecycle manager.
func New(opts ...Option) *App {
	options := options{
		ctx:   context.Background(),
		signs: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}

	if id, err := uuid.NewUUID(); err == nil {
		options.id = id.String()
	}

	for _, o := range opts {
		o(&options)
	}
	ctx, cancel := context.WithCancel(options.ctx)
	return &App{
		ctx:    ctx,
		cancel: cancel,
		opts:   options,
	}
}

// ID returns app instance id.
func (a *App) ID() string { return a.opts.id }

// Name returns service name.
func (a *App) Name() string { return a.opts.name }

// Version returns app version.
func (a *App) Version() string { return a.opts.version }

// Metadata returns service metadata.
func (a *App) Metadata() map[string]string { return a.opts.metadata }

// Endpoint returns endpoints.
func (a *App) Endpoint() []string { return a.instance.Endpoints }

func (a *App) Run() error {
	instance, err := a.buildInstance()
	if err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(a.ctx)
	wg := sync.WaitGroup{}
	for _, srv := range a.opts.servers {
		srv := srv
		eg.Go(func() error {
			<-ctx.Done() // wait for stop signal
			return srv.Stop(ctx)
		})
		wg.Add(1)
		eg.Go(func() error {
			wg.Done()
			return srv.Start(ctx)
		})
	}
	wg.Wait()
	if a.opts.registrar != nil {
		ctx, cancel := context.WithTimeout(a.opts.ctx, a.opts.registrarTimeout)
		defer cancel()
		if err := a.opts.registrar.Register(ctx, instance); err != nil {
			return err
		}
		a.instance = instance
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, a.opts.signs...)
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				err := a.Stop()
				if err != nil {
					a.opts.logger.Errorf("failed to app stop: %v", err)
				}
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

// Stop gracefully stops the application.
func (a *App) Stop() error {
	if a.opts.registrar != nil && a.instance != nil {
		ctx, cancel := context.WithTimeout(a.opts.ctx, a.opts.registrarTimeout)
		defer cancel()
		if err := a.opts.registrar.Deregister(ctx, a.instance); err != nil {
			return err
		}
	}
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}

func (a *App) buildInstance() (*registry.ServiceInstance, error) {
	var endpoints []string
	for _, e := range a.opts.endpoints {
		endpoints = append(endpoints, e.String())
	}
	if len(endpoints) == 0 {
		for _, srv := range a.opts.servers {
			if r, ok := srv.(transport.Endpointer); ok {
				e, err := r.Endpoint()
				if err != nil {
					return nil, err
				}
				endpoints = append(endpoints, e.String())
			}
		}
	}
	return &registry.ServiceInstance{
		ID:        a.opts.id,
		Name:      a.opts.name,
		Version:   a.opts.version,
		Metadata:  a.opts.metadata,
		Endpoints: endpoints,
	}, nil
}

// ID with service id.
func ID(id string) Option {
	return func(o *options) { o.id = id }
}

// Name with service name.
func Name(name string) Option {
	return func(o *options) { o.name = name }
}

// Version with service version.
func Version(version string) Option {
	return func(o *options) { o.version = version }
}

// Metadata with service metadata.
func Metadata(md map[string]string) Option {
	return func(o *options) { o.metadata = md }
}

// Endpoint with service endpoint.
func Endpoint(endpoints ...*url.URL) Option {
	return func(o *options) { o.endpoints = endpoints }
}

// Context with service context.
func Context(ctx context.Context) Option {
	return func(o *options) { o.ctx = ctx }
}

// Logger with service logger.
func Logger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = log.NewHelper(logger)
	}
}

// Server with transport servers.
func Server(srv ...transport.Server) Option {
	return func(o *options) { o.servers = srv }
}

// Signal with exit signals.
func Signal(signs ...os.Signal) Option {
	return func(o *options) { o.signs = signs }
}

// Registrar with service registry.
func Registrar(r registry.Registrar) Option {
	return func(o *options) { o.registrar = r }
}

// RegistrarTimeout with registrar timeout.
func RegistrarTimeout(t time.Duration) Option {
	return func(o *options) { o.registrarTimeout = t }
}
