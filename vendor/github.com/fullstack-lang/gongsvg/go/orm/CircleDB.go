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

	"github.com/fullstack-lang/gongsvg/go/models"
)

// dummy variable to have the import declaration wihthout compile failure (even if no code needing this import is generated)
var dummy_Circle_sql sql.NullBool
var dummy_Circle_time time.Duration
var dummy_Circle_sort sort.Float64Slice

// CircleAPI is the input in POST API
//
// for POST, API, one needs the fields of the model as well as the fields
// from associations ("Has One" and "Has Many") that are generated to
// fullfill the ORM requirements for associations
//
// swagger:model circleAPI
type CircleAPI struct {
	gorm.Model

	models.Circle_WOP

	// encoding of pointers
	CirclePointersEncoding CirclePointersEncoding
}

// CirclePointersEncoding encodes pointers to Struct and
// reverse pointers of slice of poitners to Struct
type CirclePointersEncoding struct {
	// insertion for pointer fields encoding declaration

	// field Animations is a slice of pointers to another Struct (optional or 0..1)
	Animations IntSlice `gorm:"type:TEXT"`
}

// CircleDB describes a circle in the database
//
// It incorporates the GORM ID, basic fields from the model (because they can be serialized),
// the encoded version of pointers
//
// swagger:model circleDB
type CircleDB struct {
	gorm.Model

	// insertion for basic fields declaration

	// Declation for basic field circleDB.Name
	Name_Data sql.NullString

	// Declation for basic field circleDB.CX
	CX_Data sql.NullFloat64

	// Declation for basic field circleDB.CY
	CY_Data sql.NullFloat64

	// Declation for basic field circleDB.Radius
	Radius_Data sql.NullFloat64

	// Declation for basic field circleDB.Color
	Color_Data sql.NullString

	// Declation for basic field circleDB.FillOpacity
	FillOpacity_Data sql.NullFloat64

	// Declation for basic field circleDB.Stroke
	Stroke_Data sql.NullString

	// Declation for basic field circleDB.StrokeWidth
	StrokeWidth_Data sql.NullFloat64

	// Declation for basic field circleDB.StrokeDashArray
	StrokeDashArray_Data sql.NullString

	// Declation for basic field circleDB.StrokeDashArrayWhenSelected
	StrokeDashArrayWhenSelected_Data sql.NullString

	// Declation for basic field circleDB.Transform
	Transform_Data sql.NullString
	// encoding of pointers
	CirclePointersEncoding
}

// CircleDBs arrays circleDBs
// swagger:response circleDBsResponse
type CircleDBs []CircleDB

// CircleDBResponse provides response
// swagger:response circleDBResponse
type CircleDBResponse struct {
	CircleDB
}

// CircleWOP is a Circle without pointers (WOP is an acronym for "Without Pointers")
// it holds the same basic fields but pointers are encoded into uint
type CircleWOP struct {
	ID int `xlsx:"0"`

	// insertion for WOP basic fields

	Name string `xlsx:"1"`

	CX float64 `xlsx:"2"`

	CY float64 `xlsx:"3"`

	Radius float64 `xlsx:"4"`

	Color string `xlsx:"5"`

	FillOpacity float64 `xlsx:"6"`

	Stroke string `xlsx:"7"`

	StrokeWidth float64 `xlsx:"8"`

	StrokeDashArray string `xlsx:"9"`

	StrokeDashArrayWhenSelected string `xlsx:"10"`

	Transform string `xlsx:"11"`
	// insertion for WOP pointer fields
}

var Circle_Fields = []string{
	// insertion for WOP basic fields
	"ID",
	"Name",
	"CX",
	"CY",
	"Radius",
	"Color",
	"FillOpacity",
	"Stroke",
	"StrokeWidth",
	"StrokeDashArray",
	"StrokeDashArrayWhenSelected",
	"Transform",
}

type BackRepoCircleStruct struct {
	// stores CircleDB according to their gorm ID
	Map_CircleDBID_CircleDB map[uint]*CircleDB

	// stores CircleDB ID according to Circle address
	Map_CirclePtr_CircleDBID map[*models.Circle]uint

	// stores Circle according to their gorm ID
	Map_CircleDBID_CirclePtr map[uint]*models.Circle

	db *gorm.DB

	stage *models.StageStruct
}

func (backRepoCircle *BackRepoCircleStruct) GetStage() (stage *models.StageStruct) {
	stage = backRepoCircle.stage
	return
}

func (backRepoCircle *BackRepoCircleStruct) GetDB() *gorm.DB {
	return backRepoCircle.db
}

// GetCircleDBFromCirclePtr is a handy function to access the back repo instance from the stage instance
func (backRepoCircle *BackRepoCircleStruct) GetCircleDBFromCirclePtr(circle *models.Circle) (circleDB *CircleDB) {
	id := backRepoCircle.Map_CirclePtr_CircleDBID[circle]
	circleDB = backRepoCircle.Map_CircleDBID_CircleDB[id]
	return
}

// BackRepoCircle.CommitPhaseOne commits all staged instances of Circle to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCircle *BackRepoCircleStruct) CommitPhaseOne(stage *models.StageStruct) (Error error) {

	for circle := range stage.Circles {
		backRepoCircle.CommitPhaseOneInstance(circle)
	}

	// parse all backRepo instance and checks wether some instance have been unstaged
	// in this case, remove them from the back repo
	for id, circle := range backRepoCircle.Map_CircleDBID_CirclePtr {
		if _, ok := stage.Circles[circle]; !ok {
			backRepoCircle.CommitDeleteInstance(id)
		}
	}

	return
}

// BackRepoCircle.CommitDeleteInstance commits deletion of Circle to the BackRepo
func (backRepoCircle *BackRepoCircleStruct) CommitDeleteInstance(id uint) (Error error) {

	circle := backRepoCircle.Map_CircleDBID_CirclePtr[id]

	// circle is not staged anymore, remove circleDB
	circleDB := backRepoCircle.Map_CircleDBID_CircleDB[id]
	query := backRepoCircle.db.Unscoped().Delete(&circleDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	delete(backRepoCircle.Map_CirclePtr_CircleDBID, circle)
	delete(backRepoCircle.Map_CircleDBID_CirclePtr, id)
	delete(backRepoCircle.Map_CircleDBID_CircleDB, id)

	return
}

// BackRepoCircle.CommitPhaseOneInstance commits circle staged instances of Circle to the BackRepo
// Phase One is the creation of instance in the database if it is not yet done to get the unique ID for each staged instance
func (backRepoCircle *BackRepoCircleStruct) CommitPhaseOneInstance(circle *models.Circle) (Error error) {

	// check if the circle is not commited yet
	if _, ok := backRepoCircle.Map_CirclePtr_CircleDBID[circle]; ok {
		return
	}

	// initiate circle
	var circleDB CircleDB
	circleDB.CopyBasicFieldsFromCircle(circle)

	query := backRepoCircle.db.Create(&circleDB)
	if query.Error != nil {
		log.Fatal(query.Error)
	}

	// update stores
	backRepoCircle.Map_CirclePtr_CircleDBID[circle] = circleDB.ID
	backRepoCircle.Map_CircleDBID_CirclePtr[circleDB.ID] = circle
	backRepoCircle.Map_CircleDBID_CircleDB[circleDB.ID] = &circleDB

	return
}

// BackRepoCircle.CommitPhaseTwo commits all staged instances of Circle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCircle *BackRepoCircleStruct) CommitPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	for idx, circle := range backRepoCircle.Map_CircleDBID_CirclePtr {
		backRepoCircle.CommitPhaseTwoInstance(backRepo, idx, circle)
	}

	return
}

// BackRepoCircle.CommitPhaseTwoInstance commits {{structname }} of models.Circle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCircle *BackRepoCircleStruct) CommitPhaseTwoInstance(backRepo *BackRepoStruct, idx uint, circle *models.Circle) (Error error) {

	// fetch matching circleDB
	if circleDB, ok := backRepoCircle.Map_CircleDBID_CircleDB[idx]; ok {

		circleDB.CopyBasicFieldsFromCircle(circle)

		// insertion point for translating pointers encodings into actual pointers
		// 1. reset
		circleDB.CirclePointersEncoding.Animations = make([]int, 0)
		// 2. encode
		for _, animateAssocEnd := range circle.Animations {
			animateAssocEnd_DB :=
				backRepo.BackRepoAnimate.GetAnimateDBFromAnimatePtr(animateAssocEnd)
			circleDB.CirclePointersEncoding.Animations =
				append(circleDB.CirclePointersEncoding.Animations, int(animateAssocEnd_DB.ID))
		}

		query := backRepoCircle.db.Save(&circleDB)
		if query.Error != nil {
			log.Fatalln(query.Error)
		}

	} else {
		err := errors.New(
			fmt.Sprintf("Unkown Circle intance %s", circle.Name))
		return err
	}

	return
}

// BackRepoCircle.CheckoutPhaseOne Checkouts all BackRepo instances to the Stage
//
// Phase One will result in having instances on the stage aligned with the back repo
// pointers are not initialized yet (this is for phase two)
func (backRepoCircle *BackRepoCircleStruct) CheckoutPhaseOne() (Error error) {

	circleDBArray := make([]CircleDB, 0)
	query := backRepoCircle.db.Find(&circleDBArray)
	if query.Error != nil {
		return query.Error
	}

	// list of instances to be removed
	// start from the initial map on the stage and remove instances that have been checked out
	circleInstancesToBeRemovedFromTheStage := make(map[*models.Circle]any)
	for key, value := range backRepoCircle.stage.Circles {
		circleInstancesToBeRemovedFromTheStage[key] = value
	}

	// copy orm objects to the the map
	for _, circleDB := range circleDBArray {
		backRepoCircle.CheckoutPhaseOneInstance(&circleDB)

		// do not remove this instance from the stage, therefore
		// remove instance from the list of instances to be be removed from the stage
		circle, ok := backRepoCircle.Map_CircleDBID_CirclePtr[circleDB.ID]
		if ok {
			delete(circleInstancesToBeRemovedFromTheStage, circle)
		}
	}

	// remove from stage and back repo's 3 maps all circles that are not in the checkout
	for circle := range circleInstancesToBeRemovedFromTheStage {
		circle.Unstage(backRepoCircle.GetStage())

		// remove instance from the back repo 3 maps
		circleID := backRepoCircle.Map_CirclePtr_CircleDBID[circle]
		delete(backRepoCircle.Map_CirclePtr_CircleDBID, circle)
		delete(backRepoCircle.Map_CircleDBID_CircleDB, circleID)
		delete(backRepoCircle.Map_CircleDBID_CirclePtr, circleID)
	}

	return
}

// CheckoutPhaseOneInstance takes a circleDB that has been found in the DB, updates the backRepo and stages the
// models version of the circleDB
func (backRepoCircle *BackRepoCircleStruct) CheckoutPhaseOneInstance(circleDB *CircleDB) (Error error) {

	circle, ok := backRepoCircle.Map_CircleDBID_CirclePtr[circleDB.ID]
	if !ok {
		circle = new(models.Circle)

		backRepoCircle.Map_CircleDBID_CirclePtr[circleDB.ID] = circle
		backRepoCircle.Map_CirclePtr_CircleDBID[circle] = circleDB.ID

		// append model store with the new element
		circle.Name = circleDB.Name_Data.String
		circle.Stage(backRepoCircle.GetStage())
	}
	circleDB.CopyBasicFieldsToCircle(circle)

	// in some cases, the instance might have been unstaged. It is necessary to stage it again
	circle.Stage(backRepoCircle.GetStage())

	// preserve pointer to circleDB. Otherwise, pointer will is recycled and the map of pointers
	// Map_CircleDBID_CircleDB)[circleDB hold variable pointers
	circleDB_Data := *circleDB
	preservedPtrToCircle := &circleDB_Data
	backRepoCircle.Map_CircleDBID_CircleDB[circleDB.ID] = preservedPtrToCircle

	return
}

// BackRepoCircle.CheckoutPhaseTwo Checkouts all staged instances of Circle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCircle *BackRepoCircleStruct) CheckoutPhaseTwo(backRepo *BackRepoStruct) (Error error) {

	// parse all DB instance and update all pointer fields of the translated models instance
	for _, circleDB := range backRepoCircle.Map_CircleDBID_CircleDB {
		backRepoCircle.CheckoutPhaseTwoInstance(backRepo, circleDB)
	}
	return
}

// BackRepoCircle.CheckoutPhaseTwoInstance Checkouts staged instances of Circle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCircle *BackRepoCircleStruct) CheckoutPhaseTwoInstance(backRepo *BackRepoStruct, circleDB *CircleDB) (Error error) {

	circle := backRepoCircle.Map_CircleDBID_CirclePtr[circleDB.ID]

	circleDB.DecodePointers(backRepo, circle)

	return
}

func (circleDB *CircleDB) DecodePointers(backRepo *BackRepoStruct, circle *models.Circle) {

	// insertion point for checkout of pointer encoding
	// This loop redeem circle.Animations in the stage from the encode in the back repo
	// It parses all AnimateDB in the back repo and if the reverse pointer encoding matches the back repo ID
	// it appends the stage instance
	// 1. reset the slice
	circle.Animations = circle.Animations[:0]
	for _, _Animateid := range circleDB.CirclePointersEncoding.Animations {
		circle.Animations = append(circle.Animations, backRepo.BackRepoAnimate.Map_AnimateDBID_AnimatePtr[uint(_Animateid)])
	}

	return
}

// CommitCircle allows commit of a single circle (if already staged)
func (backRepo *BackRepoStruct) CommitCircle(circle *models.Circle) {
	backRepo.BackRepoCircle.CommitPhaseOneInstance(circle)
	if id, ok := backRepo.BackRepoCircle.Map_CirclePtr_CircleDBID[circle]; ok {
		backRepo.BackRepoCircle.CommitPhaseTwoInstance(backRepo, id, circle)
	}
	backRepo.CommitFromBackNb = backRepo.CommitFromBackNb + 1
}

// CommitCircle allows checkout of a single circle (if already staged and with a BackRepo id)
func (backRepo *BackRepoStruct) CheckoutCircle(circle *models.Circle) {
	// check if the circle is staged
	if _, ok := backRepo.BackRepoCircle.Map_CirclePtr_CircleDBID[circle]; ok {

		if id, ok := backRepo.BackRepoCircle.Map_CirclePtr_CircleDBID[circle]; ok {
			var circleDB CircleDB
			circleDB.ID = id

			if err := backRepo.BackRepoCircle.db.First(&circleDB, id).Error; err != nil {
				log.Fatalln("CheckoutCircle : Problem with getting object with id:", id)
			}
			backRepo.BackRepoCircle.CheckoutPhaseOneInstance(&circleDB)
			backRepo.BackRepoCircle.CheckoutPhaseTwoInstance(backRepo, &circleDB)
		}
	}
}

// CopyBasicFieldsFromCircle
func (circleDB *CircleDB) CopyBasicFieldsFromCircle(circle *models.Circle) {
	// insertion point for fields commit

	circleDB.Name_Data.String = circle.Name
	circleDB.Name_Data.Valid = true

	circleDB.CX_Data.Float64 = circle.CX
	circleDB.CX_Data.Valid = true

	circleDB.CY_Data.Float64 = circle.CY
	circleDB.CY_Data.Valid = true

	circleDB.Radius_Data.Float64 = circle.Radius
	circleDB.Radius_Data.Valid = true

	circleDB.Color_Data.String = circle.Color
	circleDB.Color_Data.Valid = true

	circleDB.FillOpacity_Data.Float64 = circle.FillOpacity
	circleDB.FillOpacity_Data.Valid = true

	circleDB.Stroke_Data.String = circle.Stroke
	circleDB.Stroke_Data.Valid = true

	circleDB.StrokeWidth_Data.Float64 = circle.StrokeWidth
	circleDB.StrokeWidth_Data.Valid = true

	circleDB.StrokeDashArray_Data.String = circle.StrokeDashArray
	circleDB.StrokeDashArray_Data.Valid = true

	circleDB.StrokeDashArrayWhenSelected_Data.String = circle.StrokeDashArrayWhenSelected
	circleDB.StrokeDashArrayWhenSelected_Data.Valid = true

	circleDB.Transform_Data.String = circle.Transform
	circleDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromCircle_WOP
func (circleDB *CircleDB) CopyBasicFieldsFromCircle_WOP(circle *models.Circle_WOP) {
	// insertion point for fields commit

	circleDB.Name_Data.String = circle.Name
	circleDB.Name_Data.Valid = true

	circleDB.CX_Data.Float64 = circle.CX
	circleDB.CX_Data.Valid = true

	circleDB.CY_Data.Float64 = circle.CY
	circleDB.CY_Data.Valid = true

	circleDB.Radius_Data.Float64 = circle.Radius
	circleDB.Radius_Data.Valid = true

	circleDB.Color_Data.String = circle.Color
	circleDB.Color_Data.Valid = true

	circleDB.FillOpacity_Data.Float64 = circle.FillOpacity
	circleDB.FillOpacity_Data.Valid = true

	circleDB.Stroke_Data.String = circle.Stroke
	circleDB.Stroke_Data.Valid = true

	circleDB.StrokeWidth_Data.Float64 = circle.StrokeWidth
	circleDB.StrokeWidth_Data.Valid = true

	circleDB.StrokeDashArray_Data.String = circle.StrokeDashArray
	circleDB.StrokeDashArray_Data.Valid = true

	circleDB.StrokeDashArrayWhenSelected_Data.String = circle.StrokeDashArrayWhenSelected
	circleDB.StrokeDashArrayWhenSelected_Data.Valid = true

	circleDB.Transform_Data.String = circle.Transform
	circleDB.Transform_Data.Valid = true
}

// CopyBasicFieldsFromCircleWOP
func (circleDB *CircleDB) CopyBasicFieldsFromCircleWOP(circle *CircleWOP) {
	// insertion point for fields commit

	circleDB.Name_Data.String = circle.Name
	circleDB.Name_Data.Valid = true

	circleDB.CX_Data.Float64 = circle.CX
	circleDB.CX_Data.Valid = true

	circleDB.CY_Data.Float64 = circle.CY
	circleDB.CY_Data.Valid = true

	circleDB.Radius_Data.Float64 = circle.Radius
	circleDB.Radius_Data.Valid = true

	circleDB.Color_Data.String = circle.Color
	circleDB.Color_Data.Valid = true

	circleDB.FillOpacity_Data.Float64 = circle.FillOpacity
	circleDB.FillOpacity_Data.Valid = true

	circleDB.Stroke_Data.String = circle.Stroke
	circleDB.Stroke_Data.Valid = true

	circleDB.StrokeWidth_Data.Float64 = circle.StrokeWidth
	circleDB.StrokeWidth_Data.Valid = true

	circleDB.StrokeDashArray_Data.String = circle.StrokeDashArray
	circleDB.StrokeDashArray_Data.Valid = true

	circleDB.StrokeDashArrayWhenSelected_Data.String = circle.StrokeDashArrayWhenSelected
	circleDB.StrokeDashArrayWhenSelected_Data.Valid = true

	circleDB.Transform_Data.String = circle.Transform
	circleDB.Transform_Data.Valid = true
}

// CopyBasicFieldsToCircle
func (circleDB *CircleDB) CopyBasicFieldsToCircle(circle *models.Circle) {
	// insertion point for checkout of basic fields (back repo to stage)
	circle.Name = circleDB.Name_Data.String
	circle.CX = circleDB.CX_Data.Float64
	circle.CY = circleDB.CY_Data.Float64
	circle.Radius = circleDB.Radius_Data.Float64
	circle.Color = circleDB.Color_Data.String
	circle.FillOpacity = circleDB.FillOpacity_Data.Float64
	circle.Stroke = circleDB.Stroke_Data.String
	circle.StrokeWidth = circleDB.StrokeWidth_Data.Float64
	circle.StrokeDashArray = circleDB.StrokeDashArray_Data.String
	circle.StrokeDashArrayWhenSelected = circleDB.StrokeDashArrayWhenSelected_Data.String
	circle.Transform = circleDB.Transform_Data.String
}

// CopyBasicFieldsToCircle_WOP
func (circleDB *CircleDB) CopyBasicFieldsToCircle_WOP(circle *models.Circle_WOP) {
	// insertion point for checkout of basic fields (back repo to stage)
	circle.Name = circleDB.Name_Data.String
	circle.CX = circleDB.CX_Data.Float64
	circle.CY = circleDB.CY_Data.Float64
	circle.Radius = circleDB.Radius_Data.Float64
	circle.Color = circleDB.Color_Data.String
	circle.FillOpacity = circleDB.FillOpacity_Data.Float64
	circle.Stroke = circleDB.Stroke_Data.String
	circle.StrokeWidth = circleDB.StrokeWidth_Data.Float64
	circle.StrokeDashArray = circleDB.StrokeDashArray_Data.String
	circle.StrokeDashArrayWhenSelected = circleDB.StrokeDashArrayWhenSelected_Data.String
	circle.Transform = circleDB.Transform_Data.String
}

// CopyBasicFieldsToCircleWOP
func (circleDB *CircleDB) CopyBasicFieldsToCircleWOP(circle *CircleWOP) {
	circle.ID = int(circleDB.ID)
	// insertion point for checkout of basic fields (back repo to stage)
	circle.Name = circleDB.Name_Data.String
	circle.CX = circleDB.CX_Data.Float64
	circle.CY = circleDB.CY_Data.Float64
	circle.Radius = circleDB.Radius_Data.Float64
	circle.Color = circleDB.Color_Data.String
	circle.FillOpacity = circleDB.FillOpacity_Data.Float64
	circle.Stroke = circleDB.Stroke_Data.String
	circle.StrokeWidth = circleDB.StrokeWidth_Data.Float64
	circle.StrokeDashArray = circleDB.StrokeDashArray_Data.String
	circle.StrokeDashArrayWhenSelected = circleDB.StrokeDashArrayWhenSelected_Data.String
	circle.Transform = circleDB.Transform_Data.String
}

// Backup generates a json file from a slice of all CircleDB instances in the backrepo
func (backRepoCircle *BackRepoCircleStruct) Backup(dirPath string) {

	filename := filepath.Join(dirPath, "CircleDB.json")

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CircleDB, 0)
	for _, circleDB := range backRepoCircle.Map_CircleDBID_CircleDB {
		forBackup = append(forBackup, circleDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	file, err := json.MarshalIndent(forBackup, "", " ")

	if err != nil {
		log.Fatal("Cannot json Circle ", filename, " ", err.Error())
	}

	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		log.Fatal("Cannot write the json Circle file", err.Error())
	}
}

// Backup generates a json file from a slice of all CircleDB instances in the backrepo
func (backRepoCircle *BackRepoCircleStruct) BackupXL(file *xlsx.File) {

	// organize the map into an array with increasing IDs, in order to have repoductible
	// backup file
	forBackup := make([]*CircleDB, 0)
	for _, circleDB := range backRepoCircle.Map_CircleDBID_CircleDB {
		forBackup = append(forBackup, circleDB)
	}

	sort.Slice(forBackup[:], func(i, j int) bool {
		return forBackup[i].ID < forBackup[j].ID
	})

	sh, err := file.AddSheet("Circle")
	if err != nil {
		log.Fatal("Cannot add XL file", err.Error())
	}
	_ = sh

	row := sh.AddRow()
	row.WriteSlice(&Circle_Fields, -1)
	for _, circleDB := range forBackup {

		var circleWOP CircleWOP
		circleDB.CopyBasicFieldsToCircleWOP(&circleWOP)

		row := sh.AddRow()
		row.WriteStruct(&circleWOP, -1)
	}
}

// RestoreXL from the "Circle" sheet all CircleDB instances
func (backRepoCircle *BackRepoCircleStruct) RestoreXLPhaseOne(file *xlsx.File) {

	// resets the map
	BackRepoCircleid_atBckpTime_newID = make(map[uint]uint)

	sh, ok := file.Sheet["Circle"]
	_ = sh
	if !ok {
		log.Fatal(errors.New("sheet not found"))
	}

	// log.Println("Max row is", sh.MaxRow)
	err := sh.ForEachRow(backRepoCircle.rowVisitorCircle)
	if err != nil {
		log.Fatal("Err=", err)
	}
}

func (backRepoCircle *BackRepoCircleStruct) rowVisitorCircle(row *xlsx.Row) error {

	log.Printf("row line %d\n", row.GetCoordinate())
	log.Println(row)

	// skip first line
	if row.GetCoordinate() > 0 {
		var circleWOP CircleWOP
		row.ReadStruct(&circleWOP)

		// add the unmarshalled struct to the stage
		circleDB := new(CircleDB)
		circleDB.CopyBasicFieldsFromCircleWOP(&circleWOP)

		circleDB_ID_atBackupTime := circleDB.ID
		circleDB.ID = 0
		query := backRepoCircle.db.Create(circleDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoCircle.Map_CircleDBID_CircleDB[circleDB.ID] = circleDB
		BackRepoCircleid_atBckpTime_newID[circleDB_ID_atBackupTime] = circleDB.ID
	}
	return nil
}

// RestorePhaseOne read the file "CircleDB.json" in dirPath that stores an array
// of CircleDB and stores it in the database
// the map BackRepoCircleid_atBckpTime_newID is updated accordingly
func (backRepoCircle *BackRepoCircleStruct) RestorePhaseOne(dirPath string) {

	// resets the map
	BackRepoCircleid_atBckpTime_newID = make(map[uint]uint)

	filename := filepath.Join(dirPath, "CircleDB.json")
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal("Cannot restore/open the json Circle file", filename, " ", err.Error())
	}

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var forRestore []*CircleDB

	err = json.Unmarshal(byteValue, &forRestore)

	// fill up Map_CircleDBID_CircleDB
	for _, circleDB := range forRestore {

		circleDB_ID_atBackupTime := circleDB.ID
		circleDB.ID = 0
		query := backRepoCircle.db.Create(circleDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
		backRepoCircle.Map_CircleDBID_CircleDB[circleDB.ID] = circleDB
		BackRepoCircleid_atBckpTime_newID[circleDB_ID_atBackupTime] = circleDB.ID
	}

	if err != nil {
		log.Fatal("Cannot restore/unmarshall json Circle file", err.Error())
	}
}

// RestorePhaseTwo uses all map BackRepo<Circle>id_atBckpTime_newID
// to compute new index
func (backRepoCircle *BackRepoCircleStruct) RestorePhaseTwo() {

	for _, circleDB := range backRepoCircle.Map_CircleDBID_CircleDB {

		// next line of code is to avert unused variable compilation error
		_ = circleDB

		// insertion point for reindexing pointers encoding
		// update databse with new index encoding
		query := backRepoCircle.db.Model(circleDB).Updates(*circleDB)
		if query.Error != nil {
			log.Fatal(query.Error)
		}
	}

}

// BackRepoCircle.ResetReversePointers commits all staged instances of Circle to the BackRepo
// Phase Two is the update of instance with the field in the database
func (backRepoCircle *BackRepoCircleStruct) ResetReversePointers(backRepo *BackRepoStruct) (Error error) {

	for idx, circle := range backRepoCircle.Map_CircleDBID_CirclePtr {
		backRepoCircle.ResetReversePointersInstance(backRepo, idx, circle)
	}

	return
}

func (backRepoCircle *BackRepoCircleStruct) ResetReversePointersInstance(backRepo *BackRepoStruct, idx uint, circle *models.Circle) (Error error) {

	// fetch matching circleDB
	if circleDB, ok := backRepoCircle.Map_CircleDBID_CircleDB[idx]; ok {
		_ = circleDB // to avoid unused variable error if there are no reverse to reset

		// insertion point for reverse pointers reset
		// end of insertion point for reverse pointers reset
	}

	return
}

// this field is used during the restauration process.
// it stores the ID at the backup time and is used for renumbering
var BackRepoCircleid_atBckpTime_newID map[uint]uint
