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

``````- <p className="text-center">No items yet! Add one above!</p>
+ <p className="text-center">You have no todo items yet! Add one above!</p>
``````
