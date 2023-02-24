package utils

// 获取列表数据时的公共参数
type ListOptions struct {
	Page     int
	PageSize int
	Keyword  string
	Order    string // eg. id desc
}

// 对列表参数进行兜底
func (options *ListOptions) Prepare() *ListOptions {
	if options.Page <= 0 {
		options.Page = 1
	}
	if options.PageSize == 0 {
		options.PageSize = 10
	} else if options.PageSize < 0 {
		options.PageSize = -1
	}
	if options.Order == "" {
		options.Order = "id desc"
	}
	return options
}

// 是否不限制数量
func (options *ListOptions) NoLimit() bool {
	return options.PageSize == -1
}

// 对列表参数计算偏移
func (options *ListOptions) Offset() int {
	if options.NoLimit() {
		return 0
	}
	return (options.Page - 1) * options.PageSize
}

// 模糊查询关键词
func (options *ListOptions) EscKeyword() string {
	return "%" + options.Keyword + "%"
}
