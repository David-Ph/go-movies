//go:build wireinject
// +build wireinject

package main

import (
	"net/http"

	"github.com/google/wire"
)

func InitializeServer() *http.Server {
	wire.Build(
		NewServer,
	)
	return nil
}
