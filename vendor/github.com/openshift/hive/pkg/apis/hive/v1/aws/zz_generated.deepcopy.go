// +build !ignore_autogenerated

// Code generated by main. DO NOT EDIT.

package aws

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EC2RootVolume) DeepCopyInto(out *EC2RootVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EC2RootVolume.
func (in *EC2RootVolume) DeepCopy() *EC2RootVolume {
	if in == nil {
		return nil
	}
	out := new(EC2RootVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachinePoolPlatform) DeepCopyInto(out *MachinePoolPlatform) {
	*out = *in
	if in.Zones != nil {
		in, out := &in.Zones, &out.Zones
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.EC2RootVolume = in.EC2RootVolume
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachinePoolPlatform.
func (in *MachinePoolPlatform) DeepCopy() *MachinePoolPlatform {
	if in == nil {
		return nil
	}
	out := new(MachinePoolPlatform)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Platform) DeepCopyInto(out *Platform) {
	*out = *in
	out.CredentialsSecretRef = in.CredentialsSecretRef
	if in.UserTags != nil {
		in, out := &in.UserTags, &out.UserTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Platform.
func (in *Platform) DeepCopy() *Platform {
	if in == nil {
		return nil
	}
	out := new(Platform)
	in.DeepCopyInto(out)
	return out
}
