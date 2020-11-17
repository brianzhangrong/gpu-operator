/*


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

	ihomefntcomv1 "gpu/api/v1"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	sigmav1 "gpu/api/v1"
)

// GpuJobReconciler reconciles a GpuJob object
type GpuJobReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=sigma.ihomefnt.com,resources=gpujobs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=sigma.ihomefnt.com,resources=gpujobs/status,verbs=get;update;patch

func (r *GpuJobReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("gpujob", req.NamespacedName)

	// your logic here
	var gpujob ihomefntcomv1.GpuJob

	// fmt.Println("update gpujob status...", gpujob.Status.Status, gpujob.Status.UpdateLastTime, gpujob.APIVersion)
	if err := r.Get(ctx, req.NamespacedName, &gpujob); err != nil {
		if err != nil {
			if errors.IsNotFound(err) {
				// Object not found, return.  Created objects are automatically garbage collected.
				// For additional cleanup logic use finalizers.
				return reconcile.Result{}, nil
			}
			// Error reading the object - requeue the request.
			return reconcile.Result{}, err
		}
	}
	// matchedPods := &corev1.PodList{}
	// if err := r.List(context.TODO(), matchedPods, &client.ListOptions{}); err != nil {
	// 	return reconcile.Result{}, err
	// }
	// for i := range matchedPods.Items {
	// 	pod := &matchedPods.Items[i]
	// 	fmt.Println("----", pod.Kind, pod.Name, pod.Namespace)
	// }

	gpujob.Status.UpdateLastTime = metav1.Now()
	gpujob.Status.Status = "Running"
	if err := r.Update(ctx, &gpujob); err != nil {
		log.Error(err, "222222 not update gpujob  status.")
		return ctrl.Result{}, err
	}
	// log.Info("update status", gpujob.Spec.JobName, gpujob.Status.Status)

	return ctrl.Result{}, nil
}

func (r *GpuJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sigmav1.GpuJob{}).
		Complete(r)
}
