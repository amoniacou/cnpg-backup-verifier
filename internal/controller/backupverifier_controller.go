/*
Copyright 2024 Amoniac OU.

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

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/builder"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	cnpgbvv1 "github.com/amoniacou/cnpg-backup-verifier/api/v1"
	cnpgv1 "github.com/cloudnative-pg/cloudnative-pg/api/v1"
)

// BackupVerifierReconciler reconciles a BackupVerifier object
type BackupVerifierReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=cnpg.k8s.amoniac.eu,resources=backupverifiers,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cnpg.k8s.amoniac.eu,resources=backupverifiers/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=cnpg.k8s.amoniac.eu,resources=backupverifiers/finalizers,verbs=update
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=clusters,verbs=get;list;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=backups,verbs=get;list;watch
// +kubebuilder:rbac:groups=postgresql.cnpg.io,resources=backups/status,verbs=get;update;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the BackupVerifier object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.18.4/pkg/reconcile
func (r *BackupVerifierReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("Reconcile for manager started", "manager", req.NamespacedName.Name, "namespace", req.NamespacedName.Namespace)

	var backupVerifier cnpgbvv1.BackupVerifier
	if err := r.Get(ctx, req.NamespacedName, &backupVerifier); err != nil {
		log.Error(err, "Unable to fetch backupVerifier entity")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	if backupVerifier.Spec.Cluster.Name == "" {
		log.Info("Cluster name is not set so skip such varifier")
		return ctrl.Result{}, nil

	}
	clusterName := backupVerifier.Spec.Cluster.Name
	cluster := cnpgv1.Cluster{}
	err := r.Get(ctx, types.NamespacedName{Name: clusterName, Namespace: backupVerifier.Namespace}, &cluster)
	if err != nil {
		return ctrl.Result{}, err
	}

	// check if cluster have a backup section
	if cluster.Spec.Backup == nil {
		backupVerifier.Status.State = "stopped"
		backupVerifier.Status.Reason = "There no backup section in cluster"
		log.Info("No backup section in Postgres cluster")
		if err := r.Status().Update(ctx, &backupVerifier); err != nil {
			log.Error(err, "Unable to update status")
		}
		return ctrl.Result{}, err
	}

	var backups cnpgv1.BackupList

	// Get the list of backups for cluster
	if err := r.List(ctx, &backups, client.InNamespace(cluster.Namespace), client.MatchingFields{"spec.cluster.name": clusterName}); err != nil {
		log.Error(err, "Unable to fetch all backups for cluster")
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *BackupVerifierReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cnpgbvv1.BackupVerifier{}).
		Watches(
			&cnpgv1.Cluster{},
			handler.EnqueueRequestsFromMapFunc(r.requestsForClusterChanges),
			builder.WithPredicates(predicate.ResourceVersionChangedPredicate{}),
		).
		Complete(r)
}

// n
func (r *BackupVerifierReconciler) requestsForClusterChanges(ctx context.Context, o client.Object) []reconcile.Request {

}
