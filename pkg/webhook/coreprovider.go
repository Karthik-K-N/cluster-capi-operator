package webhook

import (
	"context"
	"errors"

	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/cluster-api-operator/api/v1alpha1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

type CoreProviderWebhook struct {
}

func (r *CoreProviderWebhook) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		WithValidator(r).
		For(&v1alpha1.CoreProvider{}).
		Complete()
}

var _ webhook.CustomValidator = &CoreProviderWebhook{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *CoreProviderWebhook) ValidateCreate(ctx context.Context, obj runtime.Object) error {
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *CoreProviderWebhook) ValidateUpdate(ctx context.Context, oldObj, newObj runtime.Object) error {
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *CoreProviderWebhook) ValidateDelete(_ context.Context, obj runtime.Object) error {
	return errors.New("deletion of infrastructure provider is not allowed")
}