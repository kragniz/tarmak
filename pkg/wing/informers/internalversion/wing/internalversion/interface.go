// Copyright Jetstack Ltd. See LICENSE for details.

// This file was automatically generated by informer-gen

package internalversion

import (
	internalinterfaces "github.com/jetstack/tarmak/pkg/wing/informers/internalversion/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Instances returns a InstanceInformer.
	Instances() InstanceInformer
	// Machines returns a MachineInformer.
	Machines() MachineInformer
	// PuppetTargets returns a PuppetTargetInformer.
	PuppetTargets() PuppetTargetInformer
	// WingJobs returns a WingJobInformer.
	WingJobs() WingJobInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Instances returns a InstanceInformer.
func (v *version) Instances() InstanceInformer {
	return &instanceInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Machines returns a MachineInformer.
func (v *version) Machines() MachineInformer {
	return &machineInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// PuppetTargets returns a PuppetTargetInformer.
func (v *version) PuppetTargets() PuppetTargetInformer {
	return &puppetTargetInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// WingJobs returns a WingJobInformer.
func (v *version) WingJobs() WingJobInformer {
	return &wingJobInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}
