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

type ClustersObservation struct {
}

type ClustersParameters struct {

	// The UID of the host cluster.
	// +kubebuilder:validation:Required
	ClusterUID *string `json:"clusterUid" tf:"cluster_uid,omitempty"`

	// The host DNS wildcard for the cluster. i.e. `*.dev` or `*test.com`
	// +kubebuilder:validation:Optional
	HostDNS *string `json:"hostDns,omitempty" tf:"host_dns,omitempty"`
}

type ConfigObservation struct {
}

type ConfigParameters struct {

	// The CPU limit in millicores.
	// +kubebuilder:validation:Optional
	CPUMillicore *float64 `json:"cpuMillicore,omitempty" tf:"cpu_millicore,omitempty"`

	// The host endpoint type. Allowed values are 'Ingress' or 'LoadBalancer'. Defaults to 'Ingress'.
	// +kubebuilder:validation:Optional
	HostEndpointType *string `json:"hostEndpointType,omitempty" tf:"host_endpoint_type,omitempty"`

	// The memory limit in megabytes (MB).
	// +kubebuilder:validation:Optional
	MemoryInMb *float64 `json:"memoryInMb,omitempty" tf:"memory_in_mb,omitempty"`

	// The allowed oversubscription percentage.
	// +kubebuilder:validation:Optional
	OversubscriptionPercent *float64 `json:"oversubscriptionPercent,omitempty" tf:"oversubscription_percent,omitempty"`

	// The storage limit in gigabytes (GB).
	// +kubebuilder:validation:Optional
	StorageInGb *float64 `json:"storageInGb,omitempty" tf:"storage_in_gb,omitempty"`

	// +kubebuilder:validation:Optional
	Values *string `json:"values,omitempty" tf:"values,omitempty"`
}

type GroupObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type GroupParameters struct {

	// A list of clusters to include in the cluster group.
	// +kubebuilder:validation:Optional
	Clusters []ClustersParameters `json:"clusters,omitempty" tf:"clusters,omitempty"`

	// +kubebuilder:validation:Required
	Config []ConfigParameters `json:"config" tf:"config,omitempty"`

	// The context of the Cluster group. Allowed values are 'project' or 'tenant'. Defaults to 'tenant'.
	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// A list of tags to be applied to the cluster group. Tags must be in the form of `key:value`.
	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// GroupSpec defines the desired state of Group
type GroupSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     GroupParameters `json:"forProvider"`
}

// GroupStatus defines the observed state of Group.
type GroupStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        GroupObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Group is the Schema for the Groups API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,jet-palette}
type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GroupSpec   `json:"spec"`
	Status            GroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GroupList contains a list of Groups
type GroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Group `json:"items"`
}

// Repository type metadata.
var (
	Group_Kind             = "Group"
	Group_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Group_Kind}.String()
	Group_KindAPIVersion   = Group_Kind + "." + CRDGroupVersion.String()
	Group_GroupVersionKind = CRDGroupVersion.WithKind(Group_Kind)
)

func init() {
	SchemeBuilder.Register(&Group{}, &GroupList{})
}
