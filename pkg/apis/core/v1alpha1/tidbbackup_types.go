
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
	"log"
	"context"

	"k8s.io/apimachinery/pkg/runtime"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/pingcap/apiserver/pkg/apis/core"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TidbBackup
// +k8s:openapi-gen=true
// +resource:path=tidbbackups,strategy=TidbBackupStrategy
type TidbBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TidbBackupSpec   `json:"spec,omitempty"`
	Status TidbBackupStatus `json:"status,omitempty"`
}

// TidbBackupSpec defines the desired state of TidbBackup
type TidbBackupSpec struct {
}

// TidbBackupStatus defines the observed state of TidbBackup
type TidbBackupStatus struct {
}

// Validate checks that an instance of TidbBackup is well formed
func (TidbBackupStrategy) Validate(ctx context.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*core.TidbBackup)
	log.Printf("Validating fields for TidbBackup %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default TidbBackup field values
func (TidbBackupSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*TidbBackup)
	// set default field values here
	log.Printf("Defaulting fields for TidbBackup %s\n", obj.Name)
}
