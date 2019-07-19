// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManager) DeepCopyInto(out *CertManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManager.
func (in *CertManager) DeepCopy() *CertManager {
	if in == nil {
		return nil
	}
	out := new(CertManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerList) DeepCopyInto(out *CertManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CertManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerList.
func (in *CertManagerList) DeepCopy() *CertManagerList {
	if in == nil {
		return nil
	}
	out := new(CertManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CertManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerServiceAccount) DeepCopyInto(out *CertManagerServiceAccount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerServiceAccount.
func (in *CertManagerServiceAccount) DeepCopy() *CertManagerServiceAccount {
	if in == nil {
		return nil
	}
	out := new(CertManagerServiceAccount)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerSpec) DeepCopyInto(out *CertManagerSpec) {
	*out = *in
	out.ServiceAccount = in.ServiceAccount
	out.Image = in.Image
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerSpec.
func (in *CertManagerSpec) DeepCopy() *CertManagerSpec {
	if in == nil {
		return nil
	}
	out := new(CertManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertManagerStatus) DeepCopyInto(out *CertManagerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertManagerStatus.
func (in *CertManagerStatus) DeepCopy() *CertManagerStatus {
	if in == nil {
		return nil
	}
	out := new(CertManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionManager) DeepCopyInto(out *ConnectionManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionManager.
func (in *ConnectionManager) DeepCopy() *ConnectionManager {
	if in == nil {
		return nil
	}
	out := new(ConnectionManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectionManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionManagerList) DeepCopyInto(out *ConnectionManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConnectionManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionManagerList.
func (in *ConnectionManagerList) DeepCopy() *ConnectionManagerList {
	if in == nil {
		return nil
	}
	out := new(ConnectionManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectionManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionManagerSpec) DeepCopyInto(out *ConnectionManagerSpec) {
	*out = *in
	if in.BootStrapConfig != nil {
		in, out := &in.BootStrapConfig, &out.BootStrapConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionManagerSpec.
func (in *ConnectionManagerSpec) DeepCopy() *ConnectionManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectionManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionManagerStatus) DeepCopyInto(out *ConnectionManagerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionManagerStatus.
func (in *ConnectionManagerStatus) DeepCopy() *ConnectionManagerStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectionManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeployableConfig) DeepCopyInto(out *DeployableConfig) {
	*out = *in
	out.Image = in.Image
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeployableConfig.
func (in *DeployableConfig) DeepCopy() *DeployableConfig {
	if in == nil {
		return nil
	}
	out := new(DeployableConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletConnectionManagerSpec) DeepCopyInto(out *KlusterletConnectionManagerSpec) {
	*out = *in
	if in.BootStrapConfig != nil {
		in, out := &in.BootStrapConfig, &out.BootStrapConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletConnectionManagerSpec.
func (in *KlusterletConnectionManagerSpec) DeepCopy() *KlusterletConnectionManagerSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletConnectionManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletPrometheusIntegrationSpec) DeepCopyInto(out *KlusterletPrometheusIntegrationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletPrometheusIntegrationSpec.
func (in *KlusterletPrometheusIntegrationSpec) DeepCopy() *KlusterletPrometheusIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletPrometheusIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletService) DeepCopyInto(out *KlusterletService) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletService.
func (in *KlusterletService) DeepCopy() *KlusterletService {
	if in == nil {
		return nil
	}
	out := new(KlusterletService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KlusterletService) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletServiceList) DeepCopyInto(out *KlusterletServiceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KlusterletService, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletServiceList.
func (in *KlusterletServiceList) DeepCopy() *KlusterletServiceList {
	if in == nil {
		return nil
	}
	out := new(KlusterletServiceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KlusterletServiceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletServiceSpec) DeepCopyInto(out *KlusterletServiceSpec) {
	*out = *in
	if in.ClusterLabels != nil {
		in, out := &in.ClusterLabels, &out.ClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.BootStrapConfig != nil {
		in, out := &in.BootStrapConfig, &out.BootStrapConfig
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.TillerIntegration = in.TillerIntegration
	out.PrometheusIntegration = in.PrometheusIntegration
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletServiceSpec.
func (in *KlusterletServiceSpec) DeepCopy() *KlusterletServiceSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletServiceStatus) DeepCopyInto(out *KlusterletServiceStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletServiceStatus.
func (in *KlusterletServiceStatus) DeepCopy() *KlusterletServiceStatus {
	if in == nil {
		return nil
	}
	out := new(KlusterletServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletTillerIntegrationSpec) DeepCopyInto(out *KlusterletTillerIntegrationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletTillerIntegrationSpec.
func (in *KlusterletTillerIntegrationSpec) DeepCopy() *KlusterletTillerIntegrationSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletTillerIntegrationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KlusterletWorkManagerSpec) DeepCopyInto(out *KlusterletWorkManagerSpec) {
	*out = *in
	if in.ClusterLabels != nil {
		in, out := &in.ClusterLabels, &out.ClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KlusterletWorkManagerSpec.
func (in *KlusterletWorkManagerSpec) DeepCopy() *KlusterletWorkManagerSpec {
	if in == nil {
		return nil
	}
	out := new(KlusterletWorkManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tiller) DeepCopyInto(out *Tiller) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tiller.
func (in *Tiller) DeepCopy() *Tiller {
	if in == nil {
		return nil
	}
	out := new(Tiller)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Tiller) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TillerList) DeepCopyInto(out *TillerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Tiller, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TillerList.
func (in *TillerList) DeepCopy() *TillerList {
	if in == nil {
		return nil
	}
	out := new(TillerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TillerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TillerSpec) DeepCopyInto(out *TillerSpec) {
	*out = *in
	out.Image = in.Image
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TillerSpec.
func (in *TillerSpec) DeepCopy() *TillerSpec {
	if in == nil {
		return nil
	}
	out := new(TillerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TillerStatus) DeepCopyInto(out *TillerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TillerStatus.
func (in *TillerStatus) DeepCopy() *TillerStatus {
	if in == nil {
		return nil
	}
	out := new(TillerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManager) DeepCopyInto(out *WorkManager) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManager.
func (in *WorkManager) DeepCopy() *WorkManager {
	if in == nil {
		return nil
	}
	out := new(WorkManager)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkManager) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerConfig) DeepCopyInto(out *WorkManagerConfig) {
	*out = *in
	out.Image = in.Image
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerConfig.
func (in *WorkManagerConfig) DeepCopy() *WorkManagerConfig {
	if in == nil {
		return nil
	}
	out := new(WorkManagerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerIngress) DeepCopyInto(out *WorkManagerIngress) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerIngress.
func (in *WorkManagerIngress) DeepCopy() *WorkManagerIngress {
	if in == nil {
		return nil
	}
	out := new(WorkManagerIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerList) DeepCopyInto(out *WorkManagerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkManager, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerList.
func (in *WorkManagerList) DeepCopy() *WorkManagerList {
	if in == nil {
		return nil
	}
	out := new(WorkManagerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkManagerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerPrometheusIntegration) DeepCopyInto(out *WorkManagerPrometheusIntegration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerPrometheusIntegration.
func (in *WorkManagerPrometheusIntegration) DeepCopy() *WorkManagerPrometheusIntegration {
	if in == nil {
		return nil
	}
	out := new(WorkManagerPrometheusIntegration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerService) DeepCopyInto(out *WorkManagerService) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerService.
func (in *WorkManagerService) DeepCopy() *WorkManagerService {
	if in == nil {
		return nil
	}
	out := new(WorkManagerService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerSpec) DeepCopyInto(out *WorkManagerSpec) {
	*out = *in
	if in.ClusterLabels != nil {
		in, out := &in.ClusterLabels, &out.ClusterLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.TillerIntegration = in.TillerIntegration
	out.PrometheusIntegration = in.PrometheusIntegration
	out.Service = in.Service
	out.Ingress = in.Ingress
	out.WorkManagerConfig = in.WorkManagerConfig
	out.DeployableConfig = in.DeployableConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerSpec.
func (in *WorkManagerSpec) DeepCopy() *WorkManagerSpec {
	if in == nil {
		return nil
	}
	out := new(WorkManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerStatus) DeepCopyInto(out *WorkManagerStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerStatus.
func (in *WorkManagerStatus) DeepCopy() *WorkManagerStatus {
	if in == nil {
		return nil
	}
	out := new(WorkManagerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkManagerTillerIntegration) DeepCopyInto(out *WorkManagerTillerIntegration) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkManagerTillerIntegration.
func (in *WorkManagerTillerIntegration) DeepCopy() *WorkManagerTillerIntegration {
	if in == nil {
		return nil
	}
	out := new(WorkManagerTillerIntegration)
	in.DeepCopyInto(out)
	return out
}
