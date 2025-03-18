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

	"github.com/adnpa/IM/scripts/gen_gorm/model"
)

func newGroupApply(db *gorm.DB, opts ...gen.DOOption) groupApply {
	_groupApply := groupApply{}

	_groupApply.groupApplyDo.UseDB(db, opts...)
	_groupApply.groupApplyDo.UseModel(&model.GroupApply{})

	tableName := _groupApply.groupApplyDo.TableName()
	_groupApply.ALL = field.NewAsterisk(tableName)
	_groupApply.ID = field.NewInt64(tableName, "id")
	_groupApply.GroupID = field.NewInt64(tableName, "group_id")
	_groupApply.ApplicantID = field.NewInt32(tableName, "applicant_id")
	_groupApply.Status = field.NewInt32(tableName, "status")
	_groupApply.HandlerID = field.NewInt32(tableName, "handler_id")
	_groupApply.CreatedAt = field.NewTime(tableName, "created_at")
	_groupApply.UpdatedAt = field.NewTime(tableName, "updated_at")
	_groupApply.DeletedAt = field.NewField(tableName, "deleted_at")

	_groupApply.fillFieldMap()

	return _groupApply
}

// groupApply 群聊申请表
type groupApply struct {
	groupApplyDo

	ALL         field.Asterisk
	ID          field.Int64 // 申请ID，主键，自增
	GroupID     field.Int64 // 群聊ID，外键
	ApplicantID field.Int32 // 申请人用户ID
	Status      field.Int32 // 申请状态（0:待处理，1:已通过，2:已拒绝）
	HandlerID   field.Int32 // 处理人用户ID
	CreatedAt   field.Time  // 创建时间
	UpdatedAt   field.Time  // 更新时间
	DeletedAt   field.Field // 删除时间

	fieldMap map[string]field.Expr
}

func (g groupApply) Table(newTableName string) *groupApply {
	g.groupApplyDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g groupApply) As(alias string) *groupApply {
	g.groupApplyDo.DO = *(g.groupApplyDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *groupApply) updateTableName(table string) *groupApply {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewInt64(table, "id")
	g.GroupID = field.NewInt64(table, "group_id")
	g.ApplicantID = field.NewInt32(table, "applicant_id")
	g.Status = field.NewInt32(table, "status")
	g.HandlerID = field.NewInt32(table, "handler_id")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")
	g.DeletedAt = field.NewField(table, "deleted_at")

	g.fillFieldMap()

	return g
}

func (g *groupApply) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *groupApply) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 8)
	g.fieldMap["id"] = g.ID
	g.fieldMap["group_id"] = g.GroupID
	g.fieldMap["applicant_id"] = g.ApplicantID
	g.fieldMap["status"] = g.Status
	g.fieldMap["handler_id"] = g.HandlerID
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["deleted_at"] = g.DeletedAt
}

func (g groupApply) clone(db *gorm.DB) groupApply {
	g.groupApplyDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g groupApply) replaceDB(db *gorm.DB) groupApply {
	g.groupApplyDo.ReplaceDB(db)
	return g
}

type groupApplyDo struct{ gen.DO }

type IGroupApplyDo interface {
	gen.SubQuery
	Debug() IGroupApplyDo
	WithContext(ctx context.Context) IGroupApplyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGroupApplyDo
	WriteDB() IGroupApplyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGroupApplyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGroupApplyDo
	Not(conds ...gen.Condition) IGroupApplyDo
	Or(conds ...gen.Condition) IGroupApplyDo
	Select(conds ...field.Expr) IGroupApplyDo
	Where(conds ...gen.Condition) IGroupApplyDo
	Order(conds ...field.Expr) IGroupApplyDo
	Distinct(cols ...field.Expr) IGroupApplyDo
	Omit(cols ...field.Expr) IGroupApplyDo
	Join(table schema.Tabler, on ...field.Expr) IGroupApplyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGroupApplyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGroupApplyDo
	Group(cols ...field.Expr) IGroupApplyDo
	Having(conds ...gen.Condition) IGroupApplyDo
	Limit(limit int) IGroupApplyDo
	Offset(offset int) IGroupApplyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupApplyDo
	Unscoped() IGroupApplyDo
	Create(values ...*model.GroupApply) error
	CreateInBatches(values []*model.GroupApply, batchSize int) error
	Save(values ...*model.GroupApply) error
	First() (*model.GroupApply, error)
	Take() (*model.GroupApply, error)
	Last() (*model.GroupApply, error)
	Find() ([]*model.GroupApply, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupApply, err error)
	FindInBatches(result *[]*model.GroupApply, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.GroupApply) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGroupApplyDo
	Assign(attrs ...field.AssignExpr) IGroupApplyDo
	Joins(fields ...field.RelationField) IGroupApplyDo
	Preload(fields ...field.RelationField) IGroupApplyDo
	FirstOrInit() (*model.GroupApply, error)
	FirstOrCreate() (*model.GroupApply, error)
	FindByPage(offset int, limit int) (result []*model.GroupApply, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGroupApplyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g groupApplyDo) Debug() IGroupApplyDo {
	return g.withDO(g.DO.Debug())
}

func (g groupApplyDo) WithContext(ctx context.Context) IGroupApplyDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g groupApplyDo) ReadDB() IGroupApplyDo {
	return g.Clauses(dbresolver.Read)
}

func (g groupApplyDo) WriteDB() IGroupApplyDo {
	return g.Clauses(dbresolver.Write)
}

func (g groupApplyDo) Session(config *gorm.Session) IGroupApplyDo {
	return g.withDO(g.DO.Session(config))
}

func (g groupApplyDo) Clauses(conds ...clause.Expression) IGroupApplyDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g groupApplyDo) Returning(value interface{}, columns ...string) IGroupApplyDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g groupApplyDo) Not(conds ...gen.Condition) IGroupApplyDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g groupApplyDo) Or(conds ...gen.Condition) IGroupApplyDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g groupApplyDo) Select(conds ...field.Expr) IGroupApplyDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g groupApplyDo) Where(conds ...gen.Condition) IGroupApplyDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g groupApplyDo) Order(conds ...field.Expr) IGroupApplyDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g groupApplyDo) Distinct(cols ...field.Expr) IGroupApplyDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g groupApplyDo) Omit(cols ...field.Expr) IGroupApplyDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g groupApplyDo) Join(table schema.Tabler, on ...field.Expr) IGroupApplyDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g groupApplyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGroupApplyDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g groupApplyDo) RightJoin(table schema.Tabler, on ...field.Expr) IGroupApplyDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g groupApplyDo) Group(cols ...field.Expr) IGroupApplyDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g groupApplyDo) Having(conds ...gen.Condition) IGroupApplyDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g groupApplyDo) Limit(limit int) IGroupApplyDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g groupApplyDo) Offset(offset int) IGroupApplyDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g groupApplyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupApplyDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g groupApplyDo) Unscoped() IGroupApplyDo {
	return g.withDO(g.DO.Unscoped())
}

func (g groupApplyDo) Create(values ...*model.GroupApply) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g groupApplyDo) CreateInBatches(values []*model.GroupApply, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g groupApplyDo) Save(values ...*model.GroupApply) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g groupApplyDo) First() (*model.GroupApply, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApply), nil
	}
}

func (g groupApplyDo) Take() (*model.GroupApply, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApply), nil
	}
}

func (g groupApplyDo) Last() (*model.GroupApply, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApply), nil
	}
}

func (g groupApplyDo) Find() ([]*model.GroupApply, error) {
	result, err := g.DO.Find()
	return result.([]*model.GroupApply), err
}

func (g groupApplyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.GroupApply, err error) {
	buf := make([]*model.GroupApply, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g groupApplyDo) FindInBatches(result *[]*model.GroupApply, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g groupApplyDo) Attrs(attrs ...field.AssignExpr) IGroupApplyDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g groupApplyDo) Assign(attrs ...field.AssignExpr) IGroupApplyDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g groupApplyDo) Joins(fields ...field.RelationField) IGroupApplyDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g groupApplyDo) Preload(fields ...field.RelationField) IGroupApplyDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g groupApplyDo) FirstOrInit() (*model.GroupApply, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApply), nil
	}
}

func (g groupApplyDo) FirstOrCreate() (*model.GroupApply, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.GroupApply), nil
	}
}

func (g groupApplyDo) FindByPage(offset int, limit int) (result []*model.GroupApply, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g groupApplyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g groupApplyDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g groupApplyDo) Delete(models ...*model.GroupApply) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *groupApplyDo) withDO(do gen.Dao) *groupApplyDo {
	g.DO = *do.(*gen.DO)
	return g
}
