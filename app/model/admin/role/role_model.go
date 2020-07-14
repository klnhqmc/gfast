// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package role

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"time"
)

// arModel is a active record design model for table _role operations.
type arModel struct {
	M *gdb.Model
}

var (
	// Table is the table name of _role.
	Table = "role"
	// Model is the model object of _role.
	Model = &arModel{g.DB("default").Table(Table).Safe()}
)

// TX sets the transaction for current operation.
func (m *arModel) TX(tx *gdb.TX) *arModel {
	return &arModel{m.M.TX(tx)}
}

// Master marks the following operation on master node.
func (m *arModel) Master() *arModel {
	return &arModel{m.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (m *arModel) Slave() *arModel {
	return &arModel{m.M.Slave()}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
func (m *arModel) LeftJoin(joinTable string, on string) *arModel {
	return &arModel{m.M.LeftJoin(joinTable, on)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
func (m *arModel) RightJoin(joinTable string, on string) *arModel {
	return &arModel{m.M.RightJoin(joinTable, on)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
func (m *arModel) InnerJoin(joinTable string, on string) *arModel {
	return &arModel{m.M.InnerJoin(joinTable, on)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
func (m *arModel) Fields(fields string) *arModel {
	return &arModel{m.M.Fields(fields)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
func (m *arModel) FieldsEx(fields string) *arModel {
	return &arModel{m.M.FieldsEx(fields)}
}

// Option sets the extra operation option for the model.
func (m *arModel) Option(option int) *arModel {
	return &arModel{m.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (m *arModel) OmitEmpty() *arModel {
	return &arModel{m.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (m *arModel) Filter() *arModel {
	return &arModel{m.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (m *arModel) Where(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.Where(where, args...)}
}

// And adds "AND" condition to the where statement.
func (m *arModel) And(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (m *arModel) Or(where interface{}, args ...interface{}) *arModel {
	return &arModel{m.M.Or(where, args...)}
}

// GroupBy sets the "GROUP BY" statement for the model.
func (m *arModel) GroupBy(groupBy string) *arModel {
	return &arModel{m.M.GroupBy(groupBy)}
}

// OrderBy sets the "ORDER BY" statement for the model.
func (m *arModel) OrderBy(orderBy string) *arModel {
	return &arModel{m.M.OrderBy(orderBy)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (m *arModel) Limit(limit ...int) *arModel {
	return &arModel{m.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (m *arModel) Offset(offset int) *arModel {
	return &arModel{m.M.Offset(offset)}
}

// ForPage sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (m *arModel) ForPage(page, limit int) *arModel {
	return &arModel{m.M.ForPage(page, limit)}
}

// Batch sets the batch operation number for the model.
func (m *arModel) Batch(batch int) *arModel {
	return &arModel{m.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (m *arModel) Cache(expire time.Duration, name ...string) *arModel {
	return &arModel{m.M.Cache(expire, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (m *arModel) Data(data ...interface{}) *arModel {
	return &arModel{m.M.Data(data...)}
}

// Insert does "INSERT INTO ..." statement for the model.
func (m *arModel) Insert() (result sql.Result, err error) {
	return m.M.Insert()
}

// Replace does "REPLACE INTO ..." statement for the model.
func (m *arModel) Replace() (result sql.Result, err error) {
	return m.M.Replace()
}

// Save does "INSERT INTO ... ON DUPLICATE KEY UPDATE..." statement for the model.
// It updates the record if there's primary or unique index in the saving data,
// or else it inserts a new record into the table.
func (m *arModel) Save() (result sql.Result, err error) {
	return m.M.Save()
}

// Update does "UPDATE ... " statement for the model.
func (m *arModel) Update() (result sql.Result, err error) {
	return m.M.Update()
}

// Delete does "DELETE FROM ... " statement for the model.
func (m *arModel) Delete() (result sql.Result, err error) {
	return m.M.Delete()
}

// Count does "SELECT COUNT(x) FROM ..." statement for the model.
func (m *arModel) Count() (int, error) {
	return m.M.Count()
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*Entity.
// It returns nil if there's no record retrieved with the given conditions from table.
func (m *arModel) All() ([]*Entity, error) {
	all, err := m.M.All()
	if err != nil {
		return nil, err
	}
	var entities []*Entity
	if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *Entity.
// It returns nil if there's no record retrieved with the given conditions from table.
func (m *arModel) One() (*Entity, error) {
	one, err := m.M.One()
	if err != nil {
		return nil, err
	}
	var entity *Entity
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// Value retrieves a specified record value from table and returns the result as interface type.
// It returns nil if there's no record found with the given conditions from table.
func (m *arModel) Value() (gdb.Value, error) {
	return m.M.Value()
}

// Chunk iterates the table with given size and callback function.
func (m *arModel) Chunk(limit int, callback func(entities []*Entity, err error) bool) {
	m.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*Entity
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (m *arModel) Page(page, limit int) *arModel {
	return &arModel{m.M.Page(page, limit)}
}

// Order sets the "ORDER BY" statement for the model.
func (m *arModel) Order(orderBy string) *arModel {
	return &arModel{m.M.Order(orderBy)}
}
