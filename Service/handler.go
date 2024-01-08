package Service

import (
	"github.com/gin-gonic/gin"
	"github.com/ykubernetes/GUUID/utils"
	"net/http"
	"strconv"
)

func CreateSingleUUID(isSimple bool) (Uuid string, err error) {
	if isSimple {
		Uuid = utils.CreateSimpleUUID()
	} else {
		Uuid = utils.CreateUUID()
	}
	return Uuid, nil
}

func CreateHandler(c *gin.Context) {
	resp, err := CreateSingleUUID(false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to Generate UUIDS",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Version": "v0.1",
		"uuid":    resp,
	})
}

func CreateSimpleHandler(c *gin.Context) {
	resp, err := CreateSingleUUID(true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to Generate UUIDS",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Version":     "v0.1",
		"simple_uuid": resp,
	})
}

func CreateMultiSimpleHandler(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid 'num' Parameter",
		})
		return
	}
	resp, err := utils.CreateMultiUUID(num, true)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to Generate UUIDS",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Version": "v0.1",
		"Data":    resp,
	})
}

func CreateMultiHandler(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid 'num' Parameter",
		})
		return
	}
	resp, err := utils.CreateMultiUUID(num, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Failed to Generate UUIDS",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Version": "v0.1",
		"Data":    resp,
	})
}
