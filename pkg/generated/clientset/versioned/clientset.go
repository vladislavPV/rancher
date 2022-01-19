/*
Copyright 2022 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package versioned

import (
	"fmt"

	catalogv1 "github.com/rancher/rancher/pkg/generated/clientset/versioned/typed/catalog.cattle.io/v1"
	provisioningv1 "github.com/rancher/rancher/pkg/generated/clientset/versioned/typed/provisioning.cattle.io/v1"
	upgradev1 "github.com/rancher/rancher/pkg/generated/clientset/versioned/typed/upgrade.cattle.io/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	CatalogV1() catalogv1.CatalogV1Interface
	ProvisioningV1() provisioningv1.ProvisioningV1Interface
	UpgradeV1() upgradev1.UpgradeV1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	catalogV1      *catalogv1.CatalogV1Client
	provisioningV1 *provisioningv1.ProvisioningV1Client
	upgradeV1      *upgradev1.UpgradeV1Client
}

// CatalogV1 retrieves the CatalogV1Client
func (c *Clientset) CatalogV1() catalogv1.CatalogV1Interface {
	return c.catalogV1
}

// ProvisioningV1 retrieves the ProvisioningV1Client
func (c *Clientset) ProvisioningV1() provisioningv1.ProvisioningV1Interface {
	return c.provisioningV1
}

// UpgradeV1 retrieves the UpgradeV1Client
func (c *Clientset) UpgradeV1() upgradev1.UpgradeV1Interface {
	return c.upgradeV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.catalogV1, err = catalogv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.provisioningV1, err = provisioningv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.upgradeV1, err = upgradev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.catalogV1 = catalogv1.NewForConfigOrDie(c)
	cs.provisioningV1 = provisioningv1.NewForConfigOrDie(c)
	cs.upgradeV1 = upgradev1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.catalogV1 = catalogv1.New(c)
	cs.provisioningV1 = provisioningv1.New(c)
	cs.upgradeV1 = upgradev1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
