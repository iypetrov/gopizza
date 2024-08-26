package app

import (
	"github.com/iypetrov/gopizza/pkg/app/config"
	"github.com/iypetrov/gopizza/pkg/app/logger"
	"github.com/iypetrov/gopizza/pkg/app/server"
	"net/http"
)

var (
	Log    logger.Logger
	Cfg    *config.Config
	Server *http.Server
)

func Init() {
	Cfg = config.New()
	Log = logger.New()
	Server = server.New(Cfg)
}
