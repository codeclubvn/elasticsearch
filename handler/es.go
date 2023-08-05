package handler

import (
	"elasticsearch/model"
	"elasticsearch/service"
	"encoding/json"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Elasticsearch struct {
	service service.IEsService
}

func NewESHandler(service service.IEsService) *Elasticsearch {
	return &Elasticsearch{
		service: service,
	}
}

func (h *Elasticsearch) Test(c *gin.Context) {
	infoInterface, err := h.service.Test(c)
	if err != nil {
		c.JSON(401, gin.H{
			"message": err.Error(),
		})
		return
	}
	// Type assertion để chuyển đổi interface thành cấu trúc
	infoResponse, ok := infoInterface.(*esapi.Response)
	if !ok {
		fmt.Println("Failed to convert interface to *esapi.Response")
		return
	}

	// Đọc dữ liệu từ infoResponse.Body
	bodyBytes, err := ioutil.ReadAll(infoResponse.Body)
	if err != nil {
		fmt.Println("Failed to read response body: ", err)
		return
	}

	//// Đóng nơi đọc dữ liệu
	infoResponse.Body.Close()

	// Xử lý phản hồi và chuyển đổi thành cấu trúc
	info := model.ElasticsearchInfo{}
	if err := json.Unmarshal(bodyBytes, &info); err != nil {
		fmt.Println("Failed to unmarshal response: ", err)
		return
	}

	c.JSON(200, bodyBytes)
}

func (h *Elasticsearch) Insert(c *gin.Context) {
	// Lấy thông tin từ request
	ESRequest := model.ESRequest{}
	if err := c.ShouldBindJSON(&ESRequest); err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid request",
		})
		return
	}

	// call service
	if err := h.service.Insert(c, ESRequest); err != nil {
		c.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Create index success",
	})
}
