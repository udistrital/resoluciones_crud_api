package controllers

import (
	"github.com/udistrital/resoluciones_crud/models"

	"github.com/astaxie/beego"
)

//  VigenciaContratoController operations for VigenciaContrato
type VigenciaContratoController struct {
	beego.Controller
}

// URLMapping ...
func (c *VigenciaContratoController) URLMapping() {
	c.Mapping("Vigencia_Contrato", c.VigenciaContrato)
}

// VigenciaContrato ...
// @Title VigenciaContrato
// @Description create VigenciaContrato
// @Param	body		body 	models.VigenciaContrato	true	"body for VigenciaContrato content"
// @Success 201 {int}
// @Failure 403 body is empty
// @router /
func (c *VigenciaContratoController) VigenciaContrato() {
	respuesta := models.GetVigenciaContrato()
	if len(respuesta) == 0 {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = "Error leyendo las vigencias"
		c.ServeJSON()
	} else {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = respuesta
		c.ServeJSON()
	}
}
