# Go PT ZEN Test
technical test PT ZEN - Backend Engineer using Go

## SETUP PROJECT
- **Build Script Dockerfile**:
    ```bash
     docker build --tag ptzen-app .
     ```
- **Create Network**:
    ``` bash
    docker network create ptzen-app-network
    ```
- **Run Container**:
    ```bash
    Run Docker Container Golang -> run : docker run -d -p 8080:8080 --name ptzen-app --network ptzen-app-network ptzen-app
    ```
## CREATE TABLE DATABASE:
- **Check PG Container id**:
    ``` bash
    docker ps
    ```
- **Exec PG Container**:
    ``` bash
    docker exec -it <pg_container_id> psql -U Postgres
    ```
- **Connect Database**:
    ```bash
    \c assesment_ptzen
    ```
- **Copy & Paste Query to Create Database**:
    ```bash
    querydb.sql
    ```

## SOAL
    Deskripsi Proyek:
    Buatlah sebuah sistem manajemen toko online sederhana menggunakan Go atau nodeJS dengan Frontend ReractJS. Sistem ini harus memiliki 

    fitur-fitur berikut:
    1. Manajemen produk: tambah, edit, hapus, dan lihat daftar produk.
    2. Manajemen pesanan: lihat daftar pesanan, tandai pesanan sebagai diproses atau selesai.
    3. Autentikasi: register, login, profile, reset password dan logout.
    4. Setiap entitas (produk, pesanan) harus memiliki endpoint RESTful yang sesuai.

    Tugas:
    1. Rancang struktur aplikasi yang memisahkan logika bisnis, routing, dan penyimpanan data.
    2. Implementasikan autentikasi menggunakan JWT (JSON Web Token) dan package bcrypt untuk hashing password.
    3. Buatlah endpoint-endpoint HTTP.
    4. Gunakan Go atau NodeJS untuk menghubungkan aplikasi dengan database, Anda dapat menggunakan PostgreSQL,MySQL, SQLite atau database lain sesuai preferensi Anda.
    5. Pastikan aplikasi Anda dapat menangani permintaan HTTP secara asinkron.
    6. Gunakan Go atau NodeJS untuk melakukan validasi input dari pengguna di sisi backend.
    7. Gunakan ReactJS untuk melakukan validasi input dari pengguna di sisi Frontend (Jika Fullstack).
    8. Rancang arsitektur aplikasi Anda dengan mempertimbangkan aspek keamanan, skalabilitas, dan maintainability.