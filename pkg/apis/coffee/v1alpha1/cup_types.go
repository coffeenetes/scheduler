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

// CupSpec defines the desired state of Cup
type CupSpec struct {
	// the type of drink (coffee)
	Drink DrinkType `json:"drink"`
	// the size of the drink
	// +optional
	Size DrinkSize `json:"size,omitempty"`
	// a description of the desired drink
	// +optional
	PodSpec PodSpec `json:"podSpec,omitempty"`
}

// DrinkSize defines the requests and limits for drink cup sizes
type DrinkSize struct {
	// the requested size for the cup
	Request int `json:"request"`
	// the maximum size tolerated for the cup
	// +optional
	MaxLimit int `json:"maxLimit,omitempty"`
	// the minimum size tolerated for the cup
	// +optional
	MinLimit int `json:"minLimit,omitempty"`
}

// PodSpec defines the kind of Keurig pod desired for brewing the drink
// (corresponds to descriptors on a CoffeeMaker)
type PodSpec struct {
	// the requested (want) features of the drink
	// +optional
	Requests map[string]string `json:"requests,omitempty"`
	// the required (must have) features of the drink
	// +optional
	Requirements map[string]string `json:"requirements,omitempty"`
}

// CupStatus defines the observed state of Cup
type CupStatus struct {
	// the brew information for the drink
	// +optional
	Brew BrewStatus `json:"brew,omitempty"`
	// the actual size of the drink
	// +optional
	Size int `json:"size,omitempty"`
	// the actual descriptors for the drink
	// +optional
	Descriptors map[string]string `json:"descriptors,omitempty"`
	// extra status information about the drink
	// +optional
	Conditions []CupCondition `json:"conditions,omitempty"`
}

// BrewStatus represents if and when the drink was brewed
type BrewStatus struct {
	// the state of the brewing process
	State BrewState `json:"state"`
	// when the brewing process started
	// +optional
	Start metav1.Time `json:"start,omitempty"`
	// when the brewing process ended
	// +optional
	End metav1.Time `json:"end,omitempty"`
}

// CoffeemakerCondition defines a status condition for a CoffeeMaker.
type CupCondition struct {
	// Type is the type of the condition.
	Type CupConditionType `json:"type"`
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
	CupConditionType string
	BrewState        string
	DrinkType        string
)

const (
	CoffeeDrink DrinkType = "Coffee"
	TeaDrink    DrinkType = "Tea" // HTTP 418

	NotBrewed BrewState = "NotBrewed"
	Brewed    BrewState = "Brewed"
	Brewing   BrewState = "Brewing"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cup is the Schema for the cups API
// +k8s:openapi-gen=true
type Cup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CupSpec   `json:"spec,omitempty"`
	Status CupStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CupList contains a list of Cup
type CupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Cup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Cup{}, &CupList{})
}
