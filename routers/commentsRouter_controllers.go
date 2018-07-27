package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ComponenteResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoDisponibilidadController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ContratoGeneralController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DedicacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:DisponibilidadContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:EstadoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ModificacionVinculacionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ParametrosController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "ResolucionTemplate",
			Router: `/ResolucionTemplate/:dedicacion/:nivel`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionCompletaController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionCompletaController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:idResolucion`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "CancelarResolucion",
			Router: `/CancelarResolucion/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "RestaurarResolucion",
			Router: `/RestaurarResolucion/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GenerarResolucion",
			Router: `/GenerarResolucion`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionEstadoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAllAprobada",
			Router: `/Aprobada`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAllExpedidasVigenciaPeriodo",
			Router: `/expedidas_vigencia_periodo`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionController"],
		beego.ControllerComments{
			Method: "GetAllExpedidasVigenciaPeriodoVinculacion",
			Router: `/expedidas_vigencia_periodo_vinculacion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:ResolucionVinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:SupervisorContratoController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:TipoResolucionController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VigenciaContratoController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VigenciaContratoController"],
		beego.ControllerComments{
			Method: "VigenciaContrato",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetTotalContratosXResolucion",
			Router: `/get_total_contratos_x_resolucion/:id_resolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "InsertarVinculaciones",
			Router: `/InsertarVinculaciones`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Post",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetOne",
			Router: `/:id`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Put",
			Router: `/:id`,
			AllowHTTPMethods: []string{"put"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "Delete",
			Router: `/:id`,
			AllowHTTPMethods: []string{"delete"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetVinculacionesAgrupadas",
			Router: `/get_vinculaciones_agrupadas/:id_resolucion`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"] = append(beego.GlobalControllerRouter["github.com/udistrital/resoluciones_crud/controllers:VinculacionDocenteController"],
		beego.ControllerComments{
			Method: "GetValoresTotalesPorDisponibilidad",
			Router: `/get_valores_totales_x_disponibilidad/:anio/:periodo/:id_disponibilidad`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
