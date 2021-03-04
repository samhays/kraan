/*
Ideally, we should have one `<kind>_conroller_test.go` for each controller scaffolded and called in the `test_suite.go`.
So, let's write our example test for the AddonsLayer controller (`AddonsLayer_controller_test.go.`)
*/

/*

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
// +kubebuilder:docs-gen:collapse=Apache License

/*
As usual, we start with the necessary imports. We also define some utility variables.
*/
package controllers_test

import (
	"context"
	"time"
	"encoding/json"
	"fmt"
	"io/ioutil"
	
	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
	helmctlv2 "github.com/fluxcd/helm-controller/api/v2beta1"

	kraanv1alpha1 "github.com/fidelity/kraan/api/v1alpha1"
)

const (
	AddonsLayerName = "apps"
	k8sList                           = "List"
	timeout  = time.Second * 20
	interval = time.Millisecond * 250
)

func getAddonsFromFiles(fileNames ...string) *kraanv1alpha1.AddonsLayerList {
	addonsLayersList := &kraanv1alpha1.AddonsLayerList{
		TypeMeta: metav1.TypeMeta{
			Kind:       k8sList,
			APIVersion: fmt.Sprintf("%s/%s", kraanv1alpha1.GroupVersion.Version, kraanv1alpha1.GroupVersion.Version),
		},
		Items: make([]kraanv1alpha1.AddonsLayer, 0, 10),
	}

	for _, fileName := range fileNames {
		buffer, err := ioutil.ReadFile(fileName)
		if err != nil {
			
			return nil
		}

		addons := &kraanv1alpha1.AddonsLayerList{}

		err = json.Unmarshal(buffer, addons)
		if err != nil {
			
			return nil
		}

		addonsLayersList.Items = append(addonsLayersList.Items, addons.Items...)
	}

	return addonsLayersList
}

func getAddonFromList(name string, addonList *kraanv1alpha1.AddonsLayerList) *kraanv1alpha1.AddonsLayer {
	for _, item := range addonList.Items {
		if item.ObjectMeta.Name == name {
			return &item
		}
	}

	return nil
}

func getHelmReleasesFromFiles(fileNames ...string) *helmctlv2.HelmReleaseList {
	helmReleasesList := &helmctlv2.HelmReleaseList{
		TypeMeta: metav1.TypeMeta{
			Kind:       "List",
			APIVersion: fmt.Sprintf("%s/%s", helmctlv2.GroupVersion.Version, helmctlv2.GroupVersion.Version),
		},
		Items: make([]helmctlv2.HelmRelease, 0, 10),
	}

	for _, fileName := range fileNames {
		buffer, err := ioutil.ReadFile(fileName)
		if err != nil {
		
			return nil
		}
		helmReleases := &helmctlv2.HelmReleaseList{}

		err = json.Unmarshal(buffer, helmReleases)
		if err != nil {
			
			return nil
		}

		helmReleasesList.Items = append(helmReleasesList.Items, helmReleases.Items...)
	}
	return helmReleasesList
}

func getHelmReleaseFromList(nameSpaceSlashName string, helmReleaseList *helmctlv2.HelmReleaseList) *helmctlv2.HelmRelease {
	for _, item := range helmReleaseList.Items {
		if fmt.Sprintf("%s/%s", item.Namespace, item.Name) == nameSpaceSlashName {
			return &item
		}
	}

	return nil
}


// +kubebuilder:docs-gen:collapse=Imports

var _ = Describe("AddonsLayer controller", func() {
	createAddonsLayer := func(ctx context.Context, log logr.Logger, AddonsLayer *kraanv1alpha1.AddonsLayer) *kraanv1alpha1.AddonsLayer {
		Expect(k8sClient.Create(ctx, AddonsLayer)).Should(Succeed())
		AddonsLayerLookupKey := types.NamespacedName{Name: AddonsLayerName}
		createdAddonsLayer := &kraanv1alpha1.AddonsLayer{}

		log.Info("waiting for AddonsLayer to be created")
		Eventually(func() bool {
			err := k8sClient.Get(ctx, AddonsLayerLookupKey, createdAddonsLayer)

			return err == nil
		}, timeout, interval).Should(BeTrue())
		log.Info("AddonsLayer created")

		return createdAddonsLayer
	}

	createAddonsLayers := func(ctx context.Context, log logr.Logger, AddonsLayersFileNames ...string) []*kraanv1alpha1.AddonsLayer {
		addonsLayers := []*kraanv1alpha1.AddonsLayer{}
		

		return addonsLayers
	}

	verifyAddonsLayer := func(ctx context.Context, log logr.Logger, addonsLayer *kraanv1alpha1.AddonsLayer, status string) {
		createdAddonsLayer := &kraanv1alpha1.AddonsLayer{}
		AddonsLayerLookupKey := types.NamespacedName{Name: AddonsLayerName}
		log.Info("waiting for AddonsLayer status to be expected value", "expected", status)
		Eventually(func() bool {
			err := k8sClient.Get(ctx, AddonsLayerLookupKey, createdAddonsLayer)
			if err != nil {
				return false
			}
			log.Info("AddonsLayer status", "actual", createdAddonsLayer.Status.State, "expected", status)

			return createdAddonsLayer.Status.State == status
		}, timeout, interval).Should(BeTrue())
		log.Info("AddonsLayer status achieved expected value", "expected", status)

		Expect(createdAddonsLayer.Spec.Hold).Should(Equal(addonsLayer.Spec.Hold))

		Expect(len(createdAddonsLayer.Status.Conditions)).Should(Equal(1))

		message := kraanv1alpha1.AddonsLayerDeployedMsg
		if status == kraanv1alpha1.HoldCondition {
			message = kraanv1alpha1.AddonsLayerHoldMsg
		}
		Expect(createdAddonsLayer.Status.Conditions).Should(Equal([]metav1.Condition{{
			Type:               status,
			Reason:             status,
			Status:             metav1.ConditionTrue,
			LastTransitionTime: createdAddonsLayer.Status.Conditions[0].LastTransitionTime,
			Message:            message,
		}}))

	}

	verifyAddonsLayers := func(ctx context.Context, log logr.Logger, addonsLayers []*kraanv1alpha1.AddonsLayer) {
		for _, addonsLayer := range addonsLayers {
			verifyAddonsLayer(ctx, log, addonsLayer, "Deployed")
		}
	}

	Context("When creating AddonsLayers, wait for them to be deployed state", func() {
		It("Should set AddonsLayers Status to Deployed and deploy the HelmReleases defined by each AddonsLayer", func() {
			ctx := context.Background()
			logf.SetLogger(zap.New(zap.WriteTo(GinkgoWriter)))
			log := logf.Log.WithName("test-one")

			By("By creating a new AddonsLayers")
			createdAddonsLayers := createAddonsLayers(ctx, log, AddonsLayersFileName)
			verifyAddonsLayers(ctx, log, createdAddonsLayers)

		})
	})
})