## Spring Boot Docker Image

Masuk ke laman Spring Framework [spring.io](https://start.spring.io) untuk membuat project sesuai dengan kebutuhan, dalam praktik ini hal-hal yang diperlukan diantaranya : 

1. Build : Maven 3.5.1
2. Dependenci : Spring Web
3. Java JDK 17

![create dependecie elemen](<41.Create-Snapshot-&-add-depedencies-in spring boot.PNG>)

Buka hasil generate dengan visual studio code 

![open with vs code](42.Open-the-jar-fileusing-vs-code.PNG)

Menambahkan file html kedalam folder `main/resource/static/` 

![make html file](43.Make-an-index.html-file.PNG)

Edit file html sesuai kebutuhan 

![edit html](44.edit-index.html-file.PNG)

Dalam directori booting-web buat file baru dengan nama `Dockerfile`

![make Dockerfile](45.Make-an-new-file-and-give-a-name-docker-file.PNG)

Ketikan seperti pada gambar dibawah ini kedalam file `Dockerfile` yang baru saja dibuat

![Edit Dockerfile](46.Type-like-this-add-docker-file.PNG)

Buka terminal dan masuk kedalam directory booting-web 

![open-terminal](47.Masuk-ke-directory-booting-web.PNG)