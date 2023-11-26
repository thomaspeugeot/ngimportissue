// generated by stacks/gong/go/models/orm_file_per_struct_back_repo.go
package orm

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"gorm.io/gorm"

	"github.com/tealeg/xlsx/v3"

	"github.com/fullstack-lang/gong/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_PointerToGongStructField_sql sql.NullBool
var dummy_PointerToGongStructField_time time.Duration
var dummy_PointerToGongStructField_sort sort.Float64Slice

// PointerToGongStructFieldAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model pointertogongstructfieldAPI
type PointerToGongStructFieldAPI struct {
	gorm.Model

	models.PointerToGongStructField_WOP

	// encoding of pointers
	PointerToGongStructFieldPointersEncoding PointerToGongStructFieldPointersEncoding
}

// PointerToGongStructFieldPointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type PointerToGongStructFieldPointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field GongStruct is a pointer to another Struct (optional or 0..1)
	// This field is generated into another field to enable AS ONE association
	GongStructID sql.NullInt64
}

// PointerToGongStructFieldDB describes a pointertogongstructfield in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model pointertogongstructfieldDB
type PointerToGongStructFieldDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field pointertogongstructfieldDB.Name
	Name_Data sql.NullString

	// Declation for basic field pointertogongstructfieldDB.Index
	Index_Data sql.NullInt64

	// Declation for basic field pointertogongstructfieldDB.CompositeStructName
	CompositeStructName_Data sql.NullString
	// encoding of pointers
	PointerToGongStructFieldPointersEncoding
}

// PointerToGongStructFieldDBs arrays pointertogongstructfieldDBs
// swagger:response pointertogongstructfieldDBsResponse
type PointerToGongStructFieldDBs []PointerToGongStructFieldDB

// PointerToGongStructFieldDBResponse provides response
// swagger:response pointertogongstructfieldDBResponse
type PointerToGongStructFieldDBResponse struct {
	PointerToGongStructFieldDB
}

// PointerToGongStructFieldWOP is a PointerToGongStructField without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type PointerToGongStructFieldWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	Index int `xlsx:"2"`

	CompositeStructName string `xlsx:"3"`
	// insertion for WOP pointer fields
}

var PointerToGongStructField_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"Index",
	"CompositeStructName",
}

type BackRepoPointerToGongStructFieldStruct struct {
	// stores PointerToGongStructFieldDB according to their gorm ID
	Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB map[uint]*PointerToGongStructFieldDB

	// stores PointerToGongStructFieldDB ID according to PointerToGongStructField address
	Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID map[*models.PointerToGongStructField]uint

	// stores PointerToGongStructField according to their gorm ID
	Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr map[uint]*models.PointerToGongStructField

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoPointerToGongStructField.stage
	return
}

func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) GetDB() *gorm.DB {
	return backRepoPointerToGongStructField.db
}

// GetPointerToGongStructFieldDBFromPointerToGongStructFieldPtr is a handy function to access the back repo instance from the stage instance
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) GetPointerToGongStructFieldDBFromPointerToGongStructFieldPtr(pointertogongstructfield *models.PointerToGongStructField) (pointertogongstructfieldDB *PointerToGongStructFieldDB) {
	id := backRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID[pointertogongstructfield]
	pointertogongstructfieldDB = backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB[id]
	return
}

// BackRepoPointerToGongStructField.CommitPhaseOne commits all staged instances of PointerToGongStructField to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for pointertogongstructfield := range stage.PointerToGongStructFields {
		backRepoPointerToGongStructField.CommitPhaseOneInstance(pointertogongstructfield)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, pointertogongstructfield := range backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr {
		if _, ok := stage.PointerToGongStructFields[pointertogongstructfield]; !ok {
			backRepoPointerToGongStructField.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoPointerToGongStructField.CommitDeleteInstance commits deletion of PointerToGongStructField to the BackRepo
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CommitDeleteInstance(id uint) (Error error) {

	pointertogongstructfield := backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[id]

	// pointertogongstructfield is not staged anymore, remove pointertogongstructfieldDB
	pointertogongstructfieldDB := backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB[id]
	query := backRepoPointerToGongStructField.db.Unscoped().Delete(&pointertogongstructfieldDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID, pointertogongstructfield)
	delete(backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr, id)
	delete(backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB, id)

	return
}

// BackRepoPointerToGongStructField.CommitPhaseOneInstance commits pointertogongstructfield staged instances of PointerToGongStructField to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CommitPhaseOneInstance(pointertogongstructfield *models.PointerToGongStructField) (Error error) {

	// check if the pointertogongstructfield is not commited yet
	if _, ok := backRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID[pointertogongstructfield]; ok {
		return
	}

	// initiate pointertogongstructfield
	var pointertogongstructfieldDB PointerToGongStructFieldDB
	pointertogongstructfieldDB.CopyBasicFieldsFromPointerToGongStructField(pointertogongstructfield)

	query := backRepoPointerToGongStructField.db.Create(&pointertogongstructfieldDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID[pointertogongstructfield] = pointertogongstructfieldDB.ID
	backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[pointertogongstructfieldDB.ID] = pointertogongstructfield
	backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB[pointertogongstructfieldDB.ID] = &pointertogongstructfieldDB

	return
}

// BackRepoPointerToGongStructField.CommitPhaseTwo commits all staged instances of PointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, pointertogongstructfield := range backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr {
		backRepoPointerToGongStructField.CommitPhaseTwoInstance(backRepo, idx, pointertogongstructfield)
	}

	return
}

// BackRepoPointerToGongStructField.CommitPhaseTwoInstance commits {{structname }} of models.PointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, pointertogongstructfield *models.PointerToGongStructField) (Error error) {

	// fetch matching pointertogongstructfieldDB
	if pointertogongstructfieldDB, ok := backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB[idx]; ok {

		pointertogongstructfieldDB.CopyBasicFieldsFromPointerToGongStructField(pointertogongstructfield)

		// insertion point for translating pointers encodings into actual pointers
		// commit pointer value pointertogongstructfield.GongStruct translates to updating the pointertogongstructfield.GongStructID
		pointertogongstructfieldDB.GongStructID.Valid = true // allow for a 0 value (nil association)
		if pointertogongstructfield.GongStruct != nil {
			if GongStructId, ok := backRepo.BackRepoGongStruct.Map_GongStructPtr_GongStructDBID[pointertogongstructfield.GongStruct]; ok {
				pointertogongstructfieldDB.GongStructID.Int64 = int64(GongStructId)
				pointertogongstructfieldDB.GongStructID.Valid = true
			}
		} else {
			pointertogongstructfieldDB.GongStructID.Int64 = 0
			pointertogongstructfieldDB.GongStructID.Valid = true
		}

		query := backRepoPointerToGongStructField.db.Save(&pointertogongstructfieldDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown PointerToGongStructField intance %s", pointertogongstructfield.Name))
		return err
	}

	return
}

// BackRepoPointerToGongStructField.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CheckoutPhaseOne() (Error error) {

	pointertogongstructfieldDBArray := make([]PointerToGongStructFieldDB, 0)
	query := backRepoPointerToGongStructField.db.Find(&pointertogongstructfieldDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	pointertogongstructfieldInstancesToBeRemovedFromTheStage := make(map[*models.PointerToGongStructField]any)
	for key, value := range backRepoPointerToGongStructField.stage.PointerToGongStructFields {
		pointertogongstructfieldInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, pointertogongstructfieldDB := range pointertogongstructfieldDBArray {
		backRepoPointerToGongStructField.CheckoutPhaseOneInstance(&pointertogongstructfieldDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		pointertogongstructfield, ok := backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[pointertogongstructfieldDB.ID]
		if ok {
			delete(pointertogongstructfieldInstancesToBeRemovedFromTheStage, pointertogongstructfield)
		}
	}

	// remove from stage and back repo's 3 maps all pointertogongstructfields that are not in the checkout
	for pointertogongstructfield := range pointertogongstructfieldInstancesToBeRemovedFromTheStage {
		pointertogongstructfield.Unstage(backRepoPointerToGongStructField.GetStage())

		// remove instance from the back repo 3 maps
		pointertogongstructfieldID := backRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID[pointertogongstructfield]
		delete(backRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID, pointertogongstructfield)
		delete(backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB, pointertogongstructfieldID)
		delete(backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr, pointertogongstructfieldID)
	}

	return
}

// CheckoutPhaseOneInstance takes a pointertogongstructfieldDB that has been found in the DB, updates the backRepo and stages the
// models version of the pointertogongstructfieldDB
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CheckoutPhaseOneInstance(pointertogongstructfieldDB *PointerToGongStructFieldDB) (Error error) {

	pointertogongstructfield, ok := backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[pointertogongstructfieldDB.ID]
	if !ok {
		pointertogongstructfield = new(models.PointerToGongStructField)

		backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[pointertogongstructfieldDB.ID] = pointertogongstructfield
		backRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID[pointertogongstructfield] = pointertogongstructfieldDB.ID

		// append model store with the new element
		pointertogongstructfield.Name = pointertogongstructfieldDB.Name_Data.String
		pointertogongstructfield.Stage(backRepoPointerToGongStructField.GetStage())
	}
	pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructField(pointertogongstructfield)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	pointertogongstructfield.Stage(backRepoPointerToGongStructField.GetStage())

	// preserve pointer to pointertogongstructfieldDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB)[pointertogongstructfieldDB hold variable pointers
	pointertogongstructfieldDB_Data := *pointertogongstructfieldDB
	preservedPtrToPointerToGongStructField := &pointertogongstructfieldDB_Data
	backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB[pointertogongstructfieldDB.ID] = preservedPtrToPointerToGongStructField

	return
}

// BackRepoPointerToGongStructField.CheckoutPhaseTwo Checkouts all staged instances of PointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, pointertogongstructfieldDB := range backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB {
		backRepoPointerToGongStructField.CheckoutPhaseTwoInstance(backRepo, pointertogongstructfieldDB)
	}
	return
}

// BackRepoPointerToGongStructField.CheckoutPhaseTwoInstance Checkouts staged instances of PointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, pointertogongstructfieldDB *PointerToGongStructFieldDB) (Error error) {

	pointertogongstructfield := backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr[pointertogongstructfieldDB.ID]

	pointertogongstructfieldDB.DecodePointers(backRepo, pointertogongstructfield)

	return
}

func (pointertogongstructfieldDB *PointerToGongStructFieldDB) DecodePointers(backRepo *BackRepoStruct, pointertogongstructfield *models.PointerToGongStructField) {

	// insertion point for checkout of pointer encoding
	// GongStruct field
	pointertogongstructfield.GongStruct = nil
	if pointertogongstructfieldDB.GongStructID.Int64 != 0 {
		pointertogongstructfield.GongStruct = backRepo.BackRepoGongStruct.Map_GongStructDBID_GongStructPtr[uint(pointertogongstructfieldDB.GongStructID.Int64)]
	}
	return
}

// CommitPointerToGongStructField allows commit of a single pointertogongstructfield (if already staged)
func (backRepo *BackRepoStruct) CommitPointerToGongStructField(pointertogongstructfield *models.PointerToGongStructField) {
	backRepo.BackRepoPointerToGongStructField.CommitPhaseOneInstance(pointertogongstructfield)
	if id, ok := backRepo.BackRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID[pointertogongstructfield]; ok {
		backRepo.BackRepoPointerToGongStructField.CommitPhaseTwoInstance(backRepo, id, pointertogongstructfield)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitPointerToGongStructField allows checkout of a single pointertogongstructfield (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutPointerToGongStructField(pointertogongstructfield *models.PointerToGongStructField) {
	// check if the pointertogongstructfield is staged
	if _, ok := backRepo.BackRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID[pointertogongstructfield]; ok {

		if id, ok := backRepo.BackRepoPointerToGongStructField.Map_PointerToGongStructFieldPtr_PointerToGongStructFieldDBID[pointertogongstructfield]; ok {
			var pointertogongstructfieldDB PointerToGongStructFieldDB
			pointertogongstructfieldDB.ID = id

			if err := backRepo.BackRepoPointerToGongStructField.db.First(&pointertogongstructfieldDB, id).Error; err != nil {
				log.Fatalln("CheckoutPointerToGongStructField : Problem with getting object with id:", id)
			}
			backRepo.BackRepoPointerToGongStructField.CheckoutPhaseOneInstance(&pointertogongstructfieldDB)
			backRepo.BackRepoPointerToGongStructField.CheckoutPhaseTwoInstance(backRepo, &pointertogongstructfieldDB)
		}
	}
}

// CopyBasicFieldsFromPointerToGongStructField
func (pointertogongstructfieldDB *PointerToGongStructFieldDB) CopyBasicFieldsFromPointerToGongStructField(pointertogongstructfield *models.PointerToGongStructField) {
	// insertion point for fields commit

	pointertogongstructfieldDB.Name_Data.String = pointertogongstructfield.Name
	pointertogongstructfieldDB.Name_Data.Valid = true

	pointertogongstructfieldDB.Index_Data.Int64 = int64(pointertogongstructfield.Index)
	pointertogongstructfieldDB.Index_Data.Valid = true

	pointertogongstructfieldDB.CompositeStructName_Data.String = pointertogongstructfield.CompositeStructName
	pointertogongstructfieldDB.CompositeStructName_Data.Valid = true
}

// CopyBasicFieldsFromPointerToGongStructField_WOP
func (pointertogongstructfieldDB *PointerToGongStructFieldDB) CopyBasicFieldsFromPointerToGongStructField_WOP(pointertogongstructfield *models.PointerToGongStructField_WOP) {
	// insertion point for fields commit

	pointertogongstructfieldDB.Name_Data.String = pointertogongstructfield.Name
	pointertogongstructfieldDB.Name_Data.Valid = true

	pointertogongstructfieldDB.Index_Data.Int64 = int64(pointertogongstructfield.Index)
	pointertogongstructfieldDB.Index_Data.Valid = true

	pointertogongstructfieldDB.CompositeStructName_Data.String = pointertogongstructfield.CompositeStructName
	pointertogongstructfieldDB.CompositeStructName_Data.Valid = true
}

// CopyBasicFieldsFromPointerToGongStructFieldWOP
func (pointertogongstructfieldDB *PointerToGongStructFieldDB) CopyBasicFieldsFromPointerToGongStructFieldWOP(pointertogongstructfield *PointerToGongStructFieldWOP) {
	// insertion point for fields commit

	pointertogongstructfieldDB.Name_Data.String = pointertogongstructfield.Name
	pointertogongstructfieldDB.Name_Data.Valid = true

	pointertogongstructfieldDB.Index_Data.Int64 = int64(pointertogongstructfield.Index)
	pointertogongstructfieldDB.Index_Data.Valid = true

	pointertogongstructfieldDB.CompositeStructName_Data.String = pointertogongstructfield.CompositeStructName
	pointertogongstructfieldDB.CompositeStructName_Data.Valid = true
}

// CopyBasicFieldsToPointerToGongStructField
func (pointertogongstructfieldDB *PointerToGongStructFieldDB) CopyBasicFieldsToPointerToGongStructField(pointertogongstructfield *models.PointerToGongStructField) {
	// insertion point for checkout of basic fields (back repo to stage)
	pointertogongstructfield.Name = pointertogongstructfieldDB.Name_Data.String
	pointertogongstructfield.Index = int(pointertogongstructfieldDB.Index_Data.Int64)
	pointertogongstructfield.CompositeStructName = pointertogongstructfieldDB.CompositeStructName_Data.String
}

// CopyBasicFieldsToPointerToGongStructField_WOP
func (pointertogongstructfieldDB *PointerToGongStructFieldDB) CopyBasicFieldsToPointerToGongStructField_WOP(pointertogongstructfield *models.PointerToGongStructField_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	pointertogongstructfield.Name = pointertogongstructfieldDB.Name_Data.String
	pointertogongstructfield.Index = int(pointertogongstructfieldDB.Index_Data.Int64)
	pointertogongstructfield.CompositeStructName = pointertogongstructfieldDB.CompositeStructName_Data.String
}

// CopyBasicFieldsToPointerToGongStructFieldWOP
func (pointertogongstructfieldDB *PointerToGongStructFieldDB) CopyBasicFieldsToPointerToGongStructFieldWOP(pointertogongstructfield *PointerToGongStructFieldWOP) {
	pointertogongstructfield.ID = int(pointertogongstructfieldDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	pointertogongstructfield.Name = pointertogongstructfieldDB.Name_Data.String
	pointertogongstructfield.Index = int(pointertogongstructfieldDB.Index_Data.Int64)
	pointertogongstructfield.CompositeStructName = pointertogongstructfieldDB.CompositeStructName_Data.String
}

// Backup generates a json file from a slice of all PointerToGongStructFieldDB instances in the backrepo
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "PointerToGongStructFieldDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PointerToGongStructFieldDB, 0)
	for _, pointertogongstructfieldDB := range backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB {
		forBackup = append(forBackup, pointertogongstructfieldDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json PointerToGongStructField ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json PointerToGongStructField file", err.Error())
	}
}

// Backup generates a json file from a slice of all PointerToGongStructFieldDB instances in the backrepo
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*PointerToGongStructFieldDB, 0)
	for _, pointertogongstructfieldDB := range backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB {
		forBackup = append(forBackup, pointertogongstructfieldDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("PointerToGongStructField")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&PointerToGongStructField_Fields, -1)
	for _, pointertogongstructfieldDB := range forBackup {

		var pointertogongstructfieldWOP PointerToGongStructFieldWOP
		pointertogongstructfieldDB.CopyBasicFieldsToPointerToGongStructFieldWOP(&pointertogongstructfieldWOP)

		row := sh.AddRow()
		row.WriteStruct(&pointertogongstructfieldWOP, -1)
	}
}

// RestoreXL from the "PointerToGongStructField" sheet all PointerToGongStructFieldDB instances
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoPointerToGongStructFieldid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["PointerToGongStructField"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoPointerToGongStructField.rowVisitorPointerToGongStructField)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) rowVisitorPointerToGongStructField(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var pointertogongstructfieldWOP PointerToGongStructFieldWOP
		row.ReadStruct(&pointertogongstructfieldWOP)

		// add the unmarshalled struct to the stage
		pointertogongstructfieldDB := new(PointerToGongStructFieldDB)
		pointertogongstructfieldDB.CopyBasicFieldsFromPointerToGongStructFieldWOP(&pointertogongstructfieldWOP)

		pointertogongstructfieldDB_ID_atBackupTime := pointertogongstructfieldDB.ID
		pointertogongstructfieldDB.ID = 0
		query := backRepoPointerToGongStructField.db.Create(pointertogongstructfieldDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB[pointertogongstructfieldDB.ID] = pointertogongstructfieldDB
		BackRepoPointerToGongStructFieldid_atBckpTime_newID[pointertogongstructfieldDB_ID_atBackupTime] = pointertogongstructfieldDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "PointerToGongStructFieldDB.json" in dirPath that stores an array
// of PointerToGongStructFieldDB and stores it in the database
// the map BackRepoPointerToGongStructFieldid_atBckpTime_newID is updated accordingly
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoPointerToGongStructFieldid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "PointerToGongStructFieldDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json PointerToGongStructField file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*PointerToGongStructFieldDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB
	for _, pointertogongstructfieldDB := range forRestore {

		pointertogongstructfieldDB_ID_atBackupTime := pointertogongstructfieldDB.ID
		pointertogongstructfieldDB.ID = 0
		query := backRepoPointerToGongStructField.db.Create(pointertogongstructfieldDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB[pointertogongstructfieldDB.ID] = pointertogongstructfieldDB
		BackRepoPointerToGongStructFieldid_atBckpTime_newID[pointertogongstructfieldDB_ID_atBackupTime] = pointertogongstructfieldDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json PointerToGongStructField file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<PointerToGongStructField>id_atBckpTime_newID
// to compute new index
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) RestorePhaseTwo() {

	for _, pointertogongstructfieldDB := range backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB {

		// next line of code is to avert unused variable compilation error
		_ = pointertogongstructfieldDB

		// insertion point for reindexing pointers encoding
		// reindexing GongStruct field
		if pointertogongstructfieldDB.GongStructID.Int64 != 0 {
			pointertogongstructfieldDB.GongStructID.Int64 = int64(BackRepoGongStructid_atBckpTime_newID[uint(pointertogongstructfieldDB.GongStructID.Int64)])
			pointertogongstructfieldDB.GongStructID.Valid = true
		}

		// update databse with new index encoding
		query := backRepoPointerToGongStructField.db.Model(pointertogongstructfieldDB).Updates(*pointertogongstructfieldDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoPointerToGongStructField.ResetReversePointers commits all staged instances of PointerToGongStructField to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, pointertogongstructfield := range backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldPtr {
		backRepoPointerToGongStructField.ResetReversePointersInstance(backRepo, idx, pointertogongstructfield)
	}

	return
}

func (backRepoPointerToGongStructField *BackRepoPointerToGongStructFieldStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, pointertogongstructfield *models.PointerToGongStructField) (Error error) {

	// fetch matching pointertogongstructfieldDB
	if pointertogongstructfieldDB, ok := backRepoPointerToGongStructField.Map_PointerToGongStructFieldDBID_PointerToGongStructFieldDB[idx]; ok {
		_ = pointertogongstructfieldDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoPointerToGongStructFieldid_atBckpTime_newID map[uint]uint
