package mylogger

import (
	"context"
	"errors"

	sensor "go.viam.com/rdk/components/sensor"
	"go.viam.com/rdk/logging"
	"go.viam.com/rdk/resource"
)

var (
	MyLogger         = resource.NewModel("felixr", "my-logger", "my-logger")
	errUnimplemented = errors.New("unimplemented")
)

func init() {
	resource.RegisterComponent(sensor.API, MyLogger,
		resource.Registration[sensor.Sensor, *Config]{
			Constructor: newMyLoggerMyLogger,
		},
	)
}

type Config struct {
	/*
		Put config attributes here. There should be public/exported fields
		with a `json` parameter at the end of each attribute.

		Example config struct:
			type Config struct {
				Pin   string `json:"pin"`
				Board string `json:"board"`
				MinDeg *float64 `json:"min_angle_deg,omitempty"`
			}

		If your model does not need a config, replace *Config in the init
		function with resource.NoNativeConfig
	*/
}

// Validate ensures all parts of the config are valid and important fields exist.
// Returns implicit required (first return) and optional (second return) dependencies based on the config.
// The path is the JSON path in your robot's config (not the `Config` struct) to the
// resource being validated; e.g. "components.0".
func (cfg *Config) Validate(path string) ([]string, []string, error) {
	// Add config validation code here
	return nil, nil, nil
}

type myLoggerMyLogger struct {
	resource.AlwaysRebuild

	name resource.Name

	logger logging.Logger
	cfg    *Config

	cancelCtx  context.Context
	cancelFunc func()
}

func newMyLoggerMyLogger(ctx context.Context, deps resource.Dependencies, rawConf resource.Config, logger logging.Logger) (sensor.Sensor, error) {
	conf, err := resource.NativeConfig[*Config](rawConf)
	if err != nil {
		return nil, err
	}

	return NewMyLogger(ctx, deps, rawConf.ResourceName(), conf, logger)

}

func NewMyLogger(ctx context.Context, deps resource.Dependencies, name resource.Name, conf *Config, logger logging.Logger) (sensor.Sensor, error) {

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	s := &myLoggerMyLogger{
		name:       name,
		logger:     logger,
		cfg:        conf,
		cancelCtx:  cancelCtx,
		cancelFunc: cancelFunc,
	}
	return s, nil
}

func (s *myLoggerMyLogger) Name() resource.Name {
	return s.name
}

func (s *myLoggerMyLogger) Readings(ctx context.Context, extra map[string]interface{}) (map[string]interface{}, error) {
	s.logger.Debugf("Debug message: %s", s.name)
	s.logger.Infof("Info message: %s", s.name)
	s.logger.Warnf("Warning message: %s", s.name)
	s.logger.Errorf("Error message: %s", s.name)
	//s.logger.Fatalf("Fatal message: %s", s.name) --> This will log a fatal error and then exit the program! https://medium.com/@emusbeny/when-to-use-log-fatalf-vs-log-errorf-in-go-a-friendly-guide-4ec574b50410
	return map[string]interface{}{"messages": "logged"}, nil
}
func (s *myLoggerMyLogger) DoCommand(ctx context.Context, cmd map[string]interface{}) (map[string]interface{}, error) {
	if cmd == nil {
		return nil, errUnimplemented
	}
	switch cmd["message"] {
	case "debug":
		s.logger.Debugf("Debug command received: %v", cmd)
	case "info":
		s.logger.Infof("Info command received: %v", cmd)
	case "warn":
		s.logger.Warnf("Warning command received: %v", cmd)
	case "error":
		s.logger.Errorf("Error command received: %v", cmd)
	case "fatal":
		s.logger.Fatalf("Fatal command received: %v", cmd)
	default:
		s.logger.Warnf("Unknown command received: %v", cmd)
	}
	return cmd, nil
}

func (s *myLoggerMyLogger) Close(context.Context) error {
	// Put close code here
	s.cancelFunc()
	return nil
}
