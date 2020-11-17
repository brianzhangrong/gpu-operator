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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GpuJobSpec defines the desired state of GpuJob
type GpuJobSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of GpuJob. Edit GpuJob_types.go to remove/update
	JobName string `json:"jobName,omitempty"`
	Type    string `json:"type"`
	GPU     int    `json:"gpu"`
	CPU     int    `json:"cpu"`
	Memory  int    `json:"memory"`
}

// GpuJobStatus defines the observed state of GpuJob
type GpuJobStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	UpdateLastTime metav1.Time `json:"update_last_time"`
	Status         string      `json:"status"`
}

// +kubebuilder:object:root=true

// GpuJob is the Schema for the gpujobs API
type GpuJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GpuJobSpec   `json:"spec,omitempty"`
	Status GpuJobStatus `json:"status"`
}

// +kubebuilder:object:root=true

// GpuJobList contains a list of GpuJob
type GpuJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GpuJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GpuJob{}, &GpuJobList{})
}
