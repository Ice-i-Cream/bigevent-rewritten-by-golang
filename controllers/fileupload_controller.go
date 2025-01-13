package controllers

import (
	"big_event/anno"
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type FileUploadController struct {
	web.Controller
}

func (c *FileUploadController) Upload() {
	exec := func(c *FileUploadController) (string, error) {
		f, h, err := c.GetFile("file")
		if err != nil {
			return "", nil
		}
		defer func(f multipart.File) {
			err = f.Close()
		}(f)
		if err != nil {
			return "", err
		}
		u := uuid.New()
		uuidStr := u.String()
		ext := strings.TrimPrefix(h.Filename[strings.LastIndex(h.Filename, "."):], ".")
		if ext == h.Filename {
			ext = ""
		} else {
			ext = "." + ext
		}
		fileName := uuidStr + ext
		var out *os.File
		out, err = os.Create("file/" + fileName)
		if err != nil {
			return "", err
		}
		defer func(out *os.File) {
			err = out.Close()
		}(out)
		if err != nil {
			return "", err
		}
		_, err = io.Copy(out, f)
		if err != nil {
			return "", err
		}
		var host string
		host, err = web.AppConfig.String("httphost")
		if err != nil {
			return "", err
		}
		var port int
		port, err = web.AppConfig.Int("httpport")
		if err != nil {
			return "", err
		}
		url := fmt.Sprintf("http://%s:%d/%s", host, port, "image/"+fileName)
		return url, nil
	}
	url, err := exec(c)
	anno.PostProcess((*struct{ web.Controller })(c), err, url)
}
