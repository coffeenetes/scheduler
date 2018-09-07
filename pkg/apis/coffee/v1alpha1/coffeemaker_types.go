/*
Copyright 2018 The Coffeenetes Authors.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CoffeeMakerSpec defines the desired state of CoffeeMaker
type CoffeeMakerSpec struct {
	// The supported cup sizes for this coffee maker.
	CupSizes []int `json:"cupSizes"`
	// The general descriptors for this coffemaker (like labels, could
	// be extended to support quantities later)
	// +optional
	Descriptors map[string]string `json:"descriptors,omitempty"`
}

// CoffeeMakerStatus defines the observed state of CoffeeMaker
type CoffeeMakerStatus struct {
	// extra status information about the coffeemaker
	// +optional
	Conditions []CoffeeMakerCondition `json:"conditions,omitempty"`
}

// CoffeemakerCondition defines a status condition for a CoffeeMaker.
type CoffeeMakerCondition struct {
	// Type is the type of the condition.
	Type CoffeeMakerConditionType `json:"type"`
	// Status is the status of the condition.
	// Can be True, False, Unknown.
	Status ConditionStatus `json:"status"`
	// Last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// Unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message,omitempty"`
}

type (
	ConditionStatus          string
	CoffeeMakerConditionType string
)

// These are valid condition statuses. "ConditionTrue" means a resource is in the condition.
// "ConditionFalse" means a resource is not in the condition. "ConditionUnknown" means kubernetes
// can't decide if a resource is in the condition or not. In the future, we could add other
// intermediate conditions, e.g. ConditionDegraded.
const (
	ConditionTrue    ConditionStatus = "True"
	ConditionFalse   ConditionStatus = "False"
	ConditionUnknown ConditionStatus = "Unknown"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CoffeeMaker is the Schema for the coffeemakers API
// +k8s:openapi-gen=true
type CoffeeMaker struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CoffeeMakerSpec   `json:"spec,omitempty"`
	Status CoffeeMakerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CoffeeMakerList contains a list of CoffeeMaker
type CoffeeMakerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CoffeeMaker `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CoffeeMaker{}, &CoffeeMakerList{})
}
