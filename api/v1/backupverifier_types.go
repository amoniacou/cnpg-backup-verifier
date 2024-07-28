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

package v1

import (
	cnpg "github.com/cloudnative-pg/cloudnative-pg/api/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BackupVerifierSpec defines the desired state of BackupVerifier
type BackupVerifierSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Cluster is a name of CloudNativePG Cluster which backups we need to validate
	Cluster cnpg.LocalObjectReference `json:"cluster"`
}

// BackupVerifierStatus defines the observed state of BackupVerifier
type BackupVerifierStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// What is the state of such verifier.
	// Could be a next states:
	// * validated - when we have found cluster and all backups are valid
	// * checking - when we have started to check all backups
	// * failed - when we have one or more backups invalid
	// * stopped - there is no backup section inside cluster so no need to do anything
	State string `json:"state"`

	// Reason of state
	// For example if its stopped we will add info that its stopped due to missing backup section
	Reason string `json:"reason,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// BackupVerifier is the Schema for the backupverifiers API
type BackupVerifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BackupVerifierSpec   `json:"spec,omitempty"`
	Status BackupVerifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BackupVerifierList contains a list of BackupVerifier
type BackupVerifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BackupVerifier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BackupVerifier{}, &BackupVerifierList{})
}
