# 🏥 Puskesmas API (GraphQL)

API Backend untuk sistem manajemen Puskesmas, dibangun menggunakan **Golang** dan **GraphQL**. Proyek ini menerapkan *Clean Architecture* standar industri untuk memastikan kode yang terstruktur, mudah di-maintenance, dan aman.

## 🚀 Fitur Utama
* **GraphQL API:** Menggunakan `gqlgen` untuk fleksibilitas kueri data.
* **Clean Architecture:** Pemisahan struktur yang jelas antara *Storage*, *Service*, dan *Controller/Resolver*.
* **Database Transaction:** Menjamin integritas data menggunakan `sqlx.Tx` untuk operasi kritikal (Update & Soft Delete).
* **Keamanan (JWT & RBAC):** Dilengkapi *middleware* untuk otentikasi JSON Web Token dan *Role-Based Access Control* (khusus Admin).

## 🛠️ Teknologi yang Digunakan
* **Bahasa:** Go (Golang)
* **Router:** Chi Router (`go-chi/chi/v5`)
* **GraphQL:** `99designs/gqlgen`
* **Database:** MySQL
* **DB Driver & Helper:** `go-sql-driver/mysql` & `jmoiron/sqlx`
* **Keamanan:** `golang-jwt/jwt/v5`

## 📦 Cara Menjalankan Proyek

1. **Clone repository ini:**
   ```bash
   git clone [https://github.com/Alfeus22/puskesmas-api.git](https://github.com/Alfeus22/puskesmas-api.git)
   cd puskesmas-api