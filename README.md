# 🍽️ ReservaResto - Aplikasi Reservasi Meja Restoran

[![Language](https://img.shields.io/badge/Language-Go-00ADD8?style=flat&logo=go&logoColor=white)](https://golang.org/)
[![Category](https://img.shields.io/badge/Category-Educational-blue)](https://github.com)
[![Status](https://img.shields.io/badge/Status-Finished-brightgreen)](https://github.com)

## 📖 Deskripsi Umum

**ReservaResto** adalah aplikasi reservasi meja restoran berbasis *Command Line Interface* (CLI) yang dikembangkan sebagai proyek **Tugas Besar mata kuliah Algoritma dan Pemrograman 2** (Program Studi S1 Informatika, Universitas Telkom). Aplikasi ini dirancang untuk membantu staf admin atau resepsionis restoran dalam mengelola pemesanan meja secara digital, terstruktur, dan efisien.

Aplikasi ini mengimplementasikan berbagai konsep fundamental pemrograman terstruktur, seperti:

- **Tipe Data Terstruktur** (`struct` dan array statis)
- **Operasi CRUD** (Create, Read, Update, Delete)
- **Algoritma Pencarian** (*Sequential Search* dan *Binary Search*)
- **Algoritma Pengurutan** (*Selection Sort* dan *Insertion Sort*)
- **Sistem Reservasi** dengan pengecekan bentrok jadwal
- **Statistik Reservasi** (per hari dan meja terfavorit)

---

## 🎯 Latar Belakang

Di era digital saat ini, sektor restoran dituntut untuk beradaptasi dengan sistem manajemen reservasi yang modern. Metode manual seperti buku catatan atau papan tulis seringkali menimbulkan masalah seperti *double booking*, kehilangan data, dan ketidakpuasan pelanggan. **ReservaResto** hadir sebagai solusi untuk menjawab permasalahan tersebut dengan menyediakan sistem yang:

- ✅ **Akurat** – Meminimalisir kesalahan pencatatan
- ✅ **Efisien** – Mempercepat proses pencarian dan pengurutan data
- ✅ **Informatif** – Menyediakan statistik untuk pengambilan keputusan
- ✅ **Terstruktur** – Data tersimpan rapi dalam array of struct

---

## ⚙️ Fitur Utama

| No | Fitur | Deskripsi |
|----|-------|-----------|
| 1 | **Tambah Meja** | Menambahkan data meja baru (nomor meja & kapasitas kursi) |
| 2 | **Tampil Meja** | Menampilkan seluruh data meja yang tersimpan |
| 3 | **Ubah Meja** | Mengubah data meja berdasarkan nomor meja |
| 4 | **Hapus Meja** | Menghapus data meja berdasarkan nomor meja |
| 5 | **Tambah Pelanggan** | Menambahkan data pelanggan baru (ID, nama, no. telepon) |
| 6 | **Tampil Pelanggan** | Menampilkan seluruh data pelanggan |
| 7 | **Ubah Pelanggan** | Mengubah data pelanggan berdasarkan ID |
| 8 | **Hapus Pelanggan** | Menghapus data pelanggan berdasarkan ID |
| 9 | **Cari Meja (Sequential)** | Mencari meja berdasarkan nomor atau kapasitas (Sequential Search) |
| 10 | **Cari Meja (Binary)** | Mencari meja berdasarkan nomor atau kapasitas (Binary Search) |
| 11 | **Sort Kapasitas (Selection)** | Mengurutkan meja berdasarkan kapasitas (Selection Sort - ASC/DESC) |
| 12 | **Sort Kapasitas (Insertion)** | Mengurutkan meja berdasarkan kapasitas (Insertion Sort - ASC/DESC) |
| 13 | **Tambah Reservasi** | Membuat reservasi meja dengan pengecekan bentrok jadwal |
| 14 | **Lihat Statistik** | Menampilkan total reservasi per hari & meja terfavorit |
| 15 | **Status Ketersediaan** | Menampilkan ketersediaan meja per jam pada suatu tanggal |

---

## 🧠 Konsep dan Algoritma yang Diimplementasikan

### 🔍 Pencarian (Searching)

| Algoritma | Kegunaan |
|-----------|----------|
| **Sequential Search** | Mencari meja berdasarkan nomor/kapasitas (data tidak perlu terurut) |
| **Binary Search** | Mencari meja berdasarkan nomor/kapasitas (data HARUS terurut) |

### 📊 Pengurutan (Sorting)

| Algoritma | Kegunaan |
|-----------|----------|
| **Selection Sort** | Mengurutkan meja berdasarkan kapasitas kursi (ascending/descending) |
| **Insertion Sort** | Mengurutkan meja berdasarkan kapasitas kursi (ascending/descending) |

### 📝 Operasi CRUD

- **Create** – `tambahMeja()`, `tambahPelanggan()`, `tambahReservasi()`
- **Read** – `tampilMeja()`, `tampilPelanggan()`, `tampilReservasi()`
- **Update** – `ubahMeja()`, `ubahDataPelanggan()`
- **Delete** – `hapusMeja()`, `hapusDataPelanggan()`

---

## 🛠️ Teknologi yang Digunakan

| Aspek | Keterangan |
|-------|-------------|
| **Bahasa Pemrograman** | Go (Golang) |
| **Paradigma** | Pemrograman Terstruktur |
| **Struktur Data** | Array statis dengan kapasitas NMAX = 10000 |
| **Antarmuka** | _Command Line Interface_ (CLI) |
| **Penyimpanan** | Sementara di memori (volatil) |

---

## 🚀 Cara Menjalankan Program

### Langkah-langkah

```bash
# 1. Clone repository
git clone https://github.com/darkacademiaa/ReservaResto

# 2. Masuk ke direktori project
cd ReservaResto

# 3. Jalankan program
go run reservasiMeja2.go
