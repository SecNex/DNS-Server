package resolver

import (
	"database/sql"
	"fmt"

	"github.com/secnex/dns-server/backend"
	"github.com/secnex/dns-server/config"
)

type Backend struct {
	Connection *sql.DB
	Config     config.BackendConfig
}

func NewBackend(config config.BackendConfig) (*Backend, error) {
	server, err := backend.NewServer(backend.ServerConfig{
		Host:     config.Host,
		Port:     config.Port,
		Username: config.Username,
		Password: config.Password,
		Database: config.Database,
	})
	if err != nil {
		return nil, err
	}

	connected := server.Ping()
	if connected {
		fmt.Println("Connected to database!")
		return &Backend{
			Connection: server.Connection,
			Config:     config,
		}, nil
	} else {
		return nil, err
	}
}
