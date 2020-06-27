package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

const (
	Succeeded apis.ConditionType = "Succeeded"
	Available apis.ConditionType = "DeploymeentAvailable"
)

// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KNLearning is the CRD defination
type KNLearning struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata, omitempty"`

	Spec   KNLearningSpec   `json:"spec"`
	Status KNLearningStatus `json:"status"`
}

// KNLearningSpec is the spec part in KNLearning
type KNLearningSpec struct {
	Size      *int32                      `json:"size"`
	Image     string                      `json:"image"`
	Resources corev1.ResourceRequirements `json:"resources, omitempty"`
	Ports     []corev1.ServicePort        `json:"ports, omitempty"`
}

// KNLearningStatus shows the deployment status
type KNLearningStatus struct {
	duckv1.Status `json:",inline"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KNLearningList is the format of KNLearning list
type KNLearningList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []KNLearning `json:"items"`
}
