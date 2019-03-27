/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	"context"
	"log"

	"k8s.io/apimachinery/pkg/runtime"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/pingcap/apiserver/pkg/apis/core"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TidbInstance
// +k8s:openapi-gen=true
// +resource:path=tidbinstances,strategy=TidbInstanceStrategy
type TidbInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TidbInstanceSpec   `json:"spec,omitempty"`
	Status TidbInstanceStatus `json:"status,omitempty"`
}

// TidbInstanceSpec defines the desired state of TidbInstance
type TidbInstanceSpec struct {
	Pd   PdSpec   `json:"pd,omitempty"`
	Tikv TikvSpec `json:"tikv,omitempty"`
	Tidb TidbSpec `json:"tidb,omitempty"`
}

type PdSpec struct {
	Replicas         int32                       `json:"replicas,omitempty"`
	Image            string                      `json:"image,omitempty"`
	ImagePullPolicy  *corev1.PullPolicy          `json:"imagePullPolicy,omitempty"`
	StorageClassName *string                     `json:"storageClassName,omitempty"`
	Resources        corev1.ResourceRequirements `json:"resources,omitempty"`
	Tolerations      []corev1.Toleration         `json:"tolerations,omitempty"`
}

type TikvSpec struct {
	Replicas         int32                       `json:"replicas,omitempty"`
	Image            string                      `json:"image,omitempty"`
	ImagePullPolicy  *corev1.PullPolicy          `json:"imagePullPolicy,omitempty"`
	StorageClassName *string                     `json:"storageClassName,omitempty"`
	Resources        corev1.ResourceRequirements `json:"resources,omitempty"`
	Tolerations      []corev1.Toleration         `json:"tolerations,omitempty"`
}

type TidbSpec struct {
	Replicas        int32                       `json:"replicas,omitempty"`
	Image           string                      `json:"image,omitempty"`
	ImagePullPolicy *corev1.PullPolicy          `json:"imagePullPolicy,omitempty"`
	Resources       corev1.ResourceRequirements `json:"resources,omitempty"`
	Tolerations     []corev1.Toleration         `json:"tolerations,omitempty"`
}

// TidbInstanceStatus defines the observed state of TidbInstance
type TidbInstanceStatus struct {
}

// Validate checks that an instance of TidbInstance is well formed
func (TidbInstanceStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*core.TidbInstance)
	log.Printf("Validating fields for TidbInstance %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default TidbInstance field values
func (TidbInstanceSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*TidbInstance)
	// set default field values here
	if obj.Spec.Pd.ImagePullPolicy == nil {
		f := corev1.PullIfNotPresent
		obj.Spec.Pd.ImagePullPolicy = &f
	}
	if obj.Spec.Tikv.ImagePullPolicy == nil {
		f := corev1.PullIfNotPresent
		obj.Spec.Tikv.ImagePullPolicy = &f
	}
	if obj.Spec.Tidb.ImagePullPolicy == nil {
		f := corev1.PullIfNotPresent
		obj.Spec.Tidb.ImagePullPolicy = &f
	}

	if obj.Spec.Pd.StorageClassName == nil {
		f := "local-storage"
		obj.Spec.Pd.StorageClassName = &f
	}
	if obj.Spec.Tikv.StorageClassName == nil {
		f := "local-storage"
		obj.Spec.Tikv.StorageClassName = &f
	}
	log.Printf("Defaulting fields for TidbInstance %s\n", obj.Name)
}
