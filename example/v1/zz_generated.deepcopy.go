//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CELUnion) DeepCopyInto(out *CELUnion) {
	*out = *in
	if in.RequiredMember != nil {
		in, out := &in.RequiredMember, &out.RequiredMember
		*out = new(string)
		**out = **in
	}
	if in.OptionalMember != nil {
		in, out := &in.OptionalMember, &out.OptionalMember
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CELUnion.
func (in *CELUnion) DeepCopy() *CELUnion {
	if in == nil {
		return nil
	}
	out := new(CELUnion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EvolvingUnion) DeepCopyInto(out *EvolvingUnion) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EvolvingUnion.
func (in *EvolvingUnion) DeepCopy() *EvolvingUnion {
	if in == nil {
		return nil
	}
	out := new(EvolvingUnion)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StableConfigType) DeepCopyInto(out *StableConfigType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StableConfigType.
func (in *StableConfigType) DeepCopy() *StableConfigType {
	if in == nil {
		return nil
	}
	out := new(StableConfigType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StableConfigType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StableConfigTypeList) DeepCopyInto(out *StableConfigTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StableConfigType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StableConfigTypeList.
func (in *StableConfigTypeList) DeepCopy() *StableConfigTypeList {
	if in == nil {
		return nil
	}
	out := new(StableConfigTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StableConfigTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StableConfigTypeSpec) DeepCopyInto(out *StableConfigTypeSpec) {
	*out = *in
	out.EvolvingUnion = in.EvolvingUnion
	in.CELUnion.DeepCopyInto(&out.CELUnion)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StableConfigTypeSpec.
func (in *StableConfigTypeSpec) DeepCopy() *StableConfigTypeSpec {
	if in == nil {
		return nil
	}
	out := new(StableConfigTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StableConfigTypeStatus) DeepCopyInto(out *StableConfigTypeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StableConfigTypeStatus.
func (in *StableConfigTypeStatus) DeepCopy() *StableConfigTypeStatus {
	if in == nil {
		return nil
	}
	out := new(StableConfigTypeStatus)
	in.DeepCopyInto(out)
	return out
}
