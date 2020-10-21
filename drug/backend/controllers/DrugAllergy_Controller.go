package controllers
 
import (
	"context"
	"strconv"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mmildd_s/app/ent"
	"github.com/mmildd_s/app/ent/doctor"
	"github.com/mmildd_s/app/ent/patient"
	"github.com/mmildd_s/app/ent/medicine"
	"github.com/mmildd_s/app/ent/manner"
	"github.com/mmildd_s/app/ent/drugallergy"
)
 // DrugAllergyController defines the struct for the drugAllergy controller
type DrugAllergyController struct {
	client *ent.Client
	router gin.IRouter
}
//DrugAllergy struct
type DrugAllergy struct {
	Doctor   	int
	Patient     int
	Medicine 	int
	Manner      int
}
// CreateDrugAllergy handles POST requests for adding drugAllergy entities
// @Summary Create drugAllergy
// @Description Create drugAllergy
// @ID create-drugAllergy
// @Accept   json
// @Produce  json
// @Param drugallergy body DrugAllergy true "DrugAllergy entity"
// @Success 200 {object} ent.DrugAllergy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugAllergys [post]
func (ctl *DrugAllergyController) CreateDrugAllergy(c *gin.Context) {
	obj := DrugAllergy{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "DrugAllergy binding failed",
		})
		return
	}
	d, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.Doctor))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor not found",
		})
		return
	}
	p, err := ctl.client.Patient.
		Query().
		Where(patient.IDEQ(int(obj.Patient))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Patient not found",
		})
		return
	}
	me, err := ctl.client.Medicine.
		Query().
		Where(medicine.IDEQ(int(obj.Medicine))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Medicine not found",
		})
		return
	}
	m, err := ctl.client.Manner.
		Query().
		Where(manner.IDEQ(int(obj.Manner))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Manner not found",
		})
		return
	}
	da, err := ctl.client.DrugAllergy.
		Create().
		SetDoctor(d).
		SetPatient(p).
		SetMedicine(me).
		SetManner(m).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   da,
	})
}

// GetDrugAllergy handles GET requests to retrieve a drugAllergy entity
// @Summary Get a drugAllergy entity by ID
// @Description get drugAllergy by ID
// @ID get-drugAllergy
// @Produce  json
// @Param id path int true "DrugAllergy ID"
// @Success 200 {object} ent.DrugAllergy
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugAllergys/{id} [get]
func (ctl *DrugAllergyController) GetDrugAllergy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	b, err := ctl.client.DrugAllergy.
		Query().
		Where(drugallergy.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, b)
}
 // ListDrugAllergy handles request to get a list of drugAllergy entities
// @Summary List drugAllergy entities
// @Description list drugAllergy entities
// @ID list-drugAllergy
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.DrugAllergy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugAllergys [get]
func (ctl *DrugAllergyController) ListDrugAllergy(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {limit = int(limit64)}
	}
  
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {offset = int(offset64)}
	}
  
	drugAllergys, err := ctl.client.DrugAllergy.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, drugAllergys)
	
 }

// DeleteDrugAllergy handles DELETE requests to delete a drugAllergy entity
// @Summary Delete a drugAllergy entity by ID
// @Description get drugAllergy by ID
// @ID delete-drugAllergy
// @Produce  json
// @Param id path int true "DrugAllergy ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugAllergys/{id} [delete]
func (ctl *DrugAllergyController) DeleteDrugAllergy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.DrugAllergy.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}
// UpdateDrugAllergy handles PUT requests to update a DrugAllergy entity
// @Summary Update a drugAllergy entity by ID
// @Description update drugAllergy by ID
// @ID update-drugAllergy
// @Accept   json
// @Produce  json
// @Param id path int true "DrugAllergy ID"
// @Param drugAllergy body ent.DrugAllergy true "DrugAllergy entity"
// @Success 200 {object} ent.DrugAllergy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugAllergys/{id} [put]
func (ctl *DrugAllergyController) UpdateDrugAllergy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.DrugAllergy{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "booking binding failed",
		})
		return
	}
	obj.ID = int(id)
	b, err := ctl.client.DrugAllergy.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, b)
}
 
 // NewDrugAllergyController creates and registers handles for the playlist-video controller
func NewDrugAllergyController(router gin.IRouter, client *ent.Client) *DrugAllergyController {
	da := &DrugAllergyController{
		client: client,
		router: router,
	}

	da.register()

	return da

}

func (ctl *DrugAllergyController) register() {
	drugAllergys := ctl.router.Group("/drugAllergys")

	drugAllergys.POST("", ctl.CreateDrugAllergy)
	drugAllergys.GET("", ctl.ListDrugAllergy)
	drugAllergys.DELETE(":id", ctl.DeleteDrugAllergy)

}
 
 