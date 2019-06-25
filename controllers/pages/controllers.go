package pages

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func Main(c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{})
}