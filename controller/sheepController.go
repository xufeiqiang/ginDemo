package controller

import (
	"ginDemo/config"
	"ginDemo/model"
	"ginDemo/model/vo"
	"ginDemo/service/impl"
	"github.com/gin-gonic/gin"
	"net/http"
)

type SheepController struct {
	sheepService *impl.SheepServiceImpl
}

// 构造器
func NewSheepController(ctx *config.ApplicationContext) interface{} {
	sheepService := ctx.GetBean("sheepService").(*impl.SheepServiceImpl)
	return &SheepController{
		sheepService: sheepService,
	}
}

func (controller *SheepController) GetSheepByValue(c *gin.Context) {
	// 获取url上的参数
	value := c.Param("value")
	data, err := controller.sheepService.GetSheepByCode(value)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "出错了"})
		return
	}
	c.JSON(http.StatusOK, data)
}

// 分页查询,实际开发应该封装 request,response
func (controller *SheepController) GetSheepList(c *gin.Context) {
	var req vo.SheepQueryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "【分页查询】参数错误", "detail": err.Error()})
		return
	}
	list, total, err := controller.sheepService.GetSheepByList(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "【分页查询】失败", "detail": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"list": list,
		"pagination": gin.H{
			"page":      req.PageParam.Page,
			"pageSize":  req.PageParam.PageSize,
			"total":     total,
			"totalPage": (total + int64(req.PageParam.PageSize) - 1) / int64(req.PageParam.PageSize),
		},
	})
}

// 新增
func (controller *SheepController) InsertSheep(c *gin.Context) {
	var sheep model.SheepInfos
	if err := c.ShouldBindJSON(&sheep); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "参数绑定失败", "details": err.Error()})
		return
	}
	if err := controller.sheepService.InsertSheepService(&sheep); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "新增失败", "details": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "新增成功", "details": "success"})
	}
}
