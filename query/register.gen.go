// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/adnpa/IM/model"
)

func newRegister(db *gorm.DB, opts ...gen.DOOption) register {
	_register := register{}

	_register.registerDo.UseDB(db, opts...)
	_register.registerDo.UseModel(&model.Register{})

	tableName := _register.registerDo.TableName()
	_register.ALL = field.NewAsterisk(tableName)
	_register.Account = field.NewString(tableName, "account")
	_register.Password = field.NewString(tableName, "password")

	_register.fillFieldMap()

	return _register
}

type register struct {
	registerDo

	ALL      field.Asterisk
	Account  field.String
	Password field.String

	fieldMap map[string]field.Expr
}

func (r register) Table(newTableName string) *register {
	r.registerDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r register) As(alias string) *register {
	r.registerDo.DO = *(r.registerDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *register) updateTableName(table string) *register {
	r.ALL = field.NewAsterisk(table)
	r.Account = field.NewString(table, "account")
	r.Password = field.NewString(table, "password")

	r.fillFieldMap()

	return r
}

func (r *register) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *register) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 2)
	r.fieldMap["account"] = r.Account
	r.fieldMap["password"] = r.Password
}

func (r register) clone(db *gorm.DB) register {
	r.registerDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r register) replaceDB(db *gorm.DB) register {
	r.registerDo.ReplaceDB(db)
	return r
}

type registerDo struct{ gen.DO }

type IRegisterDo interface {
	gen.SubQuery
	Debug() IRegisterDo
	WithContext(ctx context.Context) IRegisterDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRegisterDo
	WriteDB() IRegisterDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRegisterDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRegisterDo
	Not(conds ...gen.Condition) IRegisterDo
	Or(conds ...gen.Condition) IRegisterDo
	Select(conds ...field.Expr) IRegisterDo
	Where(conds ...gen.Condition) IRegisterDo
	Order(conds ...field.Expr) IRegisterDo
	Distinct(cols ...field.Expr) IRegisterDo
	Omit(cols ...field.Expr) IRegisterDo
	Join(table schema.Tabler, on ...field.Expr) IRegisterDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRegisterDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRegisterDo
	Group(cols ...field.Expr) IRegisterDo
	Having(conds ...gen.Condition) IRegisterDo
	Limit(limit int) IRegisterDo
	Offset(offset int) IRegisterDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRegisterDo
	Unscoped() IRegisterDo
	Create(values ...*model.Register) error
	CreateInBatches(values []*model.Register, batchSize int) error
	Save(values ...*model.Register) error
	First() (*model.Register, error)
	Take() (*model.Register, error)
	Last() (*model.Register, error)
	Find() ([]*model.Register, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Register, err error)
	FindInBatches(result *[]*model.Register, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Register) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRegisterDo
	Assign(attrs ...field.AssignExpr) IRegisterDo
	Joins(fields ...field.RelationField) IRegisterDo
	Preload(fields ...field.RelationField) IRegisterDo
	FirstOrInit() (*model.Register, error)
	FirstOrCreate() (*model.Register, error)
	FindByPage(offset int, limit int) (result []*model.Register, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRegisterDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r registerDo) Debug() IRegisterDo {
	return r.withDO(r.DO.Debug())
}

func (r registerDo) WithContext(ctx context.Context) IRegisterDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r registerDo) ReadDB() IRegisterDo {
	return r.Clauses(dbresolver.Read)
}

func (r registerDo) WriteDB() IRegisterDo {
	return r.Clauses(dbresolver.Write)
}

func (r registerDo) Session(config *gorm.Session) IRegisterDo {
	return r.withDO(r.DO.Session(config))
}

func (r registerDo) Clauses(conds ...clause.Expression) IRegisterDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r registerDo) Returning(value interface{}, columns ...string) IRegisterDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r registerDo) Not(conds ...gen.Condition) IRegisterDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r registerDo) Or(conds ...gen.Condition) IRegisterDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r registerDo) Select(conds ...field.Expr) IRegisterDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r registerDo) Where(conds ...gen.Condition) IRegisterDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r registerDo) Order(conds ...field.Expr) IRegisterDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r registerDo) Distinct(cols ...field.Expr) IRegisterDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r registerDo) Omit(cols ...field.Expr) IRegisterDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r registerDo) Join(table schema.Tabler, on ...field.Expr) IRegisterDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r registerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRegisterDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r registerDo) RightJoin(table schema.Tabler, on ...field.Expr) IRegisterDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r registerDo) Group(cols ...field.Expr) IRegisterDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r registerDo) Having(conds ...gen.Condition) IRegisterDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r registerDo) Limit(limit int) IRegisterDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r registerDo) Offset(offset int) IRegisterDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r registerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRegisterDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r registerDo) Unscoped() IRegisterDo {
	return r.withDO(r.DO.Unscoped())
}

func (r registerDo) Create(values ...*model.Register) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r registerDo) CreateInBatches(values []*model.Register, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r registerDo) Save(values ...*model.Register) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r registerDo) First() (*model.Register, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Register), nil
	}
}

func (r registerDo) Take() (*model.Register, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Register), nil
	}
}

func (r registerDo) Last() (*model.Register, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Register), nil
	}
}

func (r registerDo) Find() ([]*model.Register, error) {
	result, err := r.DO.Find()
	return result.([]*model.Register), err
}

func (r registerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Register, err error) {
	buf := make([]*model.Register, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r registerDo) FindInBatches(result *[]*model.Register, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r registerDo) Attrs(attrs ...field.AssignExpr) IRegisterDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r registerDo) Assign(attrs ...field.AssignExpr) IRegisterDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r registerDo) Joins(fields ...field.RelationField) IRegisterDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r registerDo) Preload(fields ...field.RelationField) IRegisterDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r registerDo) FirstOrInit() (*model.Register, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Register), nil
	}
}

func (r registerDo) FirstOrCreate() (*model.Register, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Register), nil
	}
}

func (r registerDo) FindByPage(offset int, limit int) (result []*model.Register, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r registerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r registerDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r registerDo) Delete(models ...*model.Register) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *registerDo) withDO(do gen.Dao) *registerDo {
	r.DO = *do.(*gen.DO)
	return r
}
