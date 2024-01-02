package backend

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Server struct {
	Connection *sql.DB
	Config     ServerConfig
}

type ServerConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func NewServer(config ServerConfig) (*Server, error) {
	cnxString := "user=" + config.Username + " dbname=" + config.Database + " password=" + config.Password + " host=" + config.Host + " sslmode=disable"
	db, err := sql.Open("postgres", cnxString)
	if err != nil {
		return nil, err
	}
	return &Server{
		Connection: db,
		Config:     config,
	}, nil
}

func (s *Server) Close() {
	s.Connection.Close()
}

func (s *Server) Ping() bool {
	err := s.Connection.Ping()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
