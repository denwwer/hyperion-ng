package hyperion

import (
	m "github.com/hyperion-ng/internal/model"
	"github.com/hyperion-ng/model"
)

const cmdServerInfo = "serverinfo"

// ServerInfo retrieve live state of Hyperion.
func (c *Client) ServerInfo() (*model.Information, error) {
	tan := 1
	req := m.Request{Command: cmdServerInfo, Tan: &tan}
	resp := &model.Information{}
	return resp, c.send(req, &resp)
}
