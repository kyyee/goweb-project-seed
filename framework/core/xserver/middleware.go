package xserver

import (
	"bytes"
	"fmt"
	"goweb-project-seed/framework/core/utils"
	"io"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func Debug(c *Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		begin := time.Now()
		var b bytes.Buffer
		ctx.Request.Body = io.NopCloser(io.TeeReader(ctx.Request.Body, &b))
		defer func() {
			cost := time.Since(begin)
			userAgent := ""
			if v, ok := ctx.Request.Header["User-Agent"]; ok && len(v) >= 0 {
				userAgent = v[0]
			}
			clientIp := ""
			if v, ok := ctx.Request.Header["X-Forwarded-For"]; ok && len(v) >= 0 {
				if addresses := strings.SplitN(v[0], ",", 2); len(addresses) > 0 {
					clientIp = addresses[0]
				}
			}
			if len(clientIp) <= 0 {
				clientIp = ctx.ClientIP()
			}
			headers := make([]string, 0)
			for k, v := range ctx.Request.Header {
				headers = append(headers, k+"->"+strings.Join(v, ","))
			}
			body := b.String()
			log.Printf(fmt.Sprintf("%s %s %s %s %s %s %s X-Forwarded-For: [%s], request.header: [%s], request.body: [%s]",
				utils.Green("[xgin request]"),
				utils.Red(fmt.Sprintf("[%vms]", float64(cost.Microseconds())/1000)),
				utils.Yellow(strings.Split(ctx.Request.RemoteAddr, ":")[0]),
				utils.Green(ctx.Request.Method),
				utils.Blue(ctx.Request.Host),
				utils.Blue(ctx.Request.RequestURI),
				utils.Green(userAgent),
				utils.Green(clientIp),
				utils.Green(strings.Join(headers, ";\t")),
				utils.Green(body),
			))
		}()
		ctx.Next()
	}
}
