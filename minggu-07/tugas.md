## Spring Boot Docker Image

Masuk ke laman Spring Framework [spring.io](https://start.spring.io) untuk membuat project sesuai dengan kebutuhan, dalam praktik ini hal-hal yang diperlukan diantaranya : 

1. Build : Maven
2. Dependenci : Spring Web
3. Java JDK 17
4. Maven versi : 3.5.3

Catatan : 

Konfigurasi "Dockerfile" bergantung dengan versi java yang kita gunakan, dengan demikian, pastikan dependensi elemen yang kita pakai di Spring io sesuai dengan JDK ataupun build yang ada dimesin kita.



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

Update snapshot.jar yang akan kita gunakan dengan mengetikan `mvn clean package`

Tunggu sampi proses selesai 

![clearn-package-2](<48.mvn clean package.PNG>)

![maven-clean-3](<49.mvn clean package-build sukses.PNG>)

Klik kanan pada image dan pilih "build image"

![klik-kanan-pada-image](<50.Klik Kanan-pada-Dockerfile-dan-bulid-image.PNG>)

Sesuaikan versi image yang akan kita buat

![type-version-image](51.Sesuaikan-versi-Image-yang-kita-buat.PNG)

Pastikan build image telah selesai dengan terminal yang menampilkan output sebagai berikut :

![result build image](52.Pastikan-Build-image-telah-selesai-dengan-hasil-output-sebagai-berikut.PNG)

Lihat image di docker extensions

![Buka diim](53.Buka-di-image-exstention.PNG)

Klik dan pilih "run" dalam image tersebut

![choice-run](54.Klik-run.PNG)

Cek Image yang sedang running dengan type `docker ps` bisa di terminal vscode ataupun terminal linux

![Cek image yang sedang running](<55.Cek-Image-Yang sedang-running dengan docker ps-diterminal.PNG>)

Buka localhost:8080 untuk mengecek app yang kita bangun bekerja atau tidak

Catatan :
port 8080 adalah port yang umumnya diapakai oleh framework spring boot

![Open-localhost](56.Buka-local-host-8080-dibrowser.PNG)

Stop image dengan type `docker stop <id kontainer>` bisa dilakukan di termianl vscode ataupun linux 

![stop the container](57.Stop-imgae-yang-runnig.PNG)

## Selesai