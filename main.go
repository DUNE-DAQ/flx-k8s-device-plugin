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

type Plugin struct{ Name string }

//func (p *Plugin) Start() error {
    // Set up resources if needed, initialize custom channels etc
//    return nil
//}

//func (p *Plugin) Stop() error {
    // Tear down resources if needed
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

type Lister struct{ 
    Plugins []string
}

func (l Lister) GetResourceNamespace() string {
    glog.V(3).Infof("GetResourceNamespace()")
    fmt.Print("GetResourceNamespace()\n")
    return "flx.cern"
}

// Discovery discovers all devices within the system. Monitors available resources.
func (l Lister) Discover(pluginListCh chan dpm.PluginNameList) {
        glog.V(3).Infof("Discover()")
        fmt.Print("Discover()\n")
	var plugins = make(dpm.PluginNameList, 0)

	for _, name := range l.Plugins {
                glog.V(3).Infof("Discovered %s", name)
		plugins = append(plugins, name)
	}
        pluginListCh <- plugins
}

func (l Lister) NewPlugin(name string) dpm.PluginInterface {
    glog.V(3).Infof("NewPlugin()")
    fmt.Print("NewPlugin() ", name, "\n")
    return &Plugin{ Name: name}
}

func main() {
    flag.Parse()
    var FLXPath = "/dev/flx0"

    var plugins []string
    if _, err := os.Stat(FLXPath); err == nil {
		glog.V(3).Infof("Discovered %s", FLXPath)
                fmt.Print("Discovered ", FLXPath, "\n")
		plugins = append(plugins, "FLX712_0")
    }
    //plugins = append(plugins, "FLX712_1")
    //plugins = append(plugins, "FLX712_2")
    lister := Lister{Plugins: plugins}
    manager := dpm.NewManager(lister)
    //glog.Flush()
    manager.Run()
}

