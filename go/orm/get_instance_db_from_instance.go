// generated code - do not edit
package orm

import (
	"ngimportissue/go/models"
)

type GongstructDB interface {
	// insertion point for generic types
	// "int" is present to handle the case when no struct is present
	int | CountryDB | HelloDB
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Country:
		countryInstance := any(concreteInstance).(*models.Country)
		ret2 := backRepo.BackRepoCountry.GetCountryDBFromCountryPtr(countryInstance)
		ret = any(ret2).(*T2)
	case *models.Hello:
		helloInstance := any(concreteInstance).(*models.Hello)
		ret2 := backRepo.BackRepoHello.GetHelloDBFromHelloPtr(helloInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Country:
		tmp := GetInstanceDBFromInstance[models.Country, CountryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Hello:
		tmp := GetInstanceDBFromInstance[models.Hello, HelloDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.StageStruct,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.Country:
		tmp := GetInstanceDBFromInstance[models.Country, CountryDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.Hello:
		tmp := GetInstanceDBFromInstance[models.Hello, HelloDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
