package gologger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/shirocola/go-shop/pkg/utils"
)

type IGoLogger interface {
	Print() IGoLogger
	Save()
	SetQuery(c *fiber.Ctx)
	SetBody(c *fiber.Ctx)
	SetResponse(res any)
}

type goLogger struct {
	Time       string `json:"time"`
	Ip         string `json:"ip"`
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Path       string `json:"path"`
	Query      any    `json:"query"`
	Body       any    `json:"body"`
	Response   any    `json:"response"`
}

func InitGoLogger(c *fiber.Ctx, res any) IGoLogger {
	log := &goLogger{
		Time:       time.Now().Local().Format("2006-01-02 15:04:05"),
		Ip:         c.IP(),
		Method:     c.Method(),
		StatusCode: c.Response().StatusCode(),
		Path:       c.Path(),
	}
	log.SetQuery(c)
	log.SetBody(c)
	log.SetResponse(res)
	return log
}

func (l *goLogger) Print() IGoLogger {
	utils.Debug(l)
	return l
}

func (l *goLogger) Save() {
	data := utils.Output(l)

	filename := fmt.Sprintf("./assets/logs/gologger_%v.txt", strings.ReplaceAll(time.Now().Format("2006-01-02"), "-", ""))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("open file failed: %v", err)
	}
	defer file.Close()
	file.WriteString(string(data) + "\n")
}

func (l *goLogger) SetQuery(c *fiber.Ctx) {
	var body any
	if err := c.QueryParser(&body); err != nil {
		log.Printf("query parser failed: %v", err)
	}
	l.Query = body
}

func (l *goLogger) SetBody(c *fiber.Ctx) {
	var body any
	if err := c.BodyParser(&body); err != nil {
		log.Printf("body parser failed: %v", err)
	}

	switch l.Path {
	case "/v1/users/signup":
		l.Body = "hidden"
	default:
		l.Body = body
	}

}

func (l *goLogger) SetResponse(res any) {
	l.Response = res
}
