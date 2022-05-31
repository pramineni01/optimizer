package generators


import (
	github.com/pramineni01/optimizer/generators/pyomo
	github.com/pramineni01/optimizer/generators/ortools
)

type TemplateHandler interface {

}

// TODO: Change to enums and error handling
func NewTemplateHandler(modelerType: string, model_input: string) (TemplateHandler, error) {
	modelerType = strings.ToUpper(modelerType)
	if modelerType == "PYOMO" {
		return pyomo.NewPyomoTemplateHandler(model_input), nil
	} else if modelerType == "ORTOOLS" {
		return ortools.NewORToolsTemplateHandler(model_input), nil
	} else if modelerType == "AMPL" {
		return ampl.NewAMPLTemplateHandler(model_input), nil
	} else {
		return nil, errors.New("Unknown modeler type")
	}
}

