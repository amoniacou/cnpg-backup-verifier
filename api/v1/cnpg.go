package v1

import corev1 "k8s.io/api/core/v1"

type CNPGCluster struct {
	// Name of cluster
	Name string `json:"name"`

	// CPU and Memory resources for the PostgreSQL cluster pods. Please refer to
	// https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	// for more information.
	// +kubebuilder:validation:Optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`

	// Configuration of the storage of the instances
	// +kubebuilder:validation:Optional
	StorageConfiguration StorageSpec `json:"storage,omitempty"`
}

// StorageSpec defines the storage settings for the verification cluster
type StorageSpec struct {
	// Size of the persistent volume
	// +kubebuilder:validation:Pattern=`^\d+(Gi|Mi|Ti)$`
	Size string `json:"size,omitempty"`

	// Storage class for the PVC
	StorageClass string `json:"storage_class,omitempty"`
}
