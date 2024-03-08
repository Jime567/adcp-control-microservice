# adcp-control
A microservice for controlling Sony projectors using the ADCP protocol.

## Endpoints
### Status
* `/ping` - Check if the microservice is running
* `/status` - Returns good if microservice is running
* `/:address/power` - Get the on/off status of the projector
* `/:address/blanked` - Get the blanked status of the projector
* `/:address/input ` - Get the current input of the projector
* `/:address/muted` - Get the mute status of the projector
* `/:address/volume` - Get the volume of the projector
* `/:address/hardware` - Get the hardware information of the projector
* `/:address/activesignal` - Get the active signal of the projector

### Actions
* `/:address/power/:state` - Turn the projector on or off :new_moon:/:full_moon:
    * `/00.00.00.000/power/on`
    * `/00.00.00.000/power/off`
* `/:address/blanked/:state` - Blank or unblank the projector
    * `/00.00.00.000/blanked/blank`
    * `/00.00.00.000/blanked/unblank`
* `/:address/input/:input` - Set the input of the projector :electric_plug:
    * `/00.00.00.000/input/hdmi1`
    * `/00.00.00.000/input/hdbaset1`
* `/:address/muted/:state` - Mute or unmute the projector :mute:/:speaker:
    * `/:address/muted/mute`
    * `/:address/muted/unmute`
* `/:address/volume/:volume` - Set the volume of the projector :sound:/:loud_sound:
    * `/:address/volume/50`

## Flags
* `-port`, `-p` - The port to run the microservice on. Defaults to 8012
    * `go run cmd/main.go cmd/deps.go -port 8007`

* `-log`, `-l` - The log level to run the microservice at. Defaults to info
    * `go run cmd/main.go cmd/deps.go -l debug`

## Setup
Does not require any additional environment variables.