package example

import (
	"io"

	cloudprovider "k8s.io/cloud-provider"
)

const providerName string = "example"

type cloudProvider struct {
}

func (p *cloudProvider) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {

}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (p *cloudProvider) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false

}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (p *cloudProvider) Instances() (cloudprovider.Instances, bool) {
	return nil, false

}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (p *cloudProvider) Zones() (cloudprovider.Zones, bool) {
	return nil, false

}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (p *cloudProvider) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false

}

// Routes returns a routes interface along with whether the interface is supported.
func (p *cloudProvider) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (p *cloudProvider) ProviderName() string {
	return providerName

}

// HasClusterID returns true if a ClusterID is required and set
func (p *cloudProvider) HasClusterID() bool {
	return false
}

func newCloudProvider() (cloudprovider.Interface, error) {

	return &cloudProvider{}, nil
}

func init() {
	cloudprovider.RegisterCloudProvider(providerName, func(io.Reader) (cloudprovider.Interface, error) {
		return newCloudProvider()
	})
}
