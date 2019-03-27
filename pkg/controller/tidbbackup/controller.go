
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


package tidbbackup

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"

	"github.com/pingcap/apiserver/pkg/apis/core/v1alpha1"
	"github.com/pingcap/apiserver/pkg/controller/sharedinformers"
	listers "github.com/pingcap/apiserver/pkg/client/listers_generated/core/v1alpha1"
)

// +controller:group=core,version=v1alpha1,kind=TidbBackup,resource=tidbbackups
type TidbBackupControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about TidbBackup
	lister listers.TidbBackupLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *TidbBackupControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing tidbbackups labels
	c.lister = arguments.GetSharedInformers().Factory.Core().V1alpha1().TidbBackups().Lister()
}

// Reconcile handles enqueued messages
func (c *TidbBackupControllerImpl) Reconcile(u *v1alpha1.TidbBackup) error {
	// Implement controller logic here
	log.Printf("Running reconcile TidbBackup for %s\n", u.Name)
	return nil
}

func (c *TidbBackupControllerImpl) Get(namespace, name string) (*v1alpha1.TidbBackup, error) {
	return c.lister.TidbBackups(namespace).Get(name)
}
