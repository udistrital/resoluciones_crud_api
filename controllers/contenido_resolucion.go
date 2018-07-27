package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/udistrital/resoluciones_crud/models"

	"github.com/astaxie/beego"
)

type ResolucionCompletaController struct {
	beego.Controller
}

func (c *ResolucionCompletaController) URLMapping() {
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("Put", c.Put)
	c.Mapping("ResolucionTemplate", c.ResolucionTemplate)
}

// GetOne ...
// @Title Get Template
// @Description get ResolucionCompleta by id
// @Param	dedicacion	path 	string	true		"nombre de la dedicacion"
// @Success 200 {object} models.ResolucionCompleta
// @Failure 403
// @router /ResolucionTemplate/:dedicacion/:nivel [get]
func (c *ResolucionCompletaController) ResolucionTemplate() {
	dedicacion := c.Ctx.Input.Param(":dedicacion")
	nivel := c.Ctx.Input.Param(":nivel")
	fmt.Println("dedicacion", dedicacion, nivel)
	resolucion := models.GetTemplateResolucion(dedicacion, nivel)
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = resolucion
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get ResolucionCompleta by id
// @Param	idResolucion		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ResolucionCompleta
// @Failure 403 :idResolucion is empty
// @router /:idResolucion [get]
func (c *ResolucionCompletaController) GetOne() {
	idResolucion := c.Ctx.Input.Param(":idResolucion")
	resolucion := models.GetOneResolucionCompleta(idResolucion)
	c.Ctx.Output.SetStatus(201)
	c.Data["json"] = resolucion
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the ResolucionCompleta
// @Param	idResolucion		path 	string	true		"The id you want to update"
// @Success 200 {object} models.ResolucionCompleta
// @Failure 403 :idResolucion is not int
// @router /:idResolucion [put]
func (c *ResolucionCompletaController) Put() {
	idResolucionStr := c.Ctx.Input.Param(":idResolucion")
	idResolucion, _ := strconv.Atoi(idResolucionStr)
	v := models.ResolucionCompleta{Id: idResolucion}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateResolucionCompletaById(&v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		fmt.Println(err.Error())
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
