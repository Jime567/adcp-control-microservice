package device

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DeviceManager struct {
	Log *zap.Logger
}

func (d *DeviceManager) GetLogger() *zap.Logger {
	return d.Log
}

func (d *DeviceManager) RunHTTPServer(router *gin.Engine, port string) error {
	d.Log.Info("registering http endpoints")

	// status endpoints
	route := router.Group("")
	route.GET("/:address/power", d.GetPower)
	route.GET("/:address/blanked", d.GetBlanked)
	route.GET("/:address/input", d.GetInput)
	route.GET("/:address/muted", d.GetMuted)
	route.GET("/:address/volume", d.GetVolume)
	route.GET("/:address/hardware", d.GetHardwareInfo)
	route.GET("/:address/activesignal", d.GetActiveSignal)

	// action endpoints
	route.GET("/:address/power/:state", d.SetPower)
	route.GET("/:address/blanked/:state", d.SetBlanked)
	route.GET("/:address/input/:port", d.SetInput)
	route.GET("/:address/muted/:state", d.SetMuted)
	route.GET("/:address/volume/:level", d.SetVolume)

	server := &http.Server{
		Addr:           port,
		MaxHeaderBytes: 1021 * 10,
	}

	d.Log.Info("running http server", zap.String("port", port))
	err := router.Run(server.Addr)

	d.Log.Error("http server stopped", zap.Error(err))

	return fmt.Errorf("http server stopped")
}
