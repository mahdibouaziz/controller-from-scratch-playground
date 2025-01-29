// +kubebuilder:rbac:groups=sample.example.com,resources=myresources,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sample.example.com,resources=myresources/status,verbs=get;update;patch
package controllers

import (
	"context"

	samplev1 "github.com/mahdibouaziz/controller-from-scratch-playground/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MyResourceReconciler reconciles a MyResource object
type MyResourceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// SetupWithManager registers the controller with the manager
func (r *MyResourceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&samplev1.MyResource{}).
		Complete(r)
}

func (r *MyResourceReconciler) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	// Fetch the MyResource instance
	myResource := &samplev1.MyResource{}

	if err := r.Get(ctx, req.NamespacedName, myResource); err != nil {
		return reconcile.Result{}, client.IgnoreNotFound(err)
	}

	// Define the desired Deployment
	deployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      myResource.Name + "-deployment",
			Namespace: myResource.Namespace,
		},

		Spec: appsv1.DeploymentSpec{
			Replicas: &myResource.Spec.ReplicaCount,
			Selector: &metav1.LabelSelector{
				MatchLabels: map[string]string{"app": myResource.Name},
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"app": myResource.Name},
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  myResource.Name,
							Image: "nginx",
						},
					},
				},
			},
		},
	}

	// Set MyResource as the owner of the Deployment
	if err := controllerutil.SetControllerReference(myResource, deployment, r.Scheme); err != nil {
		return reconcile.Result{}, err
	}

	// Check if the Deployment already exists
	found := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, found)
	if err != nil && client.IgnoreNotFound(err) != nil {
		return reconcile.Result{}, err
	} else if err == nil {
		// Update existing deployment if necessary
		if *found.Spec.Replicas != myResource.Spec.ReplicaCount {
			found.Spec.Replicas = &myResource.Spec.ReplicaCount
			if err := r.Update(ctx, found); err != nil {
				return reconcile.Result{}, err
			}
		}
	} else {
		// Create the Deployment
		if err := r.Create(ctx, deployment); err != nil {
			return reconcile.Result{}, err
		}
	}

	// Update the status of myResource
	myResource.Status.AvailableReplicas = int(*deployment.Spec.Replicas)
	if err := r.Status().Update(ctx, myResource); err != nil {
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil

}
