// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/fullstack-lang/gongdoc/go/models"
	"github.com/fullstack-lang/gongdoc/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Field__dummysDeclaration__ models.Field
var __Field_time__dummyDeclaration time.Duration

var mutexField sync.Mutex

// An FieldID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getField updateField deleteField
type FieldID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// FieldInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postField updateField
type FieldInput struct {
	// The Field to submit or modify
	// in: body
	Field *orm.FieldAPI
}

// GetFields
//
// swagger:route GET /fields fields getFields
//
// # Get all fields
//
// Responses:
// default: genericError
//
//	200: fieldDBResponse
func (controller *Controller) GetFields(c *gin.Context) {

	// source slice
	var fieldDBs []orm.FieldDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoField.GetDB()

	query := db.Find(&fieldDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	fieldAPIs := make([]orm.FieldAPI, 0)

	// for each field, update fields from the database nullable fields
	for idx := range fieldDBs {
		fieldDB := &fieldDBs[idx]
		_ = fieldDB
		var fieldAPI orm.FieldAPI

		// insertion point for updating fields
		fieldAPI.ID = fieldDB.ID
		fieldDB.CopyBasicFieldsToField_WOP(&fieldAPI.Field_WOP)
		fieldAPI.FieldPointersEncoding = fieldDB.FieldPointersEncoding
		fieldAPIs = append(fieldAPIs, fieldAPI)
	}

	c.JSON(http.StatusOK, fieldAPIs)
}

// PostField
//
// swagger:route POST /fields fields postField
//
// Creates a field
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostField(c *gin.Context) {

	mutexField.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostFields", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoField.GetDB()

	// Validate input
	var input orm.FieldAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create field
	fieldDB := orm.FieldDB{}
	fieldDB.FieldPointersEncoding = input.FieldPointersEncoding
	fieldDB.CopyBasicFieldsFromField_WOP(&input.Field_WOP)

	query := db.Create(&fieldDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoField.CheckoutPhaseOneInstance(&fieldDB)
	field := backRepo.BackRepoField.Map_FieldDBID_FieldPtr[fieldDB.ID]

	if field != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), field)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, fieldDB)

	mutexField.Unlock()
}

// GetField
//
// swagger:route GET /fields/{ID} fields getField
//
// Gets the details for a field.
//
// Responses:
// default: genericError
//
//	200: fieldDBResponse
func (controller *Controller) GetField(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoField.GetDB()

	// Get fieldDB in DB
	var fieldDB orm.FieldDB
	if err := db.First(&fieldDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var fieldAPI orm.FieldAPI
	fieldAPI.ID = fieldDB.ID
	fieldAPI.FieldPointersEncoding = fieldDB.FieldPointersEncoding
	fieldDB.CopyBasicFieldsToField_WOP(&fieldAPI.Field_WOP)

	c.JSON(http.StatusOK, fieldAPI)
}

// UpdateField
//
// swagger:route PATCH /fields/{ID} fields updateField
//
// # Update a field
//
// Responses:
// default: genericError
//
//	200: fieldDBResponse
func (controller *Controller) UpdateField(c *gin.Context) {

	mutexField.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoField.GetDB()

	// Validate input
	var input orm.FieldAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var fieldDB orm.FieldDB

	// fetch the field
	query := db.First(&fieldDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	fieldDB.CopyBasicFieldsFromField_WOP(&input.Field_WOP)
	fieldDB.FieldPointersEncoding = input.FieldPointersEncoding

	query = db.Model(&fieldDB).Updates(fieldDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	fieldNew := new(models.Field)
	fieldDB.CopyBasicFieldsToField(fieldNew)

	// redeem pointers
	fieldDB.DecodePointers(backRepo, fieldNew)

	// get stage instance from DB instance, and call callback function
	fieldOld := backRepo.BackRepoField.Map_FieldDBID_FieldPtr[fieldDB.ID]
	if fieldOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), fieldOld, fieldNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the fieldDB
	c.JSON(http.StatusOK, fieldDB)

	mutexField.Unlock()
}

// DeleteField
//
// swagger:route DELETE /fields/{ID} fields deleteField
//
// # Delete a field
//
// default: genericError
//
//	200: fieldDBResponse
func (controller *Controller) DeleteField(c *gin.Context) {

	mutexField.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteField", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack github.com/fullstack-lang/gongdoc/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoField.GetDB()

	// Get model if exist
	var fieldDB orm.FieldDB
	if err := db.First(&fieldDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&fieldDB)

	// get an instance (not staged) from DB instance, and call callback function
	fieldDeleted := new(models.Field)
	fieldDB.CopyBasicFieldsToField(fieldDeleted)

	// get stage instance from DB instance, and call callback function
	fieldStaged := backRepo.BackRepoField.Map_FieldDBID_FieldPtr[fieldDB.ID]
	if fieldStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), fieldStaged, fieldDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexField.Unlock()
}