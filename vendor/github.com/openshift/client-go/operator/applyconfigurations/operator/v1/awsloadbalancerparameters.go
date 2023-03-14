// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// AWSLoadBalancerParametersApplyConfiguration represents an declarative configuration of the AWSLoadBalancerParameters type for use
// with apply.
type AWSLoadBalancerParametersApplyConfiguration struct {
	Type                          *v1.AWSLoadBalancerType                             `json:"type,omitempty"`
	ClassicLoadBalancerParameters *AWSClassicLoadBalancerParametersApplyConfiguration `json:"classicLoadBalancer,omitempty"`
	NetworkLoadBalancerParameters *v1.AWSNetworkLoadBalancerParameters                `json:"networkLoadBalancer,omitempty"`
}

// AWSLoadBalancerParametersApplyConfiguration constructs an declarative configuration of the AWSLoadBalancerParameters type for use with
// apply.
func AWSLoadBalancerParameters() *AWSLoadBalancerParametersApplyConfiguration {
	return &AWSLoadBalancerParametersApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *AWSLoadBalancerParametersApplyConfiguration) WithType(value v1.AWSLoadBalancerType) *AWSLoadBalancerParametersApplyConfiguration {
	b.Type = &value
	return b
}

// WithClassicLoadBalancerParameters sets the ClassicLoadBalancerParameters field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ClassicLoadBalancerParameters field is set to the value of the last call.
func (b *AWSLoadBalancerParametersApplyConfiguration) WithClassicLoadBalancerParameters(value *AWSClassicLoadBalancerParametersApplyConfiguration) *AWSLoadBalancerParametersApplyConfiguration {
	b.ClassicLoadBalancerParameters = value
	return b
}

// WithNetworkLoadBalancerParameters sets the NetworkLoadBalancerParameters field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NetworkLoadBalancerParameters field is set to the value of the last call.
func (b *AWSLoadBalancerParametersApplyConfiguration) WithNetworkLoadBalancerParameters(value v1.AWSNetworkLoadBalancerParameters) *AWSLoadBalancerParametersApplyConfiguration {
	b.NetworkLoadBalancerParameters = &value
	return b
}
