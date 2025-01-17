/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type AttachedDisksObservation struct {
}

type AttachedDisksParameters struct {

	// +kubebuilder:validation:Optional
	Managed *bool `json:"managed,omitempty" tf:"managed,omitempty"`

	// +kubebuilder:validation:Required
	SizeInGb *float64 `json:"sizeInGb" tf:"size_in_gb,omitempty"`
}

type GpuConfigObservation struct {
}

type GpuConfigParameters struct {

	// +kubebuilder:validation:Optional
	Addresses map[string]*string `json:"addresses,omitempty" tf:"addresses,omitempty"`

	// +kubebuilder:validation:Required
	DeviceModel *string `json:"deviceModel" tf:"device_model,omitempty"`

	// +kubebuilder:validation:Required
	NumGpus *float64 `json:"numGpus" tf:"num_gpus,omitempty"`

	// +kubebuilder:validation:Required
	Vendor *string `json:"vendor" tf:"vendor,omitempty"`
}

type LibvirtBackupPolicyObservation struct {
}

type LibvirtBackupPolicyParameters struct {

	// +kubebuilder:validation:Required
	BackupLocationID *string `json:"backupLocationId" tf:"backup_location_id,omitempty"`

	// +kubebuilder:validation:Required
	ExpiryInHour *float64 `json:"expiryInHour" tf:"expiry_in_hour,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeClusterResources *bool `json:"includeClusterResources,omitempty" tf:"include_cluster_resources,omitempty"`

	// +kubebuilder:validation:Optional
	IncludeDisks *bool `json:"includeDisks,omitempty" tf:"include_disks,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []*string `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Required
	Prefix *string `json:"prefix" tf:"prefix,omitempty"`

	// +kubebuilder:validation:Required
	Schedule *string `json:"schedule" tf:"schedule,omitempty"`
}

type LibvirtCloudConfigObservation struct {
}

type LibvirtCloudConfigParameters struct {

	// +kubebuilder:validation:Optional
	NtpServers []*string `json:"ntpServers,omitempty" tf:"ntp_servers,omitempty"`

	// +kubebuilder:validation:Required
	SSHKey *string `json:"sshKey" tf:"ssh_key,omitempty"`

	// +kubebuilder:validation:Required
	Vip *string `json:"vip" tf:"vip,omitempty"`
}

type LibvirtClusterProfileObservation struct {
}

type LibvirtClusterProfilePackManifestObservation struct {
}

type LibvirtClusterProfilePackManifestParameters struct {

	// +kubebuilder:validation:Required
	Content *string `json:"content" tf:"content,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type LibvirtClusterProfilePackObservation struct {
}

type LibvirtClusterProfilePackParameters struct {

	// +kubebuilder:validation:Optional
	Manifest []LibvirtClusterProfilePackManifestParameters `json:"manifest,omitempty" tf:"manifest,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Required
	Values *string `json:"values" tf:"values,omitempty"`
}

type LibvirtClusterProfileParameters struct {

	// +kubebuilder:validation:Required
	ID *string `json:"id" tf:"id,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []LibvirtClusterProfilePackParameters `json:"pack,omitempty" tf:"pack,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

type LibvirtClusterRbacBindingObservation struct {
}

type LibvirtClusterRbacBindingParameters struct {

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Optional
	Role map[string]*string `json:"role,omitempty" tf:"role,omitempty"`

	// +kubebuilder:validation:Optional
	Subjects []LibvirtClusterRbacBindingSubjectsParameters `json:"subjects,omitempty" tf:"subjects,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LibvirtClusterRbacBindingSubjectsObservation struct {
}

type LibvirtClusterRbacBindingSubjectsParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	Namespace *string `json:"namespace,omitempty" tf:"namespace,omitempty"`

	// +kubebuilder:validation:Required
	Type *string `json:"type" tf:"type,omitempty"`
}

type LibvirtHostConfigObservation struct {
}

type LibvirtHostConfigParameters struct {

	// The external traffic policy for the cluster.
	// +kubebuilder:validation:Optional
	ExternalTrafficPolicy *string `json:"externalTrafficPolicy,omitempty" tf:"external_traffic_policy,omitempty"`

	// The type of endpoint for the cluster. Can be either 'Ingress' or 'LoadBalancer'. The default is 'Ingress'.
	// +kubebuilder:validation:Optional
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// The host for the Ingress endpoint. Required if 'host_endpoint_type' is set to 'Ingress'.
	// +kubebuilder:validation:Optional
	IngressHost *string `json:"ingressHost,omitempty" tf:"ingress_host,omitempty"`

	// The source ranges for the load balancer. Required if 'host_endpoint_type' is set to 'LoadBalancer'.
	// +kubebuilder:validation:Optional
	LoadBalancerSourceRanges *string `json:"loadBalancerSourceRanges,omitempty" tf:"load_balancer_source_ranges,omitempty"`
}

type LibvirtLocationConfigObservation struct {
}

type LibvirtLocationConfigParameters struct {

	// The country code of the country the cluster is located in.
	// +kubebuilder:validation:Optional
	CountryCode *string `json:"countryCode,omitempty" tf:"country_code,omitempty"`

	// The name of the country.
	// +kubebuilder:validation:Optional
	CountryName *string `json:"countryName,omitempty" tf:"country_name,omitempty"`

	// The latitude coordinates value.
	// +kubebuilder:validation:Required
	Latitude *float64 `json:"latitude" tf:"latitude,omitempty"`

	// The longitude coordinates value.
	// +kubebuilder:validation:Required
	Longitude *float64 `json:"longitude" tf:"longitude,omitempty"`

	// The region code of where the cluster is located in.
	// +kubebuilder:validation:Optional
	RegionCode *string `json:"regionCode,omitempty" tf:"region_code,omitempty"`

	// The name of the region.
	// +kubebuilder:validation:Optional
	RegionName *string `json:"regionName,omitempty" tf:"region_name,omitempty"`
}

type LibvirtMachinePoolObservation struct {
}

type LibvirtMachinePoolParameters struct {

	// +kubebuilder:validation:Optional
	AdditionalLabels map[string]*string `json:"additionalLabels,omitempty" tf:"additional_labels,omitempty"`

	// +kubebuilder:validation:Optional
	ControlPlane *bool `json:"controlPlane,omitempty" tf:"control_plane,omitempty"`

	// +kubebuilder:validation:Optional
	ControlPlaneAsWorker *bool `json:"controlPlaneAsWorker,omitempty" tf:"control_plane_as_worker,omitempty"`

	// +kubebuilder:validation:Required
	Count *float64 `json:"count" tf:"count,omitempty"`

	// +kubebuilder:validation:Required
	InstanceType []MachinePoolInstanceTypeParameters `json:"instanceType" tf:"instance_type,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Placements []MachinePoolPlacementsParameters `json:"placements" tf:"placements,omitempty"`

	// +kubebuilder:validation:Optional
	Taints []LibvirtMachinePoolTaintsParameters `json:"taints,omitempty" tf:"taints,omitempty"`

	// +kubebuilder:validation:Optional
	UpdateStrategy *string `json:"updateStrategy,omitempty" tf:"update_strategy,omitempty"`
}

type LibvirtMachinePoolTaintsObservation struct {
}

type LibvirtMachinePoolTaintsParameters struct {

	// +kubebuilder:validation:Required
	Effect *string `json:"effect" tf:"effect,omitempty"`

	// +kubebuilder:validation:Required
	Key *string `json:"key" tf:"key,omitempty"`

	// +kubebuilder:validation:Required
	Value *string `json:"value" tf:"value,omitempty"`
}

type LibvirtNamespacesObservation struct {
}

type LibvirtNamespacesParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	ResourceAllocation map[string]*string `json:"resourceAllocation" tf:"resource_allocation,omitempty"`
}

type LibvirtObservation struct {
	CloudConfigID *string `json:"cloudConfigId,omitempty" tf:"cloud_config_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Kubeconfig *string `json:"kubeconfig,omitempty" tf:"kubeconfig,omitempty"`
}

type LibvirtPackObservation struct {
}

type LibvirtPackParameters struct {

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Optional
	RegistryUID *string `json:"registryUid,omitempty" tf:"registry_uid,omitempty"`

	// +kubebuilder:validation:Required
	Tag *string `json:"tag" tf:"tag,omitempty"`

	// +kubebuilder:validation:Required
	Values *string `json:"values" tf:"values,omitempty"`
}

type LibvirtParameters struct {

	// +kubebuilder:validation:Optional
	ApplySetting *string `json:"applySetting,omitempty" tf:"apply_setting,omitempty"`

	// +kubebuilder:validation:Optional
	BackupPolicy []LibvirtBackupPolicyParameters `json:"backupPolicy,omitempty" tf:"backup_policy,omitempty"`

	// +kubebuilder:validation:Optional
	CloudAccountID *string `json:"cloudAccountId,omitempty" tf:"cloud_account_id,omitempty"`

	// +kubebuilder:validation:Required
	CloudConfig []LibvirtCloudConfigParameters `json:"cloudConfig" tf:"cloud_config,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterProfile []LibvirtClusterProfileParameters `json:"clusterProfile,omitempty" tf:"cluster_profile,omitempty"`

	// +kubebuilder:validation:Optional
	ClusterRbacBinding []LibvirtClusterRbacBindingParameters `json:"clusterRbacBinding,omitempty" tf:"cluster_rbac_binding,omitempty"`

	// +kubebuilder:validation:Optional
	HostConfig []LibvirtHostConfigParameters `json:"hostConfig,omitempty" tf:"host_config,omitempty"`

	// +kubebuilder:validation:Optional
	LocationConfig []LibvirtLocationConfigParameters `json:"locationConfig,omitempty" tf:"location_config,omitempty"`

	// +kubebuilder:validation:Required
	MachinePool []LibvirtMachinePoolParameters `json:"machinePool" tf:"machine_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Namespaces []LibvirtNamespacesParameters `json:"namespaces,omitempty" tf:"namespaces,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchAfter *string `json:"osPatchAfter,omitempty" tf:"os_patch_after,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchOnBoot *bool `json:"osPatchOnBoot,omitempty" tf:"os_patch_on_boot,omitempty"`

	// +kubebuilder:validation:Optional
	OsPatchSchedule *string `json:"osPatchSchedule,omitempty" tf:"os_patch_schedule,omitempty"`

	// +kubebuilder:validation:Optional
	Pack []LibvirtPackParameters `json:"pack,omitempty" tf:"pack,omitempty"`

	// +kubebuilder:validation:Optional
	ScanPolicy []LibvirtScanPolicyParameters `json:"scanPolicy,omitempty" tf:"scan_policy,omitempty"`

	// +kubebuilder:validation:Optional
	SkipCompletion *bool `json:"skipCompletion,omitempty" tf:"skip_completion,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type LibvirtScanPolicyObservation struct {
}

type LibvirtScanPolicyParameters struct {

	// +kubebuilder:validation:Required
	ConfigurationScanSchedule *string `json:"configurationScanSchedule" tf:"configuration_scan_schedule,omitempty"`

	// +kubebuilder:validation:Required
	ConformanceScanSchedule *string `json:"conformanceScanSchedule" tf:"conformance_scan_schedule,omitempty"`

	// +kubebuilder:validation:Required
	PenetrationScanSchedule *string `json:"penetrationScanSchedule" tf:"penetration_scan_schedule,omitempty"`
}

type MachinePoolInstanceTypeObservation struct {
}

type MachinePoolInstanceTypeParameters struct {

	// +kubebuilder:validation:Optional
	AttachedDisks []AttachedDisksParameters `json:"attachedDisks,omitempty" tf:"attached_disks,omitempty"`

	// +kubebuilder:validation:Required
	CPU *float64 `json:"cpu" tf:"cpu,omitempty"`

	// +kubebuilder:validation:Optional
	CachePassthrough *bool `json:"cachePassthrough,omitempty" tf:"cache_passthrough,omitempty"`

	// +kubebuilder:validation:Optional
	CpusSets *string `json:"cpusSets,omitempty" tf:"cpus_sets,omitempty"`

	// +kubebuilder:validation:Required
	DiskSizeGb *float64 `json:"diskSizeGb" tf:"disk_size_gb,omitempty"`

	// +kubebuilder:validation:Optional
	GpuConfig []GpuConfigParameters `json:"gpuConfig,omitempty" tf:"gpu_config,omitempty"`

	// +kubebuilder:validation:Required
	MemoryMb *float64 `json:"memoryMb" tf:"memory_mb,omitempty"`
}

type MachinePoolPlacementsObservation struct {
}

type MachinePoolPlacementsParameters struct {

	// +kubebuilder:validation:Required
	ApplianceID *string `json:"applianceId" tf:"appliance_id,omitempty"`

	// +kubebuilder:validation:Required
	DataStoragePool *string `json:"dataStoragePool" tf:"data_storage_pool,omitempty"`

	// +kubebuilder:validation:Required
	ImageStoragePool *string `json:"imageStoragePool" tf:"image_storage_pool,omitempty"`

	// +kubebuilder:validation:Optional
	Network *string `json:"network,omitempty" tf:"network,omitempty"`

	// +kubebuilder:validation:Required
	NetworkNames *string `json:"networkNames" tf:"network_names,omitempty"`

	// +kubebuilder:validation:Required
	NetworkType *string `json:"networkType" tf:"network_type,omitempty"`

	// +kubebuilder:validation:Required
	TargetStoragePool *string `json:"targetStoragePool" tf:"target_storage_pool,omitempty"`
}

// LibvirtSpec defines the desired state of Libvirt
type LibvirtSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     LibvirtParameters `json:"forProvider"`
}

// LibvirtStatus defines the observed state of Libvirt.
type LibvirtStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        LibvirtObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Libvirt is the Schema for the Libvirts API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,jet-palette}
type Libvirt struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LibvirtSpec   `json:"spec"`
	Status            LibvirtStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LibvirtList contains a list of Libvirts
type LibvirtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Libvirt `json:"items"`
}

// Repository type metadata.
var (
	Libvirt_Kind             = "Libvirt"
	Libvirt_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Libvirt_Kind}.String()
	Libvirt_KindAPIVersion   = Libvirt_Kind + "." + CRDGroupVersion.String()
	Libvirt_GroupVersionKind = CRDGroupVersion.WithKind(Libvirt_Kind)
)

func init() {
	SchemeBuilder.Register(&Libvirt{}, &LibvirtList{})
}
