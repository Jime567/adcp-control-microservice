package device

import (
	"net/http"
	"strconv"

	"github.com/byuoitav/adcp-control-microservice/device/adcp"
	"github.com/byuoitav/common/status"
	"github.com/gin-gonic/gin"
)

// GetPower .
func (d *DeviceManager) GetPower(context *gin.Context) {
	address := context.Param("address")

	resp, err := adcp.GetPower(address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetBlanked .
func (d *DeviceManager) GetBlanked(context *gin.Context) {
	address := context.Param("address")

	resp, err := adcp.GetBlanked(address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetInput .
func (d *DeviceManager) GetInput(context *gin.Context) {
	address := context.Param("address")

	resp, err := adcp.GetInput(address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetMuted .
func (d *DeviceManager) GetMuted(context *gin.Context) {
	address := context.Param("address")

	resp, err := adcp.GetMuted(address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetVolume .
func (d *DeviceManager) GetVolume(context *gin.Context) {
	address := context.Param("address")

	resp, err := adcp.GetVolume(address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetHardwareInfo .
func (d *DeviceManager) GetHardwareInfo(context *gin.Context) {
	address := context.Param("address")

	resp, err := adcp.GetHardwareInfo(address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// GetActiveSignal .
func (d *DeviceManager) GetActiveSignal(context *gin.Context) {
	address := context.Param("address")

	resp, err := adcp.GetActiveSignal(address)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, resp)
}

// SetPower .
func (d *DeviceManager) SetPower(context *gin.Context) {
	address := context.Param("address")
	state := status.Power{
		Power: context.Param("state"),
	}

	err := adcp.SetPower(address, state)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, state)
}

// SetBlanked .
func (d *DeviceManager) SetBlanked(context *gin.Context) {
	address := context.Param("address")
	state := status.Blanked{
		Blanked: context.Param("state") == "blank",
	}

	err := adcp.SetBlanked(address, state)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, state)
}

// SetInput .
func (d *DeviceManager) SetInput(context *gin.Context) {
	address := context.Param("address")
	state := status.Input{
		Input: context.Param("port"),
	}

	err := adcp.SetInput(address, state)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, state)
}

// SetMuted .
func (d *DeviceManager) SetMuted(context *gin.Context) {
	address := context.Param("address")
	state := status.Mute{
		Muted: context.Param("state") == "mute",
	}

	err := adcp.SetMuted(address, state)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, state)
}

// SetVolume .
func (d *DeviceManager) SetVolume(context *gin.Context) {
	address := context.Param("address")

	vol, err := strconv.Atoi(context.Param("level"))
	if err != nil {
		context.String(http.StatusBadRequest, err.Error())
	}

	volume := status.Volume{
		Volume: vol,
	}

	err = adcp.SetVolume(address, volume)
	if err != nil {
		context.JSON(http.StatusInternalServerError, status.Error{
			Error: err.Error(),
		})
	}

	context.JSON(http.StatusOK, volume)
}
