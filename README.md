# Kubernetes Cloud Controller Manager for Example

A simple kubernetes cloud controller manager for example.


## Developing Cloud Controller Manager Guide

### [1/3] implement interface below

```go

// Interface is an abstract, pluggable interface for cloud providers.
type Interface interface {
	// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
	// to perform housekeeping or run custom controllers specific to the cloud provider.
	// Any tasks started here should be cleaned up when the stop channel closes.
	Initialize(clientBuilder ControllerClientBuilder, stop <-chan struct{})
	// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
	LoadBalancer() (LoadBalancer, bool)
	// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
	Instances() (Instances, bool)
	// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
	Zones() (Zones, bool)
	// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
	Clusters() (Clusters, bool)
	// Routes returns a routes interface along with whether the interface is supported.
	Routes() (Routes, bool)
	// ProviderName returns the cloud provider ID.
	ProviderName() string
	// HasClusterID returns true if a ClusterID is required and set
	HasClusterID() bool
}

```


### [2/3] register provider

register provider in package.

```go

func init() {
	cloudprovider.RegisterCloudProvider(providerName, func(io.Reader) (cloudprovider.Interface, error) {
		return newCloudProvider()
	})
}
```

### [3/3] deploy cloud controller manager

kubectl apply 


## Cloud Controller Manager mechanism

cloud controller manager 是一个非常简单的中间件（或者说是客户端）， 将云厂商提供的 API 映射到 K8S  的对象，有点像 ORM 了。



## Reference

- https://kubernetes.io/docs/tasks/administer-cluster/developing-cloud-controller-manager/
- https://github.com/digitalocean/digitalocean-cloud-controller-manager/









