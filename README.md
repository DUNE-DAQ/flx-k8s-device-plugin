### General info

Discover and lists `/dev/flx*` devices. Handles allocation when `felix.cern/flx` resources are requested by K8s pods.

Requires privileged access to `/dev` and `/var/lib/kubelet/device-plugins` in a K8s node.

Based on the library https://github.com/kubevirt/device-plugin-manager

### Development cheat sheet

```
sudo docker run -it --rm --name flx-dev-plugin --privileged -v /var/lib/kubelet/device-plugins/:/var/lib/kubelet/device-plugins/ -v /nfs/home/engamber/flx-dev-plugin:/usr/src/app egamberini/flx-dev-plugin:latest /bin/sh
```

Inside container:
```
cd /usr/src/app
go run main.go
```
