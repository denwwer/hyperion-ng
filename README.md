![Static Badge](https://img.shields.io/badge/go-1.23-027c9c)
![GitHub last commit](https://img.shields.io/github/last-commit/denwwer/hyperion-ng)
[![Go Reference](https://pkg.go.dev/badge/github.com/denwwer/hyperion-ng.svg)](https://pkg.go.dev/github.com/denwwer/hyperion-ng)

**hyperion-ng** is library that helps communicate with [Hyperion-NG](https://hyperion-project.org/) JSON API.

## Install
```
go get -u github.com/denwwer/hyperion-ng
```

## Examples
```go
import (
    "log"
    "github.com/denwwer/hyperion-ng"
    "github.com/denwwer/hyperion-ng/model"
)

func main() {
    // Create configuration
    conf := hyperion.Config{
    VerboseLog: false,
    Connection: hyperion.Connection{
        Token:   "6c224a4c-6ebf-491a-9d70-fb7681ca2a59",
        Type:    hyperion.ConnectHTTP,
        Host:    "192.168.53.130",
        Port:    8090,
        SSL:     false,
        Timeout: 10,
        },
    }

    // Create client with custom header options
    cl := hyperion.NewClient(conf, hyperion.WithHeader(map[string]string{"my-header": "value"}))
    
    // Get full information
    info, err := cl.ServerInfo()
    if err != nil {
        log.Panic(err)
    }
    
    // Manage Instances
    for _, instance := range info.Instances {
        log.Printf(`name: "%s" active: %t`, instance.Name, instance.Running)
    
        err = cl.Instance(instance.Instance, model.InstanceCmdStop)
        if err != nil {
            log.Fatalln(err)
        }
    }

    // Manage Components
    for _, comp := range info.Components {
        if comp.Switchable() {
            err = cl.ComponentState(comp.Name, false)
            if err != nil {
                log.Fatalln(err)
            }
        }
    }
    
    // List user defined effects
    for _, effect := range info.Effects.Users() {
        log.Printf(`name: "%s"`, effect.Name)
    }
    
    // Change video mode
    err = cl.VideoMode(model.VideoMode2D)
    if err != nil {
        log.Fatalln(err)
    }
}
```

Additional doumentation on [pkg.go.dev](https://pkg.go.dev/github.com/denwwer/hyperion-ng)
