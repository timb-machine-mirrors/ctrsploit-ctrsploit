# env/graphdriver

graphdriver

存储驱动

```
root@b1f9f8da70c3:~/app# ./bin/release/env_linux_amd64 graphdriver
===========Overlay===========
[Y]  Enabled	
[Y]  Used	
The number of graph driver mounted	# equal to the number of containers
6
The host path of container's rootfs	
/var/lib/docker/overlay2/254080c7a7418fb2db4d5d27e03bc86f3fe4ae52b5c15b36ad0196adfc887f9e/merged
===========DeviceMapper===========
[Y]  Enabled	
[N]  Used	
The number of graph driver mounted	# equal to the number of containers
0
The host path of container's rootfs	


root@b1f9f8da70c3:~/app# ./bin/release/env_linux_amd64 --colorful graphdriver
===========Overlay===========
✔  Enabled	
✔  Used	
The number of graph driver mounted	# equal to the number of containers
6
The host path of container's rootfs	
/var/lib/docker/overlay2/254080c7a7418fb2db4d5d27e03bc86f3fe4ae52b5c15b36ad0196adfc887f9e/merged
===========DeviceMapper===========
✔  Enabled	
✘  Used	
The number of graph driver mounted	# equal to the number of containers
0
The host path of container's rootfs	


root@b1f9f8da70c3:~/app# ./bin/release/env_linux_amd64 --json graphdriver
{"overlay":{"loaded":true,"used":true,"refcnt":6,"hostPath":"/var/lib/docker/overlay2/254080c7a7418fb2db4d5d27e03bc86f3fe4ae52b5c15b36ad0196adfc887f9e/merged"},"device_mapper":{"loaded":true,"used":false,"refcnt":0,"hostPath":""}}
root@b1f9f8da70c3:~/app# ./bin/release/env_linux_amd64 --json graphdriver |jq
{
  "overlay": {
    "loaded": true,
    "used": true,
    "refcnt": 6,
    "hostPath": "/var/lib/docker/overlay2/254080c7a7418fb2db4d5d27e03bc86f3fe4ae52b5c15b36ad0196adfc887f9e/merged"
  },
  "device_mapper": {
    "loaded": true,
    "used": false,
    "refcnt": 0,
    "hostPath": ""
  }
}
```