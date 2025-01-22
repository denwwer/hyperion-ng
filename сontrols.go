package hyperion

import (
	"errors"

	m "github.com/denwwer/hyperion-ng/internal/model"
	"github.com/denwwer/hyperion-ng/model"
)

const (
	cmdColor          = "color"
	cmdEffect         = "effect"
	cmdImage          = "image"
	cmdClear          = "clear"
	cmdSourceselect   = "sourceselect"
	cmdAdjustment     = "adjustment"
	cmdProcessing     = "processing"
	cmdVideomode      = "videomode"
	cmdComponentstate = "componentstate"
	cmdInstance       = "instance"
)

// SetColor for all LEDs.
func (c Client) SetColor(color []int, priority int, origin string, duration *int) error {
	// [R, G, B] or [R, G, B, R, G, B ...]
	if len(color) < 2 {
		return errors.New(m.ColorRequired)
	}

	if err := validate(priority, origin, duration); err != nil {
		return err
	}

	req := struct {
		m.Request
		Color []int `json:"color"`
	}{
		Request: m.Request{
			Command:  cmdColor,
			Priority: &priority,
			Origin:   origin,
			Duration: duration,
		},
		Color: color,
	}

	return c.send(req, nil)
}

// SetEffect by name with optional overridden arguments.
func (c Client) SetEffect(effect model.Effect, priority int, origin string, duration *int) error {
	if err := validate(priority, origin, duration); err != nil {
		return err
	}

	req := struct {
		m.Request
		Effect model.Effect `json:"effect"`
	}{
		Request: m.Request{
			Command:  cmdEffect,
			Priority: &priority,
			Origin:   origin,
			Duration: duration,
		},
		Effect: effect,
	}

	return c.send(req, nil)
}

// SetImage a single image.
func (c Client) SetImage(image model.Image, priority int, origin string, duration *int) error {
	if err := validate(priority, origin, duration); err != nil {
		return err
	}

	req := struct {
		m.Request
		model.Image
	}{
		Request: m.Request{
			Command:  cmdImage,
			Priority: &priority,
			Origin:   origin,
			Duration: duration,
		},
		Image: image,
	}

	if req.Format == nil {
		f := "auto"
		req.Format = &f
	}

	return c.send(req, nil)
}

// ClearPriority used to revert SetColor, SetEffect or SetImage.
func (c Client) ClearPriority(priority int) error {
	req := struct {
		m.Request
	}{
		Request: m.Request{
			Command:  cmdClear,
			Priority: &priority,
		},
	}

	return c.send(req, nil)
}

// SetSource priority manually.
func (c Client) SetSource(priority int) error {
	req := struct {
		m.Request
	}{
		Request: m.Request{
			Command:  cmdSourceselect,
			Priority: &priority,
		},
	}

	return c.send(req, nil)
}

// SetSourceAuto visible source is determined by priority.
func (c Client) SetSourceAuto() error {
	req := struct {
		m.Request
		Auto bool `json:"auto"`
	}{
		Request: m.Request{Command: cmdSourceselect},
		Auto:    true,
	}

	return c.send(req, nil)
}

// SetAdjustment to color calibration.
func (c Client) SetAdjustment(adj model.Adjustment) error {
	req := struct {
		m.Request
		Adjustment model.Adjustment `json:"adjustment"`
	}{
		Request:    m.Request{Command: cmdAdjustment},
		Adjustment: adj,
	}

	return c.send(req, nil)
}

// LEDMode switched the LED mapping mode for the incoming image.
func (c Client) LEDMode(mode model.LEDMode) error {
	req := struct {
		m.Request
		Type model.LEDMode `json:"mappingType"`
	}{
		Request: m.Request{Command: cmdProcessing},
		Type:    mode,
	}

	return c.send(req, nil)
}

// VideoMode switching.
func (c Client) VideoMode(mode model.VideoMode) error {
	req := struct {
		m.Request
		Mode model.VideoMode `json:"videoMode"`
	}{
		Request: m.Request{Command: cmdVideomode},
		Mode:    mode,
	}

	return c.send(req, nil)
}

// ComponentState enabled or disabled at runtime.
func (c Client) ComponentState(name string, enable bool) error {
	req := struct {
		m.Request
		Component map[string]interface{} `json:"componentstate"`
	}{
		Request: m.Request{
			Command: cmdComponentstate,
		},
		Component: map[string]interface{}{"component": name, "state": enable},
	}

	return c.send(req, nil)
}

// Instance controlling.
func (c Client) Instance(instance int, command model.InstanceCmd) error {
	req := struct {
		m.Request
		Instance int `json:"instance"`
	}{
		Request: m.Request{
			Command:    cmdInstance,
			Subcommand: string(command),
		},
		Instance: instance,
	}

	return c.send(req, nil)
}

func validate(priority int, origin string, duration *int) error {
	if priority < 1 {
		return errors.New(m.PriorityRequired)
	}

	if len(origin) < 3 {
		return errors.New(m.OriginRequired)
	}

	if duration != nil && *duration < 0 {
		return errors.New(m.DurationRequired)
	}

	return nil
}
