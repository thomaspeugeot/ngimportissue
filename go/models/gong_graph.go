// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Country:
		ok = stage.IsStagedCountry(target)

	case *Hello:
		ok = stage.IsStagedHello(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
	func (stage *StageStruct) IsStagedCountry(country *Country) (ok bool) {

		_, ok = stage.Countrys[country]
	
		return
	}

	func (stage *StageStruct) IsStagedHello(hello *Hello) (ok bool) {

		_, ok = stage.Hellos[hello]
	
		return
	}


// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Country:
		stage.StageBranchCountry(target)

	case *Hello:
		stage.StageBranchHello(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *StageStruct) StageBranchCountry(country *Country) {

	// check if instance is already staged
	if IsStaged(stage, country) {
		return
	}

	country.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if country.Hello != nil {
		StageBranch(stage, country.Hello)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _hello := range country.AlternateHellos {
		StageBranch(stage, _hello)
	}

}

func (stage *StageStruct) StageBranchHello(hello *Hello) {

	// check if instance is already staged
	if IsStaged(stage, hello) {
		return
	}

	hello.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}


// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Country:
		stage.UnstageBranchCountry(target)

	case *Hello:
		stage.UnstageBranchHello(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *StageStruct) UnstageBranchCountry(country *Country) {

	// check if instance is already staged
	if ! IsStaged(stage, country) {
		return
	}

	country.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if country.Hello != nil {
		UnstageBranch(stage, country.Hello)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _hello := range country.AlternateHellos {
		UnstageBranch(stage, _hello)
	}

}

func (stage *StageStruct) UnstageBranchHello(hello *Hello) {

	// check if instance is already staged
	if ! IsStaged(stage, hello) {
		return
	}

	hello.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

