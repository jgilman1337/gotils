package cfg

import (
	"sync"
)

var (
	//Config *Config
	once *sync.Once = &sync.Once{}
)

// Initializes the config singleton.
func Singleton() {
	once.Do(func() {
		/*
			cfg, err := InitConfig()
			if err != nil {
				log.Fatalf("Failed to initialize config: %s", err)
			}
			EnvConfig = cfg
		*/
	})
}
