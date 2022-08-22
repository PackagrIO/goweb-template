package pkg

import (
	"github.com/packagrio/goweb-template/backend/pkg/config"
)

type Pipeline struct {
	Config config.Interface
}

func (p *Pipeline) Start(configData config.Interface) error {
	// Initialize Pipeline.
	p.Config = configData
	return nil
}
