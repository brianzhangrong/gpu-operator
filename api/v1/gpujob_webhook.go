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

package v1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var gpujoblog = logf.Log.WithName("gpujob-resource")

func (r *GpuJob) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-sigma-ihomefnt-com-v1-gpujob,mutating=true,failurePolicy=fail,groups=sigma.ihomefnt.com,resources=GpuJob,verbs=create;update;delete,versions=v1,name=mgpujob.kb.io

var _ webhook.Defaulter = &GpuJob{}

// Default implements webhook.Defaulter so a webhook will be registered for the type
func (r *GpuJob) Default() {
	fmt.Println("-------------------")

	gpujoblog.Info("default", "name", r.Name, r.Kind, r.Spec.CPU)
	r.Status.Status = "Pending"
	r.Status.UpdateLastTime = metav1.Now()
	// TODO(user): fill in your defaulting logic.
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-sigma-ihomefnt-com-v1-gpujob,mutating=false,failurePolicy=fail,groups=sigma.ihomefnt.com,resources=GpuJob,versions=v1,name=vgpujob.kb.io

var _ webhook.Validator = &GpuJob{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *GpuJob) ValidateCreate() error {
	fmt.Println("1111111111111111")
	gpujoblog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *GpuJob) ValidateUpdate(old runtime.Object) error {
	fmt.Println("1111111111111111")
	gpujoblog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *GpuJob) ValidateDelete() error {
	fmt.Println("1111111111111111")
	gpujoblog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
