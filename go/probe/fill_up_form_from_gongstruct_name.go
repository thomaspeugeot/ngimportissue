// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"ngimportissue/go/models"
)

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()
	formStage.Commit()

	var prefix string

	if isNewInstance {
		prefix = "New"
	} else {
		prefix = "Update"
	}

	switch gongstructName {
	// insertion point
	case "Country":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Country Form",
			OnSave: __gong__New__CountryFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		country := new(models.Country)
		FillUpForm(country, formGroup, probe)
	case "Hello":
		formGroup := (&form.FormGroup{
			Name:  form.FormGroupDefaultName.ToString(),
			Label: prefix + " Hello Form",
			OnSave: __gong__New__HelloFormCallback(
				nil,
				probe,
			),
		}).Stage(formStage)
		hello := new(models.Hello)
		FillUpForm(hello, formGroup, probe)
	}
	formStage.Commit()
}
