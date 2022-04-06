package main

import (
    "os"
    "flag"
    "fmt"

    "golang.org/x/net/context"
    "github.com/golang/glog"
    "github.com/kubevirt/device-plugin-manager/pkg/dpm"
    pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// FLXLister is the object responsible for discovering initial pool of devices and their allocation
type FLXLister struct {}

type message struct {}

type Plugin struct{ 
  counter int
  devs []*pluginapi.Device
  update chan message
}

// Set up resources if needed, initialize custom channels etc
//func (p *Plugin) Start() error {
//    return nil
//}

// Tear down resources if needed
//func (p *Plugin) Stop() error {
//    return nil
//}

// Monitors available resource's devices and notifies Kubernetes
func (p *Plugin) ListAndWatch(e *pluginapi.Empty, s pluginapi.DevicePlugin_ListAndWatchServer) error {
    fmt.Print("ListAndWatch()\n")
    return nil
}

// Allocates a device requested by one of Pods
func (p *Plugin) Allocate(ctx context.Context, r *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
    fmt.Print("Allocate\n")
    return nil, nil
}

// GetDevicePluginOptions returns options to be communicated with Device Manager
func (Plugin) GetDevicePluginOptions(context.Context, *pluginapi.Empty) (*pluginapi.DevicePluginOptions, error) {
    return nil, nil
}

// PreStartContainer is called, if indicated by Device Plugin during registeration phase,
// before each container start. Device plugin can run device specific operations
// such as reseting the device before making devices available to the container
func (Plugin) PreStartContainer(context.Context, *pluginapi.PreStartContainerRequest) (*pluginapi.PreStartContainerResponse, error) {
    return nil, nil
}

func (dp *Plugin) GetPreferredAllocation(ctx context.Context, request *pluginapi.PreferredAllocationRequest) (*pluginapi.PreferredAllocationResponse, error) {
    return nil, nil
}

//type Lister struct{ 
//    Plugins []string
//}

func (l FLXLister) GetResourceNamespace() string {
    glog.V(3).Infof("GetResourceNamespace()")
    fmt.Print("GetResourceNamespace()\n")
    return "flx.cern"
}

// Discovery discovers all devices within the system. Monitors available resources.
func (l FLXLister) Discover(pluginListCh chan dpm.PluginNameList) {
    glog.V(3).Infof("Discover()")
    fmt.Print("Discover()\n")
    var plugins = make(dpm.PluginNameList, 0)

    var FLXPath = "/dev/flx0"

    if _, err := os.Stat(FLXPath); err == nil {
        glog.V(3).Infof("Discovered %s", FLXPath)
        fmt.Print("Discovered ", FLXPath, "\n")
	plugins = append(plugins, "FLX712_0")
    }

    pluginListCh <- plugins
}

func (l FLXLister) NewPlugin(name string) dpm.PluginInterface {
    glog.V(3).Infof("NewPlugin()")
    fmt.Print("NewPlugin() ", name, "\n")
    return &Plugin{
        counter: 0,
        devs: make([]*pluginapi.Device, 0),
        update: make(chan message),
    }
}

func main() {
    flag.Parse()

    manager := dpm.NewManager(FLXLister{})
    manager.Run()
}

