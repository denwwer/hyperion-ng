package model

import (
	"slices"
	"strings"
)

var componentsNonSwitchable = []string{"color", "effect", "image", "flatbufserver", "protoserver"}

// Information response provides data about the live state of Hyperion.
type Information struct {
	ActiveEffects         []ActiveEffect           `json:"activeEffects"`
	ActiveLedColor        []map[string]interface{} `json:"activeLedColor"`
	Components            []Component              `json:"components"`
	Adjustments           []Adjustment             `json:"adjustment"`
	Effects               Effects                  `json:"effects"`
	ImageToLedMappingType string                   `json:"imageToLedMappingType"` // More info at https://docs.hyperion-project.org/json/Control.html#led-mapping
	VideoMode             string                   `json:"videomode"`             // More info at https://docs.hyperion-project.org/json/Control.html#video-mode
	Priorities            []Priority               `json:"priorities"`
	PrioritiesAutoselect  bool                     `json:"priorities_autoselect"`
	Instances             Instances                `json:"instance"`

	Grabbers struct {
		Audio  Grabber `json:"audio"`
		Screen Grabber `json:"screen"`
		Video  Grabber `json:"video"`
	} `json:"grabbers"`

	LedDevices struct {
		Available []string `json:"available"`
	} `json:"ledDevices"`

	Leds     []Led    `json:"leds"`
	Services []string `json:"services"`
}

// Effects list of Effect's.
type Effects []Effect

// Users returns user created effects.
func (e Effects) Users() []Effect {
	effects := []Effect{}

	for _, effect := range e {
		if strings.HasPrefix(effect.File, "/") {
			effects = append(effects, effect)
		}
	}

	return effects
}

// System returns system provided effects.
func (e Effects) System() []Effect {
	effects := []Effect{}

	for _, effect := range e {
		if strings.HasPrefix(effect.File, ":") {
			effects = append(effects, effect)
		}
	}

	return effects
}

// Component of Hyperion.
type Component struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name"`
}

// Switchable determine that component can be enabled/disabled.
func (c Component) Switchable() bool {
	return !slices.Contains(componentsNonSwitchable, strings.ToLower(c.Name))
}

// Adjustment values of color.
type Adjustment struct {
	BacklightColored       *bool    `json:"backlightColored,omitempty"`
	BacklightThreshold     *int     `json:"backlightThreshold,omitempty"`
	Brightness             *int     `json:"brightness,omitempty"`
	BrightnessCompensation *int     `json:"brightnessCompensation,omitempty"`
	BrightnessGain         *float64 `json:"brightnessGain,omitempty"`
	Blue                   []int    `json:"blue,omitempty"`
	Cyan                   []int    `json:"cyan,omitempty"`
	GammaBlue              *float64 `json:"gammaBlue,omitempty"`
	GammaGreen             *float64 `json:"gammaGreen,omitempty"`
	GammaRed               *float64 `json:"gammaRed,omitempty"`
	Green                  []int    `json:"green,omitempty"`
	ID                     string   `json:"id,omitempty"`
	Magenta                []int    `json:"magenta,omitempty"`
	Red                    []int    `json:"red,omitempty"`
	SaturationGain         *float64 `json:"saturationGain,omitempty"`
	White                  []int    `json:"white,omitempty"`
	Yellow                 []int    `json:"yellow,omitempty"`
}

// Effect object of named effect.
type Effect struct {
	Args   map[string]interface{} `json:"args,omitempty"`   // Optional object with additional properties
	File   string                 `json:"file,omitempty"`   // Optional
	Name   string                 `json:"name"`             // Required
	Script string                 `json:"script,omitempty"` // Optional
}

// ActiveEffect active effect info.
type ActiveEffect struct {
	Script   string                 `json:"script"`
	Name     string                 `json:"name"`
	Priority int                    `json:"priority"`
	Timeout  int                    `json:"timeout"`
	Args     map[string]interface{} `json:"args"`
}

// Priority info of the registered/active sources.
type Priority struct {
	Active      bool   `json:"active"`
	Visible     bool   `json:"visible"`
	ComponentID string `json:"componentId"`
	Origin      string `json:"origin"`
	Owner       string `json:"owner"`
	Priority    int    `json:"priority"`
	Value       struct {
		HSL []float64 `json:"HSL"`
		RGB []int     `json:"RGB"`
	} `json:"value"`
	Duration int `json:"duration_ms"`
}

// Instances list of Instance's.
type Instances []Instance

// Find instance by id.
func (i Instances) Find(instance int) *Instance {
	for _, ins := range i {
		if ins.Instance == instance {
			return &ins
		}
	}
	return nil
}

// Instance information and their state.
type Instance struct {
	Instance int    `json:"instance"`
	Running  bool   `json:"running"`
	Name     string `json:"friendly_name"`
}

// Led layout information.
type Led struct {
	HMin float64 `json:"hmin"`
	HMax float64 `json:"hmax"`
	VMin float64 `json:"vmin"`
	VMax float64 `json:"vmax"`
}

// Grabber state information.
type Grabber struct {
	Active    []string `json:"active"`
	Available []string `json:"available"`
}
