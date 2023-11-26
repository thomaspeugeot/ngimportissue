// generated code - do not edit
package probe

import (
	"log"
	"slices"
	"time"

	table "github.com/fullstack-lang/gongtable/go/models"

	"ngimportissue/go/models"
	"ngimportissue/go/orm"
)

const __dummmy__time = time.Nanosecond

var __dummmy__letters = slices.Delete([]string{"a"}, 0, 1)
var __dummy_orm = orm.BackRepoStruct{}

// insertion point
func __gong__New__CountryFormCallback(
	country *models.Country,
	probe *Probe,
) (countryFormCallback *CountryFormCallback) {
	countryFormCallback = new(CountryFormCallback)
	countryFormCallback.probe = probe
	countryFormCallback.country = country

	countryFormCallback.CreationMode = (country == nil)

	return
}

type CountryFormCallback struct {
	country *models.Country

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (countryFormCallback *CountryFormCallback) OnSave() {

	log.Println("CountryFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	countryFormCallback.probe.formStage.Checkout()

	if countryFormCallback.country == nil {
		countryFormCallback.country = new(models.Country).Stage(countryFormCallback.probe.stageOfInterest)
	}
	country_ := countryFormCallback.country
	_ = country_

	// get the formGroup
	formGroup := countryFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(country_.Name), formDiv)
		case "Hello":
			FormDivSelectFieldToField(&(country_.Hello), countryFormCallback.probe.stageOfInterest, formDiv)
		}
	}

	countryFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Country](
		countryFormCallback.probe,
	)
	countryFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if countryFormCallback.CreationMode {
		countryFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__CountryFormCallback(
				nil,
				countryFormCallback.probe,
			),
		}).Stage(countryFormCallback.probe.formStage)
		country := new(models.Country)
		FillUpForm(country, newFormGroup, countryFormCallback.probe)
		countryFormCallback.probe.formStage.Commit()
	}

	fillUpTree(countryFormCallback.probe)
}
func __gong__New__HelloFormCallback(
	hello *models.Hello,
	probe *Probe,
) (helloFormCallback *HelloFormCallback) {
	helloFormCallback = new(HelloFormCallback)
	helloFormCallback.probe = probe
	helloFormCallback.hello = hello

	helloFormCallback.CreationMode = (hello == nil)

	return
}

type HelloFormCallback struct {
	hello *models.Hello

	// If the form call is called on the creation of a new instnace
	CreationMode bool

	probe *Probe
}

func (helloFormCallback *HelloFormCallback) OnSave() {

	log.Println("HelloFormCallback, OnSave")

	// checkout formStage to have the form group on the stage synchronized with the
	// back repo (and front repo)
	helloFormCallback.probe.formStage.Checkout()

	if helloFormCallback.hello == nil {
		helloFormCallback.hello = new(models.Hello).Stage(helloFormCallback.probe.stageOfInterest)
	}
	hello_ := helloFormCallback.hello
	_ = hello_

	// get the formGroup
	formGroup := helloFormCallback.probe.formStage.FormGroups_mapString[table.FormGroupDefaultName.ToString()]

	for _, formDiv := range formGroup.FormDivs {
		switch formDiv.Name {
		// insertion point per field
		case "Name":
			FormDivBasicFieldToField(&(hello_.Name), formDiv)
		case "Country:AlternateHellos":
			// we need to retrieve the field owner before the change
			var pastCountryOwner *models.Country
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Country"
			rf.Fieldname = "AlternateHellos"
			reverseFieldOwner := orm.GetReverseFieldOwner(
				helloFormCallback.probe.stageOfInterest,
				helloFormCallback.probe.backRepoOfInterest,
				hello_,
				&rf)

			if reverseFieldOwner != nil {
				pastCountryOwner = reverseFieldOwner.(*models.Country)
			}
			if formDiv.FormFields[0].FormFieldSelect.Value == nil {
				if pastCountryOwner != nil {
					idx := slices.Index(pastCountryOwner.AlternateHellos, hello_)
					pastCountryOwner.AlternateHellos = slices.Delete(pastCountryOwner.AlternateHellos, idx, idx+1)
				}
			} else {
				// we need to retrieve the field owner after the change
				// parse all astrcut and get the one with the name in the
				// div
				for _country := range *models.GetGongstructInstancesSet[models.Country](helloFormCallback.probe.stageOfInterest) {

					// the match is base on the name
					if _country.GetName() == formDiv.FormFields[0].FormFieldSelect.Value.GetName() {
						newCountryOwner := _country // we have a match
						if pastCountryOwner != nil {
							if newCountryOwner != pastCountryOwner {
								idx := slices.Index(pastCountryOwner.AlternateHellos, hello_)
								pastCountryOwner.AlternateHellos = slices.Delete(pastCountryOwner.AlternateHellos, idx, idx+1)
								newCountryOwner.AlternateHellos = append(newCountryOwner.AlternateHellos, hello_)
							}
						} else {
							newCountryOwner.AlternateHellos = append(newCountryOwner.AlternateHellos, hello_)
						}
					}
				}
			}
		}
	}

	helloFormCallback.probe.stageOfInterest.Commit()
	fillUpTable[models.Hello](
		helloFormCallback.probe,
	)
	helloFormCallback.probe.tableStage.Commit()

	// display a new form by reset the form stage
	if helloFormCallback.CreationMode {
		helloFormCallback.probe.formStage.Reset()
		newFormGroup := (&table.FormGroup{
			Name: table.FormGroupDefaultName.ToString(),
			OnSave: __gong__New__HelloFormCallback(
				nil,
				helloFormCallback.probe,
			),
		}).Stage(helloFormCallback.probe.formStage)
		hello := new(models.Hello)
		FillUpForm(hello, newFormGroup, helloFormCallback.probe)
		helloFormCallback.probe.formStage.Commit()
	}

	fillUpTree(helloFormCallback.probe)
}
