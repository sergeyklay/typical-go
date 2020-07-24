package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/kelseyhightower/envconfig"
	"github.com/typical-go/typical-go/examples/use-config/internal/app"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() {
	typapp.AppendCtor(
		&typapp.Constructor{
			Name: "",
			Fn: func() (*app.ServerCfg, error) {
				var cfg app.ServerCfg
				if err := envconfig.Process("SERVER", &cfg); err != nil {
					return nil, err
				}
				return &cfg, nil
			},
		},
	)
}
