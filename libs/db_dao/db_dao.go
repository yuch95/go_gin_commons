package db_dao

type QueryType map[string]any

// QueryParams 查询参数配置
type QueryParams struct {
	QueryData      QueryType // 查询条件
	Entities       []string  // 返回字段限制
	OrderFields    []string  // 排序字段
	DistinctFields []string  // 去重字段
	Page           int       // 页码
	PageSize       int       // 条目数
}

type DBDao interface {
	Insert(obj any) error       // 插入一条数据 也可以插入多条
	InsertMany(obj []any) error // 插入多条数据
	Update(obj any, queryParam *QueryParams, updateData map[string]any) error
	UpdateInsert(obj any, queryParam *QueryParams, updateData map[string]any) error
	QueryOne(obj any, queryParam *QueryParams)  // 查询一条数据
	QueryAll(obj any, queryParam *QueryParams)  // 查询全部数据
	PageQuery(obj any, queryParam *QueryParams) // 分页查询
	Commit()
	Rollback()
}
