package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/mmildd_s/app/ent"
   "github.com/mmildd_s/app/ent/medicine"
   "github.com/gin-gonic/gin"
)
 
// MedicineController defines the struct for the medicine controller
type MedicineController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateMedicine handles POST requests for adding medicine entities
// @Summary Create medicine
// @Description Create medicine
// @ID create-medicine
// @Accept   json
// @Produce  json
// @Param medicine body ent.Medicine true "Medicine entity"
// @Success 200 {object} ent.Medicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicines [post]
func (ctl *MedicineController) CreateMedicine(c *gin.Context) {
	obj := ent.Medicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Medicine binding failed",
		})
		return
	}
  
	me, err := ctl.client.Medicine.
		Create().
		SetMedicineName(obj.MedicineName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
  
	c.JSON(200, me)
 }
 // GetMedicine handles GET requests to retrieve a medicine entity
// @Summary Get a medicine entity by ID
// @Description get medicine by ID
// @ID get-medicine
// @Produce  json
// @Param id path int true "Medicine ID"
// @Success 200 {object} ent.Medicine
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicines/{id} [get]
func (ctl *MedicineController) GetMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	me, err := ctl.client.Medicine.
		Query().
		Where(medicine.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	c.JSON(200, me)
 }
 // ListMedicine handles request to get a list of medicine entities
// @Summary List medicine entities
// @Description list medicine entities
// @ID list-medicine
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Medicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicines [get]
func (ctl *MedicineController) ListMedicine(c *gin.Context) {
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
  
	medicines, err := ctl.client.Medicine.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, medicines)
 }
 // DeleteMedicine handles DELETE requests to delete a medicine entity
// @Summary Delete a medicine entity by ID
// @Description get medicine by ID
// @ID delete-medicine
// @Produce  json
// @Param id path int true "Medicine ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicines/{id} [delete]
func (ctl *MedicineController) DeleteMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Medicine.
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
 // UpdateMedicine handles PUT requests to update a medicine entity
// @Summary Update a medicine entity by ID
// @Description update medicine by ID
// @ID update-medicine
// @Accept   json
// @Produce  json
// @Param id path int true "Medicine ID"
// @Param medicine body ent.Medicine true "Medicine entity"
// @Success 200 {object} ent.Medicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /medicines/{id} [put]
func (ctl *MedicineController) UpdateMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Medicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Medicine binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	me, err := ctl.client.Medicine.
		UpdateOne(&obj).
		SetMedicineName(obj.MedicineName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, me)
 }
 // NewMedicineController creates and registers handles for the medicine controller
func NewMedicineController(router gin.IRouter, client *ent.Client) *MedicineController {
	uc := &MedicineController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitMedicineController registers routes to the main engine
 func (ctl *MedicineController) register() {
	medicines := ctl.router.Group("/medicines")
  
	medicines.GET("", ctl.ListMedicine)
  
	// CRUD
	medicines.POST("", ctl.CreateMedicine)
	medicines.GET(":id", ctl.GetMedicine)
	medicines.PUT(":id", ctl.UpdateMedicine)
	medicines.DELETE(":id", ctl.DeleteMedicine)
 }
 
 