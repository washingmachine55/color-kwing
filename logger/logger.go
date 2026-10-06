package logger

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

type Config struct {
	mu         sync.RWMutex
	level      zerolog.Level
	isPretty   bool
	writer     io.Writer
	timeFormat string
	withCaller bool
}

var (
	// Changed to a pointer so everyone shares the exact same memory address
	instance *zerolog.Logger
	once     sync.Once
	cfg      *Config
)

type Option func(*Config)

func initDefaults() {
	cfg = &Config{
		level:      zerolog.DebugLevel,
		isPretty:   true,
		writer:     os.Stdout,
		timeFormat: time.Kitchen,
		withCaller: true,
	}
}

// Get now returns a POINTER (*zerolog.Logger)
func Get() *zerolog.Logger {
	once.Do(func() {
		if cfg == nil {
			initDefaults()
		}
		// Initialize the pointer once
		l := zerolog.New(os.Stdout)
		instance = &l
		rebuildLogger()
	})

	return instance
}

func rebuildLogger() {
	var out io.Writer = cfg.writer

	if cfg.isPretty {
		out = zerolog.ConsoleWriter{
			Out:        cfg.writer,
			TimeFormat: cfg.timeFormat,
		}
	}

	zerolog.SetGlobalLevel(cfg.level)

	ctx := zerolog.New(out).With().Timestamp()
	if cfg.withCaller {
		ctx = ctx.Caller()
	}

	// Mutate the actual underlying logger instead of replacing the pointer
	cfg.mu.Lock()
	*instance = ctx.Logger()
	cfg.mu.Unlock()
}

func Configure(opts ...Option) {
	once.Do(func() {
		// If Configure is called first, initialize everything
		l := zerolog.New(os.Stdout)
		instance = &l
		initDefaults()
	})

	for _, opt := range opts {
		opt(cfg)
	}
	rebuildLogger()
}

// ... Keep your WithJSON, WithPrettyPrint, etc. functions exactly the same ...
