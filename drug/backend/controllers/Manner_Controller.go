package controllers
 
import (
   "context"
   "fmt"
   "strconv"
   "github.com/mmildd_s/app/ent"
   "github.com/mmildd_s/app/ent/manner"
   "github.com/gin-gonic/gin"
)
 
// MannerController defines the struct for the manner controller
type MannerController struct {
   client *ent.Client
   router gin.IRouter
}
// CreateManner handles POST requests for adding manner entities
// @Summary Create manner
// @Description Create manner
// @ID create-manner
// @Accept   json
// @Produce  json
// @Param manner body ent.Manner true "Manner entity"
// @Success 200 {object} ent.Manner
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /manners [post]
func (ctl *MannerController) CreateManner(c *gin.Context) {
	obj := ent.Manner{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "manner binding failed",
		})
		return
	}
  
	m, err := ctl.client.Manner.
		Create().
		SetMannerName(obj.MannerName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	
	c.JSON(200, gin.H{
		"data" : m,
		"id" : m.ID,
	})
	
 }
 // GetManner handles GET requests to retrieve a manner entity
// @Summary Get a manner entity by ID
// @Description get manner by ID
// @ID get-manner
// @Produce  json
// @Param id path int true "Manner ID"
// @Success 200 {object} ent.Manner
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /manners/{id} [get]
func (ctl *MannerController) GetManner(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	
	m, err := ctl.client.Manner.
		Query().
		Where(manner.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, m,)
 }

 // ListManner handles request to get a list of manner entities
// @Summary List manner entities
// @Description list manner entities
// @ID list-manner
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Manner
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /manners [get]
func (ctl *MannerController) ListManner(c *gin.Context) {
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
  
	manners, err := ctl.client.Manner.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
		if err != nil {
		c.JSON(400, gin.H{"error": err.Error(),})
		return
	}
  
	c.JSON(200, manners)
 }
 // DeleteManner handles DELETE requests to delete a manner entity
// @Summary Delete a manner entity by ID
// @Description get manner by ID
// @ID delete-manner
// @Produce  json
// @Param id path int true "Manner ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /manners/{id} [delete]
func (ctl *MannerController) DeleteManner(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	err = ctl.client.Manner.
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
 // UpdateManner handles PUT requests to update a manner entity
// @Summary Update a manner entity by ID
// @Description update manner by ID
// @ID update-manner
// @Accept   json
// @Produce  json
// @Param id path int true "Manner ID"
// @Param manner body ent.Manner true "Manner entity"
// @Success 200 {object} ent.Manner
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /manners/{id} [put]
func (ctl *MannerController) UpdateManner(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
  
	obj := ent.Manner{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Manner binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	u, err := ctl.client.Manner.
		UpdateOne(&obj).
		SetMannerName(obj.MannerName).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed",})
		return
	}
  
	c.JSON(200, u)
 }
 // NewMannerController creates and registers handles for the manner controller
func NewMannerController(router gin.IRouter, client *ent.Client) *MannerController {
	uc := &MannerController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
 }
  
 // InitMannerController registers routes to the main engine
 func (ctl *MannerController) register() {
	manners := ctl.router.Group("/manners")
  
	manners.GET("", ctl.ListManner)
  
	// CRUD
	manners.POST("", ctl.CreateManner)
	manners.GET(":id", ctl.GetManner)
	manners.PUT(":id", ctl.UpdateManner)
	manners.DELETE(":id", ctl.DeleteManner)
 }
 
 