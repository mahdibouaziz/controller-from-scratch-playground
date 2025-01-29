package main

import (
	"os"

	samplev1 "github.com/mahdibouaziz/controller-from-scratch-playground/api/v1"
	"github.com/mahdibouaziz/controller-from-scratch-playground/controllers"
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	scheme   = runtime.NewScheme()
	setupLog = ctrl.Log.WithName("setup")
)

func init() {
	// Register the MyResource CRD with the scheme
	utilruntime.Must(samplev1.AddToScheme(scheme))
}

func main() {
	ctrl.SetLogger(ctrl.Log.WithName("controller"))

	// Create new manager
	mgr, err := manager.New(config.GetConfigOrDie(), manager.Options{
		Scheme: scheme,
	})
	if err != nil {
		setupLog.Error(err, "Unable to start manager")
		os.Exit(1)
	}

	// Add the reconciler to the manager
	if err = (&controllers.MyResourceReconciler{
		Client: mgr.GetClient(),
		Scheme: mgr.GetScheme(),
	}).SetupWithManager(mgr); err != nil {
		setupLog.Error(err, "Unable to create controller", "controller", "MyResource")
		os.Exit(1)
	}

	// Start the manager
	setupLog.Info("Starting controller manager")
	if err := mgr.Start(ctrl.SetupSignalHandler()); err != nil {
		setupLog.Error(err, "Problem running manager")
		os.Exit(1)
	}

}
