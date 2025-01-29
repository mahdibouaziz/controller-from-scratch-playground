package main

import (
	"os"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

func main() {
	ctrl.SetLogger(zap.New(zap.UseDevMode(true)))

	// Create new manager
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{})
	if err != nil {
		os.Exit(1)
	}

	// Add your controller to the manager
	if err = (&controllers.MyResourceReconciler{
		Client: mgr.GetClient(),
	}).SetupWithManager(mgr); err != nil {
		os.Exit(1)
	}

}
