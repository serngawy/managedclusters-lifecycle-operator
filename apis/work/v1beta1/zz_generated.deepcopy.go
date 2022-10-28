//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterWorkState) DeepCopyInto(out *ClusterWorkState) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterWorkState.
func (in *ClusterWorkState) DeepCopy() *ClusterWorkState {
	if in == nil {
		return nil
	}
	out := new(ClusterWorkState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterGroupWork) DeepCopyInto(out *ManagedClusterGroupWork) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterGroupWork.
func (in *ManagedClusterGroupWork) DeepCopy() *ManagedClusterGroupWork {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterGroupWork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterGroupWork) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterGroupWorkList) DeepCopyInto(out *ManagedClusterGroupWorkList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedClusterGroupWork, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterGroupWorkList.
func (in *ManagedClusterGroupWorkList) DeepCopy() *ManagedClusterGroupWorkList {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterGroupWorkList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ManagedClusterGroupWorkList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterGroupWorkSpec) DeepCopyInto(out *ManagedClusterGroupWorkSpec) {
	*out = *in
	out.Placement = in.Placement
	in.GenericPlacementFields.DeepCopyInto(&out.GenericPlacementFields)
	in.ManifestWork.DeepCopyInto(&out.ManifestWork)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterGroupWorkSpec.
func (in *ManagedClusterGroupWorkSpec) DeepCopy() *ManagedClusterGroupWorkSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterGroupWorkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedClusterGroupWorkStatus) DeepCopyInto(out *ManagedClusterGroupWorkStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Clusters != nil {
		in, out := &in.Clusters, &out.Clusters
		*out = make([]*ClusterWorkState, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ClusterWorkState)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedClusterGroupWorkStatus.
func (in *ManagedClusterGroupWorkStatus) DeepCopy() *ManagedClusterGroupWorkStatus {
	if in == nil {
		return nil
	}
	out := new(ManagedClusterGroupWorkStatus)
	in.DeepCopyInto(out)
	return out
}
