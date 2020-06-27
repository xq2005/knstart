package main

import (
	// The set of controllers this controller process runs.
	"knstart/pkg/reconciler"

	// This defines the shared main for injected controllers.
	"knative.dev/pkg/injection/sharedmain"
)

func main() {
	sharedmain.Main("knative_start", reconciler.NewController)
}
