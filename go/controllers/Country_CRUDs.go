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
var __Country__dummysDeclaration__ models.Country
var __Country_time__dummyDeclaration time.Duration

var mutexCountry sync.Mutex

// An CountryID parameter model.
//
// This is used for operations that want the ID of an order in the path
// swagger:parameters getCountry updateCountry deleteCountry
type CountryID struct {
	// The ID of the order
	//
	// in: path
	// required: true
	ID int64
}

// CountryInput is a schema that can validate the user’s
// input to prevent us from getting invalid data
// swagger:parameters postCountry updateCountry
type CountryInput struct {
	// The Country to submit or modify
	// in: body
	Country *orm.CountryAPI
}

// GetCountrys
//
// swagger:route GET /countrys countrys getCountrys
//
// # Get all countrys
//
// Responses:
// default: genericError
//
//	200: countryDBResponse
func (controller *Controller) GetCountrys(c *gin.Context) {

	// source slice
	var countryDBs []orm.CountryDB

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCountrys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCountry.GetDB()

	query := db.Find(&countryDBs)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// slice that will be transmitted to the front
	countryAPIs := make([]orm.CountryAPI, 0)

	// for each country, update fields from the database nullable fields
	for idx := range countryDBs {
		countryDB := &countryDBs[idx]
		_ = countryDB
		var countryAPI orm.CountryAPI

		// insertion point for updating fields
		countryAPI.ID = countryDB.ID
		countryDB.CopyBasicFieldsToCountry_WOP(&countryAPI.Country_WOP)
		countryAPI.CountryPointersEncoding = countryDB.CountryPointersEncoding
		countryAPIs = append(countryAPIs, countryAPI)
	}

	c.JSON(http.StatusOK, countryAPIs)
}

// PostCountry
//
// swagger:route POST /countrys countrys postCountry
//
// Creates a country
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
//	Responses:
//	  200: nodeDBResponse
func (controller *Controller) PostCountry(c *gin.Context) {

	mutexCountry.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("PostCountrys", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCountry.GetDB()

	// Validate input
	var input orm.CountryAPI

	err := c.ShouldBindJSON(&input)
	if err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// Create country
	countryDB := orm.CountryDB{}
	countryDB.CountryPointersEncoding = input.CountryPointersEncoding
	countryDB.CopyBasicFieldsFromCountry_WOP(&input.Country_WOP)

	query := db.Create(&countryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	backRepo.BackRepoCountry.CheckoutPhaseOneInstance(&countryDB)
	country := backRepo.BackRepoCountry.Map_CountryDBID_CountryPtr[countryDB.ID]

	if country != nil {
		models.AfterCreateFromFront(backRepo.GetStage(), country)
	}

	// a POST is equivalent to a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, countryDB)

	mutexCountry.Unlock()
}

// GetCountry
//
// swagger:route GET /countrys/{ID} countrys getCountry
//
// Gets the details for a country.
//
// Responses:
// default: genericError
//
//	200: countryDBResponse
func (controller *Controller) GetCountry(c *gin.Context) {

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("GetCountry", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCountry.GetDB()

	// Get countryDB in DB
	var countryDB orm.CountryDB
	if err := db.First(&countryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	var countryAPI orm.CountryAPI
	countryAPI.ID = countryDB.ID
	countryAPI.CountryPointersEncoding = countryDB.CountryPointersEncoding
	countryDB.CopyBasicFieldsToCountry_WOP(&countryAPI.Country_WOP)

	c.JSON(http.StatusOK, countryAPI)
}

// UpdateCountry
//
// swagger:route PATCH /countrys/{ID} countrys updateCountry
//
// # Update a country
//
// Responses:
// default: genericError
//
//	200: countryDBResponse
func (controller *Controller) UpdateCountry(c *gin.Context) {

	mutexCountry.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("UpdateCountry", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCountry.GetDB()

	// Validate input
	var input orm.CountryAPI
	if err := c.ShouldBindJSON(&input); err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get model if exist
	var countryDB orm.CountryDB

	// fetch the country
	query := db.First(&countryDB, c.Param("id"))

	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// update
	countryDB.CopyBasicFieldsFromCountry_WOP(&input.Country_WOP)
	countryDB.CountryPointersEncoding = input.CountryPointersEncoding

	query = db.Model(&countryDB).Updates(countryDB)
	if query.Error != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = query.Error.Error()
		log.Println(query.Error.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// get an instance (not staged) from DB instance, and call callback function
	countryNew := new(models.Country)
	countryDB.CopyBasicFieldsToCountry(countryNew)

	// redeem pointers
	countryDB.DecodePointers(backRepo, countryNew)

	// get stage instance from DB instance, and call callback function
	countryOld := backRepo.BackRepoCountry.Map_CountryDBID_CountryPtr[countryDB.ID]
	if countryOld != nil {
		models.AfterUpdateFromFront(backRepo.GetStage(), countryOld, countryNew)
	}

	// an UPDATE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	// in some cases, with the marshalling of the stage, this operation might
	// generates a checkout
	backRepo.IncrementPushFromFrontNb()

	// return status OK with the marshalling of the the countryDB
	c.JSON(http.StatusOK, countryDB)

	mutexCountry.Unlock()
}

// DeleteCountry
//
// swagger:route DELETE /countrys/{ID} countrys deleteCountry
//
// # Delete a country
//
// default: genericError
//
//	200: countryDBResponse
func (controller *Controller) DeleteCountry(c *gin.Context) {

	mutexCountry.Lock()

	values := c.Request.URL.Query()
	stackPath := ""
	if len(values) == 1 {
		value := values["GONG__StackPath"]
		if len(value) == 1 {
			stackPath = value[0]
			// log.Println("DeleteCountry", "GONG__StackPath", stackPath)
		}
	}
	backRepo := controller.Map_BackRepos[stackPath]
	if backRepo == nil {
		log.Panic("Stack ngimportissue/go/models, Unkown stack", stackPath)
	}
	db := backRepo.BackRepoCountry.GetDB()

	// Get model if exist
	var countryDB orm.CountryDB
	if err := db.First(&countryDB, c.Param("id")).Error; err != nil {
		var returnError GenericError
		returnError.Body.Code = http.StatusBadRequest
		returnError.Body.Message = err.Error()
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, returnError.Body)
		return
	}

	// with gorm.Model field, default delete is a soft delete. Unscoped() force delete
	db.Unscoped().Delete(&countryDB)

	// get an instance (not staged) from DB instance, and call callback function
	countryDeleted := new(models.Country)
	countryDB.CopyBasicFieldsToCountry(countryDeleted)

	// get stage instance from DB instance, and call callback function
	countryStaged := backRepo.BackRepoCountry.Map_CountryDBID_CountryPtr[countryDB.ID]
	if countryStaged != nil {
		models.AfterDeleteFromFront(backRepo.GetStage(), countryStaged, countryDeleted)
	}

	// a DELETE generates a back repo commit increase
	// (this will be improved with implementation of unit of work design pattern)
	backRepo.IncrementPushFromFrontNb()

	c.JSON(http.StatusOK, gin.H{"data": true})

	mutexCountry.Unlock()
}
