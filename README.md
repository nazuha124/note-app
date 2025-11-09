# 📝 Note App — Next.js + Go Fiber + PostgreSQL (Dockerized)

Aplikasi catatan multi-user dengan **JWT Authentication**, **CRUD Notes (Create, Read, Delete)**, dan **Upload Gambar dengan Auto Resize & Kompresi**.  
Dibangun menggunakan **Next.js** (Frontend), **Go Fiber** (Backend), dan **PostgreSQL** (Database) dengan **Docker Compose** untuk deployment.

---

## 🚀 Fitur Utama

✅ **Autentikasi JWT**
- Register & Login user
- Token JWT disimpan di browser dan diverifikasi di backend

✅ **Manajemen Catatan (CRUD)**
- Tambah, tampilkan, hapus catatan milik user sendiri  
- Upload gambar (.png/.jpg) → otomatis resize 800px & kompres

✅ **Keamanan & Logging**
- Hash password dengan bcrypt
- Middleware JWT untuk proteksi endpoint `/api/notes`
- Logger mencatat setiap request dan response

✅ **Dockerized**
- Semua service (frontend, backend, database) berjalan otomatis dengan `docker compose up`

---

## 🧱 Arsitektur Sistem

📦 **Komunikasi antar service (Docker Compose):**
- `frontend` → port **3000**
- `backend` → port **8081**
- `db` (PostgreSQL) → port **5432**

---

## 🛠️ Tech Stack

| Layer | Teknologi | Fungsi |
|--------|------------|--------|
| **Frontend** | Next.js (App Router), TailwindCSS, Axios | UI & komunikasi ke API |
| **Backend** | Go Fiber, JWT, bcrypt, imaging | REST API, autentikasi, upload gambar |
| **Database** | PostgreSQL | Penyimpanan data relasional |
| **Deployment** | Docker Compose | Menjalankan semua service terintegrasi |

---

## 🧩 Struktur Proyek

note-app/
├── backend/
│ ├── main.go
│ ├── routes/
│ │ ├── auth.go
│ │ └── notes.go
│ ├── middleware/
│ │ └── jwtmw.go
│ └── database/
│ └── db.go
│
├── frontend/
│ ├── app/
│ │ ├── login/page.tsx
│ │ ├── register/page.tsx
│ │ └── notes/page.tsx
│ ├── lib/api.ts
│ └── ...
│
├── docker-compose.yml
└── README.md
