//
// GENERATED SOURCE - DO NOT EDIT
//
package entities

import kapeta "github.com/kapetacom/sdk-go-config"

// Task type
type Task struct {
	Idable
	UserId string `json:"userId" xml:"userId" yaml:"userId"`

	// Name of the task
	Title string `json:"title" xml:"title" yaml:"title"`

	// Longer description
	Description string `json:"description" xml:"description" yaml:"description"`

	// Defines if the task is done or not
	Done bool `json:"done" xml:"done" yaml:"done"`

	// Age of the task
	Age float64 `json:"age" xml:"age" yaml:"age"`

	// Created date
	Created kapeta.Epoch `json:"created" xml:"created" yaml:"created"`

	Metadata any `json:"metadata" xml:"metadata" yaml:"metadata"`

	Details struct {
		InnerProp string `json:"innerProp" xml:"innerProp" yaml:"innerProp"`

		MoreDetails struct {
			InnerProp2 string `json:"innerProp2" xml:"innerProp2" yaml:"innerProp2"`
		} `json:"moreDetails" xml:"moreDetails" yaml:"moreDetails"`
	} `json:"details" xml:"details" yaml:"details"`
}
