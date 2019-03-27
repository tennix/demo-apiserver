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

package tidbinstance

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder-alpha/pkg/builders"

	"github.com/pingcap/apiserver/pkg/apis/core/v1alpha1"
	listers "github.com/pingcap/apiserver/pkg/client/listers_generated/core/v1alpha1"
	"github.com/pingcap/apiserver/pkg/controller/sharedinformers"
)

// +controller:group=core,version=v1alpha1,kind=TidbInstance,resource=tidbinstances
type TidbInstanceControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about TidbInstance
	lister listers.TidbInstanceLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *TidbInstanceControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing tidbinstances labels
	c.lister = arguments.GetSharedInformers().Factory.Core().V1alpha1().TidbInstances().Lister()
}

// Reconcile handles enqueued messages
func (c *TidbInstanceControllerImpl) Reconcile(u *v1alpha1.TidbInstance) error {
	// Implement controller logic here
	log.Printf("Running reconcile TidbInstance for %s: %v\n", u.Name, u.Spec)
	// resp, err := helmCli.ReleaseContent(relName)
	// if err != nil {
	// // if not exist, then helm install the release
	// 	return err
	// }
	// cfg, err := chartutil.CoalesceValues(resp.Release.Chart, resp.Release.Config)
	// if err != nil {
	// 	return err
	// }
	// // cfg is the computed values, the controller compares cfg and tidbinstance.spec
	// // if different, then helm upgrade the release, otherwise do nothing

	// // we may also compare the user supplied values: resp.Release.Config.Raw with tidbinstance.spec
	// // if we only use override values other than replace values.yaml totally
	return nil
}

func (c *TidbInstanceControllerImpl) Get(namespace, name string) (*v1alpha1.TidbInstance, error) {
	return c.lister.TidbInstances(namespace).Get(name)
}
