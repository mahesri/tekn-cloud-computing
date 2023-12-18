# Docker Orchestration Hands-on Lab

## Section 1: What is Orchestration

Sesi 1 merupakan penjelasan mengenai Orkrestasi, Orkrestasi secara umum dapat dijelaskan "misalkan" kita memiliki aplikasi dengan trafik yang tinggi dan memiliki high-availability requirements.

Salah satu manfaat dari Orkrestasi adalah kita dapat menghilangkan sebagian besar pekerjaan manual dalam mendeploy aplikasi dan membiarkan otomatisasi melakukan pekerjaan beratnya. Salah satu fitur keren dari Orkestrasi dengan Docker Swarm, adalah bahwa kita dapat mendeploy aplikasi di banyak host hanya dengan satu perintah (setelah Swarm mode diaktifkan). Selain itu, jika salah satu node/host pendukung mati dalam Docker Swarm kita, node/host lain secara otomatis akan mengambil beban, dan aplikasi Anda akan terus berjalan seperti biasa.

## Section 2: Configure Swarm Mode

Menambahkan container baru perintah berikut ke node 1 `docker run -dt ubuntu sleep infinity` dan didapati output sebagai berikut 

```
$ docker run -dt ubuntu sleep i
nfinity
Unable to find image 'ubuntu:la
test' locally
latest: Pulling from library/ubuntu
a48641193673: Pull complete
Digest: sha256:6042500cf4b44023ea1894effe7890666b0c5c7871ed83a97c36c76ae560bb9b
Status: Downloaded newer imagefor ubuntu:latest
14942d3d2e2a285e6be200c8b1d466d14a0a1966e106bf9a20694762258cb2d9
```

Cek container dengan perintah `docker ps` dan didapati output sebagai berikut :

[cek container](01.cek-container.PNG)

## Step 2.1 - Create a Manager node

Inisiasikan docker swarm di node1 dengan perintah `docker swarm init --advertise-addr $(hostname -i)` dan cek apakah node1 sudah terkonfigurasi menjadi manager dengan perintah `docker info` dan didapati output sebagai berikut :

```
Client:
 Version:    24.0.7
 Context:    default
 Debug Mode: false
 Plugins:
  buildx: Docker Buildx (Docker Inc.)
    Version:  v0.11.2
    Path:     /usr/local/libexec/docker/cli-plugins/docker-buildx
  compose: Docker Compose (Docker Inc.)
    Version:  v2.23.0
    Path:     /usr/local/libexec/docker/cli-plugins/docker-compose
  scout: Docker Scout (Docker Inc.)
    Version:  v1.0.9
    Path:     /usr/lib/docker/cli-plugins/docker-scout

Server:
 Containers: 1
  Running: 1
  Paused: 0
  Stopped: 0
 Images: 1
 Server Version: 24.0.7
 Storage Driver: overlay2
  Backing Filesystem: xfs
  Supports d_type: true
  Using metacopy: false
  Native Overlay Diff: true
  userxattr: false
 Logging Driver: json-file
 Cgroup Driver: cgroupfs
 Cgroup Version: 1
 Plugins:
  Volume: local
  Network: bridge host ipvlan macvlan null overlay
  Log: awslogs fluentd gcplogs gelf journald json-file local logentries splunk syslog
 Swarm: active
  NodeID: arejcb3u21gmvy65ikiqh87lj
  Is Manager: true
  ClusterID: dzobsuoed7gmwg2iyace3u86f
  Managers: 1
  Nodes: 1
  Default Address Pool: 10.0.0.0/8
  SubnetSize: 24
  Data Path Port: 4789
  Orchestration:
   Task History Retention Limit: 5
  Raft:
   Snapshot Interval: 10000
   Number of Old Snapshots to Retain: 0
   Heartbeat Tick: 1
   Election Tick: 10
  Dispatcher:
   Heartbeat Period: 5 seconds
  CA Configuration:
   Expiry Duration: 3 months
   Force Rotate: 0
  Autolock Managers: false
  Root Rotation In Progress: false
  Node Address: 192.168.0.27
  Manager Addresses:
   192.168.0.27:2377
 Runtimes: io.containerd.runc.v2 runc
 Default Runtime: runc
 Init Binary: docker-init
 containerd version: 091922f03c2762540fd057fba91260237ff86acb
 runc version: v1.1.9-0-gccaecfc
 init version: de40ad0
 Security Options:
  apparmor
  seccomp
   Profile: builtin
 Kernel Version: 4.4.0-210-generic
 Operating System: Alpine Linux v3.18 (containerized)
 OSType: linux
 Architecture: x86_64
 CPUs: 8
 Total Memory: 31.42GiB
 Name: node1
 ID: dfe970e0-375f-404a-be13-f23f67ac4b2c
 Docker Root Dir: /var/lib/docker
 Debug Mode: true
  File Descriptors: 46
  Goroutines: 163
  System Time: 2023-12-18T14:22:54.40698414Z
  EventsListeners: 0
 Experimental: true
 Insecure Registries:
  127.0.0.1
  127.0.0.0/8
 Live Restore Enabled: false
 Product License: Community Engine

WARNING: API is accessible on http://0.0.0.0:2375 without encryption.
         Access to the remote API is equivalent to root access on the host. Refer
         to the 'Docker daemon attack surface' section in the documentation for
         more information: https://docs.docker.com/go/attack-surface/
WARNING: No swap limit support
WARNING: bridge-nf-call-iptables is disabled
WARNING: bridge-nf-call-ip6tables is disabled
```
## Step 2.2 - Join Worker nodes to the Swarm

Jalankan perintah berikut kedalam node1 dan node2 `docker swarm join --token SWMTKN-1-2wljeoxo0wsu14a58ki9dhsao3o8fdf4pq8y1g2j2vjacxlcgf-8yxql4d6dw6gvh7126qaup4n5 192.168.0.7:2377` dan dalam praktikum ini didapai output sebagai berikut dalam masing-masing node :

Display output node2 

[node2 display](02.display-node2.PNG)

Display output node3

[display node3](03.display-node3.PNG)

Kembali ke node1, dan periksa apakah node1 dan node2 sudah menjadi worker nodes dari node1 sebagai managernya, dengan perintah berikut `docker node ls` dan didapati output sebagai berikut :

```

ID                            HOSTNAME   STATUS    AVAILABILITY   MANAGER STATUS   ENGINE VERSION
ghknskfz1o5z8qgrz1zjct2yt *   node1      Ready     Active         Leader          24.0.7
i5sbfs6zpsngfj3vax6guem6f     node2      Ready     Active          24.0.7
h2tylrd6wckobix2bh0j31izs     node3      Ready     Active          24.0.7
```

## Section 3: Deploy applications across multiple hosts

## Step 3.1 - Deploy the application components as Docker services

Deploy sleep as a Service across Docker Swarm dengan perintah berikut `docker service create --name sleep-app ubuntu sleep infinity` dan didapati output sebagai berikut :




