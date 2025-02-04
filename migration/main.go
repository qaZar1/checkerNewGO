package main

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"

	"github.com/Impisigmatus/service_core/postgres"
	"github.com/jmoiron/sqlx"
	"github.com/qaZar1/checkerNewGO/migration/internal/parser"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			file := f.File
			file = file[len(path.Dir(os.Args[0]))+1:]
			return "", fmt.Sprintf("%s:%d", file, f.Line)
		},
	})
}

func main() {
	const (
		base        = 10
		size        = 64
		address     = "ADDRESS"
		auth        = "APIS_AUTH_BASIC"
		pgHost      = "POSTGRES_HOSTNAME"
		pgPort      = "POSTGRES_PORT"
		pgDB        = "POSTGRES_DATABASE"
		pgUser      = "POSTGRES_USER"
		pgPassword  = "POSTGRES_PASSWORD"
		siteAddress = "ADDRESS_SITE"
	)

	port, err := strconv.ParseUint(os.Getenv(pgPort), base, size)
	if err != nil {
		logrus.Panicf("Invalid postgres port: %s", err)
	}

	site := parser.NewSite(os.Getenv(siteAddress),
		sqlx.NewDb(postgres.NewPostgres(
			postgres.Config{
				Hostname: os.Getenv(pgHost),
				Port:     port,
				Database: os.Getenv(pgDB),
				User:     os.Getenv(pgUser),
				Password: os.Getenv(pgPassword),
			},
		), "pgx"))

	if err := site.ParseReleases(); err != nil {
		logrus.Errorf("Invalid parse: %s", err)
	}
}
