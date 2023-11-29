# Docker Networking

## Task 1 : Networking Basics

Melihat daftar commands utama dengan perintah `docker network` daftar commands utama berisi command untuk mengkonfigurasi dan mengatur network 
container network.

## Step 2: List networks

Melihat Docker Network yang tersedia saat ini di host dengan perintah `docker network ls`

![cek available network](01.cek-available-network.PNG)

## Step 3: Inspect a network

Melihat detail sebuah network dari container di docker host kita dengan perintah `docker network inspect bridge` dari prompt tersebut maka akan didapati output sebagai berikut : 

```
[
    {
        "Name": "bridge",
        "Id": "b836373449ed285a95cb41724a347b37144bb33be7fd428849c5af9
276ec23fa",
        "Created": "2023-11-28T23:23:25.906925887Z",
        "Scope": "local",
        "Driver": "bridge",
        "EnableIPv6": false,
        "IPAM": {
            "Driver": "default",
            "Options": null,
            "Config": [
                {
                    "Subnet": "172.17.0.0/16",
                    "Gateway": "172.17.0.1"
                }
            ]
        },
        "Internal": false,
        "Attachable": false,
        "Ingress": false,
        "ConfigFrom": {
            "Network": ""
        },
        "ConfigOnly": false,
        "Containers": {},
        "Options": {
            "com.docker.network.bridge.default_bridge": "true",
            "com.docker.network.bridge.enable_icc": "true",
            "com.docker.network.bridge.enable_ip_masquerade": "true",
            "com.docker.network.bridge.host_binding_ipv4": "0.0.0.0",
            "com.docker.network.bridge.name": "docker0",
            "com.docker.network.driver.mtu": "1500"
        },
        "Labels": {}
    }
]
```

Catatan :

`docker network inspect bridge` adalah perintah untuk mengecek detail dari jaringan bernama `bridge` oleh sebab itu, perintah yang umum digunakan untuk mnegecek sebuah jaringan untuk sebuah container adalah `docker network inspect <nama/id sebuah jaringan>`

## Step 4: List network driver plugins

Cek Installasi Docker di host dengan perintah `docker info`


