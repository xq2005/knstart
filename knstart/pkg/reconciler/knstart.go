package reconciler

import (
	"context"
	"fmt"
	"reflect"

	// k8s
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	apierrs "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/kubernetes"
	v1listers "k8s.io/client-go/listers/apps/v1"

	// knative

	"knative.dev/pkg/logging"
	pkgreconciler "knative.dev/pkg/reconciler"

	// knstart defined interface
	v1 "knstart/pkg/apis/operator/v1"
	reconcilekindinterface "knstart/pkg/client/injection/reconciler/operator/v1/knlearning"
)

// Reconciler reconciles a knstart object
type Reconciler struct {
	DeploymentLister v1listers.DeploymentLister
	K8ClientSet      kubernetes.Interface
}

var _ reconcilekindinterface.Interface = (*Reconciler)(nil)

// ReconcileKind implements Interface.ReconcileKind.
func (r *Reconciler) ReconcileKind(ctx context.Context, req *v1.KNLearning) pkgreconciler.Event {
	logging.FromContext(ctx).Info("Enter ReconcleKind, NameSpace:" + req.Namespace + " with name: " + req.Name)

	req.Status.InitializeConditions()

	// reconciler Depoyment
	return r.ReconcileDeployment(ctx, req)
}

// ReconcileDeployment create or update deployment based on the new req and old deployment
func (r *Reconciler) ReconcileDeployment(ctx context.Context, req *v1.KNLearning) error {
	ns := req.Namespace
	dpName := req.Name + "-dep"

	deployment, err := r.DeploymentLister.Deployments(ns).Get(dpName)
	if apierrs.IsNotFound(err) {
		req.Status.SetUnavailable(dpName)
		// can not find the old deployment, create a new one
		deployment = r.buildDeployment(req, dpName)
		deployment, err = r.K8ClientSet.AppsV1().Deployments(ns).Create(deployment)
		if err != nil {
			return fmt.Errorf("Failed to create deployment in %q: %w", dpName, err)
		}
	} else if err != nil {
		// meet error when querying deployment
		req.Status.SetUnavailable(dpName)
		return fmt.Errorf("Failed to query deployment %q: %w", deployment, err)
	} else {
		// check the new spec and old spec diff
		oldSpec := deployment.Spec
		newSpec := r.buildDeployment(req, dpName).Spec

		// new spec is different from old spec, update the deployment
		if !reflect.DeepEqual(oldSpec, newSpec) {
			deployment.Spec = newSpec
			deployment, err = r.K8ClientSet.AppsV1().Deployments(ns).Update(deployment)
			if err != nil {
				return fmt.Errorf("Failed to update the deployment %w", err)
			}
		} else {
			if r.checkDeployment(deployment) && !r.checkKNLearning(req) {
				req.Status.SetAvailable()
				req.Status.SetReady()
			}
		}
	}

	return nil
}

func (r *Reconciler) buildDeployment(req *v1.KNLearning, deploymentName string) *appsv1.Deployment {
	replica := req.Spec.Size
	labels := map[string]string{
		"app": deploymentName,
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: req.Namespace,
			OwnerReferences: []metav1.OwnerReference{
				*metav1.NewControllerRef(req, schema.GroupVersionKind{
					Group:   v1.SchemeGroupVersion.Group,
					Version: v1.SchemeGroupVersion.Version,
					Kind:    "KNLearning",
				}),
			},
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: replica,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: r.buildContainers(req),
				},
			},
		},
	}
}

func (r *Reconciler) buildContainers(req *v1.KNLearning) []corev1.Container {
	containerPorts := []corev1.ContainerPort{}
	for _, svcPort := range req.Spec.Ports {
		cport := corev1.ContainerPort{}
		cport.ContainerPort = svcPort.TargetPort.IntVal
		containerPorts = append(containerPorts, cport)
	}

	return []corev1.Container{
		{
			Name:            req.Name,
			Image:           req.Spec.Image,
			Resources:       req.Spec.Resources,
			Ports:           containerPorts,
			ImagePullPolicy: corev1.PullIfNotPresent,
		},
	}
}

func (r *Reconciler) checkDeployment(dep *appsv1.Deployment) bool {
	for _, c := range dep.Status.Conditions {
		if c.Type == appsv1.DeploymentAvailable && c.Status == corev1.ConditionTrue {
			return true
		}
	}

	return false
}

func (r *Reconciler) checkKNLearning(req *v1.KNLearning) bool {
	for _, c := range req.Status.Conditions {
		if c.Type == v1.Succeeded && c.Status == corev1.ConditionTrue {
			return true
		}
	}

	return false
}
