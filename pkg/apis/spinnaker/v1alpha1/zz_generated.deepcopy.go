// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerService) DeepCopyInto(out *SpinnakerService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerService.
func (in *SpinnakerService) DeepCopy() *SpinnakerService {
	if in == nil {
		return nil
	}
	out := new(SpinnakerService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpinnakerService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerServiceList) DeepCopyInto(out *SpinnakerServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SpinnakerService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerServiceList.
func (in *SpinnakerServiceList) DeepCopy() *SpinnakerServiceList {
	if in == nil {
		return nil
	}
	out := new(SpinnakerServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SpinnakerServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerServiceSpec) DeepCopyInto(out *SpinnakerServiceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerServiceSpec.
func (in *SpinnakerServiceSpec) DeepCopy() *SpinnakerServiceSpec {
	if in == nil {
		return nil
	}
	out := new(SpinnakerServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpinnakerServiceStatus) DeepCopyInto(out *SpinnakerServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpinnakerServiceStatus.
func (in *SpinnakerServiceStatus) DeepCopy() *SpinnakerServiceStatus {
	if in == nil {
		return nil
	}
	out := new(SpinnakerServiceStatus)
	in.DeepCopyInto(out)
	return out
}