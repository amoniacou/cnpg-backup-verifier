/*
Copyright 2025 Amoniac OU.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// WatchVerifierSpec defines the desired state of WatchVerifier.
type WatchVerifierSpec struct {
	// Cluster reference
	Cluster CNPGCluster `json:"cluster"`
}

// WatchVerifierStatus defines the observed state of WatchVerifier.
type WatchVerifierStatus struct {
	// Track all VerifyJobs created by this CronVerifier
	// +kubebuilder:validation:Optional
	VerifyJobs []corev1.ObjectReference `json:"verify_jobs,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// WatchVerifier is the Schema for the watches API.
type WatchVerifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WatchVerifierSpec   `json:"spec,omitempty"`
	Status WatchVerifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WatchVerifierList contains a list of WatchVerifier.
type WatchVerifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WatchVerifier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&WatchVerifier{}, &WatchVerifierList{})
}
