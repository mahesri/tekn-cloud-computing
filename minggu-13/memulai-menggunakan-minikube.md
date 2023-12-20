# Hello Minikube

Dalam praktikum ini akan belajar menjalankan contoh aplikasi di Kubernetes menggunakan minikube. Dan didalam modul praktikum ini juga disedaikan image container yang menggunakan NGINX untuk mengembalikan semua permintaan.

Start minikube

![start minikube](02.Start-Minikube.PNG)

Buka dashboard minikube dengan mengetikan `minikube dashboard --url` dan copy-kan url didalam output.

![copy the url](03.Open-url-and-copy.PNG)

Buka browser dan Paste-kan url ditab baru.

![paste the url](01.open-dashboard.PNG)

Install kubectl 

![kubectl install](04.Install-kubectl.PNG)

Membuat Deployment untuk me-manages Pod

![create deploeyment](05.Create-helonode.PNG)

Lihat Deployment 

![see deployment](06.See-the-deployment.PNG)

View the Pods

![view the pods](07.View-the-pods.PNG)

View Cluster Event

![view culster event](09.View-cluster-events.PNG)

## Create Service 

Publish Pod ke internet

![publish pod](10.publish-pod.PNG)

View the Service

![View the service](11.view-the-service.PNG)

Run the following command

![run the](12.make-the-service-visible.PNG)

## Enable addons 

View minikube addon

![minikube](08.View-addons.PNG)

Enable addon metrics-server

![addon 1](13.addon-1.PNG)

Installing addon

![addon 3](14.addon-2.PNG)

## CLEANUP

Ketikan perintah berikut :

```
kubectl delete service hello-node
kubectl delete deployment hello-node
minikube stop
```