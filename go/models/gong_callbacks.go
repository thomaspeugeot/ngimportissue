// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Country:
		if stage.OnAfterCountryCreateCallback != nil {
			stage.OnAfterCountryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Hello:
		if stage.OnAfterHelloCreateCallback != nil {
			stage.OnAfterHelloCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Country:
		newTarget := any(new).(*Country)
		if stage.OnAfterCountryUpdateCallback != nil {
			stage.OnAfterCountryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Hello:
		newTarget := any(new).(*Hello)
		if stage.OnAfterHelloUpdateCallback != nil {
			stage.OnAfterHelloUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Country:
		if stage.OnAfterCountryDeleteCallback != nil {
			staged := any(staged).(*Country)
			stage.OnAfterCountryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Hello:
		if stage.OnAfterHelloDeleteCallback != nil {
			staged := any(staged).(*Hello)
			stage.OnAfterHelloDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Country:
		if stage.OnAfterCountryReadCallback != nil {
			stage.OnAfterCountryReadCallback.OnAfterRead(stage, target)
		}
	case *Hello:
		if stage.OnAfterHelloReadCallback != nil {
			stage.OnAfterHelloReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Country:
		stage.OnAfterCountryUpdateCallback = any(callback).(OnAfterUpdateInterface[Country])
	
	case *Hello:
		stage.OnAfterHelloUpdateCallback = any(callback).(OnAfterUpdateInterface[Hello])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Country:
		stage.OnAfterCountryCreateCallback = any(callback).(OnAfterCreateInterface[Country])
	
	case *Hello:
		stage.OnAfterHelloCreateCallback = any(callback).(OnAfterCreateInterface[Hello])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Country:
		stage.OnAfterCountryDeleteCallback = any(callback).(OnAfterDeleteInterface[Country])
	
	case *Hello:
		stage.OnAfterHelloDeleteCallback = any(callback).(OnAfterDeleteInterface[Hello])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Country:
		stage.OnAfterCountryReadCallback = any(callback).(OnAfterReadInterface[Country])
	
	case *Hello:
		stage.OnAfterHelloReadCallback = any(callback).(OnAfterReadInterface[Hello])
	
	}
}
