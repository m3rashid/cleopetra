package app

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/m3rashid/cleopatra/config"
)

var Http *config.AppConfig

var Version = "develop"

func Load(configFile string) {
	Http = &config.AppConfig{ConfigFile: configFile}
	Http.Setup()
	LoadBuiltInMiddlewares(Http)
}

func LoadAdditionalServices() {
	Http.PayPal.Connect(Http.Server.Env)
}

func LoadBuiltInMiddlewares(app *config.AppConfig) {
	app.Server.Use(recover.New())
	app.Server.Use(etag.New())
	app.Server.Use(compress.New(compress.Config{
		Level: 1,
	}))
	if app.Server.Debug {
		app.Server.Use(pprof.New())
	}
}

func Location(c *fiber.Ctx) (string, error) {
	ip := IP(c)
	_, err := Http.GeoIP.GetLocation(ip)
	if err != nil {
		return "127.0.0.1", err
	}
	return "127.0.0.1", err
}

var fetchIpFromString = regexp.MustCompile(`(\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3})`)
var possibleHeaders = []string{
	"X-Original-Forwarded-For",
	"X-Forwarded-For",
	"X-Real-Ip",
	"X-Client-Ip",
	"Forwarded-For",
	"Forwarded",
	"Remote-Addr",
	"Client-Ip",
	"CF-Connecting-IP",
}

// determine user ip
func IP(c *fiber.Ctx) string {
	headerValue := []byte{}
	if Http.Server.Config().ProxyHeader == "*" {
		for _, headerName := range possibleHeaders {
			headerValue = c.Request().Header.Peek(headerName)
			if len(headerValue) > 3 {
				return string(fetchIpFromString.Find(headerValue))
			}
		}
	}
	headerValue = []byte(c.IP())
	if len(headerValue) <= 3 {
		headerValue = []byte("0.0.0.0")
	}

	// find ip address in string
	return string(fetchIpFromString.Find(headerValue))
}
