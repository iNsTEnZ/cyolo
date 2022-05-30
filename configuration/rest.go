package configuration

import "time"

type RestConfiguration struct {
	Enabled               bool
	MaxConcurrentRequests int
	Port                  string
	GracefulShutdownSec   time.Duration
	ServerReadTimeoutSec  time.Duration
	ServerWriteTimeoutSec time.Duration
	ServerIdleTimeoutSec  time.Duration
}

func GetRestConfiguration() RestConfiguration {
	return RestConfiguration{
		Enabled:               EnvVerBool("ENABLE_REST_API", true),
		MaxConcurrentRequests: EnvVerInt("MAX_CONCURRENT_REQ", 100),
		Port:                  EnvVerStr("SERVICE_PORT", ":8080"),
		GracefulShutdownSec:   time.Second * time.Duration(EnvVerInt("GRACEFUL_SHUTDOWN_SEC", 25)),
		ServerReadTimeoutSec:  time.Second * time.Duration(EnvVerInt("SERVER_READ_TIMEOUT_SEC", 30)),
		ServerWriteTimeoutSec: time.Second * time.Duration(EnvVerInt("SERVER_WRITE_TIMEOUT_SEC", 30)),
		ServerIdleTimeoutSec:  time.Second * time.Duration(EnvVerInt("SERVER_IDLE_TIMEOUT_SEC", 30)),
	}
}
