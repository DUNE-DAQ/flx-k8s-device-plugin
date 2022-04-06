### Development cheat sheet

```
sudo docker run -it --rm --name flx-dev-plugin --privileged -v /var/lib/kubelet/device-plugins/:/var/lib/kubelet/device-plugins/ -v /nfs/home/engamber/flx-dev-plugin:/usr/src/app egamberini/flx-dev-plugin:latest /bin/sh
```

Inside container:
```
cd /usr/src/app
go run main.go
```
