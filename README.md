# Withdraw Apps
## Demo
https://bearded-minister-51298.herokuapp.com
## Requirement
1. Docker dan docker compose
2. Go 1.16
3. Postgresql 13
4. Heroku CLI

## Konfigurasi
Konfigurasi dapat dilakukkan dengan mengubah nilai yang ada di file **flip.env** . Berikut penjelasan tiap fieldnya.

|field|description  |
|---|---|
| DBHost | alamat database server |
| DBPort | port database server yang digunakan |
| DBUser | username database server |
| DBPassword | password database server |
| DBName |nama database yang digunakan  |
| DBSsl |Jika akses database menggunakan ssl maka diset nilai **required**, jika tidak maka set dengan nilai **disable**   |
| FlipHost |alamat server flip  |
| FlipSecret | kode rahasia untuk mengakses server flip |

## Inisialisasi Database
Untuk membuat table yang dibutuhkan kalian bisa menjalankan file **database.sql**. Setelah kalian jalankan, akan muncul 3 tabel yaitu withdrawals, bigflip_logs dan bigflip_response. Proses ini bisa kalian lewati jika menggunakan postgres yang telah disediakan (dengan menggunakan docker).

## Petunjuk Installasi di Komputer Lokal
Installasi ini ditujukkan agar aplikasi ini dapat dijalankan di komputer local. Proses installasi ini hanya memerlukan docker karena Komponen yang digunakan yaitu docker sebagai container. Berikut langkah-langkahnya:
1. Nyalakan database postgres. Jika ingin menggunakan postgres di dalam docker, maka jalankan perintah `make start-db`. Namun jika tidak ingin menggunakan postgres, ubah konfigurasi di file **flip.env**
2. Jalankan aplikasi utama dengan menggunakan perintah `make start-app`.
3. Buka browser dan akses ke `localhost:8000` untuk mengakses aplikasi utama.
4. Untuk menghentikan aplikasi, cukup jalankan `make stop`

## Petunjuk Installasi di Heroku
Installasi ini ditujukkan jika ingin aplikasi dijalankan di platform heroku. Berikut hal-hal yang harus dipersiapkan sebelum proses installasi.
1. Install heroku cli di komputer kalian. Jika belum, dapat lihat petunjuk di link berikut https://devcenter.heroku.com/articles/heroku-cli
2. Login akun heroku melalui heroku cli
3. Buat container heroku. Jika belum, bisa jalankan perintah `heroku stack:set container` di direktori aplikasi.
4. Buat database di heroku. Jika belum, maka jalankan perintah `make heroku-create-db` lalu ubah konfigurasi **file.env** sesuai kredensial database yang kalian peroleh saat membuat database di heroku.
5. Jalankan proses inisialisasi database untuk membuat tabel-tabel yang diperlukan.

Setelah proses pra installasi diatas, berikut pentunjuk untuk proses installasi:
1.  Jalankan perintah `make heroku-deploy`. Perintah tersebut akan menjalankan proses deployment ke heroku.
2. Jika proses deployment sudah selesai, jalankan perintah `heroku open` untuk mengakses aplikasi.

## Withdraw
Proses withdraw bisa dilakukkan dengan mengakses aplikasi melalui browser. Selain itu, proses withdraw juga dapat dilakukan dengan menggunakan command line melalui terminal. Berikut contoh perintah untuk melakukan withdraw menggunakan command line.

    ACCOUNT_NUMBER=11122 AMOUNT=10000 REMARK=thanks BANK_CODE=bni make withdraw
**Note**: Perlu diingat, proses withdraw dengan command line hanya bisa digunakan jika kita melakukan proses installasi di komputer lokal.

