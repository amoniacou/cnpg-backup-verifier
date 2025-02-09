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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CronVerifierStatusType represents the state of a CronVerifier
// +kubebuilder:validation:Enum=Active;Suspended;Failed
type CronVerifierStatusType string

const (
	CronVerifierActive    CronVerifierStatusType = "Active"    // Cron verifier is active
	CronVerifierFailed    CronVerifierStatusType = "Failed"    // Cron verifier have an issues
	CronVerifierSuspended CronVerifierStatusType = "Suspended" // Cron verifier is non-active
)

// CronVerifierSpec defines the desired state of CronVerifier
type CronVerifierSpec struct {
	// Cron expression to determine when to run verification jobs
	// +kubebuilder:validation:Pattern=`^(((\d+,)+\d+|(\d+(\/|-)\d+)|\d+|\*) ?){5,7}$`
	Cron string `json:"schedule"`

	// Cluster reference
	Cluster CNPGCluster `json:"cluster"`
}

// CronVerifierStatus defines the observed state of CronVerifier
type CronVerifierStatus struct {

	// status represents the current state of the verification job.
	// Possible values: Pending, Running, Succeeded, Failed
	Status CronVerifierStatusType `json:"status,omitempty"`

	// error_message contains details in case of failure.
	// This field is populated only when status is "Failed".
	ErrorMessage string `json:"error_message,omitempty"`

	// Last time a verification job was triggered
	// +kubebuilder:validation:Optional
	LastRunTime *metav1.Time `json:"lastRunTime,omitempty"`

	// Status of the last verification job
	// +kubebuilder:validation:Optional
	LastRunStatus string `json:"lastRunStatus,omitempty"`

	// Track all VerifyJobs created by this CronVerifier
	// +kubebuilder:validation:Optional
	VerifyJobs []corev1.ObjectReference `json:"verify_jobs,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CronVerifier is the Schema for the backupverifiers API
type CronVerifier struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CronVerifierSpec   `json:"spec,omitempty"`
	Status CronVerifierStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CronVerifierList contains a list of CronVerifier
type CronVerifierList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CronVerifier `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CronVerifier{}, &CronVerifierList{})
}
