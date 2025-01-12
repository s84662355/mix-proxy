package  common

import (
	"context"
	"github.com/spf13/viper"
	  "mix-proxy/uitl/Map"
)

func New(ctx context.Context) *Common{
	c:=&Common{
		ctx:ctx,
	}
	c.initConfig()
	c.initConnMap() 
	return c
}

type Common struct{
    ctx context.Context
    viperTool *viper.Viper
    viperConfig  *Config
    connConcurrentMap *Map.ConcurrentMap[*ConnContext]
}

