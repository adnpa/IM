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

func newFriendApply(db *gorm.DB, opts ...gen.DOOption) friendApply {
	_friendApply := friendApply{}

	_friendApply.friendApplyDo.UseDB(db, opts...)
	_friendApply.friendApplyDo.UseModel(&model.FriendApply{})

	tableName := _friendApply.friendApplyDo.TableName()
	_friendApply.ALL = field.NewAsterisk(tableName)
	_friendApply.ID = field.NewInt32(tableName, "id")
	_friendApply.FromID = field.NewInt32(tableName, "from_id")
	_friendApply.ToID = field.NewInt32(tableName, "to_id")
	_friendApply.Status = field.NewInt32(tableName, "status")
	_friendApply.ApplyReason = field.NewString(tableName, "apply_reason")
	_friendApply.CreatedAt = field.NewTime(tableName, "created_at")
	_friendApply.UpdatedAt = field.NewTime(tableName, "updated_at")
	_friendApply.DeletedAt = field.NewField(tableName, "deleted_at")

	_friendApply.fillFieldMap()

	return _friendApply
}

// friendApply 好友申请表
type friendApply struct {
	friendApplyDo

	ALL         field.Asterisk
	ID          field.Int32  // 申请ID
	FromID      field.Int32  // 申请者ID
	ToID        field.Int32  // 被申请者ID
	Status      field.Int32  // 申请状态
	ApplyReason field.String // 申请理由
	CreatedAt   field.Time   // 创建时间
	UpdatedAt   field.Time   // 更新时间
	DeletedAt   field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (f friendApply) Table(newTableName string) *friendApply {
	f.friendApplyDo.UseTable(newTableName)
	return f.updateTableName(newTableName)
}

func (f friendApply) As(alias string) *friendApply {
	f.friendApplyDo.DO = *(f.friendApplyDo.As(alias).(*gen.DO))
	return f.updateTableName(alias)
}

func (f *friendApply) updateTableName(table string) *friendApply {
	f.ALL = field.NewAsterisk(table)
	f.ID = field.NewInt32(table, "id")
	f.FromID = field.NewInt32(table, "from_id")
	f.ToID = field.NewInt32(table, "to_id")
	f.Status = field.NewInt32(table, "status")
	f.ApplyReason = field.NewString(table, "apply_reason")
	f.CreatedAt = field.NewTime(table, "created_at")
	f.UpdatedAt = field.NewTime(table, "updated_at")
	f.DeletedAt = field.NewField(table, "deleted_at")

	f.fillFieldMap()

	return f
}

func (f *friendApply) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := f.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (f *friendApply) fillFieldMap() {
	f.fieldMap = make(map[string]field.Expr, 8)
	f.fieldMap["id"] = f.ID
	f.fieldMap["from_id"] = f.FromID
	f.fieldMap["to_id"] = f.ToID
	f.fieldMap["status"] = f.Status
	f.fieldMap["apply_reason"] = f.ApplyReason
	f.fieldMap["created_at"] = f.CreatedAt
	f.fieldMap["updated_at"] = f.UpdatedAt
	f.fieldMap["deleted_at"] = f.DeletedAt
}

func (f friendApply) clone(db *gorm.DB) friendApply {
	f.friendApplyDo.ReplaceConnPool(db.Statement.ConnPool)
	return f
}

func (f friendApply) replaceDB(db *gorm.DB) friendApply {
	f.friendApplyDo.ReplaceDB(db)
	return f
}

type friendApplyDo struct{ gen.DO }

type IFriendApplyDo interface {
	gen.SubQuery
	Debug() IFriendApplyDo
	WithContext(ctx context.Context) IFriendApplyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IFriendApplyDo
	WriteDB() IFriendApplyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IFriendApplyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IFriendApplyDo
	Not(conds ...gen.Condition) IFriendApplyDo
	Or(conds ...gen.Condition) IFriendApplyDo
	Select(conds ...field.Expr) IFriendApplyDo
	Where(conds ...gen.Condition) IFriendApplyDo
	Order(conds ...field.Expr) IFriendApplyDo
	Distinct(cols ...field.Expr) IFriendApplyDo
	Omit(cols ...field.Expr) IFriendApplyDo
	Join(table schema.Tabler, on ...field.Expr) IFriendApplyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IFriendApplyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IFriendApplyDo
	Group(cols ...field.Expr) IFriendApplyDo
	Having(conds ...gen.Condition) IFriendApplyDo
	Limit(limit int) IFriendApplyDo
	Offset(offset int) IFriendApplyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IFriendApplyDo
	Unscoped() IFriendApplyDo
	Create(values ...*model.FriendApply) error
	CreateInBatches(values []*model.FriendApply, batchSize int) error
	Save(values ...*model.FriendApply) error
	First() (*model.FriendApply, error)
	Take() (*model.FriendApply, error)
	Last() (*model.FriendApply, error)
	Find() ([]*model.FriendApply, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FriendApply, err error)
	FindInBatches(result *[]*model.FriendApply, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.FriendApply) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IFriendApplyDo
	Assign(attrs ...field.AssignExpr) IFriendApplyDo
	Joins(fields ...field.RelationField) IFriendApplyDo
	Preload(fields ...field.RelationField) IFriendApplyDo
	FirstOrInit() (*model.FriendApply, error)
	FirstOrCreate() (*model.FriendApply, error)
	FindByPage(offset int, limit int) (result []*model.FriendApply, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IFriendApplyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (f friendApplyDo) Debug() IFriendApplyDo {
	return f.withDO(f.DO.Debug())
}

func (f friendApplyDo) WithContext(ctx context.Context) IFriendApplyDo {
	return f.withDO(f.DO.WithContext(ctx))
}

func (f friendApplyDo) ReadDB() IFriendApplyDo {
	return f.Clauses(dbresolver.Read)
}

func (f friendApplyDo) WriteDB() IFriendApplyDo {
	return f.Clauses(dbresolver.Write)
}

func (f friendApplyDo) Session(config *gorm.Session) IFriendApplyDo {
	return f.withDO(f.DO.Session(config))
}

func (f friendApplyDo) Clauses(conds ...clause.Expression) IFriendApplyDo {
	return f.withDO(f.DO.Clauses(conds...))
}

func (f friendApplyDo) Returning(value interface{}, columns ...string) IFriendApplyDo {
	return f.withDO(f.DO.Returning(value, columns...))
}

func (f friendApplyDo) Not(conds ...gen.Condition) IFriendApplyDo {
	return f.withDO(f.DO.Not(conds...))
}

func (f friendApplyDo) Or(conds ...gen.Condition) IFriendApplyDo {
	return f.withDO(f.DO.Or(conds...))
}

func (f friendApplyDo) Select(conds ...field.Expr) IFriendApplyDo {
	return f.withDO(f.DO.Select(conds...))
}

func (f friendApplyDo) Where(conds ...gen.Condition) IFriendApplyDo {
	return f.withDO(f.DO.Where(conds...))
}

func (f friendApplyDo) Order(conds ...field.Expr) IFriendApplyDo {
	return f.withDO(f.DO.Order(conds...))
}

func (f friendApplyDo) Distinct(cols ...field.Expr) IFriendApplyDo {
	return f.withDO(f.DO.Distinct(cols...))
}

func (f friendApplyDo) Omit(cols ...field.Expr) IFriendApplyDo {
	return f.withDO(f.DO.Omit(cols...))
}

func (f friendApplyDo) Join(table schema.Tabler, on ...field.Expr) IFriendApplyDo {
	return f.withDO(f.DO.Join(table, on...))
}

func (f friendApplyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IFriendApplyDo {
	return f.withDO(f.DO.LeftJoin(table, on...))
}

func (f friendApplyDo) RightJoin(table schema.Tabler, on ...field.Expr) IFriendApplyDo {
	return f.withDO(f.DO.RightJoin(table, on...))
}

func (f friendApplyDo) Group(cols ...field.Expr) IFriendApplyDo {
	return f.withDO(f.DO.Group(cols...))
}

func (f friendApplyDo) Having(conds ...gen.Condition) IFriendApplyDo {
	return f.withDO(f.DO.Having(conds...))
}

func (f friendApplyDo) Limit(limit int) IFriendApplyDo {
	return f.withDO(f.DO.Limit(limit))
}

func (f friendApplyDo) Offset(offset int) IFriendApplyDo {
	return f.withDO(f.DO.Offset(offset))
}

func (f friendApplyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IFriendApplyDo {
	return f.withDO(f.DO.Scopes(funcs...))
}

func (f friendApplyDo) Unscoped() IFriendApplyDo {
	return f.withDO(f.DO.Unscoped())
}

func (f friendApplyDo) Create(values ...*model.FriendApply) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Create(values)
}

func (f friendApplyDo) CreateInBatches(values []*model.FriendApply, batchSize int) error {
	return f.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (f friendApplyDo) Save(values ...*model.FriendApply) error {
	if len(values) == 0 {
		return nil
	}
	return f.DO.Save(values)
}

func (f friendApplyDo) First() (*model.FriendApply, error) {
	if result, err := f.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.FriendApply), nil
	}
}

func (f friendApplyDo) Take() (*model.FriendApply, error) {
	if result, err := f.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.FriendApply), nil
	}
}

func (f friendApplyDo) Last() (*model.FriendApply, error) {
	if result, err := f.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.FriendApply), nil
	}
}

func (f friendApplyDo) Find() ([]*model.FriendApply, error) {
	result, err := f.DO.Find()
	return result.([]*model.FriendApply), err
}

func (f friendApplyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.FriendApply, err error) {
	buf := make([]*model.FriendApply, 0, batchSize)
	err = f.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (f friendApplyDo) FindInBatches(result *[]*model.FriendApply, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return f.DO.FindInBatches(result, batchSize, fc)
}

func (f friendApplyDo) Attrs(attrs ...field.AssignExpr) IFriendApplyDo {
	return f.withDO(f.DO.Attrs(attrs...))
}

func (f friendApplyDo) Assign(attrs ...field.AssignExpr) IFriendApplyDo {
	return f.withDO(f.DO.Assign(attrs...))
}

func (f friendApplyDo) Joins(fields ...field.RelationField) IFriendApplyDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Joins(_f))
	}
	return &f
}

func (f friendApplyDo) Preload(fields ...field.RelationField) IFriendApplyDo {
	for _, _f := range fields {
		f = *f.withDO(f.DO.Preload(_f))
	}
	return &f
}

func (f friendApplyDo) FirstOrInit() (*model.FriendApply, error) {
	if result, err := f.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.FriendApply), nil
	}
}

func (f friendApplyDo) FirstOrCreate() (*model.FriendApply, error) {
	if result, err := f.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.FriendApply), nil
	}
}

func (f friendApplyDo) FindByPage(offset int, limit int) (result []*model.FriendApply, count int64, err error) {
	result, err = f.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = f.Offset(-1).Limit(-1).Count()
	return
}

func (f friendApplyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = f.Count()
	if err != nil {
		return
	}

	err = f.Offset(offset).Limit(limit).Scan(result)
	return
}

func (f friendApplyDo) Scan(result interface{}) (err error) {
	return f.DO.Scan(result)
}

func (f friendApplyDo) Delete(models ...*model.FriendApply) (result gen.ResultInfo, err error) {
	return f.DO.Delete(models)
}

func (f *friendApplyDo) withDO(do gen.Dao) *friendApplyDo {
	f.DO = *do.(*gen.DO)
	return f
}
