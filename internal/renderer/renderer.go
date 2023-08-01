package renderer

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	FormatQueryParam = "format"
	XMLFormat        = "xml"
	YAMLFormat       = "yaml"
)

type Render struct{}

func NewRender() *Render {
	return &Render{}
}

func (that *Render) SendSuccess(c *gin.Context, data interface{}) {
	format := c.Query(FormatQueryParam)
	switch format {
	case XMLFormat:
		c.XML(http.StatusOK, data)
	case YAMLFormat:
		c.YAML(http.StatusOK, data)
	default:
		c.IndentedJSON(http.StatusOK, data)
	}
}

func (that *Render) SendError(c *gin.Context, statusCode int, errMsg string) {
	format := c.Query(FormatQueryParam)
	switch format {
	case XMLFormat:
		c.XML(statusCode, gin.H{"error": errMsg})
	case YAMLFormat:
		c.YAML(statusCode, gin.H{"error": errMsg})
	default:
		c.IndentedJSON(statusCode, gin.H{"error": errMsg})
	}
}
