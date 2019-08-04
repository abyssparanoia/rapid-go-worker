package main

import (
	"github.com/abyssparanoia/rapid-go-worker/src/handler"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/cloudfirestore"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/cloudpubsub"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/deploy"
	"github.com/abyssparanoia/rapid-go-worker/src/lib/log"
	"github.com/abyssparanoia/rapid-go-worker/src/repository"
	"github.com/abyssparanoia/rapid-go-worker/src/service"
)

// Dependency ... 依存性
type Dependency struct {
	Log          *log.Middleware
	Handle       *handler.MessageHandler
	Subscription cloudpubsub.Subscription
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

	svc := service.NewSample(repo)

	// Middleware
	d.Log = log.NewMiddleware(lCli, e.MinLogSeverity)

	d.Handle = handler.NewHandler(svc)

	// Subscription
	d.Subscription = cloudpubsub.NewSubscription(e.ProjectID, e.CredentialsPath, e.PubSubSubsriptionID)
}
