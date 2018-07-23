// +build !ignore_autogenerated

// Copyright Jetstack Ltd. See LICENSE for details.

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	wing "github.com/jetstack/tarmak/pkg/apis/wing"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_FileManifestSource_To_wing_FileManifestSource,
		Convert_wing_FileManifestSource_To_v1alpha1_FileManifestSource,
		Convert_v1alpha1_Instance_To_wing_Instance,
		Convert_wing_Instance_To_v1alpha1_Instance,
		Convert_v1alpha1_InstanceList_To_wing_InstanceList,
		Convert_wing_InstanceList_To_v1alpha1_InstanceList,
		Convert_v1alpha1_InstanceSpec_To_wing_InstanceSpec,
		Convert_wing_InstanceSpec_To_v1alpha1_InstanceSpec,
		Convert_v1alpha1_InstanceStatus_To_wing_InstanceStatus,
		Convert_wing_InstanceStatus_To_v1alpha1_InstanceStatus,
		Convert_v1alpha1_ManifestSource_To_wing_ManifestSource,
		Convert_wing_ManifestSource_To_v1alpha1_ManifestSource,
		Convert_v1alpha1_PuppetTarget_To_wing_PuppetTarget,
		Convert_wing_PuppetTarget_To_v1alpha1_PuppetTarget,
		Convert_v1alpha1_PuppetTargetList_To_wing_PuppetTargetList,
		Convert_wing_PuppetTargetList_To_v1alpha1_PuppetTargetList,
		Convert_v1alpha1_S3ManifestSource_To_wing_S3ManifestSource,
		Convert_wing_S3ManifestSource_To_v1alpha1_S3ManifestSource,
		Convert_v1alpha1_WingJob_To_wing_WingJob,
		Convert_wing_WingJob_To_v1alpha1_WingJob,
		Convert_v1alpha1_WingJobList_To_wing_WingJobList,
		Convert_wing_WingJobList_To_v1alpha1_WingJobList,
		Convert_v1alpha1_WingJobSpec_To_wing_WingJobSpec,
		Convert_wing_WingJobSpec_To_v1alpha1_WingJobSpec,
		Convert_v1alpha1_WingJobStatus_To_wing_WingJobStatus,
		Convert_wing_WingJobStatus_To_v1alpha1_WingJobStatus,
	)
}

func autoConvert_v1alpha1_FileManifestSource_To_wing_FileManifestSource(in *FileManifestSource, out *wing.FileManifestSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_v1alpha1_FileManifestSource_To_wing_FileManifestSource is an autogenerated conversion function.
func Convert_v1alpha1_FileManifestSource_To_wing_FileManifestSource(in *FileManifestSource, out *wing.FileManifestSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_FileManifestSource_To_wing_FileManifestSource(in, out, s)
}

func autoConvert_wing_FileManifestSource_To_v1alpha1_FileManifestSource(in *wing.FileManifestSource, out *FileManifestSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_wing_FileManifestSource_To_v1alpha1_FileManifestSource is an autogenerated conversion function.
func Convert_wing_FileManifestSource_To_v1alpha1_FileManifestSource(in *wing.FileManifestSource, out *FileManifestSource, s conversion.Scope) error {
	return autoConvert_wing_FileManifestSource_To_v1alpha1_FileManifestSource(in, out, s)
}

func autoConvert_v1alpha1_Instance_To_wing_Instance(in *Instance, out *wing.Instance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.InstanceID = in.InstanceID
	out.InstancePool = in.InstancePool
	out.Spec = (*wing.InstanceSpec)(unsafe.Pointer(in.Spec))
	out.Status = (*wing.InstanceStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_v1alpha1_Instance_To_wing_Instance is an autogenerated conversion function.
func Convert_v1alpha1_Instance_To_wing_Instance(in *Instance, out *wing.Instance, s conversion.Scope) error {
	return autoConvert_v1alpha1_Instance_To_wing_Instance(in, out, s)
}

func autoConvert_wing_Instance_To_v1alpha1_Instance(in *wing.Instance, out *Instance, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.InstanceID = in.InstanceID
	out.InstancePool = in.InstancePool
	out.Spec = (*InstanceSpec)(unsafe.Pointer(in.Spec))
	out.Status = (*InstanceStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_wing_Instance_To_v1alpha1_Instance is an autogenerated conversion function.
func Convert_wing_Instance_To_v1alpha1_Instance(in *wing.Instance, out *Instance, s conversion.Scope) error {
	return autoConvert_wing_Instance_To_v1alpha1_Instance(in, out, s)
}

func autoConvert_v1alpha1_InstanceList_To_wing_InstanceList(in *InstanceList, out *wing.InstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]wing.Instance)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_InstanceList_To_wing_InstanceList is an autogenerated conversion function.
func Convert_v1alpha1_InstanceList_To_wing_InstanceList(in *InstanceList, out *wing.InstanceList, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceList_To_wing_InstanceList(in, out, s)
}

func autoConvert_wing_InstanceList_To_v1alpha1_InstanceList(in *wing.InstanceList, out *InstanceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Instance)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_wing_InstanceList_To_v1alpha1_InstanceList is an autogenerated conversion function.
func Convert_wing_InstanceList_To_v1alpha1_InstanceList(in *wing.InstanceList, out *InstanceList, s conversion.Scope) error {
	return autoConvert_wing_InstanceList_To_v1alpha1_InstanceList(in, out, s)
}

func autoConvert_v1alpha1_InstanceSpec_To_wing_InstanceSpec(in *InstanceSpec, out *wing.InstanceSpec, s conversion.Scope) error {
	out.PuppetTargetRef = in.PuppetTargetRef
	return nil
}

// Convert_v1alpha1_InstanceSpec_To_wing_InstanceSpec is an autogenerated conversion function.
func Convert_v1alpha1_InstanceSpec_To_wing_InstanceSpec(in *InstanceSpec, out *wing.InstanceSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceSpec_To_wing_InstanceSpec(in, out, s)
}

func autoConvert_wing_InstanceSpec_To_v1alpha1_InstanceSpec(in *wing.InstanceSpec, out *InstanceSpec, s conversion.Scope) error {
	out.PuppetTargetRef = in.PuppetTargetRef
	return nil
}

// Convert_wing_InstanceSpec_To_v1alpha1_InstanceSpec is an autogenerated conversion function.
func Convert_wing_InstanceSpec_To_v1alpha1_InstanceSpec(in *wing.InstanceSpec, out *InstanceSpec, s conversion.Scope) error {
	return autoConvert_wing_InstanceSpec_To_v1alpha1_InstanceSpec(in, out, s)
}

func autoConvert_v1alpha1_InstanceStatus_To_wing_InstanceStatus(in *InstanceStatus, out *wing.InstanceStatus, s conversion.Scope) error {
	out.PuppetTargetRef = in.PuppetTargetRef
	out.Converged = in.Converged
	return nil
}

// Convert_v1alpha1_InstanceStatus_To_wing_InstanceStatus is an autogenerated conversion function.
func Convert_v1alpha1_InstanceStatus_To_wing_InstanceStatus(in *InstanceStatus, out *wing.InstanceStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InstanceStatus_To_wing_InstanceStatus(in, out, s)
}

func autoConvert_wing_InstanceStatus_To_v1alpha1_InstanceStatus(in *wing.InstanceStatus, out *InstanceStatus, s conversion.Scope) error {
	out.PuppetTargetRef = in.PuppetTargetRef
	out.Converged = in.Converged
	return nil
}

// Convert_wing_InstanceStatus_To_v1alpha1_InstanceStatus is an autogenerated conversion function.
func Convert_wing_InstanceStatus_To_v1alpha1_InstanceStatus(in *wing.InstanceStatus, out *InstanceStatus, s conversion.Scope) error {
	return autoConvert_wing_InstanceStatus_To_v1alpha1_InstanceStatus(in, out, s)
}

func autoConvert_v1alpha1_ManifestSource_To_wing_ManifestSource(in *ManifestSource, out *wing.ManifestSource, s conversion.Scope) error {
	out.S3 = (*wing.S3ManifestSource)(unsafe.Pointer(in.S3))
	out.File = (*wing.FileManifestSource)(unsafe.Pointer(in.File))
	return nil
}

// Convert_v1alpha1_ManifestSource_To_wing_ManifestSource is an autogenerated conversion function.
func Convert_v1alpha1_ManifestSource_To_wing_ManifestSource(in *ManifestSource, out *wing.ManifestSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManifestSource_To_wing_ManifestSource(in, out, s)
}

func autoConvert_wing_ManifestSource_To_v1alpha1_ManifestSource(in *wing.ManifestSource, out *ManifestSource, s conversion.Scope) error {
	out.S3 = (*S3ManifestSource)(unsafe.Pointer(in.S3))
	out.File = (*FileManifestSource)(unsafe.Pointer(in.File))
	return nil
}

// Convert_wing_ManifestSource_To_v1alpha1_ManifestSource is an autogenerated conversion function.
func Convert_wing_ManifestSource_To_v1alpha1_ManifestSource(in *wing.ManifestSource, out *ManifestSource, s conversion.Scope) error {
	return autoConvert_wing_ManifestSource_To_v1alpha1_ManifestSource(in, out, s)
}

func autoConvert_v1alpha1_PuppetTarget_To_wing_PuppetTarget(in *PuppetTarget, out *wing.PuppetTarget, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ManifestSource_To_wing_ManifestSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	out.Hash = in.Hash
	return nil
}

// Convert_v1alpha1_PuppetTarget_To_wing_PuppetTarget is an autogenerated conversion function.
func Convert_v1alpha1_PuppetTarget_To_wing_PuppetTarget(in *PuppetTarget, out *wing.PuppetTarget, s conversion.Scope) error {
	return autoConvert_v1alpha1_PuppetTarget_To_wing_PuppetTarget(in, out, s)
}

func autoConvert_wing_PuppetTarget_To_v1alpha1_PuppetTarget(in *wing.PuppetTarget, out *PuppetTarget, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_wing_ManifestSource_To_v1alpha1_ManifestSource(&in.Source, &out.Source, s); err != nil {
		return err
	}
	out.Hash = in.Hash
	return nil
}

// Convert_wing_PuppetTarget_To_v1alpha1_PuppetTarget is an autogenerated conversion function.
func Convert_wing_PuppetTarget_To_v1alpha1_PuppetTarget(in *wing.PuppetTarget, out *PuppetTarget, s conversion.Scope) error {
	return autoConvert_wing_PuppetTarget_To_v1alpha1_PuppetTarget(in, out, s)
}

func autoConvert_v1alpha1_PuppetTargetList_To_wing_PuppetTargetList(in *PuppetTargetList, out *wing.PuppetTargetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]wing.PuppetTarget)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_PuppetTargetList_To_wing_PuppetTargetList is an autogenerated conversion function.
func Convert_v1alpha1_PuppetTargetList_To_wing_PuppetTargetList(in *PuppetTargetList, out *wing.PuppetTargetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_PuppetTargetList_To_wing_PuppetTargetList(in, out, s)
}

func autoConvert_wing_PuppetTargetList_To_v1alpha1_PuppetTargetList(in *wing.PuppetTargetList, out *PuppetTargetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]PuppetTarget)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_wing_PuppetTargetList_To_v1alpha1_PuppetTargetList is an autogenerated conversion function.
func Convert_wing_PuppetTargetList_To_v1alpha1_PuppetTargetList(in *wing.PuppetTargetList, out *PuppetTargetList, s conversion.Scope) error {
	return autoConvert_wing_PuppetTargetList_To_v1alpha1_PuppetTargetList(in, out, s)
}

func autoConvert_v1alpha1_S3ManifestSource_To_wing_S3ManifestSource(in *S3ManifestSource, out *wing.S3ManifestSource, s conversion.Scope) error {
	out.BucketName = in.BucketName
	out.Path = in.Path
	out.Region = in.Region
	return nil
}

// Convert_v1alpha1_S3ManifestSource_To_wing_S3ManifestSource is an autogenerated conversion function.
func Convert_v1alpha1_S3ManifestSource_To_wing_S3ManifestSource(in *S3ManifestSource, out *wing.S3ManifestSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_S3ManifestSource_To_wing_S3ManifestSource(in, out, s)
}

func autoConvert_wing_S3ManifestSource_To_v1alpha1_S3ManifestSource(in *wing.S3ManifestSource, out *S3ManifestSource, s conversion.Scope) error {
	out.BucketName = in.BucketName
	out.Path = in.Path
	out.Region = in.Region
	return nil
}

// Convert_wing_S3ManifestSource_To_v1alpha1_S3ManifestSource is an autogenerated conversion function.
func Convert_wing_S3ManifestSource_To_v1alpha1_S3ManifestSource(in *wing.S3ManifestSource, out *S3ManifestSource, s conversion.Scope) error {
	return autoConvert_wing_S3ManifestSource_To_v1alpha1_S3ManifestSource(in, out, s)
}

func autoConvert_v1alpha1_WingJob_To_wing_WingJob(in *WingJob, out *wing.WingJob, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Spec = (*wing.WingJobSpec)(unsafe.Pointer(in.Spec))
	out.Status = (*wing.WingJobStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_v1alpha1_WingJob_To_wing_WingJob is an autogenerated conversion function.
func Convert_v1alpha1_WingJob_To_wing_WingJob(in *WingJob, out *wing.WingJob, s conversion.Scope) error {
	return autoConvert_v1alpha1_WingJob_To_wing_WingJob(in, out, s)
}

func autoConvert_wing_WingJob_To_v1alpha1_WingJob(in *wing.WingJob, out *WingJob, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Spec = (*WingJobSpec)(unsafe.Pointer(in.Spec))
	out.Status = (*WingJobStatus)(unsafe.Pointer(in.Status))
	return nil
}

// Convert_wing_WingJob_To_v1alpha1_WingJob is an autogenerated conversion function.
func Convert_wing_WingJob_To_v1alpha1_WingJob(in *wing.WingJob, out *WingJob, s conversion.Scope) error {
	return autoConvert_wing_WingJob_To_v1alpha1_WingJob(in, out, s)
}

func autoConvert_v1alpha1_WingJobList_To_wing_WingJobList(in *WingJobList, out *wing.WingJobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]wing.WingJob)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1alpha1_WingJobList_To_wing_WingJobList is an autogenerated conversion function.
func Convert_v1alpha1_WingJobList_To_wing_WingJobList(in *WingJobList, out *wing.WingJobList, s conversion.Scope) error {
	return autoConvert_v1alpha1_WingJobList_To_wing_WingJobList(in, out, s)
}

func autoConvert_wing_WingJobList_To_v1alpha1_WingJobList(in *wing.WingJobList, out *WingJobList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]WingJob)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_wing_WingJobList_To_v1alpha1_WingJobList is an autogenerated conversion function.
func Convert_wing_WingJobList_To_v1alpha1_WingJobList(in *wing.WingJobList, out *WingJobList, s conversion.Scope) error {
	return autoConvert_wing_WingJobList_To_v1alpha1_WingJobList(in, out, s)
}

func autoConvert_v1alpha1_WingJobSpec_To_wing_WingJobSpec(in *WingJobSpec, out *wing.WingJobSpec, s conversion.Scope) error {
	out.InstanceName = in.InstanceName
	out.PuppetTargetRef = in.PuppetTargetRef
	out.Operation = in.Operation
	out.RequestTimestamp = in.RequestTimestamp
	return nil
}

// Convert_v1alpha1_WingJobSpec_To_wing_WingJobSpec is an autogenerated conversion function.
func Convert_v1alpha1_WingJobSpec_To_wing_WingJobSpec(in *WingJobSpec, out *wing.WingJobSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_WingJobSpec_To_wing_WingJobSpec(in, out, s)
}

func autoConvert_wing_WingJobSpec_To_v1alpha1_WingJobSpec(in *wing.WingJobSpec, out *WingJobSpec, s conversion.Scope) error {
	out.InstanceName = in.InstanceName
	out.PuppetTargetRef = in.PuppetTargetRef
	out.Operation = in.Operation
	out.RequestTimestamp = in.RequestTimestamp
	return nil
}

// Convert_wing_WingJobSpec_To_v1alpha1_WingJobSpec is an autogenerated conversion function.
func Convert_wing_WingJobSpec_To_v1alpha1_WingJobSpec(in *wing.WingJobSpec, out *WingJobSpec, s conversion.Scope) error {
	return autoConvert_wing_WingJobSpec_To_v1alpha1_WingJobSpec(in, out, s)
}

func autoConvert_v1alpha1_WingJobStatus_To_wing_WingJobStatus(in *WingJobStatus, out *wing.WingJobStatus, s conversion.Scope) error {
	out.Messages = in.Messages
	out.ExitCode = in.ExitCode
	out.Completed = in.Completed
	out.LastUpdateTimestamp = in.LastUpdateTimestamp
	return nil
}

// Convert_v1alpha1_WingJobStatus_To_wing_WingJobStatus is an autogenerated conversion function.
func Convert_v1alpha1_WingJobStatus_To_wing_WingJobStatus(in *WingJobStatus, out *wing.WingJobStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_WingJobStatus_To_wing_WingJobStatus(in, out, s)
}

func autoConvert_wing_WingJobStatus_To_v1alpha1_WingJobStatus(in *wing.WingJobStatus, out *WingJobStatus, s conversion.Scope) error {
	out.Messages = in.Messages
	out.ExitCode = in.ExitCode
	out.Completed = in.Completed
	out.LastUpdateTimestamp = in.LastUpdateTimestamp
	return nil
}

// Convert_wing_WingJobStatus_To_v1alpha1_WingJobStatus is an autogenerated conversion function.
func Convert_wing_WingJobStatus_To_v1alpha1_WingJobStatus(in *wing.WingJobStatus, out *WingJobStatus, s conversion.Scope) error {
	return autoConvert_wing_WingJobStatus_To_v1alpha1_WingJobStatus(in, out, s)
}
