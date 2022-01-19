// @generated Do not edit this file, which is automatically generated by the generator.

package dbschema

import (
	"fmt"

	"time"

	"github.com/webx-top/com"
	"github.com/webx-top/db"
	"github.com/webx-top/db/lib/factory"
	"github.com/webx-top/echo"
	"github.com/webx-top/echo/param"
)

type Slice_NgingSendingLog []*NgingSendingLog

func (s Slice_NgingSendingLog) Range(fn func(m factory.Model) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingSendingLog) RangeRaw(fn func(m *NgingSendingLog) error) error {
	for _, v := range s {
		if err := fn(v); err != nil {
			return err
		}
	}
	return nil
}

func (s Slice_NgingSendingLog) GroupBy(keyField string) map[string][]*NgingSendingLog {
	r := map[string][]*NgingSendingLog{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		if _, y := r[vkey]; !y {
			r[vkey] = []*NgingSendingLog{}
		}
		r[vkey] = append(r[vkey], row)
	}
	return r
}

func (s Slice_NgingSendingLog) KeyBy(keyField string) map[string]*NgingSendingLog {
	r := map[string]*NgingSendingLog{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = row
	}
	return r
}

func (s Slice_NgingSendingLog) AsKV(keyField string, valueField string) param.Store {
	r := param.Store{}
	for _, row := range s {
		dmap := row.AsMap()
		vkey := fmt.Sprint(dmap[keyField])
		r[vkey] = dmap[valueField]
	}
	return r
}

func (s Slice_NgingSendingLog) Transform(transfers map[string]param.Transfer) []param.Store {
	r := make([]param.Store, len(s))
	for idx, row := range s {
		r[idx] = row.AsMap().Transform(transfers)
	}
	return r
}

func (s Slice_NgingSendingLog) FromList(data interface{}) Slice_NgingSendingLog {
	values, ok := data.([]*NgingSendingLog)
	if !ok {
		for _, value := range data.([]interface{}) {
			row := &NgingSendingLog{}
			row.FromRow(value.(map[string]interface{}))
			s = append(s, row)
		}
		return s
	}
	s = append(s, values...)

	return s
}

func NewNgingSendingLog(ctx echo.Context) *NgingSendingLog {
	m := &NgingSendingLog{}
	m.SetContext(ctx)
	return m
}

// NgingSendingLog 邮件短信等发送日志
type NgingSendingLog struct {
	base    factory.Base
	objects []*NgingSendingLog

	Id              uint64 `db:"id,omitempty,pk" bson:"id,omitempty" comment:"ID" json:"id" xml:"id"`
	Created         uint   `db:"created" bson:"created" comment:"创建时间" json:"created" xml:"created"`
	SentAt          uint   `db:"sent_at" bson:"sent_at" comment:"发送时间" json:"sent_at" xml:"sent_at"`
	SourceId        uint64 `db:"source_id" bson:"source_id" comment:"来源ID" json:"source_id" xml:"source_id"`
	SourceType      string `db:"source_type" bson:"source_type" comment:"来源类型" json:"source_type" xml:"source_type"`
	Disabled        string `db:"disabled" bson:"disabled" comment:"是否禁用" json:"disabled" xml:"disabled"`
	Method          string `db:"method" bson:"method" comment:"发送方式(mobile-手机;email-邮箱)" json:"method" xml:"method"`
	To              string `db:"to" bson:"to" comment:"发送目标" json:"to" xml:"to"`
	Provider        string `db:"provider" bson:"provider" comment:"发送平台" json:"provider" xml:"provider"`
	Result          string `db:"result" bson:"result" comment:"发送结果描述" json:"result" xml:"result"`
	Status          string `db:"status" bson:"status" comment:"发送状态(none-无需发送)" json:"status" xml:"status"`
	Retries         uint   `db:"retries" bson:"retries" comment:"重试次数" json:"retries" xml:"retries"`
	Content         string `db:"content" bson:"content" comment:"发送消息内容" json:"content" xml:"content"`
	Params          string `db:"params" bson:"params" comment:"发送消息参数(JSON)" json:"params" xml:"params"`
	AppointmentTime uint   `db:"appointment_time" bson:"appointment_time" comment:"预约发送时间" json:"appointment_time" xml:"appointment_time"`
}

// - base function

func (a *NgingSendingLog) Trans() *factory.Transaction {
	return a.base.Trans()
}

func (a *NgingSendingLog) Use(trans *factory.Transaction) factory.Model {
	a.base.Use(trans)
	return a
}

func (a *NgingSendingLog) SetContext(ctx echo.Context) factory.Model {
	a.base.SetContext(ctx)
	return a
}

func (a *NgingSendingLog) EventON(on ...bool) factory.Model {
	a.base.EventON(on...)
	return a
}

func (a *NgingSendingLog) EventOFF(off ...bool) factory.Model {
	a.base.EventOFF(off...)
	return a
}

func (a *NgingSendingLog) Context() echo.Context {
	return a.base.Context()
}

func (a *NgingSendingLog) SetConnID(connID int) factory.Model {
	a.base.SetConnID(connID)
	return a
}

func (a *NgingSendingLog) SetNamer(namer func(factory.Model) string) factory.Model {
	a.base.SetNamer(namer)
	return a
}

func (a *NgingSendingLog) Namer() func(factory.Model) string {
	return a.base.Namer()
}

func (a *NgingSendingLog) SetParam(param *factory.Param) factory.Model {
	a.base.SetParam(param)
	return a
}

func (a *NgingSendingLog) Param(mw func(db.Result) db.Result, args ...interface{}) *factory.Param {
	if a.base.Param() == nil {
		return a.NewParam().SetMiddleware(mw).SetArgs(args...)
	}
	return a.base.Param().SetMiddleware(mw).SetArgs(args...)
}

// - current function

func (a *NgingSendingLog) New(structName string, connID ...int) factory.Model {
	if len(connID) > 0 {
		return factory.NewModel(structName, connID[0]).Use(a.base.Trans())
	}
	return factory.NewModel(structName, a.base.ConnID()).Use(a.base.Trans())
}

func (a *NgingSendingLog) Objects() []*NgingSendingLog {
	if a.objects == nil {
		return nil
	}
	return a.objects[:]
}

func (a *NgingSendingLog) XObjects() Slice_NgingSendingLog {
	return Slice_NgingSendingLog(a.Objects())
}

func (a *NgingSendingLog) NewObjects() factory.Ranger {
	return &Slice_NgingSendingLog{}
}

func (a *NgingSendingLog) InitObjects() *[]*NgingSendingLog {
	a.objects = []*NgingSendingLog{}
	return &a.objects
}

func (a *NgingSendingLog) NewParam() *factory.Param {
	return factory.NewParam(factory.DefaultFactory).SetIndex(a.base.ConnID()).SetTrans(a.base.Trans()).SetCollection(a.Name_()).SetModel(a)
}

func (a *NgingSendingLog) Short_() string {
	return "nging_sending_log"
}

func (a *NgingSendingLog) Struct_() string {
	return "NgingSendingLog"
}

func (a *NgingSendingLog) Name_() string {
	if a.base.Namer() != nil {
		return WithPrefix(a.base.Namer()(a))
	}
	return WithPrefix(factory.TableNamerGet(a.Short_())(a))
}

func (a *NgingSendingLog) CPAFrom(source factory.Model) factory.Model {
	a.SetContext(source.Context())
	a.Use(source.Trans())
	a.SetNamer(source.Namer())
	return a
}

func (a *NgingSendingLog) Get(mw func(db.Result) db.Result, args ...interface{}) (err error) {
	base := a.base
	if !a.base.Eventable() {
		err = a.Param(mw, args...).SetRecv(a).One()
		a.base = base
		return
	}
	queryParam := a.Param(mw, args...).SetRecv(a)
	if err = DBI.FireReading(a, queryParam); err != nil {
		return
	}
	err = queryParam.One()
	a.base = base
	if err == nil {
		err = DBI.FireReaded(a, queryParam)
	}
	return
}

func (a *NgingSendingLog) List(recv interface{}, mw func(db.Result) db.Result, page, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetPage(page).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingSendingLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingSendingLog(*v))
		case []*NgingSendingLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingSendingLog(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingSendingLog) GroupBy(keyField string, inputRows ...[]*NgingSendingLog) map[string][]*NgingSendingLog {
	var rows Slice_NgingSendingLog
	if len(inputRows) > 0 {
		rows = Slice_NgingSendingLog(inputRows[0])
	} else {
		rows = Slice_NgingSendingLog(a.Objects())
	}
	return rows.GroupBy(keyField)
}

func (a *NgingSendingLog) KeyBy(keyField string, inputRows ...[]*NgingSendingLog) map[string]*NgingSendingLog {
	var rows Slice_NgingSendingLog
	if len(inputRows) > 0 {
		rows = Slice_NgingSendingLog(inputRows[0])
	} else {
		rows = Slice_NgingSendingLog(a.Objects())
	}
	return rows.KeyBy(keyField)
}

func (a *NgingSendingLog) AsKV(keyField string, valueField string, inputRows ...[]*NgingSendingLog) param.Store {
	var rows Slice_NgingSendingLog
	if len(inputRows) > 0 {
		rows = Slice_NgingSendingLog(inputRows[0])
	} else {
		rows = Slice_NgingSendingLog(a.Objects())
	}
	return rows.AsKV(keyField, valueField)
}

func (a *NgingSendingLog) ListByOffset(recv interface{}, mw func(db.Result) db.Result, offset, size int, args ...interface{}) (func() int64, error) {
	if recv == nil {
		recv = a.InitObjects()
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv).List()
	}
	queryParam := a.Param(mw, args...).SetOffset(offset).SetSize(size).SetRecv(recv)
	if err := DBI.FireReading(a, queryParam); err != nil {
		return nil, err
	}
	cnt, err := queryParam.List()
	if err == nil {
		switch v := recv.(type) {
		case *[]*NgingSendingLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingSendingLog(*v))
		case []*NgingSendingLog:
			err = DBI.FireReaded(a, queryParam, Slice_NgingSendingLog(v))
		case factory.Ranger:
			err = DBI.FireReaded(a, queryParam, v)
		}
	}
	return cnt, err
}

func (a *NgingSendingLog) Add() (pk interface{}, err error) {
	a.Created = uint(time.Now().Unix())
	a.Id = 0
	if len(a.SourceType) == 0 {
		a.SourceType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Method) == 0 {
		a.Method = "mobile"
	}
	if len(a.Status) == 0 {
		a.Status = "waiting"
	}
	if a.base.Eventable() {
		err = DBI.Fire("creating", a, nil)
		if err != nil {
			return
		}
	}
	pk, err = a.Param(nil).SetSend(a).Insert()
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		err = DBI.Fire("created", a, nil)
	}
	return
}

func (a *NgingSendingLog) Edit(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if len(a.SourceType) == 0 {
		a.SourceType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Method) == 0 {
		a.Method = "mobile"
	}
	if len(a.Status) == 0 {
		a.Status = "waiting"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(a).Update(); err != nil {
		return
	}
	return DBI.Fire("updated", a, mw, args...)
}

func (a *NgingSendingLog) Editx(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if len(a.SourceType) == 0 {
		a.SourceType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Method) == 0 {
		a.Method = "mobile"
	}
	if len(a.Status) == 0 {
		a.Status = "waiting"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(a).Update()
	}
	if err = DBI.Fire("updating", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).SetSend(a).Updatex(); err != nil {
		return
	}
	err = DBI.Fire("updated", a, mw, args...)
	return
}

func (a *NgingSendingLog) EditByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (err error) {

	if len(a.SourceType) == 0 {
		a.SourceType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Method) == 0 {
		a.Method = "mobile"
	}
	if len(a.Status) == 0 {
		a.Status = "waiting"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdateByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).UpdateByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingSendingLog) EditxByFields(mw func(db.Result) db.Result, fields []string, args ...interface{}) (affected int64, err error) {

	if len(a.SourceType) == 0 {
		a.SourceType = "user"
	}
	if len(a.Disabled) == 0 {
		a.Disabled = "N"
	}
	if len(a.Method) == 0 {
		a.Method = "mobile"
	}
	if len(a.Status) == 0 {
		a.Status = "waiting"
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).UpdatexByStruct(a, fields...)
	}
	editColumns := make([]string, len(fields))
	for index, field := range fields {
		editColumns[index] = com.SnakeCase(field)
	}
	if err = DBI.FireUpdate("updating", a, editColumns, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).UpdatexByStruct(a, fields...); err != nil {
		return
	}
	err = DBI.FireUpdate("updated", a, editColumns, mw, args...)
	return
}

func (a *NgingSendingLog) SetField(mw func(db.Result) db.Result, field string, value interface{}, args ...interface{}) (err error) {
	return a.SetFields(mw, map[string]interface{}{
		field: value,
	}, args...)
}

func (a *NgingSendingLog) SetFields(mw func(db.Result) db.Result, kvset map[string]interface{}, args ...interface{}) (err error) {

	if val, ok := kvset["source_type"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["source_type"] = "user"
		}
	}
	if val, ok := kvset["disabled"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["disabled"] = "N"
		}
	}
	if val, ok := kvset["method"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["method"] = "mobile"
		}
	}
	if val, ok := kvset["status"]; ok && val != nil {
		if v, ok := val.(string); ok && len(v) == 0 {
			kvset["status"] = "waiting"
		}
	}
	if !a.base.Eventable() {
		return a.Param(mw, args...).SetSend(kvset).Update()
	}
	m := *a
	m.FromRow(kvset)
	var editColumns []string
	for column := range kvset {
		editColumns = append(editColumns, column)
	}
	if err = DBI.FireUpdate("updating", &m, editColumns, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).SetSend(kvset).Update(); err != nil {
		return
	}
	return DBI.FireUpdate("updated", &m, editColumns, mw, args...)
}

func (a *NgingSendingLog) Upsert(mw func(db.Result) db.Result, args ...interface{}) (pk interface{}, err error) {
	pk, err = a.Param(mw, args...).SetSend(a).Upsert(func() error {
		if len(a.SourceType) == 0 {
			a.SourceType = "user"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Method) == 0 {
			a.Method = "mobile"
		}
		if len(a.Status) == 0 {
			a.Status = "waiting"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("updating", a, mw, args...)
	}, func() error {
		a.Created = uint(time.Now().Unix())
		a.Id = 0
		if len(a.SourceType) == 0 {
			a.SourceType = "user"
		}
		if len(a.Disabled) == 0 {
			a.Disabled = "N"
		}
		if len(a.Method) == 0 {
			a.Method = "mobile"
		}
		if len(a.Status) == 0 {
			a.Status = "waiting"
		}
		if !a.base.Eventable() {
			return nil
		}
		return DBI.Fire("creating", a, nil)
	})
	if err == nil && pk != nil {
		if v, y := pk.(uint64); y {
			a.Id = v
		} else if v, y := pk.(int64); y {
			a.Id = uint64(v)
		}
	}
	if err == nil && a.base.Eventable() {
		if pk == nil {
			err = DBI.Fire("updated", a, mw, args...)
		} else {
			err = DBI.Fire("created", a, nil)
		}
	}
	return
}

func (a *NgingSendingLog) Delete(mw func(db.Result) db.Result, args ...interface{}) (err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Delete()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if err = a.Param(mw, args...).Delete(); err != nil {
		return
	}
	return DBI.Fire("deleted", a, mw, args...)
}

func (a *NgingSendingLog) Deletex(mw func(db.Result) db.Result, args ...interface{}) (affected int64, err error) {

	if !a.base.Eventable() {
		return a.Param(mw, args...).Deletex()
	}
	if err = DBI.Fire("deleting", a, mw, args...); err != nil {
		return
	}
	if affected, err = a.Param(mw, args...).Deletex(); err != nil {
		return
	}
	err = DBI.Fire("deleted", a, mw, args...)
	return
}

func (a *NgingSendingLog) Count(mw func(db.Result) db.Result, args ...interface{}) (int64, error) {
	return a.Param(mw, args...).Count()
}

func (a *NgingSendingLog) Exists(mw func(db.Result) db.Result, args ...interface{}) (bool, error) {
	return a.Param(mw, args...).Exists()
}

func (a *NgingSendingLog) Reset() *NgingSendingLog {
	a.Id = 0
	a.Created = 0
	a.SentAt = 0
	a.SourceId = 0
	a.SourceType = ``
	a.Disabled = ``
	a.Method = ``
	a.To = ``
	a.Provider = ``
	a.Result = ``
	a.Status = ``
	a.Retries = 0
	a.Content = ``
	a.Params = ``
	a.AppointmentTime = 0
	return a
}

func (a *NgingSendingLog) AsMap(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["Id"] = a.Id
		r["Created"] = a.Created
		r["SentAt"] = a.SentAt
		r["SourceId"] = a.SourceId
		r["SourceType"] = a.SourceType
		r["Disabled"] = a.Disabled
		r["Method"] = a.Method
		r["To"] = a.To
		r["Provider"] = a.Provider
		r["Result"] = a.Result
		r["Status"] = a.Status
		r["Retries"] = a.Retries
		r["Content"] = a.Content
		r["Params"] = a.Params
		r["AppointmentTime"] = a.AppointmentTime
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "Id":
			r["Id"] = a.Id
		case "Created":
			r["Created"] = a.Created
		case "SentAt":
			r["SentAt"] = a.SentAt
		case "SourceId":
			r["SourceId"] = a.SourceId
		case "SourceType":
			r["SourceType"] = a.SourceType
		case "Disabled":
			r["Disabled"] = a.Disabled
		case "Method":
			r["Method"] = a.Method
		case "To":
			r["To"] = a.To
		case "Provider":
			r["Provider"] = a.Provider
		case "Result":
			r["Result"] = a.Result
		case "Status":
			r["Status"] = a.Status
		case "Retries":
			r["Retries"] = a.Retries
		case "Content":
			r["Content"] = a.Content
		case "Params":
			r["Params"] = a.Params
		case "AppointmentTime":
			r["AppointmentTime"] = a.AppointmentTime
		}
	}
	return r
}

func (a *NgingSendingLog) FromRow(row map[string]interface{}) {
	for key, value := range row {
		switch key {
		case "id":
			a.Id = param.AsUint64(value)
		case "created":
			a.Created = param.AsUint(value)
		case "sent_at":
			a.SentAt = param.AsUint(value)
		case "source_id":
			a.SourceId = param.AsUint64(value)
		case "source_type":
			a.SourceType = param.AsString(value)
		case "disabled":
			a.Disabled = param.AsString(value)
		case "method":
			a.Method = param.AsString(value)
		case "to":
			a.To = param.AsString(value)
		case "provider":
			a.Provider = param.AsString(value)
		case "result":
			a.Result = param.AsString(value)
		case "status":
			a.Status = param.AsString(value)
		case "retries":
			a.Retries = param.AsUint(value)
		case "content":
			a.Content = param.AsString(value)
		case "params":
			a.Params = param.AsString(value)
		case "appointment_time":
			a.AppointmentTime = param.AsUint(value)
		}
	}
}

func (a *NgingSendingLog) Set(key interface{}, value ...interface{}) {
	switch k := key.(type) {
	case map[string]interface{}:
		for kk, vv := range k {
			a.Set(kk, vv)
		}
	default:
		var (
			kk string
			vv interface{}
		)
		if k, y := key.(string); y {
			kk = k
		} else {
			kk = fmt.Sprint(key)
		}
		if len(value) > 0 {
			vv = value[0]
		}
		switch kk {
		case "Id":
			a.Id = param.AsUint64(vv)
		case "Created":
			a.Created = param.AsUint(vv)
		case "SentAt":
			a.SentAt = param.AsUint(vv)
		case "SourceId":
			a.SourceId = param.AsUint64(vv)
		case "SourceType":
			a.SourceType = param.AsString(vv)
		case "Disabled":
			a.Disabled = param.AsString(vv)
		case "Method":
			a.Method = param.AsString(vv)
		case "To":
			a.To = param.AsString(vv)
		case "Provider":
			a.Provider = param.AsString(vv)
		case "Result":
			a.Result = param.AsString(vv)
		case "Status":
			a.Status = param.AsString(vv)
		case "Retries":
			a.Retries = param.AsUint(vv)
		case "Content":
			a.Content = param.AsString(vv)
		case "Params":
			a.Params = param.AsString(vv)
		case "AppointmentTime":
			a.AppointmentTime = param.AsUint(vv)
		}
	}
}

func (a *NgingSendingLog) AsRow(onlyFields ...string) param.Store {
	r := param.Store{}
	if len(onlyFields) == 0 {
		r["id"] = a.Id
		r["created"] = a.Created
		r["sent_at"] = a.SentAt
		r["source_id"] = a.SourceId
		r["source_type"] = a.SourceType
		r["disabled"] = a.Disabled
		r["method"] = a.Method
		r["to"] = a.To
		r["provider"] = a.Provider
		r["result"] = a.Result
		r["status"] = a.Status
		r["retries"] = a.Retries
		r["content"] = a.Content
		r["params"] = a.Params
		r["appointment_time"] = a.AppointmentTime
		return r
	}
	for _, field := range onlyFields {
		switch field {
		case "id":
			r["id"] = a.Id
		case "created":
			r["created"] = a.Created
		case "sent_at":
			r["sent_at"] = a.SentAt
		case "source_id":
			r["source_id"] = a.SourceId
		case "source_type":
			r["source_type"] = a.SourceType
		case "disabled":
			r["disabled"] = a.Disabled
		case "method":
			r["method"] = a.Method
		case "to":
			r["to"] = a.To
		case "provider":
			r["provider"] = a.Provider
		case "result":
			r["result"] = a.Result
		case "status":
			r["status"] = a.Status
		case "retries":
			r["retries"] = a.Retries
		case "content":
			r["content"] = a.Content
		case "params":
			r["params"] = a.Params
		case "appointment_time":
			r["appointment_time"] = a.AppointmentTime
		}
	}
	return r
}

func (a *NgingSendingLog) BatchValidate(kvset map[string]interface{}) error {
	if kvset == nil {
		kvset = a.AsRow()
	}
	return DBI.Fields.BatchValidate(a.Short_(), kvset)
}

func (a *NgingSendingLog) Validate(field string, value interface{}) error {
	return DBI.Fields.Validate(a.Short_(), field, value)
}
