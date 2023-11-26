// generated code - do not edit
package controllers

import (
	"log"
	"net/http"
	"sync"
	"time"

	"ngimportissue/go/models"
	"ngimportissue/go/orm"

	"github.com/gin-gonic/gin"
)

// declaration in order to justify use of the models import
var __Hello__dummysDeclaration__ models.Hello
var __Hello_time__dummyDeclaration time.Duration

var mutexHello sync.Mutex

// An HelloID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getHello updateHello deleteHello
type HelloID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// HelloInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postHello updateHello
type HelloInput struct {
	// The Hello to submit or modify
	// in: body
	Hello *orm.HelloAPI
}

// GetHellos
//
// swagger:route GET /hellos hellos getHellos
//
// # Get all hellos
//
// Responses:
// default: genericError
//
//	200: helloDBResponse
func (controller *Controller) GetHellos(c *gin.Context) {

	// source slice
	var helloDBs []orm.HelloDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHellos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHello.GetDB()

	query := db.Find(&helloDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	helloAPIs := make([]orm.HelloAPI, 0)

	// for each hello, update fields from the database nullable fields
	for idx := range helloDBs {
		helloDB := &helloDBs[idx]
		_ = helloDB
		var helloAPI orm.HelloAPI

		// insertion point for updating fields
		helloAPI.ID = helloDB.ID
		helloDB.CopyBasicFieldsToHello_WOP(&helloAPI.Hello_WOP)
		helloAPI.HelloPointersEncoding = helloDB.HelloPointersEncoding
		helloAPIs = append(helloAPIs, helloAPI)
	}

	c.JSON(http.StatusOK, helloAPIs)
}

// PostHello
//
// swagger:route POST /hellos hellos postHello
//
// Creates a hello
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostHello(c *gin.Context) {

	mutexHello.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostHellos", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHello.GetDB()

	// Validate input
	var input orm.HelloAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create hello
	helloDB := orm.HelloDB{}
	helloDB.HelloPointersEncoding = input.HelloPointersEncoding
	helloDB.CopyBasicFieldsFromHello_WOP(&input.Hello_WOP)

	query := db.Create(&helloDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoHello.CheckoutPhaseOneInstance(&helloDB)
	hello := backRepo.BackRepoHello.Map_HelloDBID_HelloPtr[helloDB.ID]

	if hello != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), hello)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, helloDB)

	mutexHello.Unlock()
}

// GetHello
//
// swagger:route GET /hellos/{ID} hellos getHello
//
// Gets the details for a hello.
//
// Responses:
// default: genericError
//
//	200: helloDBResponse
func (controller *Controller) GetHello(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetHello", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHello.GetDB()

	// Get helloDB in DB
	var helloDB orm.HelloDB
	if err := db.First(&helloDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var helloAPI orm.HelloAPI
	helloAPI.ID = helloDB.ID
	helloAPI.HelloPointersEncoding = helloDB.HelloPointersEncoding
	helloDB.CopyBasicFieldsToHello_WOP(&helloAPI.Hello_WOP)

	c.JSON(http.StatusOK, helloAPI)
}

// UpdateHello
//
// swagger:route PATCH /hellos/{ID} hellos updateHello
//
// # Update a hello
//
// Responses:
// default: genericError
//
//	200: helloDBResponse
func (controller *Controller) UpdateHello(c *gin.Context) {

	mutexHello.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateHello", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHello.GetDB()

	// Validate input
	var input orm.HelloAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var helloDB orm.HelloDB

	// fetch the hello
	query := db.First(&helloDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	helloDB.CopyBasicFieldsFromHello_WOP(&input.Hello_WOP)
	helloDB.HelloPointersEncoding = input.HelloPointersEncoding

	query = db.Model(&helloDB).Updates(helloDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	helloNew := new(models.Hello)
	helloDB.CopyBasicFieldsToHello(helloNew)

	// redeem pointers
	helloDB.DecodePointers(backRepo, helloNew)

	// get stage instance from DB instance, and call callback function
	helloOld := backRepo.BackRepoHello.Map_HelloDBID_HelloPtr[helloDB.ID]
	if helloOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), helloOld, helloNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the helloDB
	c.JSON(http.StatusOK, helloDB)

	mutexHello.Unlock()
}

// DeleteHello
//
// swagger:route DELETE /hellos/{ID} hellos deleteHello
//
// # Delete a hello
//
// default: genericError
//
//	200: helloDBResponse
func (controller *Controller) DeleteHello(c *gin.Context) {

	mutexHello.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteHello", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoHello.GetDB()

	// Get model if exist
	var helloDB orm.HelloDB
	if err := db.First(&helloDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&helloDB)

	// get an instance (not staged) from DB instance, and call callback function
	helloDeleted := new(models.Hello)
	helloDB.CopyBasicFieldsToHello(helloDeleted)

	// get stage instance from DB instance, and call callback function
	helloStaged := backRepo.BackRepoHello.Map_HelloDBID_HelloPtr[helloDB.ID]
	if helloStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), helloStaged, helloDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexHello.Unlock()
}
