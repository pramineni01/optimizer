package templates

type TemplateHandler interface {

}

// TODO: Change to enums and error handling
func NewTemplateHandler(modelerType: string) (TemplateHandler, error) {
	modelerType = strings.ToUpper(modelerType)
	if modelerType == "PYOMO" {
		return pyomo.NewPyomoTemplateHandler(), nil
	} else if modelerType == "ORTOOLS" {
		return ortools.NewORToolsTemplateHandler(), nil
	} else if modelerType == "AMPL" {
		return ampl.NewAMPLTemplateHandler(), nil
	} else {
		return nil, errors.New("Unknown modeler type")
	}
}

