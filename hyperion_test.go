package hyperion

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/denwwer/hyperion-ng/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const imageB64 = `iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAABGdBTUEAALGPC/xhBQAAACBjSFJNAAB6JgAAgIQAAPoAAACA6AAAdTAAAOpgAAA6mAAAF3CculE8AAAAUGVYSWZNTQAqAAAACAACARIAAwAAAAEAAQAAh2kABAAAAAEAAAAmAAAAAAADoAEAAwAAAAEAAQAAoAIABAAAAAEAAAABoAMABAAAAAEAAAABAAAAAOv/s+AAAAIwaVRYdFhNTDpjb20uYWRvYmUueG1wAAAAAAA8eDp4bXBtZXRhIHhtbG5zOng9ImFkb2JlOm5zOm1ldGEvIiB4OnhtcHRrPSJYTVAgQ29yZSA2LjAuMCI+CiAgIDxyZGY6UkRGIHhtbG5zOnJkZj0iaHR0cDovL3d3dy53My5vcmcvMTk5OS8wMi8yMi1yZGYtc3ludGF4LW5zIyI+CiAgICAgIDxyZGY6RGVzY3JpcHRpb24gcmRmOmFib3V0PSIiCiAgICAgICAgICAgIHhtbG5zOmV4aWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20vZXhpZi8xLjAvIgogICAgICAgICAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29tL3RpZmYvMS4wLyI+CiAgICAgICAgIDxleGlmOlBpeGVsWURpbWVuc2lvbj4xMTwvZXhpZjpQaXhlbFlEaW1lbnNpb24+CiAgICAgICAgIDxleGlmOlBpeGVsWERpbWVuc2lvbj4xMDwvZXhpZjpQaXhlbFhEaW1lbnNpb24+CiAgICAgICAgIDxleGlmOkNvbG9yU3BhY2U+MTwvZXhpZjpDb2xvclNwYWNlPgogICAgICAgICA8dGlmZjpPcmllbnRhdGlvbj4xPC90aWZmOk9yaWVudGF0aW9uPgogICAgICA8L3JkZjpEZXNjcmlwdGlvbj4KICAgPC9yZGY6UkRGPgo8L3g6eG1wbWV0YT4KegZxpgAAAAxJREFUCB1juKFtAwADHgFAgWYtFwAAAABJRU5ErkJggg==`

func TestMain(m *testing.M) {
	s := testServer()
	defer s.Close()
	code := m.Run()
	os.Exit(code)
}

func TestServerInfo(t *testing.T) {
	t.Parallel()
	c := testClient()

	resp, err := c.ServerInfo()
	require.Nil(t, err)
	assert.NotEmpty(t, resp.Effects)
	assert.NotEmpty(t, resp.Components)
	assert.True(t, resp.Components[0].Switchable())
	assert.NotEmpty(t, resp.Adjustments)
	assert.NotEmpty(t, resp.Effects[0].File)
	assert.NotEmpty(t, resp.Effects[0].Name)
	assert.NotEmpty(t, resp.Effects[0].Args)
	assert.Empty(t, resp.Effects.Users())
	assert.NotEmpty(t, resp.Effects.System())
	assert.NotEmpty(t, resp.LedDevices.Available)
	assert.Equal(t, "First LED Hardware instance", resp.Instances[0].Name)
}

func TestSystemInfo(t *testing.T) {
	t.Parallel()
	c := testClient()

	resp, err := c.SystemInfo()
	require.Nil(t, err)
	assert.NotEmpty(t, resp.System.Architecture)
	assert.Equal(t, "linux", resp.System.KernelType)
	assert.NotEmpty(t, resp.Hyperion.Version)
}

func TestSetColor(t *testing.T) {
	t.Parallel()
	c := testClient()

	err := c.SetColor([]int{0, 0, 0}, 20, "test 1", nil)
	require.Nil(t, err)
}

func TestSetEffect(t *testing.T) {
	t.Parallel()
	c := testClient()

	duration := 50
	err := c.SetEffect(model.Effect{
		Args: map[string]interface{}{"color-end": []int{200, 200, 200}},
		Name: "Blue mood blobs",
	}, 20, "test 1", &duration)
	require.Nil(t, err)
}

func TestSetImage(t *testing.T) {
	t.Parallel()
	c := testClient()

	err := c.SetImage(model.Image{
		ImageB64: imageB64,
		Name:     "New image",
	}, 20, "test 1", nil)
	require.Nil(t, err)
}

func TestClearPriority(t *testing.T) {
	t.Parallel()
	c := testClient()

	err := c.ClearPriority(10)
	require.Nil(t, err)
}

func TestSetSource(t *testing.T) {
	t.Parallel()
	c := testClient()

	testCases := []struct {
		Name     string
		Priority int
		Expect   func(err error)
	}{
		{
			Name:     "success",
			Priority: 20,
			Expect: func(err error) {
				assert.Nil(t, err)
			},
		},
		{
			Name:     "error",
			Priority: -1,
			Expect: func(err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(t *testing.T) {
			err := c.SetSource(tc.Priority)
			tc.Expect(err)
		})
	}
}

func TestSetSourceAuto(t *testing.T) {
	t.Parallel()
	c := testClient()

	err := c.SetSourceAuto()
	require.Nil(t, err)
}

func TestSetAdjustment(t *testing.T) {
	t.Parallel()
	c := testClient()

	b := true
	err := c.SetAdjustment(model.Adjustment{
		BacklightColored: &b,
		Green:            []int{0, 236, 0},
	})
	require.Nil(t, err)
}

func TestLEDMap(t *testing.T) {
	t.Parallel()
	c := testClient()

	err := c.LEDMode(model.LEDModeAdvanced)
	require.Nil(t, err)
}

func TestVideoMode(t *testing.T) {
	t.Parallel()
	c := testClient()

	err := c.VideoMode(model.VideoMode3DS)
	require.Nil(t, err)
}

func TestComponentState(t *testing.T) {
	t.Parallel()
	c := testClient()

	err := c.ComponentState("LEDDEVICE", true)
	require.Nil(t, err)
}

func TestInstance(t *testing.T) {
	t.Parallel()
	c := testClient()

	err := c.Instance(0, model.InstanceCmdSwitch)
	require.Nil(t, err)
}

var testURL string

func testServer() *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := map[string]interface{}{}

		d := json.NewDecoder(r.Body)
		if err := d.Decode(&req); err != nil {
			log.Fatalln(err)
		}

		cmd := req["command"].(string)
		if cmd == "sourceselect" && req["priority"] != nil && req["priority"].(float64) == -1 {
			cmd = "sourceselecterror" // returns error for sourceselect
		}

		fb, err := os.ReadFile(fmt.Sprintf("testdata/%s.json", cmd))
		if err != nil {
			if os.IsNotExist(err) {
				fb, _ = os.ReadFile("testdata/success.json")
			} else {
				log.Fatalln(err)
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(fb)
	}))

	testURL = s.URL
	return s
}

func testClient() *Client {
	u, _ := url.Parse(testURL)
	host := u.Hostname()
	port, _ := strconv.Atoi(u.Port())

	// Hyperion
	// host = "192.168.53.130"
	// port = 8090

	return NewClient(Config{
		// VerboseLog: true,
		Connection: Connection{
			Type:  ConnectHTTP,
			Host:  host,
			Port:  port,
			Token: "6c224a4c-6ebf-491a-9d70-fb7681ca2a59",
		},
	})
}
