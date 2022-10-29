// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/dbresolver"

	"github.com/go-sonic/sonic/model/entity"
)

func newCommentBlack(db *gorm.DB) commentBlack {
	_commentBlack := commentBlack{}

	_commentBlack.commentBlackDo.UseDB(db)
	_commentBlack.commentBlackDo.UseModel(&entity.CommentBlack{})

	tableName := _commentBlack.commentBlackDo.TableName()
	_commentBlack.ALL = field.NewAsterisk(tableName)
	_commentBlack.ID = field.NewInt64(tableName, "id")
	_commentBlack.CreateTime = field.NewTime(tableName, "create_time")
	_commentBlack.UpdateTime = field.NewTime(tableName, "update_time")
	_commentBlack.BanTime = field.NewTime(tableName, "ban_time")
	_commentBlack.IPAddress = field.NewString(tableName, "ip_address")

	_commentBlack.fillFieldMap()

	return _commentBlack
}

type commentBlack struct {
	commentBlackDo commentBlackDo

	ALL        field.Asterisk
	ID         field.Int64
	CreateTime field.Time
	UpdateTime field.Time
	BanTime    field.Time
	IPAddress  field.String

	fieldMap map[string]field.Expr
}

func (c commentBlack) Table(newTableName string) *commentBlack {
	c.commentBlackDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c commentBlack) As(alias string) *commentBlack {
	c.commentBlackDo.DO = *(c.commentBlackDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *commentBlack) updateTableName(table string) *commentBlack {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.CreateTime = field.NewTime(table, "create_time")
	c.UpdateTime = field.NewTime(table, "update_time")
	c.BanTime = field.NewTime(table, "ban_time")
	c.IPAddress = field.NewString(table, "ip_address")

	c.fillFieldMap()

	return c
}

func (c *commentBlack) WithContext(ctx context.Context) *commentBlackDo {
	return c.commentBlackDo.WithContext(ctx)
}

func (c commentBlack) TableName() string { return c.commentBlackDo.TableName() }

func (c commentBlack) Alias() string { return c.commentBlackDo.Alias() }

func (c *commentBlack) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *commentBlack) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["create_time"] = c.CreateTime
	c.fieldMap["update_time"] = c.UpdateTime
	c.fieldMap["ban_time"] = c.BanTime
	c.fieldMap["ip_address"] = c.IPAddress
}

func (c commentBlack) clone(db *gorm.DB) commentBlack {
	c.commentBlackDo.ReplaceDB(db)
	return c
}

type commentBlackDo struct{ gen.DO }

func (c commentBlackDo) Debug() *commentBlackDo {
	return c.withDO(c.DO.Debug())
}

func (c commentBlackDo) WithContext(ctx context.Context) *commentBlackDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c commentBlackDo) ReadDB() *commentBlackDo {
	return c.Clauses(dbresolver.Read)
}

func (c commentBlackDo) WriteDB() *commentBlackDo {
	return c.Clauses(dbresolver.Write)
}

func (c commentBlackDo) Clauses(conds ...clause.Expression) *commentBlackDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c commentBlackDo) Returning(value interface{}, columns ...string) *commentBlackDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c commentBlackDo) Not(conds ...gen.Condition) *commentBlackDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c commentBlackDo) Or(conds ...gen.Condition) *commentBlackDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c commentBlackDo) Select(conds ...field.Expr) *commentBlackDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c commentBlackDo) Where(conds ...gen.Condition) *commentBlackDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c commentBlackDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *commentBlackDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c commentBlackDo) Order(conds ...field.Expr) *commentBlackDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c commentBlackDo) Distinct(cols ...field.Expr) *commentBlackDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c commentBlackDo) Omit(cols ...field.Expr) *commentBlackDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c commentBlackDo) Join(table schema.Tabler, on ...field.Expr) *commentBlackDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c commentBlackDo) LeftJoin(table schema.Tabler, on ...field.Expr) *commentBlackDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c commentBlackDo) RightJoin(table schema.Tabler, on ...field.Expr) *commentBlackDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c commentBlackDo) Group(cols ...field.Expr) *commentBlackDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c commentBlackDo) Having(conds ...gen.Condition) *commentBlackDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c commentBlackDo) Limit(limit int) *commentBlackDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c commentBlackDo) Offset(offset int) *commentBlackDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c commentBlackDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *commentBlackDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c commentBlackDo) Unscoped() *commentBlackDo {
	return c.withDO(c.DO.Unscoped())
}

func (c commentBlackDo) Create(values ...*entity.CommentBlack) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c commentBlackDo) CreateInBatches(values []*entity.CommentBlack, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c commentBlackDo) Save(values ...*entity.CommentBlack) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c commentBlackDo) First() (*entity.CommentBlack, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*entity.CommentBlack), nil
	}
}

func (c commentBlackDo) Take() (*entity.CommentBlack, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*entity.CommentBlack), nil
	}
}

func (c commentBlackDo) Last() (*entity.CommentBlack, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*entity.CommentBlack), nil
	}
}

func (c commentBlackDo) Find() ([]*entity.CommentBlack, error) {
	result, err := c.DO.Find()
	return result.([]*entity.CommentBlack), err
}

func (c commentBlackDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*entity.CommentBlack, err error) {
	buf := make([]*entity.CommentBlack, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c commentBlackDo) FindInBatches(result *[]*entity.CommentBlack, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c commentBlackDo) Attrs(attrs ...field.AssignExpr) *commentBlackDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c commentBlackDo) Assign(attrs ...field.AssignExpr) *commentBlackDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c commentBlackDo) Joins(fields ...field.RelationField) *commentBlackDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c commentBlackDo) Preload(fields ...field.RelationField) *commentBlackDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c commentBlackDo) FirstOrInit() (*entity.CommentBlack, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*entity.CommentBlack), nil
	}
}

func (c commentBlackDo) FirstOrCreate() (*entity.CommentBlack, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*entity.CommentBlack), nil
	}
}

func (c commentBlackDo) FindByPage(offset int, limit int) (result []*entity.CommentBlack, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c commentBlackDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c commentBlackDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c commentBlackDo) Delete(models ...*entity.CommentBlack) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *commentBlackDo) withDO(do gen.Dao) *commentBlackDo {
	c.DO = *do.(*gen.DO)
	return c
}
