## 1 Install Docker Dekstop

Proses install Docker dekstop di Windows

![install dekstop docker](02.Install-docker-dekstop-windows.PNG)

Memilih pengaturan pengistallan yang direkomendasikan 

![memilih pengaturan yang direkomendasikan](03.Install-docker-dekstop-windows.PNG)

Sign Up ke akun Docker 

![sign-up docker](04.sign_up-docker.PNG)

Siap membuat sebuah image dengan Docker

![siap membuat image dengan docker](05.ready-built-an-image-with-docker.PNG)

## 2 Get Started 

Cloning get-started docker repository 

![Cloning-get-started-docker-Repository](06.Cloning-get-started-docker-Repository.PNG)

Mengecek hasil cloningan 

![Cek Hasil cloningan](07.Make_Sure_That_Cloning_Proses_Is_Done.PNG)

Menambahkan konten konfigurasi image kedalam docker file 

![add content to docker-file](08.Add-Content-to-Docker-File.PNG)

Ketikan `docker run -dp 127.0.0.1:3000:3000 getting-started`

![Login ke docker](09.Login-to-Docker.PNG)

Masuk web browser dan ketikan `127.0.0.1:3000`

![Masuk-ke-web-browser](10.Login-to-web-browser.PNG)

Menambahkan item kedalam aplikasi to do

![add-item](11.add-item.PNG)

Mengecek container yang sedang running

![cek-running-containers](12.Cek-containers.PNG)

Memodifikasi `src/static/js/app.js' ketikan : 

``````
- <p className="text-center">No items yet! Add one above!</p>
+ <p className="text-center">You have no todo items yet! Add one above!</p>
``````
Menghapus container yang sedang berjalan dengan `docker rf -f <id / nama container>`

![Hapus container yang sedang running](13.Stop-Continer-Who-Running.PNG)

Berbagi repository image 

![share repository image](15.Share-repository-image.PNG)

Login docker di CLI 

![login-docker-cli](16.Login-Docker.PNG)

Cek Images yang sedang running 

![cek image yang sedang runnimng](17.Cek-Images.PNG)

Push image ke reposytory yang tadi dibuat 

![push image to repo](18.push-image-latest-tag.PNG)

Cek hasil image yang berhasil di push ke repository

![cek the repo](19.Cek-hasil-push-didocker.PNG)

Belajar docker secara online dengan klik [play-with-docker](https://labs.play-with-docker.com/)

![go-to play with docker](20.goto-play-with-docker-create-new-instance.PNG)

Manjalankan image yang tadi dibuat direpository dengan play with docker 

![play-with-docker](21.Run-Image-that-we've-been-push-from-our-repo.PNG)


Tekan port 300 maka akan dialihkan ke todo app seperti digambar berikut 

![todo app](22.Klik-3000-port-to-open-to-do-app.PNG)

Membuat database todo di docker 

![create db](23.Persist-DB1.PNG)

<b>Persist DB</b> dalam langkah ini, ditujukan untuk belajar agar data yang kita buat dapat menyimpan file di host dan membuatnya tersedia untuk kontainer berikutnya, kontainer berikutnya seharusnya dapat melanjutkan dari tempat yang terakhir ditinggalkan. Dengan membuat volume dan menautkannya dengan kontainer yang baru.

![Persist DB 1](23.Persist-DB1.PNG)

Membuat database 

![Create database](24.Menghapus-image-sebelumnya.PNG)

Menghapus image sebelumnya yang masih berjalan 

![new imaghe](25.Menghapus-image-sebelumnya-dan-membuat-baru.PNG)

Buka kembali aplikasi todo di lokalhost dan tambahkan beberapa item 

![add item todo app](26.Buka-kembali-to-do-APP.PNG)

Mengetahui letak todo-db

![Mengetahui letak todo-db](27.Mengetahui-letak-to-do-db.PNG)

Memberi tahu docker membuat --mount

![bind mount](30.Tells-docker-to-create-mount.PNG)

Create myfile.txt

![create myfile.txt](31.Create-myfile.txt-in-src-directori.PNG)

Membuat docker network 

![make an network](32.make-network.PNG)

Menghubungkan Kontainer kedalam network

![connect kontainer to network](33.menghubungkan-container-kedalam-network.PNG)

Melihat database

![melihat database](34.Melihat-databases.PNG)

Mneginstall netsnod 

![install netsnod](35.Install-netshot.PNG)

Dig mysql

![dig mysql](36.Dig-mySql.PNG)

Melihat log yang sedang berlangsung dari kontainer todo-app  

![see](37.melihat-log-dari-container-todo-app.PNG)

Melihat item yang kita tambahkan diitem todo yang kita tambahka di localhost tadi ke database terakhir kita buat di CLI 

![show db](<38.Melihat-apakah data-yang-ada-didalam-database-sama-dengan-yang-dimasukan.PNG>)

<b>Docker compose</b>

Docker compose memungkinkan kita mendefinisikan tumpukan aplikasi Anda dalam satu file.

![docker compose](39.Docker-Compose-1.PNG)

Langkah terakhir mengetahui daftar aplikasi yang ter-compose 

![compose-docker](40.Docker-Compose-2-Cek.PNG)