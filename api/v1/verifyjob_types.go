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

// VerifyJobStatusType represents the state of a VerifyJob
// +kubebuilder:validation:Enum=Pending;Running;Succeeded;Failed
type VerifyJobStatusType string

const (
	VerifyJobPending   VerifyJobStatusType = "Pending"   // Job is created but not yet started
	VerifyJobRunning   VerifyJobStatusType = "Running"   // Job is actively restoring & verifying the backup
	VerifyJobSucceeded VerifyJobStatusType = "Succeeded" // Backup verification completed successfully
	VerifyJobFailed    VerifyJobStatusType = "Failed"    // Verification failed due to an error
)

// VerifyJobSpec defines the desired state of VerifyJob.
type VerifyJobSpec struct {
	// CronVerifier reference
	CronVerifierRef corev1.ObjectReference `json:"schedule"`

	// children CNPG cluster which we will create to validate backup
	CNPGClusterRef corev1.ObjectReference `json:"cluster"`

	// BackupName specifies the name of the backup that is being verified.
	// This is required for tracking which backup is used for restoration.
	BackupName string `json:"backup_name"`
}

// VerifyJobStatus defines the observed state of VerifyJob.
type VerifyJobStatus struct {
	// status represents the current state of the verification job.
	// Possible values: Pending, Running, Succeeded, Failed
	Status VerifyJobStatusType `json:"status,omitempty"`

	// last_updated records the last time this job's status was modified.
	LastUpdated *metav1.Time `json:"last_updated,omitempty"`

	// error_message contains details in case of failure.
	// This field is populated only when status is "Failed".
	ErrorMessage string `json:"error_message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VerifyJob is the Schema for the verifyjobs API.
type VerifyJob struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VerifyJobSpec   `json:"spec,omitempty"`
	Status VerifyJobStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VerifyJobList contains a list of VerifyJob.
type VerifyJobList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VerifyJob `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VerifyJob{}, &VerifyJobList{})
}
