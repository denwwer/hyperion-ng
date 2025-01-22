package model

// System information about Hyperion server.
type System struct {
	Hyperion struct {
		Build        string `json:"build"`
		GitRemote    string `json:"gitremote"`
		ID           string `json:"id"`
		IsGuiMode    bool   `json:"isGuiMode"`
		ReadOnlyMode bool   `json:"readOnlyMode"`
		RootPath     string `json:"rootPath"`
		Time         string `json:"time"`
		Version      string `json:"version"`
	} `json:"hyperion"`
	System struct {
		Architecture   string `json:"architecture"`
		CPUHardware    string `json:"cpuHardware"`
		CPUModelName   string `json:"cpuModelName"`
		CPUModelType   string `json:"cpuModelType"`
		CPURevision    string `json:"cpuRevision"`
		DomainName     string `json:"domainName"`
		HostName       string `json:"hostName"`
		IsUserAdmin    bool   `json:"isUserAdmin"`
		KernelType     string `json:"kernelType"`
		KernelVersion  string `json:"kernelVersion"`
		PrettyName     string `json:"prettyName"`
		ProductType    string `json:"productType"`
		ProductVersion string `json:"productVersion"`
		PyVersion      string `json:"pyVersion"`
		QtVersion      string `json:"qtVersion"`
		WordSize       string `json:"wordSize"`
	} `json:"system"`
}
