package reconciler

import (
	"context"

	// k8s native interface
	"k8s.io/client-go/tools/cache"

	// knative develop interface
	"knative.dev/pkg/configmap"
	"knative.dev/pkg/controller"

	// knative injection interface
	kubek8client "knative.dev/pkg/client/injection/kube/client"
	kndeploymentinformer "knative.dev/pkg/client/injection/kube/informers/apps/v1/deployment"

	// knstart injection interface
	knlearninginfomer "knstart/pkg/client/injection/informers/operator/v1/knlearning"
	"knstart/pkg/client/injection/reconciler/operator/v1/knlearning"

	// knstart defined interface
	knstartv1 "knstart/pkg/apis/operator/v1"
)

func NewController(
	ctx context.Context,
	cmw configmap.Watcher,
) *controller.Impl {
	dpInformer := kndeploymentinformer.Get(ctx)
	knlearningInformer := knlearninginfomer.Get(ctx)
	r := &Reconciler{
		K8ClientSet:      kubek8client.Get(ctx),
		DeploymentLister: dpInformer.Lister(),
	}

	impl := knlearning.NewImpl(ctx, r)

	// binding all knlearning event to controller HandleAll
	// HandleAll contains 3 functions: add, update, delete
	// add, delete calls impl.Enqueue with the requirement
	// update only call impl.Enqueue with the new requirement, the old req is ignored.
	knlearningInformer.Informer().AddEventHandler(controller.HandleAll(impl.Enqueue))

	dpInformer.Informer().AddEventHandler(cache.FilteringResourceEventHandler{
		FilterFunc: controller.FilterGroupKind(knstartv1.Kind("KNLearning")),
		Handler:    controller.HandleAll(impl.EnqueueControllerOf),
	})

	return impl
}
