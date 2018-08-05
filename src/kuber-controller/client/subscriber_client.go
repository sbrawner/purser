package client

import (
	"kuber-controller/crd"

	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type SubscriberCrdClient struct {
	cl     *rest.RESTClient
	ns     string
	plural string
	codec  runtime.ParameterCodec
}

// This file implement all the (CRUD) client methods we need to access Subscriber CRD object
func CreateSubscriberCrdClient(cl *rest.RESTClient, scheme *runtime.Scheme, namespace string) *SubscriberCrdClient {
	return &SubscriberCrdClient{cl: cl, ns: namespace, plural: crd.SubscriberPlural,
		codec: runtime.NewParameterCodec(scheme)}
}

func (f *SubscriberCrdClient) CreateSubscriber(obj *crd.Subscriber) (*crd.Subscriber, error) {
	var result crd.Subscriber
	err := f.cl.Post().
		Namespace(f.ns).Resource(f.plural).
		Body(obj).Do().Into(&result)
	return &result, err
}

func (f *SubscriberCrdClient) UpdateSubscriber(obj *crd.Subscriber) (*crd.Subscriber, error) {
	var result crd.Subscriber
	err := f.cl.
	//Put().
		Put().Name((obj.Name)).
	//Patch(types.JSONPatchType).Name(obj.Name).
		Namespace(f.ns).Resource(f.plural).
		Body(obj).Do().Into(&result)
	return &result, err
}

func (f *SubscriberCrdClient) DeleteSubscriber(name string, options *meta_v1.DeleteOptions) error {
	return f.cl.Delete().
		Namespace(f.ns).Resource(f.plural).
		Name(name).Body(options).Do().
		Error()
}

func (f *SubscriberCrdClient) GetSubscriber(name string) (*crd.Subscriber, error) {
	var result crd.Subscriber
	err := f.cl.Get().
		Namespace(f.ns).Resource(f.plural).
		Name(name).Do().Into(&result)
	return &result, err
}

func (f *SubscriberCrdClient) ListSubscriber(opts meta_v1.ListOptions) (*crd.SubscriberList, error) {
	var result crd.SubscriberList
	err := f.cl.Get().
		Namespace(f.ns).Resource(f.plural).
		VersionedParams(&opts, f.codec).
		Do().Into(&result)
	return &result, err
}

// Create a new List watch for our TPR
func (f *SubscriberCrdClient) NewListWatchSubscriber() *cache.ListWatch {
	return cache.NewListWatchFromClient(f.cl, f.plural, f.ns, fields.Everything())
}