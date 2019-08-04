package main

import (
	"github.com/abyssparanoia/rapid-go-worker/src/handler/api"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/cloudfirestore"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/deploy"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/firebaseauth"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/jsonrpc2"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/log"
	"github.com/abyssparanoia/rapid-go-worker/src/repository"
	"github.com/abyssparanoia/rapid-go-worker/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	Log             *log.Middleware
	FirebaseAuth    *firebaseauth.Middleware
	SampleHandler   *api.SampleHandler
	JSONRPC2Handler *jsonrpc2.Handler
}

// Inject ... 依存性を注入する
func (d *Dependency) Inject(e *Environment) {
	// Client
	fCli := cloudfirestore.NewClient(e.CredentialsPath)
	var lCli log.Writer
	if deploy.IsLocal() {
		lCli = log.NewWriterStdout()
	} else {
		lCli = log.NewWriterStackdriver(e.ProjectID)
	}

	// Repository
	repo := repository.NewSample(fCli)

	// Service
	var faSvc firebaseauth.Service
	if deploy.IsProduction() {
		faSvc = firebaseauth.NewService()
	} else {
		faSvc = firebaseauth.NewDebugService()
	}
	svc := service.NewSample(repo)

	// Middleware
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)
	d.FirebaseAuth = firebaseauth.NewMiddleware(faSvc)

	// Handler
	d.SampleHandler = api.NewSampleHandler(svc)
	d.JSONRPC2Handler = jsonrpc2.NewHandler()
}
