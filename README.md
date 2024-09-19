Enigma Laundry

Enigma Laundry adalah aplikasi berbasis Go yang menggunakan PostgreSQL sebagai database untuk mengelola data pelanggan (customer), layanan (service), dan pesanan (order).

Prasyarat

    Go (versi 1.18 ke atas)
    PostgreSQL (versi 12 ke atas)

Struktur Folder

Berikut adalah struktur folder aplikasi ini:

go

challenge-godb/
│
├── entity/
│   ├── customer_enrollment.go
│   ├── service_enrollment.go
│   ├── order_enrollment.go
│   ├── order_detail_enrollment.go
│
├── sql/
│   ├── DDL.sql
│   ├── DML.sql
│
├── main.go
├── go.mod
├── go.sum
└── README.md

Langkah-Langkah Instalasi
1. Clone Repository

Clone repository ini ke dalam direktori lokal Anda:

bash

git clone https://github.com/AndikaPrasetia/challenge-godb.git

2. Buat Database di PostgreSQL

Login ke PostgreSQL dan buat database baru:

sql

CREATE DATABASE enigma_laundry;

Setelah itu, buat tabel-tabel yang dibutuhkan:

sql

File SQL tertera di 

1. Konfigurasi Database di main.go

Buka file main.go dan sesuaikan kredensial database Anda:

go

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "password_database_anda"
    dbname   = "enigma_laundry"
)

4. Instal Dependensi

Jalankan perintah berikut untuk menginstal dependensi yang dibutuhkan:

bash

go mod tidy

5. Jalankan Aplikasi

Setelah semua dependensi terpasang, jalankan aplikasi dengan perintah berikut:

bash

go run main.go

6. Fitur Aplikasi

Aplikasi ini menyediakan fitur-fitur berikut:

    Customer Management:
        Tambah pelanggan
        Lihat daftar pelanggan
        Lihat detail pelanggan berdasarkan ID
        Update data pelanggan
        Hapus pelanggan

    Service Management:
        Tambah layanan
        Lihat daftar layanan
        Lihat detail layanan berdasarkan ID
        Update layanan
        Hapus layanan

    Order Management:
        Tambah pesanan
        Lengkapi pesanan
        Lihat daftar pesanan
        Lihat detail pesanan berdasarkan ID

Sumber

    Golang
    PostgreSQL