package device

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/adcp-control-microservice/device/adcp"
	"github.com/byuoitav/common/status"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetPower .
func (d *DeviceManager) GetPower(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Getting power status", zap.String("address", address))
	resp, err := adcp.GetPower(address)
	if err != nil {
		d.Log.Error("failed to get power status", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetBlanked .
func (d *DeviceManager) GetBlanked(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Getting blanked status", zap.String("address", address))
	resp, err := adcp.GetBlanked(address)
	if err != nil {
		d.Log.Error("failed to get blanked status", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetInput .
func (d *DeviceManager) GetInput(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Getting input status", zap.String("address", address))
	resp, err := adcp.GetInput(address)
	if err != nil {
		d.Log.Error("failed to get input status", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetMuted .
func (d *DeviceManager) GetMuted(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Getting muted status", zap.String("address", address))
	resp, err := adcp.GetMuted(address)
	if err != nil {
		d.Log.Error("failed to get muted status", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetVolume .
func (d *DeviceManager) GetVolume(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Getting volume status", zap.String("address", address))
	resp, err := adcp.GetVolume(address)
	if err != nil {
		d.Log.Error("failed to get volume status", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetHardwareInfo .
func (d *DeviceManager) GetHardwareInfo(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Getting hardware info", zap.String("address", address))
	resp, err := adcp.GetHardwareInfo(address)
	if err != nil {
		d.Log.Error("failed to get hardware info", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetActiveSignal .
func (d *DeviceManager) GetActiveSignal(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Getting active signal", zap.String("address", address))
	resp, err := adcp.GetActiveSignal(address)
	if err != nil {
		d.Log.Error("failed to get active signal", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// SetPower .
func (d *DeviceManager) SetPower(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Setting power", zap.String("address", address), zap.String("state", context.Param("state")))
	state := status.Power{
		Power: context.Param("state"),
	}

	err := adcp.SetPower(address, state)
	if err != nil {
		d.Log.Error("failed to set power", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, state)
}

// SetBlanked .
func (d *DeviceManager) SetBlanked(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Setting blanked", zap.String("address", address), zap.String("state", context.Param("state")))
	state := status.Blanked{
		Blanked: context.Param("state") == "blank",
	}

	err := adcp.SetBlanked(address, state)
	if err != nil {
		d.Log.Error("failed to set blanked", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, state)
}

// SetInput .
func (d *DeviceManager) SetInput(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Setting input", zap.String("address", address), zap.String("input", context.Param("port")))
	state := status.Input{
		Input: context.Param("port"),
	}

	err := adcp.SetInput(address, state)
	if err != nil {
		d.Log.Error("failed to set input", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, state)
}

// SetMuted .
func (d *DeviceManager) SetMuted(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Setting muted", zap.String("address", address), zap.String("state", context.Param("state")))
	state := status.Mute{
		Muted: context.Param("state") == "mute",
	}

	err := adcp.SetMuted(address, state)
	if err != nil {
		d.Log.Error("failed to set muted", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, state)
}

// SetVolume .
func (d *DeviceManager) SetVolume(context *gin.Context) {
	address := context.Param("address")
	d.Log.Debug("Setting volume", zap.String("address", address), zap.String("level", context.Param("level")))
	vol, err := strconv.Atoi(context.Param("level"))
	if err != nil {
		d.Log.Error("failed to convert volume to int", zap.Error(err))
		context.String(http.StatusBadRequest, err.Error())
	}

	volume := status.Volume{
		Volume: vol,
	}

	err = adcp.SetVolume(address, volume)
	if err != nil {
		d.Log.Error("failed to set volume", zap.Error(err))
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, volume)
}
