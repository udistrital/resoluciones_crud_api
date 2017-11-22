package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type TrPagoMensual struct {
	PagoMensual []*PagoMensual
	PagoMensualEstadoPagoMensual []*PagoMensualEstadoPagoMensual
}

/*
	Función para la transaccion de solicitudes de pagos mensuales
*/
func AddSolPagoMensual(m *TrPagoMensual) (alerta []string, err error) {
	o := orm.NewOrm()
	o.Begin()
	alerta = append(alerta, "success")
	var id int64
	//Seteo fecha de creación transacción
	m.PagoMensualEstadoPagoMensual.Fecha = time.Now()
	m.PagoMensual.Mes = float64(m.PagoMensualEstadoPagoMensual.Fecha.Month())


	if id, err = o.Insert(m.Necesidad); err == nil {
		//m.Necesidad.Id = int(id)
		fmt.Println("Fuentes", m.Ffapropiacion)
		for _, v := range m.Ffapropiacion {
			v.Necesidad = &Necesidad{Id: int(id)}
			//---
			if _, err = o.Insert(v); err != nil {
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Error: ¡Ocurrió un error al insertar las fuentes de financiamiento!")
				return
			}
		}
		//--
		fmt.Println("Marco", m.MarcoLegalNecesidad)
		for _, vm := range m.MarcoLegalNecesidad {
			vm.Necesidad = &Necesidad{Id: int(id)}
			//---
			if _, err = o.Insert(vm); err != nil {
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Error: ¡Ocurrió un error al insertar los marcos legales!")
				return
			}
		}

		m.DependenciaNecesidad.Necesidad = &Necesidad{Id: int(id)}
		if _, err = o.Insert(m.DependenciaNecesidad); err != nil {
			o.Rollback()
			alerta[0] = "error"
			alerta = append(alerta, "Error: ¡Ocurrió un error al insertar los datos de las dependencias y responsables!")
			return
		}
		if m.Necesidad.TipoContratoNecesidad.Id == 1 {

			for _, ve := range m.Especificacion {
				ve.EspecificacionTecnica.Necesidad = &Necesidad{Id: int(id)}
				//---
				if _, err = o.Insert(ve.EspecificacionTecnica); err != nil {
					o.Rollback()
					alerta[0] = "error"
					alerta = append(alerta, "Error: ¡Ocurrió un error al insertar las especificaciones técnicas!")
					return
				} else {
					for _, vr := range ve.RequisitoMinimo {
						vr.EspecificacionTecnica = ve.EspecificacionTecnica
						//---
						if _, err = o.Insert(vr); err != nil {
							o.Rollback()
							alerta[0] = "error"
							alerta = append(alerta, "Error: ¡Ocurrió un error al insertar los requisitos mínimos!")
							return
						}
					}
				}
			}
		}
		if m.Necesidad.TipoContratoNecesidad.Id == 2 {
			for _, va := range m.ActividadEconomicaNecesidad {
				va.Necesidad = &Necesidad{Id: int(id)}
				//---
				if _, err = o.Insert(va); err != nil {
					o.Rollback()
					alerta[0] = "error"
					alerta = append(alerta, "Error: ¡Ocurrió un error al insertar las actividades económicas!")
					return
				}
			}
			for _, vp := range m.ActividadEspecifica {
				vp.Necesidad = &Necesidad{Id: int(id)}
				//---
				if _, err = o.Insert(vp); err != nil {
					o.Rollback()
					alerta[0] = "error"
					alerta = append(alerta, "Error: ¡Ocurrió un error al insertar las actividades específicas!")
					return
				}
			}
			m.DetalleServicioNecesidad.Necesidad = &Necesidad{Id: int(id)}
			if _, err = o.Insert(m.DetalleServicioNecesidad); err != nil {
				o.Rollback()
				alerta[0] = "error"
				alerta = append(alerta, "Error: ¡Ocurrió un error al insertar servicio necesidad!")
				return
			}
		}
		o.Commit()
		alerta = append(alerta, "La solicitud con No. de elaboración "+strconv.Itoa(m.Necesidad.NumeroElaboracion)+" del "+strconv.Itoa((m.Necesidad.FechaSolicitud).Year())+" fue creada exitosamente")
		return
	} else {
		o.Rollback()
		alerta[0] = "error"
		alerta = append(alerta, "Error: ¡Ocurrió un error al insertar la necesidad!")
		return
	}
	return alerta, err
}
