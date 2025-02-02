// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type ConfigSpec -type EndpointSpec -type IdentitySpec -type PeerSpecSpec -type PeerStatusSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package kubespan

import (
	"net/netip"
)

// DeepCopy generates a deep copy of ConfigSpec.
func (o ConfigSpec) DeepCopy() ConfigSpec {
	var cp ConfigSpec = o
	return cp
}

// DeepCopy generates a deep copy of EndpointSpec.
func (o EndpointSpec) DeepCopy() EndpointSpec {
	var cp EndpointSpec = o
	return cp
}

// DeepCopy generates a deep copy of IdentitySpec.
func (o IdentitySpec) DeepCopy() IdentitySpec {
	var cp IdentitySpec = o
	return cp
}

// DeepCopy generates a deep copy of PeerSpecSpec.
func (o PeerSpecSpec) DeepCopy() PeerSpecSpec {
	var cp PeerSpecSpec = o
	if o.AllowedIPs != nil {
		cp.AllowedIPs = make([]netip.Prefix, len(o.AllowedIPs))
		copy(cp.AllowedIPs, o.AllowedIPs)
	}
	if o.Endpoints != nil {
		cp.Endpoints = make([]netip.AddrPort, len(o.Endpoints))
		copy(cp.Endpoints, o.Endpoints)
	}
	return cp
}

// DeepCopy generates a deep copy of PeerStatusSpec.
func (o PeerStatusSpec) DeepCopy() PeerStatusSpec {
	var cp PeerStatusSpec = o
	return cp
}
