package paginator

import (
	"fmt"
	"liu/pkg/config"
	"liu/pkg/logger"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Paging struct {
	CurrentPage int    // 当前页
	PerPage     int    // 每页条数
	TotalPage   int    // 总页数
	TotalCount  int64  // 总条数
	NextPageURL string // 下一页的链接
	PrevPageURL string // 上一页的链接
}

// Paginator 分页操作类
type Paginator struct {
	BaseURL    string       // 用以拼接 URL
	PerPage    int          // 每页条数
	Page       int          // 当前页
	TotalCount int64        // 总条数
	TotalPage  int          // 总页数 = TotalCount/PerPage
	Offset     int          // 数据库读取数据时 Offset 的值
	Sort       string       // 排序规则
	Order      string       // 排序顺序
	query      *gorm.DB     // db query 句柄
	ctx        *gin.Context // gin context，方便调用
}

func Paginate(ctx *gin.Context, db *gorm.DB, data interface{}, baseURL string, perPage int) Paging {
	p := &Paginator{
		query: db,
		ctx:   ctx,
	}

	// 初始化参数
	p.initProperties(perPage, baseURL)

	// 查询数据
	err := p.query.Preload(clause.Associations).
		Order(p.Sort + " " + p.Order).
		Limit(p.PerPage).
		Offset(p.Offset).
		Find(data).
		Error

	if err != nil {
		logger.LogIf(err)
		return Paging{}
	}

	return Paging{
		CurrentPage: p.Page,
		PerPage:     p.PerPage,
		TotalPage:   p.TotalPage,
		TotalCount:  p.TotalCount,
		NextPageURL: p.getNextPageURL(),
		PrevPageURL: p.getPrevPageURL(),
	}
}

func (p *Paginator) initProperties(perPage int, baseURL string) {
	p.BaseURL = p.formatBaseURL(baseURL)
	p.PerPage = p.getPerPage(perPage)

	// 时间处理
	created_at := p.ctx.Query("created_at")
	updated_at := p.ctx.Query("updated_at")
	p.getTimeWhere(created_at, updated_at)

	// 搜索字段
	search := p.ctx.Query(config.Get("paging.url_query_search"))
	if len(search) > 0 {
		p.getCurrentWhere(search)
	}

	p.Order = p.ctx.DefaultQuery(config.Get("paging.url_query_order"), "asc")
	p.Sort = p.ctx.DefaultQuery(config.Get("paging.url_query_sort"), "id")

	p.TotalCount = p.getTotalCount()
	p.TotalPage = p.getTotalPage()

	p.Page = p.getCurrentPage()

	p.Offset = (p.Page - 1) * p.PerPage
}

func (p Paginator) getPerPage(perPage int) int {
	queryPage := p.ctx.Query(config.Get("paging.url_query_per_page"))
	if len(queryPage) > 0 {
		perPage = cast.ToInt(queryPage)
	}
	if perPage <= 0 {
		perPage = config.GetInt("paging.per_page")
	}
	return perPage
}

// 返回当前页码
func (p Paginator) getCurrentPage() int {
	page := cast.ToInt(p.ctx.Query(config.Get("paging.url_query_page")))
	// 默认分页
	if page <= 0 {
		page = 1
	}
	// TotalPage 等于 0 ，意味着数据不够分页
	if p.TotalPage == 0 {
		return 0
	}
	// 请求页数大于总页数，返回总页数
	if page > p.TotalPage {
		return p.TotalPage
	}
	return page
}

// 返回搜索条件( 目前支持 name， user，title等)
func (p Paginator) getCurrentWhere(search string) interface{} {
	return p.query.Where("user LIKE ? OR name LIKE ? ", "%"+search+"%", "%"+search+"%")
}

// 时间条件搜索
func (p Paginator) getTimeWhere(created_at string, updated_at string) interface{} {
	if len(created_at) > 0 && len(updated_at) > 0 {
		p.query.Where("created_at BETWEEN ? AND ?", created_at, updated_at)
	} else if len(created_at) > 0 {
		p.query.Where("created_at <= ?", created_at)
	} else if len(updated_at) > 0 {
		p.query.Where("created_at <= ?", updated_at)
	}
	return p.query
}

// 返回的是数据库里的条数
func (p *Paginator) getTotalCount() int64 {
	var count int64
	if err := p.query.Count(&count).Error; err != nil {
		return 0
	}
	return count
}

// 计算总页数
func (p Paginator) getTotalPage() int {
	if p.TotalCount == 0 {
		return 0
	}
	nums := int64(math.Ceil(float64(p.TotalCount) / float64(p.PerPage)))
	if nums == 0 {
		nums = 1
	}
	return int(nums)
}

// 兼容 URL 带与不带 `?` 的情况
func (p *Paginator) formatBaseURL(baseURL string) string {
	if strings.Contains(baseURL, "?") {
		baseURL = baseURL + "&" + config.Get("paging.url_query_page") + "="
	} else {
		baseURL = baseURL + "?" + config.Get("paging.url_query_page") + "="
	}
	return baseURL
}

// 拼接分页链接
func (p Paginator) getPageLink(page int) string {
	return fmt.Sprintf(
		"%v%v&%s=%s&%s=%s&%s=%v",
		p.BaseURL,
		page,
		config.Get("paging.url_query_sort"),
		p.Sort,
		config.Get("paging.url_query_order"),
		p.Order,
		config.Get("paging.url_query_per_page"),
		p.PerPage,
	)
}

// 返回下一页的链接
func (p Paginator) getNextPageURL() string {
	if p.TotalPage > p.Page {
		return p.getPageLink(p.Page + 1)
	}
	return ""
}

// 返回下一页的链接
func (p Paginator) getPrevPageURL() string {
	if p.Page <= 1 || p.Page > p.TotalPage {
		return ""
	}
	return p.getPageLink(p.Page - 1)
}
