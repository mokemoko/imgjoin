package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"imgjoin/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
)

const (
	defaultHeight = 100
)

func RegisterJoin(c *gin.RouterGroup) {
	c.GET("/", indexHandler)
}

func indexHandler(c *gin.Context) {
	urls, height := parseRequest(c.Request)

	// get specified images
	var images []io.Reader
	for _, u := range urls {
		res, err := http.Get(u)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		images = append(images, res.Body)
	}

	// join images
	buf := bytes.Buffer{}
	if err := utils.Join(&buf, images, height); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// res image
	c.Data(http.StatusOK, "image/png", buf.Bytes())
}

func parseRequest(r *http.Request) ([]string, int) {
	uri, _ := url.Parse(r.RequestURI)
	q := uri.Query()
	height, err := strconv.Atoi(q.Get("height"))
	if err != nil {
		height = defaultHeight
	}
	return q["target"], height
}