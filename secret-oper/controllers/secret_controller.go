/*
Copyright 2022.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-logr/logr"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// SecretReconciler reconciles a Secret object
type SecretReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
}

//+kubebuilder:rbac:groups=haithamoon.me,resources=secrets,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=haithamoon.me,resources=secrets/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=haithamoon.me,resources=secrets/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Secret object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
var globalLog = logf.Log.WithName("global")

func (r *SecretReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	opts := zap.Options{}
	logger := zap.New(zap.UseFlagOptions(&opts))
	logf.SetLogger(logger)
	// globalLog.Info("Printing at INFO level")
	//log := globalLog
	secret := &corev1.Secret{}
	r.Get(ctx, req.NamespacedName, secret)

	dep := &appv1.DeploymentList{}
	listOps := []client.ListOption{
		client.InNamespace("default")}
	r.List(ctx, dep, listOps...)

	for _, d := range dep.Items {

		fmt.Println(d.Name, d.Namespace)

	}
	fmt.Println("Doneeeeeeee")

	for i := range dep.Items {
		secretChanged := strings.Contains(dep.Items[i].Annotations["sadafnoor.me/pod-delete-on-secret-change"], secret.Name)

		if secretChanged {
			podt := &corev1.Pod{}
			fmt.Println("Trying to delete all pods that has been using secret.Name: " + secret.Name)
			r.DeleteAllOf(ctx, podt, client.InNamespace(req.NamespacedName.Namespace), client.MatchingLabels(dep.Items[i].Spec.Selector.MatchLabels))
		}

	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *SecretReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Secret{}).
		Complete(r)

}
