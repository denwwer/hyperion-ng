package model

// LEDMode for the incoming image.
type LEDMode string

// List of LEDMode's.
const (
	LEDModeMulticolor LEDMode = "multicolor_mean"         // Simple per LED
	LEDModeUnicolor   LEDMode = "unicolor_mean"           //  Applied on all LEDs
	LEDModeSquared    LEDMode = "multicolor_mean_squared" //  Squared per LED
	LEDModeDominant   LEDMode = "dominant_color"          //  Dominant per LED
	LEDModeAdvanced   LEDMode = "dominant_color_advanced" //  Dominant advanced per LED
)

// VideoMode switch.
type VideoMode string

// List of VideoMode's.
const (
	VideoMode2D  VideoMode = "2D"
	VideoMode3DS VideoMode = "3DSBS"
	VideoMode3DT VideoMode = "3DTAB"
)

// InstanceCmd
type InstanceCmd string

// List of InstanceCmd's.
const (
	InstanceCmdStart  InstanceCmd = "startInstance"
	InstanceCmdStop   InstanceCmd = "stopInstance"
	InstanceCmdSwitch InstanceCmd = "switchTo"
)

type Image struct {
	ImageB64 string  `json:"imagedata"` // Data of image as Base64
	Format   *string `json:"format"`    // Default is "auto"
	Name     string  `json:"name"`      // The name of the image
}
