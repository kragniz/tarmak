# Calico

#### Table of Contents

1. [Description](#description)
1. [Setup - The basics of getting started with calico](#setup)
    * [What calico does](#what-etcd-affects)
    * [Setup requirements](#setup-requirements)
    * [Beginning with calico](#beginning-with-etcd)
1. [Usage - Configuration options and additional functionality](#usage)
1. [Reference - An under-the-hood peek at what the module is doing and how](#reference)
1. [Limitations - OS compatibility, etc.](#limitations)
1. [Development - Guide for contributing to the module](#development)

## Description

This module installs [Calico](https://www.projectcalico.org/) on servers running a
binary distribution of Kubernetes and was initially developed to do this on CentOS 7 
running across multiple AZs on AWS EC2.

It should be parameterised and customisable in such a way that use on other cloud
providers (or bare metal) are in scope. It's probably harder to use this module for
non-Kubernetes orchestrators, at least for now.

## Setup

### What calico does

* Installs the calico and calico-ipam binaries used for CNI (Container Network Integration).
* Installs the 'localhost' CNI binary used by the kubelet.
* Installs a systemd service that runs the calico-node container that creates the routing mesh
* Configures all of these to communicate with a (dedicated or shared) etcd cluster
* **Optionally** disables AWS source/destination checking on the instance (given correct IAM permissions)
* **Optionally** hacks the shipped calico filters to use L2 networking within the AZ and L3 between AZs
* **Can** be used to add an IP pool for calico (via the ip_pool defined function)
* **Can** be used to deploy the calico policy controller for Kubernetes, which is necessary 
(via the policy_controller class)
* It requires (and includes, by way of ensure_packages) docker

### Setup Requirements **OPTIONAL**

If your module requires anything extra before setting up (pluginsync enabled,
etc.), mention it here.

If your most recent release breaks compatibility or requires particular steps
for upgrading, you might want to include an additional "Upgrading" section
here.

### Beginning with calico

```puppet
class { 'calico':
  etcd_cluster      => [ 'etcd1.mydomain', 'etcd2.mydomain', 'etcd3.mydomain' ],
  aws               => true,
  aws_filter_hack   => true,
  tls               => false,
  etcd_overlay_port => 2379,
}

calico::ip_pool { '10.234.235.0/24':
  ip_pool      => '10.234.235.0',
  ip_mask      => 24,
  ipip_enabled => 'false',
}

class { 'calico::policy_controller':
  require => Service['kubernetes_apiserver']
}
```

## Usage

This section is where you describe how to customize, configure, and do the
fancy stuff with your module here. It's especially helpful if you include usage
examples and code samples for doing things with your module.

## Reference

Here, include a complete list of your module's classes, types, providers,
facts, along with the parameters for each. Users refer to this section (thus
the name "Reference") to find specific details; most users don't read it per
se.

## Limitations

This is where you list OS compatibility, version compatibility, etc. If there
are Known Issues, you might want to include them under their own heading here.

## Development

Since your module is awesome, other users will want to play with it. Let them
know what the ground rules for contributing are.

## Release Notes/Contributors/Etc. **Optional**

If you aren't using changelog, put your release notes here (though you should
consider using changelog). You can also add any additional sections you feel
are necessary or important to include here. Please use the `## ` header.
