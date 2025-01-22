package hyperion

import (
	m "github.com/hyperion-ng/internal/model"
	"github.com/hyperion-ng/model"
)

const cmdSysInfo = "sysinfo"

// SystemInfo retrieve basic system information about Hyperion server.
func (c *Client) SystemInfo() (*model.System, error) {
	tan := 1
	req := m.Request{Command: cmdSysInfo, Tan: &tan}
	resp := &model.System{}
	return resp, c.send(req, &resp)
}
