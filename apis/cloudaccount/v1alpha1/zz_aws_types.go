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

type AwsObservation struct {
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type AwsParameters struct {

	// +kubebuilder:validation:Optional
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// +kubebuilder:validation:Optional
	AwsAccessKey *string `json:"awsAccessKey,omitempty" tf:"aws_access_key,omitempty"`

	// +kubebuilder:validation:Optional
	AwsSecretKeySecretRef *v1.SecretKeySelector `json:"awsSecretKeySecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Context *string `json:"context,omitempty" tf:"context,omitempty"`

	// +kubebuilder:validation:Optional
	ExternalIDSecretRef *v1.SecretKeySelector `json:"externalIdSecretRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`
}

// AwsSpec defines the desired state of Aws
type AwsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AwsParameters `json:"forProvider"`
}

// AwsStatus defines the observed state of Aws.
type AwsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AwsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Aws is the Schema for the Awss API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,jet-palette}
type Aws struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSpec   `json:"spec"`
	Status            AwsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AwsList contains a list of Awss
type AwsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Aws `json:"items"`
}

// Repository type metadata.
var (
	Aws_Kind             = "Aws"
	Aws_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Aws_Kind}.String()
	Aws_KindAPIVersion   = Aws_Kind + "." + CRDGroupVersion.String()
	Aws_GroupVersionKind = CRDGroupVersion.WithKind(Aws_Kind)
)

func init() {
	SchemeBuilder.Register(&Aws{}, &AwsList{})
}
