# 🔐 PDF Encryptor API
 
PDF Encryptor API adalah layanan HTTP untuk melakukan enkripsi dan dekripsi file PDF menggunakan algoritma **AES-GCM (256-bit)** dan **ChaCha20-Poly1305 (256-bit)**. File disimpan di disk lokal tanpa menggunakan database, dan aman untuk ditimpa (overwrite). Cocok digunakan dalam alur tanda tangan elektronik untuk menjaga kerahasiaan dan integritas dokumen PDF.

## ✨ Fitur

- 🔐 Enkripsi dan dekripsi file PDF dengan AES-GCM dan ChaCha20-Poly1305

- 🧾 Mendukung file overwrite (src == dest)

- 🔑 Endpoint untuk generate kunci enkripsi 256-bit

- 📂 Penyimpanan file di local disk, tanpa database

- 📜 Dokumentasi Swagger

- 🧱 Struktur aman dan ringan, cocok untuk container

- 🐳 Siap digunakan di Docker dan Docker Compose
  

## ⚙️ Requirement  

- Go 1.21 atau lebih tinggi (jika dijalankan secara lokal)

- Docker & Docker Compose (jika dijalankan sebagai container)

- [`swag` CLI](https://github.com/swaggo/swag) untuk generate dokumentasi Swagger:

```bash

go install github.com/swaggo/swag/cmd/swag@v1.16.4

```

## 🐳 Cara Menjalankan (Docker)

### Build & Run

```bash

docker  build  -t  yourname/pdf-encryptor:latest  .

docker  run  -p  7082:7082  yourname/pdf-encryptor:latest

```

### Dengan Docker Compose

```bash

docker-compose  up  -d

```  

Swagger dapat diakses di:

📍 `http://localhost:7082/swagger/index.html`


## 💻 Cara Menjalankan (Go Langsung)

```bash

git  clone  https://github.com/yourname/pdf-encryptor.git

cd  pdf-encryptor

go  mod  tidy

swag  init

go  run  main.go

```

## 📘 Dokumentasi API (Swagger)

 
Swagger UI:

```

http://localhost:7082/swagger/index.html

```

Generate dokumentasi:

```bash

swag  init

```

## 🔁 Format & Error Response  

### Contoh Response Sukses:

```json

{

"status": true,

"message": "Success"

}

```


### Contoh Error Response:

```json

{

"status": false,

"message": "invalid key: must be 256-bit hex string"

}

```


| Status Code | Keterangan |
|------------|-----------------------------------|
| 200 | Berhasil |
| 400 | Input tidak valid (key, path) |
| 403 | Akses file ditolak (permission) |
| 404 | File tidak ditemukan |
| 500 | Error sistem / enkripsi gagal |

  

## 🔐 Keamanan


- 🧱 Tidak menggunakan root user di dalam container
- 🔒 Source code tidak disertakan dalam final image
- 📦 File terenkripsi disimpan dalam binary-only image
- 🔁 Overwrite aman menggunakan file sementara (`.tmp`)
- 🔐 Nonce (IV) selalu unik dan disisipkan di awal file  

## 🪪 Lisensi

MIT License © 2025