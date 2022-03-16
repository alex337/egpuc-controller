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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EGPUC is a specification for a EGPUC resource
type EGPUC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EGPUCSpec   `json:"spec"`
	Status EGPUCStatus `json:"status"`
}

// EGPUCSpec is the spec for a EGPUC resource
type EGPUCSpec struct {
	PodName string `json:"podName"`
	NameSpace string `json:"nameSpace"`
	Resources EGPUCResource `json:"resources"`
}

type EGPUCResource struct {
	Requests EGPUCRequest `json:"requests"`
}

type EGPUCRequest struct {
	QGPUCore string `json:"QGPUCore"`
	QGPUMemory string `json:"QGPUMemory"`
}

// EGPUCStatus is the status for a EGPUC resource
type EGPUCStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// EGPUCList is a list of EGPUC resources
type EGPUCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []EGPUC `json:"items"`
}


