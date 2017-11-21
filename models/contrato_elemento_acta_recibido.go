package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type ContratoElementoActaRecibido struct {
	Id               int               `orm:"column(id);pk"`
	Cantidad         int               `orm:"column(cantidad)"`
	Elemento         string            `orm:"column(elemento)"`
	Marca            string            `orm:"column(marca);null"`
	Serie            string            `orm:"column(serie);null"`
	VidaUtil         int               `orm:"column(vida_util);null"`
	SubgrupoCatalogo *SubgrupoCatalogo `orm:"column(subgrupo_catalogo);rel(fk)"`
	ContratoElemento int               `orm:"column(contrato_elemento)"`
	Unidad           int               `orm:"column(unidad)"`
	ActaRecibido     *ActaRecibido     `orm:"column(acta_recibido);rel(fk)"`
}

func (t *ContratoElementoActaRecibido) TableName() string {
	return "contrato_elemento_acta_recibido"
}

func init() {
	orm.RegisterModel(new(ContratoElementoActaRecibido))
}

// AddContratoElementoActaRecibido insert a new ContratoElementoActaRecibido into database and returns
// last inserted Id on success.
func AddContratoElementoActaRecibido(m *ContratoElementoActaRecibido) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetContratoElementoActaRecibidoById retrieves ContratoElementoActaRecibido by Id. Returns error if
// Id doesn't exist
func GetContratoElementoActaRecibidoById(id int) (v *ContratoElementoActaRecibido, err error) {
	o := orm.NewOrm()
	v = &ContratoElementoActaRecibido{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllContratoElementoActaRecibido retrieves all ContratoElementoActaRecibido matches certain condition. Returns empty list if
// no records exist
func GetAllContratoElementoActaRecibido(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ContratoElementoActaRecibido))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []ContratoElementoActaRecibido
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateContratoElementoActaRecibido updates ContratoElementoActaRecibido by Id and returns error if
// the record to be updated doesn't exist
func UpdateContratoElementoActaRecibidoById(m *ContratoElementoActaRecibido) (err error) {
	o := orm.NewOrm()
	v := ContratoElementoActaRecibido{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteContratoElementoActaRecibido deletes ContratoElementoActaRecibido by Id and returns error if
// the record to be deleted doesn't exist
func DeleteContratoElementoActaRecibido(id int) (err error) {
	o := orm.NewOrm()
	v := ContratoElementoActaRecibido{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ContratoElementoActaRecibido{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
