package middleware

import (
	"crypto/tls"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"golang.org/x/crypto/acme/autocert"
	"net"
	"time"
)

func UseCsrf(app *fiber.App) {
	app.Use(csrf.New(csrf.Config{
		KeyLookup:      "header:X-Csrf-Token-Algo-Api",
		CookieName:     "csrf_token",
		CookieSameSite: "Strict",
		Expiration:     7 * time.Hour,
		KeyGenerator:   utils.UUID,
	}))
}

func AutoCert() net.Listener {
	certManager := &autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("api.2084team.com"),
		Cache:      autocert.DirCache("./cert"),
	}
	certConfig := &tls.Config{
		GetCertificate: certManager.GetCertificate,
		NextProtos: []string{
			"http/1.1", "acme-tls/1",
		},
	}
	listenPort, err := tls.Listen("tcp", ":443", certConfig)
	if err != nil {
		panic(err)
	}
	return listenPort
}
