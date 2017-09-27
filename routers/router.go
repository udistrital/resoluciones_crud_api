// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/udistrital/administrativa_local/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/solicitud_disponibilidad",
			beego.NSInclude(
				&controllers.SolicitudDisponibilidadController{},
			),
		),

		beego.NSNamespace("/periodo_pago",
			beego.NSInclude(
				&controllers.PeriodoPagoController{},
			),
		),

		beego.NSNamespace("/resolucion_vinculacion_docente",
			beego.NSInclude(
				&controllers.ResolucionVinculacionDocenteController{},
			),
		),

		beego.NSNamespace("/tipo_upc",
			beego.NSInclude(
				&controllers.TipoUpcController{},
			),
		),

		beego.NSNamespace("/tipo_resolucion",
			beego.NSInclude(
				&controllers.TipoResolucionController{},
			),
		),

		beego.NSNamespace("/necesidad_otro_si",
			beego.NSInclude(
				&controllers.NecesidadOtroSiController{},
			),
		),

		beego.NSNamespace("/preliquidacion",
			beego.NSInclude(
				&controllers.PreliquidacionController{},
			),
		),

		beego.NSNamespace("/relacion_parametro_contraloria",
			beego.NSInclude(
				&controllers.RelacionParametroContraloriaController{},
			),
		),

		beego.NSNamespace("/catalogo_elemento_grupo",
			beego.NSInclude(
				&controllers.CatalogoElementoGrupoController{},
			),
		),

		beego.NSNamespace("/pago",
			beego.NSInclude(
				&controllers.PagoController{},
			),
		),

		beego.NSNamespace("/solicitud_rp",
			beego.NSInclude(
				&controllers.SolicitudRpController{},
			),
		),

		beego.NSNamespace("/asignacion_pension_user_role",
			beego.NSInclude(
				&controllers.AsignacionPensionUserRoleController{},
			),
		),

		beego.NSNamespace("/estado_resolucion",
			beego.NSInclude(
				&controllers.EstadoResolucionController{},
			),
		),

		beego.NSNamespace("/necesidad_proceso_externo",
			beego.NSInclude(
				&controllers.NecesidadProcesoExternoController{},
			),
		),

		beego.NSNamespace("/pension",
			beego.NSInclude(
				&controllers.PensionController{},
			),
		),

		beego.NSNamespace("/dedicacion",
			beego.NSInclude(
				&controllers.DedicacionController{},
			),
		),

		beego.NSNamespace("/dependencia_necesidad",
			beego.NSInclude(
				&controllers.DependenciaNecesidadController{},
			),
		),

		beego.NSNamespace("/entrada",
			beego.NSInclude(
				&controllers.EntradaController{},
			),
		),

		beego.NSNamespace("/necesidad_rechazada",
			beego.NSInclude(
				&controllers.NecesidadRechazadaController{},
			),
		),

		beego.NSNamespace("/actividad_especifica",
			beego.NSInclude(
				&controllers.ActividadEspecificaController{},
			),
		),

		beego.NSNamespace("/concepto_nomina_por_persona",
			beego.NSInclude(
				&controllers.ConceptoNominaPorPersonaController{},
			),
		),

		beego.NSNamespace("/naturaleza_concepto_nomina",
			beego.NSInclude(
				&controllers.NaturalezaConceptoNominaController{},
			),
		),

		beego.NSNamespace("/resolucion",
			beego.NSInclude(
				&controllers.ResolucionController{},
			),
		),

		beego.NSNamespace("/contrato",
			beego.NSInclude(
				&controllers.ContratoController{},
			),
		),

		beego.NSNamespace("/detalle_preliquidacion",
			beego.NSInclude(
				&controllers.DetallePreliquidacionController{},
			),
		),

		beego.NSNamespace("/disponibilidad_apropiacion_solicitud_rp",
			beego.NSInclude(
				&controllers.DisponibilidadApropiacionSolicitudRpController{},
			),
		),

		beego.NSNamespace("/disponibilidad_contrato",
			beego.NSInclude(
				&controllers.DisponibilidadContratoController{},
			),
		),

		beego.NSNamespace("/poliza_contrato",
			beego.NSInclude(
				&controllers.PolizaContratoController{},
			),
		),

		beego.NSNamespace("/relacion_poliza_amparo",
			beego.NSInclude(
				&controllers.RelacionPolizaAmparoController{},
			),
		),

		beego.NSNamespace("/tipo_contrato_necesidad",
			beego.NSInclude(
				&controllers.TipoContratoNecesidadController{},
			),
		),

		beego.NSNamespace("/estado_rol_pension",
			beego.NSInclude(
				&controllers.EstadoRolPensionController{},
			),
		),

		beego.NSNamespace("/tipo_necesidad",
			beego.NSInclude(
				&controllers.TipoNecesidadController{},
			),
		),

		beego.NSNamespace("/tipo_nomina",
			beego.NSInclude(
				&controllers.TipoNominaController{},
			),
		),

		beego.NSNamespace("/tipo_preliquidacion",
			beego.NSInclude(
				&controllers.TipoPreliquidacionController{},
			),
		),

		beego.NSNamespace("/amparo_contractual",
			beego.NSInclude(
				&controllers.AmparoContractualController{},
			),
		),

		beego.NSNamespace("/detalle_servicio_necesidad",
			beego.NSInclude(
				&controllers.DetalleServicioNecesidadController{},
			),
		),

		beego.NSNamespace("/estado_preliquidacion",
			beego.NSInclude(
				&controllers.EstadoPreliquidacionController{},
			),
		),

		beego.NSNamespace("/tipo_pension",
			beego.NSInclude(
				&controllers.TipoPensionController{},
			),
		),

		beego.NSNamespace("/upc_adicional",
			beego.NSInclude(
				&controllers.UpcAdicionalController{},
			),
		),

		beego.NSNamespace("/escalafon",
			beego.NSInclude(
				&controllers.EscalafonController{},
			),
		),

		beego.NSNamespace("/rango_edad_upc",
			beego.NSInclude(
				&controllers.RangoEdadUpcController{},
			),
		),

		beego.NSNamespace("/concepto_nomina",
			beego.NSInclude(
				&controllers.ConceptoNominaController{},
			),
		),

		beego.NSNamespace("/nomina",
			beego.NSInclude(
				&controllers.NominaController{},
			),
		),

		beego.NSNamespace("/parametro_contraloria",
			beego.NSInclude(
				&controllers.ParametroContraloriaController{},
			),
		),

		beego.NSNamespace("/tipo_concepto_nomina",
			beego.NSInclude(
				&controllers.TipoConceptoNominaController{},
			),
		),

		beego.NSNamespace("/catalogo_elemento",
			beego.NSInclude(
				&controllers.CatalogoElementoController{},
			),
		),

		beego.NSNamespace("/componente_resolucion",
			beego.NSInclude(
				&controllers.ComponenteResolucionController{},
			),
		),

		beego.NSNamespace("/parametro_contraloria_contrato",
			beego.NSInclude(
				&controllers.ParametroContraloriaContratoController{},
			),
		),

		beego.NSNamespace("/resolucion_estado",
			beego.NSInclude(
				&controllers.ResolucionEstadoController{},
			),
		),

		beego.NSNamespace("/tipo_pago",
			beego.NSInclude(
				&controllers.TipoPagoController{},
			),
		),

		beego.NSNamespace("/zona_upc",
			beego.NSInclude(
				&controllers.ZonaUpcController{},
			),
		),

		beego.NSNamespace("/escalafon_persona",
			beego.NSInclude(
				&controllers.EscalafonPersonaController{},
			),
		),

		beego.NSNamespace("/estado_contrato",
			beego.NSInclude(
				&controllers.EstadoContratoController{},
			),
		),

		beego.NSNamespace("/relacion_contrato_estado",
			beego.NSInclude(
				&controllers.RelacionContratoEstadoController{},
			),
		),

		beego.NSNamespace("/tipo_valor_amparo",
			beego.NSInclude(
				&controllers.TipoValorAmparoController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
