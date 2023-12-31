// generated code - do not edit
package orm

import (
	"ngimportissue/go/models"
)

func GetReverseFieldOwnerName[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res string) {

	res = ""
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Country:
		tmp := GetInstanceDBFromInstance[models.Country, CountryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Country":
			switch reverseField.Fieldname {
			}
		case "Hello":
			switch reverseField.Fieldname {
			}
		}

	case *models.Hello:
		tmp := GetInstanceDBFromInstance[models.Hello, HelloDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Country":
			switch reverseField.Fieldname {
			case "AlternateHellos":
				if _country, ok := stage.Country_AlternateHellos_reverseMap[inst]; ok {
					res = _country.Name
				}
			}
		case "Hello":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return
}

func GetReverseFieldOwner[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T,
	reverseField *models.ReverseField) (res any) {

	res = nil
	switch inst := any(instance).(type) {
	// insertion point
	case *models.Country:
		tmp := GetInstanceDBFromInstance[models.Country, CountryDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Country":
			switch reverseField.Fieldname {
			}
		case "Hello":
			switch reverseField.Fieldname {
			}
		}

	case *models.Hello:
		tmp := GetInstanceDBFromInstance[models.Hello, HelloDB](
			stage, backRepo, inst,
		)
		_ = tmp
		switch reverseField.GongstructName {
		// insertion point
		case "Country":
			switch reverseField.Fieldname {
			case "AlternateHellos":
				res = stage.Country_AlternateHellos_reverseMap[inst]
			}
		case "Hello":
			switch reverseField.Fieldname {
			}
		}

	default:
		_ = inst
	}
	return res
}
