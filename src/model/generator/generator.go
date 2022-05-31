package generator

import (
	"github.com/pramineni01/optimizer/model/generator/templates"
)

type ModelSettings struct {
	ModelType string `json:"model_type"`
	ModelName string `json:"model_name"`
}

type ModelData struct {

}

func (mg *appGenerator) Generate() (AppComponent, error) {
	// Read the main template.
	tmplHandler := templates.NewTemplateHandler(mg.ModelSettings.ModelType)
	return tmplHandler.Hydrate(mg.ModelSettings, mg.ModelData)
	// Start populating it
	// Output
}