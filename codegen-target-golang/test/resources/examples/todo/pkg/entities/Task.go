//
// GENERATED SOURCE - DO NOT EDIT
//
package entities

import "time"

// Task type
type Task struct {
	Idable
	userId string

	// Name of the task
	title string

	// Longer description
	description string

	// Defines if the task is done or not
	done bool

	// Age of the task
	age float64

	// Created date
	created time.Time

	metadata any

	details struct {
		innerProp string

		moreDetails struct {
			innerProp2 string
		}
	}
}
