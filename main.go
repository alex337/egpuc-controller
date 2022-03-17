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

package main

import (
	"context"
	"flag"
	"fmt"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	egpucController "github.com/alex337/egpuc-controller/pkg/apis/egpucController/v1alpha1"

	//"time"
	//
	//kubeinformers "k8s.io/client-go/informers"
	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/alex337/egpuc-controller/pkg/generated/clientset/versioned"
	//informers "k8s.io/sample-controller/pkg/generated/informers/externalversions"
	//"k8s.io/sample-controller/pkg/signals"
)

var (
	masterURL  string
	kubeconfig string
)

func main() {
	klog.InitFlags(nil)
	flag.Parse()

	//// set up signals so we handle the first shutdown signal gracefully
	//stopCh := signals.SetupSignalHandler()
	//
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		klog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	//kubeClient, err := kubernetes.NewForConfig(cfg)
	//if err != nil {
	//	klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	//}

	exampleClient, err := clientset.NewForConfig(cfg)
	if err != nil {
		klog.Fatalf("Error building example clientset: %s", err.Error())
	}
	//deploymentsClient := exampleClient.Deployments(apiv1.NamespaceDefault)





	//kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	//exampleInformerFactory := informers.NewSharedInformerFactory(exampleClient, time.Second*30)
	//
	//controller := NewController(kubeClient, exampleClient,
	//	kubeInformerFactory.Apps().V1().Deployments(),
	//	exampleInformerFactory.Samplecontroller().V1alpha1().Foos())
	//
	//// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	//// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	//kubeInformerFactory.Start(stopCh)
	//exampleInformerFactory.Start(stopCh)
	//
	//if err = controller.Run(2, stopCh); err != nil {
	//	klog.Fatalf("Error running controller: %s", err.Error())
	//}
	result, err := exampleClient.EgpuccontrollerV1alpha1().EGPUCs(apiv1.NamespaceDefault).Create(context.TODO(),newEGPUC("test", int32Ptr(1)), metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}
	klog.Info("Created EGPUC:", result.GetObjectMeta().GetName())

}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
}

func newEGPUC(name string, replicas *int32) *egpucController.EGPUC {
	return &egpucController.EGPUC{
		TypeMeta: metav1.TypeMeta{APIVersion: egpucController.SchemeGroupVersion.String()},
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: metav1.NamespaceDefault,
		},
		Spec: egpucController.EGPUCSpec{
			PodName: fmt.Sprintf("%s-deployment", name),
			NameSpace: "default",
			Container: egpucController.Container{
				ContainerName: "cont1",
				Resource: egpucController.EGPUCResource{
					Requests: egpucController.EGPUCRequest{
						QGPUMemory: "1",
						QGPUCore: "1",
					},
				},
			},
		},
	}
}

func int32Ptr(i int32) *int32 { return &i }
