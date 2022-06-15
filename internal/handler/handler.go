package handler

import (
	"config-server/internal/resource"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetResources(c *gin.Context) {
	namespace := c.Param("namespace")
	kind := c.Param("kind")
	limitStr := c.Query("limit")
	if limitStr == "" {
		limitStr = "1000"
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		errResponse(c, http.StatusBadRequest, err)
		return
	}

	rl := resource.ResourceList{}
	err = rl.Get(namespace, kind, limit)
	if err != nil {
		errResponse(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, rl)
}

func CreateResource(c *gin.Context) {
	namespace := c.Param("namespace")
	kind := c.Param("kind")

	r := resource.Resource{}
	if err := r.Create(namespace, kind, c.Request.Body); err != nil {
		errResponse(c, http.StatusBadRequest, err)
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
		errResponse(c, http.StatusBadRequest, err)
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
		errResponse(c, http.StatusBadRequest, err)
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
		errResponse(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
