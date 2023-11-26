// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gongtable/go/models"

	"ngimportissue/go/models"
	"ngimportissue/go/orm"
)

var __dummy_orm_fillup_form = orm.BackRepoStruct{}

func FillUpForm[T models.Gongstruct](
	instance *T,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Country:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		AssociationFieldToForm("Hello", instanceWithInferedType.Hello, formGroup, probe)
		AssociationSliceToForm("AlternateHellos", instanceWithInferedType, &instanceWithInferedType.AlternateHellos, formGroup, probe)

	case *models.Hello:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup, false)
		{
			var rf models.ReverseField
			_ = rf
			rf.GongstructName = "Country"
			rf.Fieldname = "AlternateHellos"
			reverseFieldOwner := orm.GetReverseFieldOwner(probe.stageOfInterest, probe.backRepoOfInterest, instanceWithInferedType, &rf)
			if reverseFieldOwner != nil {
				AssociationReverseFieldToForm(
					reverseFieldOwner.(*models.Country),
					"AlternateHellos",
					instanceWithInferedType,
					formGroup,
					probe)
			} else {
				AssociationReverseFieldToForm[*models.Country, *models.Hello](
					nil,
					"AlternateHellos",
					instanceWithInferedType,
					formGroup,
					probe)
			}	
		}

	default:
		_ = instanceWithInferedType
	}
}
