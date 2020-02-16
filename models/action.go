package models

import (
	"fmt"

	"github.com/asalih/guardian/waf/transformations"
)

//Action definition for Rules
type Action struct {
	ID               string
	Phase            int
	Transformations  []string
	DisruptiveAction DisruptiveAction
	LogAction        LogAction
}

func (a *Action) ExecuteTransformation(variableData interface{}) interface{} {
	if len(a.Transformations) == 0 {
		return variableData
	}

	for _, t := range a.Transformations {
		fn := transformations.TransformationMaps.Get(t)

		if fn == nil {
			//TODO Handle unrecognized fn
			fmt.Println("Unrecognized Transformation fn " + t)
			continue
		}

		variableData = fn(variableData)
	}

	return variableData
}
