package common


type PageParam struct {
	Page int64 `json:"page"`
	TotalPage int64 `json:"total_page"`
	TotalRecord int64 `json:"total_record"`
	PageSize int64 `json:"page_size"`
}

func NewPageParam(page, totalRecord, pageSize int64) *PageParam {
	res := &PageParam{
		Page: page,
		TotalRecord: totalRecord,
		PageSize: pageSize,
	}
	totalPage := totalRecord / pageSize
	if totalRecord % pageSize > 0 {
		totalPage++
	}
	res.TotalPage = totalPage
	return res
}