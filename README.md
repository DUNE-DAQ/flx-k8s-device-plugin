### General info

Discovers and lists `/dev/flx*` devices. Handles allocation when `felix.cern/flx` resources are requested by K8s pods.

Requires access to `/dev` and `/var/lib/kubelet/device-plugins` in a K8s node.

K8s Device Plugins https://kubernetes.io/docs/concepts/extend-kubernetes/compute-storage-net/device-plugins/

Based on the library https://github.com/kubevirt/device-plugin-manager

### Development cheat sheet

```
sudo docker run -it --rm --name flx-dev-plugin -v /dev:/dev -v /var/lib/kubelet/device-plugins/:/var/lib/kubelet/device-plugins/ -v <abs_path>/flx-dev-plugin:/usr/src/app egamberini/flx-dev-plugin:latest /bin/sh
```

Inside container:
```
cd /usr/src/app
go run main.go
```
