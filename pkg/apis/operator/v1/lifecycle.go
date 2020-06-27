package v1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
)

var condSet = apis.NewLivingConditionSet(
	Available,
	Succeeded,
)

func (ks *KNLearningStatus) GetCondition(t apis.ConditionType) *apis.Condition {
	return condSet.Manage(ks).GetCondition(t)
}

func (ks *KNLearning) GetStatus() *duckv1.Status {
	return &ks.Status.Status
}

// GetConditionSet returns SampleSource ConditionSet.
func (*KNLearning) GetConditionSet() apis.ConditionSet {
	return condSet
}

// GetGroupVersionKind implements kmeta.OwnerRefable
func (knl *KNLearning) GetGroupVersionKind() schema.GroupVersionKind {
	return SchemeGroupVersion.WithKind("KNLearning")
}

func (ks *KNLearningStatus) InitializeConditions() {
	condSet.Manage(ks).InitializeConditions()
}

func (ks *KNLearningStatus) SetUnavailable(name string) {
	condSet.Manage(ks).MarkFalse(
		Available,
		"DeploymentUnavailable",
		"Deployment %q wasn't found.", name)
}

func (ks *KNLearningStatus) SetReady() {
	condSet.Manage(ks).MarkTrue(Succeeded)
}

func (ks *KNLearningStatus) SetAvailable() {
	condSet.Manage(ks).MarkTrue(Available)
}
