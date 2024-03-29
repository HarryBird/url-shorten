// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"github.com/HarryBird/url-shorten/app/shorten/internal/data/dao/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newURLShortened(db *gorm.DB) uRLShortened {
	_uRLShortened := uRLShortened{}

	_uRLShortened.uRLShortenedDo.UseDB(db)
	_uRLShortened.uRLShortenedDo.UseModel(&model.URLShortened{})

	tableName := _uRLShortened.uRLShortenedDo.TableName()
	_uRLShortened.ALL = field.NewField(tableName, "*")
	_uRLShortened.ID = field.NewInt64(tableName, "id")
	_uRLShortened.URLFull = field.NewString(tableName, "url_full")
	_uRLShortened.URLHost = field.NewString(tableName, "url_host")
	_uRLShortened.URLURI = field.NewString(tableName, "url_uri")
	_uRLShortened.URLQuery = field.NewString(tableName, "url_query")
	_uRLShortened.URLCode = field.NewString(tableName, "url_code")
	_uRLShortened.CreatedAt = field.NewInt64(tableName, "created_at")
	_uRLShortened.UpdatedAt = field.NewInt64(tableName, "updated_at")
	_uRLShortened.DeletedAt = field.NewInt64(tableName, "deleted_at")

	_uRLShortened.fillFieldMap()

	return _uRLShortened
}

type uRLShortened struct {
	uRLShortenedDo uRLShortenedDo

	ALL       field.Field
	ID        field.Int64
	URLFull   field.String
	URLHost   field.String
	URLURI    field.String
	URLQuery  field.String
	URLCode   field.String
	CreatedAt field.Int64
	UpdatedAt field.Int64
	DeletedAt field.Int64

	fieldMap map[string]field.Expr
}

func (u uRLShortened) Table(newTableName string) *uRLShortened {
	u.uRLShortenedDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uRLShortened) As(alias string) *uRLShortened {
	u.uRLShortenedDo.DO = *(u.uRLShortenedDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uRLShortened) updateTableName(table string) *uRLShortened {
	u.ALL = field.NewField(table, "*")
	u.ID = field.NewInt64(table, "id")
	u.URLFull = field.NewString(table, "url_full")
	u.URLHost = field.NewString(table, "url_host")
	u.URLURI = field.NewString(table, "url_uri")
	u.URLQuery = field.NewString(table, "url_query")
	u.URLCode = field.NewString(table, "url_code")
	u.CreatedAt = field.NewInt64(table, "created_at")
	u.UpdatedAt = field.NewInt64(table, "updated_at")
	u.DeletedAt = field.NewInt64(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *uRLShortened) WithContext(ctx context.Context) *uRLShortenedDo {
	return u.uRLShortenedDo.WithContext(ctx)
}

func (u uRLShortened) TableName() string { return u.uRLShortenedDo.TableName() }

func (u *uRLShortened) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uRLShortened) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 9)
	u.fieldMap["id"] = u.ID
	u.fieldMap["url_full"] = u.URLFull
	u.fieldMap["url_host"] = u.URLHost
	u.fieldMap["url_uri"] = u.URLURI
	u.fieldMap["url_query"] = u.URLQuery
	u.fieldMap["url_code"] = u.URLCode
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u uRLShortened) clone(db *gorm.DB) uRLShortened {
	u.uRLShortenedDo.ReplaceDB(db)
	return u
}

type uRLShortenedDo struct{ gen.DO }

func (u uRLShortenedDo) Debug() *uRLShortenedDo {
	return u.withDO(u.DO.Debug())
}

func (u uRLShortenedDo) WithContext(ctx context.Context) *uRLShortenedDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uRLShortenedDo) Clauses(conds ...clause.Expression) *uRLShortenedDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uRLShortenedDo) Not(conds ...gen.Condition) *uRLShortenedDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uRLShortenedDo) Or(conds ...gen.Condition) *uRLShortenedDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uRLShortenedDo) Select(conds ...field.Expr) *uRLShortenedDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uRLShortenedDo) Where(conds ...gen.Condition) *uRLShortenedDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uRLShortenedDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *uRLShortenedDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u uRLShortenedDo) Order(conds ...field.Expr) *uRLShortenedDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uRLShortenedDo) Distinct(cols ...field.Expr) *uRLShortenedDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uRLShortenedDo) Omit(cols ...field.Expr) *uRLShortenedDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uRLShortenedDo) Join(table schema.Tabler, on ...field.Expr) *uRLShortenedDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uRLShortenedDo) LeftJoin(table schema.Tabler, on ...field.Expr) *uRLShortenedDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uRLShortenedDo) RightJoin(table schema.Tabler, on ...field.Expr) *uRLShortenedDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uRLShortenedDo) Group(cols ...field.Expr) *uRLShortenedDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uRLShortenedDo) Having(conds ...gen.Condition) *uRLShortenedDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uRLShortenedDo) Limit(limit int) *uRLShortenedDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uRLShortenedDo) Offset(offset int) *uRLShortenedDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uRLShortenedDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *uRLShortenedDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uRLShortenedDo) Unscoped() *uRLShortenedDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uRLShortenedDo) Create(values ...*model.URLShortened) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uRLShortenedDo) CreateInBatches(values []*model.URLShortened, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uRLShortenedDo) Save(values ...*model.URLShortened) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uRLShortenedDo) First() (*model.URLShortened, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.URLShortened), nil
	}
}

func (u uRLShortenedDo) Take() (*model.URLShortened, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.URLShortened), nil
	}
}

func (u uRLShortenedDo) Last() (*model.URLShortened, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.URLShortened), nil
	}
}

func (u uRLShortenedDo) Find() ([]*model.URLShortened, error) {
	result, err := u.DO.Find()
	return result.([]*model.URLShortened), err
}

func (u uRLShortenedDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.URLShortened, err error) {
	buf := make([]*model.URLShortened, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uRLShortenedDo) FindInBatches(result *[]*model.URLShortened, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uRLShortenedDo) Attrs(attrs ...field.AssignExpr) *uRLShortenedDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uRLShortenedDo) Assign(attrs ...field.AssignExpr) *uRLShortenedDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uRLShortenedDo) Joins(field field.RelationField) *uRLShortenedDo {
	return u.withDO(u.DO.Joins(field))
}

func (u uRLShortenedDo) Preload(field field.RelationField) *uRLShortenedDo {
	return u.withDO(u.DO.Preload(field))
}

func (u uRLShortenedDo) FirstOrInit() (*model.URLShortened, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.URLShortened), nil
	}
}

func (u uRLShortenedDo) FirstOrCreate() (*model.URLShortened, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.URLShortened), nil
	}
}

func (u uRLShortenedDo) FindByPage(offset int, limit int) (result []*model.URLShortened, count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = u.Offset(offset).Limit(limit).Find()
	return
}

func (u uRLShortenedDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u *uRLShortenedDo) withDO(do gen.Dao) *uRLShortenedDo {
	u.DO = *do.(*gen.DO)
	return u
}
