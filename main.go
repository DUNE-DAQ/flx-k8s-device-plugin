// enrico.gamberini@cern.ch April 2022
// Adapted from:
// https://github.com/kubevirt/kubernetes-device-plugins/blob/master/pkg/kvm/kvm.go
// https://github.com/kubevirt/kubernetes-device-plugins/blob/master/pkg/pci/plugin.go
// and more...

package main

import (
    "os"
    "flag"
    "fmt"
    "strconv"
    "strings"

    "golang.org/x/net/context"
    "github.com/kubevirt/device-plugin-manager/pkg/dpm"
    pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

// FLXLister is the object responsible for discovering initial pool of devices and their allocation
type FLXLister struct {}

type message struct {}

type Plugin struct{
  name string
  update chan message
}

const (
    DevPath           = "/dev/"
    FLXName           = "flx"
    resourceNamespace = "felix.cern"
)

// Set up resources if needed, initialize custom channels etc
//func (p *Plugin) Start() error {
//    return nil
//}

// Tear down resources if needed
//func (p *Plugin) Stop() error {
//    return nil
//}

// ListAndWatch returns a stream of List of Devices
// Whenever a Device state changes or a Device disappears, ListAndWatch
// returns the new list
func (p *Plugin) ListAndWatch(e *pluginapi.Empty, s pluginapi.DevicePlugin_ListAndWatchServer) error {
    fmt.Println("ListAndWatch()", p.name)

    var devs []*pluginapi.Device

    var dev0 = DevPath + "flx0"
    var dev1 = DevPath + "flx1"
    if _, err0 := os.Stat(dev0); err0 == nil {
        if _, err1 := os.Stat(dev1); err1 == nil {
            devs = append(devs, &pluginapi.Device {
                ID: "flx0",
                Health: pluginapi.Healthy,
            })
        }
    }

    s.Send(&pluginapi.ListAndWatchResponse{Devices: devs})

    for {
        select {
        case <-p.update:
            fmt.Println("is this ever called?")
            s.Send(&pluginapi.ListAndWatchResponse{Devices: devs})
        }
    }
}

// Allocate is called during container creation so that the Device
// Plugin can run device specific operations and instruct Kubelet
// of the steps to make the Device available in the container
func (p *Plugin) Allocate(ctx context.Context, r *pluginapi.AllocateRequest) (*pluginapi.AllocateResponse, error) {
    fmt.Println("Allocate()", p.name)

    var response pluginapi.AllocateResponse

    for _, req := range r.ContainerRequests {
        var devices []*pluginapi.DeviceSpec
        for _, id := range req.DevicesIDs {
            fmt.Println("Allocate id", id)
            dev0 := new(pluginapi.DeviceSpec)
            fmt.Println("dev path", DevPath + id)
            dev0.HostPath = DevPath + id
            dev0.ContainerPath = DevPath + id
            dev0.Permissions = "rw"
            devices = append(devices, dev0)
            dev1 := new(pluginapi.DeviceSpec)
            dev1.HostPath = DevPath + "flx1"
            dev1.ContainerPath = DevPath + "flx1"
            dev1.Permissions = "rw"
            devices = append(devices, dev1)
            devmem := new(pluginapi.DeviceSpec)
            devmem.HostPath = DevPath + "cmem_rcc"
            devmem.ContainerPath = DevPath + "cmem_rcc"
            devmem.Permissions = "rw"
            devices = append(devices, devmem)
        }
        response.ContainerResponses = append(response.ContainerResponses, &pluginapi.ContainerAllocateResponse{
			Devices: devices})
    }

    return &response, nil
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

func (l FLXLister) GetResourceNamespace() string {
    return resourceNamespace
}

// Discovery discovers all devices within the system. Monitors available resources.
func (l FLXLister) Discover(pluginListCh chan dpm.PluginNameList) {
    fmt.Println("Discover()")
    var plugins = make(dpm.PluginNameList, 0)

    f, _ := os.Open(DevPath)
    files, _ := f.Readdir(0)
    var count = 0
    for _, v := range files {
        // looking for /dev/flx* but not /dev/flx (soft link)
        if strings.Contains(v.Name(), FLXName) && v.Name() != FLXName {
            fmt.Println(v.Name(),"contains flx")
            count = count + 1
        }
    }
    fmt.Println("flx count", count)
    for i := 0; i < count; i+=2 {
        fmt.Println(i)
        plugins = append(plugins, FLXName + strconv.Itoa(i) + "-data")
        plugins = append(plugins, FLXName + strconv.Itoa(i) + "-ctrl")
    }

    pluginListCh <- plugins
}

func (l FLXLister) NewPlugin(name string) dpm.PluginInterface {
    fmt.Println("NewPlugin()", name)
    return &Plugin{
        name: name,
        update: make(chan message),
    }
}

func main() {
    flag.Parse()

    manager := dpm.NewManager(FLXLister{})
    manager.Run()
}

