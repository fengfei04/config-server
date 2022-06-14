package handler

import (
	"config-server/internal/resource"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetResources(c *gin.Context) {
	// namespace := c.Param("namespace")
	// kind := c.Param("kind")

	c.JSON(http.StatusOK, gin.H{})
}

func CreateResource(c *gin.Context) {
	namespace := c.Param("namespace")
	kind := c.Param("kind")

	r := resource.Resource{}
	if err := r.Create(namespace, kind, c.Request.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func GetResource(c *gin.Context) {
	namespace := c.Param("namespace")
	kind := c.Param("kind")
	name := c.Param("name")

	r := resource.Resource{}
	if err := r.Get(namespace, kind, name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func UpdateResource(c *gin.Context) {
	namespace := c.Param("namespace")
	kind := c.Param("kind")
	name := c.Param("name")

	r := resource.Resource{}
	if err := r.Update(namespace, kind, name, c.Request.Body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, r)
}

func DeleteResource(c *gin.Context) {
	namespace := c.Param("namespace")
	kind := c.Param("kind")
	name := c.Param("name")

	r := resource.Resource{}
	if err := r.Delete(namespace, kind, name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
