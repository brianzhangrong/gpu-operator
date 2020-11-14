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
	"fmt"

	ihomefntcomv1 "gpu/api/v1"

	"github.com/go-logr/logr"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

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
	if err := r.Get(ctx, req.NamespacedName, &gpujob); err != nil {
		log.Error(err, "unable to get gpujob")
	} else {
		fmt.Println("Get gpujob spec info success, ", gpujob.Spec.JobName, gpujob.Spec.Type, gpujob.Spec.CPU, gpujob.Spec.Memory, gpujob.Spec.GPU)

	}

	gpujob.Status.UpdateLastTime = metav1.Now()
	gpujob.Status.Status = "Running"
	// fmt.Println("update gpujob status...", gpujob.Status.Status, gpujob.Status.UpdateLastTime, gpujob.APIVersion)

	if err := r.Update(ctx, &gpujob); err != nil {
		log.Error(err, "not update gpujob  status.")
	}

	return ctrl.Result{}, nil
}

func (r *GpuJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&sigmav1.GpuJob{}).
		Complete(r)
}
