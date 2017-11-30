package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"github.com/udistrital/administrativa_crud_api/models"
)

// oprations for Necesidad
type TrNecesidadDocenteController struct {
	beego.Controller
}

func (c *TrNecesidadDocenteController) URLMapping() {
	c.Mapping("Post", c.Post)
}

// @Title TrNecesidadDocente
// @Description insert the TrNecesidadDocente in the tables Necesidad, FuenteFinanciacionRubroNecesidad, MarcoLegalNecesidad, DependenciaNecesidad
// @Param	body		body 	models.TrNecesidadDocente	true	"body for TrNecesidadDocente content"
// @Success 201 {object} msg
// @Failure 403 :id is not int
// @router / [post]
func (c *TrNecesidadDocenteController) Post() {
	var v models.TrNecesidadDocente
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if alerta, err := models.AddTrNecesidadDocente(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = alerta
		} else {
			c.Data["json"] = alerta
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()

}
