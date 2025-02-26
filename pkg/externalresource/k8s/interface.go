package k8s

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
)

const (
	// ExternalResourceResource defines the resource of ExternalResource
	ExternalResourceResource = "externals"

	// ExternalResourceKind defines the kind of the resource external
	ExternalResourceKind = "External"
)

type externalClient struct {
	restClient rest.Interface
	scheme     *runtime.Scheme
	ns         string
}

// ExternalResourceInterface defines the operations for the external resource item Kubernetes client
type ExternalResourceInterface interface {
	Update(ctx context.Context, external *External) (*External, error)
	Get(ctx context.Context, name string, options metav1.GetOptions) (*External, error)
	Create(ctx context.Context, external *External, options metav1.CreateOptions) (*External, error)
	List(ctx context.Context, options metav1.ListOptions) (*ExternalList, error)
}

func (c *externalClient) Create(ctx context.Context, external *External, _ metav1.CreateOptions) (*External, error) {
	result := External{}
	err := c.restClient.
		Post().
		Namespace(c.ns).
		Resource(ExternalResourceResource).
		Body(external).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *externalClient) Get(ctx context.Context, name string, opts metav1.GetOptions) (*External, error) {
	result := External{}
	err := c.restClient.
		Get().
		Namespace(c.ns).
		Resource(ExternalResourceResource).
		Name(name).
		VersionedParams(&opts, runtime.NewParameterCodec(c.scheme)).
		Do(ctx).
		Into(&result)
	return &result, err
}

func (c *externalClient) Update(ctx context.Context, external *External) (*External, error) {
	result := External{}
	err := c.restClient.
		Put().
		Namespace(c.ns).
		Resource(ExternalResourceResource).
		Name(external.Name).
		Body(external).
		Do(ctx).
		Into(&result)

	return &result, err
}

func (c *externalClient) List(ctx context.Context, opts metav1.ListOptions) (*ExternalList, error) {
	result := ExternalList{}
	err := c.restClient.Get().
		Namespace(c.ns).
		Resource(ExternalResourceResource).
		VersionedParams(&opts, runtime.NewParameterCodec(c.scheme)).
		Do(ctx).
		Into(&result)
	return &result, err
}
