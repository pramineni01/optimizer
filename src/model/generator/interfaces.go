package generator

/*
	Model: A representation of problem to be solved
	Modeler: Usually libraries like pyomo and or-tools which are used to model the linear programming problems
	
	AppComponent: An representation of an app built around a modeler library. A ModelerApp programmed as a component, which ideally should be useable in a lambda or a microservice or a stand-alone.
	AppGenerator: Generates AppComponent.
*/

import (
	"fmt"
	"json"
)

// Represents Application code based on chosen Modeler (pyomo/or-tools).
interface AppComponent {
	func Print(format: string) ([]byte, error)
	func ToBytes(out: *[]bytes) error
	func Save(format: string, loc: string): (bool, error)
}

// Represents App generator
interface AppGenerator {
	// Generates Modeler, which inturn when run, will generate model
	func Generate() (AppComponent, error)
}

func NewAppGenerator(settings: string, data: string) AppGenerator {
	return &appGenerator{
		ModelSettings: json.loads(settings)
		ModelData: json.loads(data)
	}
} 

type appGenerator struct {
	// actual implementation
	settings ModelSettings
	data ModelData
}
