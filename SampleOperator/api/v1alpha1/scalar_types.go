/*
Copyright 2024 Phani Aalla.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	SUCCESS = "Success"
	FAILED  = "Failed"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ScalarSpec defines the desired state of Scalar
type ScalarSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Scalar. Edit scalar_types.go to remove/update
	Start       int              `json:"start"`
	End         int              `json:"end"`
	Replicas    int32            `json:"replicas"`
	Deployments []NamespacedName `json:"deployments"`
}

type NamespacedName struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// ScalarStatus defines the observed state of Scalar
type ScalarStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Status string `json:"status,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Scalar is the Schema for the scalars API
type Scalar struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ScalarSpec   `json:"spec,omitempty"`
	Status ScalarStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ScalarList contains a list of Scalar
type ScalarList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Scalar `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Scalar{}, &ScalarList{})
}
