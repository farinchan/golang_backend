# Golang Backend

Proyek ini adalah backend sederhana menggunakan **Golang**, **Fiber**, dan **GORM**.

## Fitur

- RESTful API dengan Fiber
- ORM menggunakan GORM
- Struktur kode yang mudah dikembangkan

## Prasyarat

- Go 1.18+
- Database (MySQL/PostgreSQL/SQLite, sesuai konfigurasi GORM)

## Instalasi

```bash
git clone https://github.com/username/golang_backend.git
cd golang_backend
go mod tidy
```

## Menjalankan Aplikasi

```bash
go run main.go
```

## Struktur Direktori

```
.
├── main.go
├── config/
├── models/
├── routes/
└── controllers/
```

## Konfigurasi Database

Edit file konfigurasi database di `config/` sesuai kebutuhan Anda.

## Lisensi

MIT License