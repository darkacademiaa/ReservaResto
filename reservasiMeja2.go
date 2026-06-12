package main
import "fmt"
const NMAX int = 10000

type Meja struct {
	nomor int
	kapasitasKursi int
	count int
}

type Pelanggan struct {
	id   int
	nama string
	telp string
}

type Reservasi struct {
	idReservasi int
	idPelanggan int
	nomorMeja int
	tanggal string
	jam int
}

type tabMeja [NMAX]Meja
type tabPelanggan [NMAX]Pelanggan
type tabReservasi [NMAX]Reservasi

func main() {
	var M tabMeja
	var P tabPelanggan
	var R tabReservasi
	var nMeja, nPelanggan, nReservasi int
	var pilihan, i int

	for pilihan != 99 {
		fmt.Println("\n================================================================")
		fmt.Println("                    RESERVARESTO - RESERVASI")
		fmt.Println("================================================================")
		fmt.Println("1. Tambah Meja")
		fmt.Println("2. Tampil Meja")
		fmt.Println("3. Ubah Meja")
		fmt.Println("4. Hapus Meja")
		fmt.Println("5. Tambah Pelanggan")
		fmt.Println("6. Tampil Pelanggan")
		fmt.Println("7. Ubah Pelanggan")
		fmt.Println("8. Hapus Pelanggan")
		fmt.Println("9. Cari Meja (Sequential Search)")
		fmt.Println("10. Cari Meja (Binary Search)")
		fmt.Println("11. Sort Kapasitas (Selection Sort)")
		fmt.Println("12. Sort Kapasitas (Insertion Sort)")
		fmt.Println("13. Tambah Reservasi")
		fmt.Println("14. Lihat Statistik Reservasi")
		fmt.Println("15. Lihat Status Ketersediaan Meja")
		fmt.Println("99. Keluar")
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			tambahMeja(&M, &nMeja)
		} else if pilihan == 2 {
			tampilMeja(M, nMeja)
		} else if pilihan == 3 {
			var nomor int
			fmt.Print("Nomor meja yang akan diubah: ")
			fmt.Scan(&nomor)
			ubahMeja(&M, nMeja, nomor)
		} else if pilihan == 4 {
			var nomor int
			fmt.Print("Nomor meja yang akan dihapus: ")
			fmt.Scan(&nomor)
			hapusMeja(&M, &nMeja, nomor)
		} else if pilihan == 5 {
			tambahPelanggan(&P, &nPelanggan)
		} else if pilihan == 6 {
			tampilPelanggan(P, nPelanggan)
		} else if pilihan == 7 {
			var id int
			fmt.Print("ID pelanggan yang akan diubah: ")
			fmt.Scan(&id)
			ubahDataPelanggan(&P, nPelanggan, id)
		} else if pilihan == 8 {
			var id int
			fmt.Print("ID pelanggan yang akan dihapus: ")
			fmt.Scan(&id)
			hapusDataPelanggan(&P, &nPelanggan, id)
		} else if pilihan == 9 {
			var cariBerdasarkan, nilaiCari int
			fmt.Println("Cari Meja Berdasarkan:")
			fmt.Println("1. Nomor Meja")
			fmt.Println("2. Kapasitas Kursi")
			fmt.Print("Pilihan: ")
			fmt.Scan(&cariBerdasarkan)
			if cariBerdasarkan == 1 {
				fmt.Print("Nomor meja yang dicari: ")
				fmt.Scan(&nilaiCari)
				i = cariMeja(M, nMeja, nilaiCari, 1, "sequential")
			} else if cariBerdasarkan == 2 {
				fmt.Print("Kapasitas kursi yang dicari: ")
				fmt.Scan(&nilaiCari)
				i = cariMeja(M, nMeja, nilaiCari, 2, "sequential")
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
			if i != -1 {
				fmt.Println("Nomor Meja     :", M[i].nomor)
				fmt.Println("Kapasitas Kursi:", M[i].kapasitasKursi)
			} else {
				fmt.Println("Meja tidak ditemukan")
			}
		} else if pilihan == 10 {
			var cariBerdasarkan, nilaiCari int
			fmt.Println("Cari Meja Berdasarkan (Binary Search):")
			fmt.Println("1. Nomor Meja")
			fmt.Println("2. Kapasitas Kursi")
			fmt.Print("Pilihan: ")
			fmt.Scan(&cariBerdasarkan)
			if cariBerdasarkan == 1 {
				fmt.Print("Nomor meja yang dicari: ")
				fmt.Scan(&nilaiCari)
				var tempTab tabMeja
				for j := 0; j < nMeja; j++ {
					tempTab[j] = M[j]
				}
				selectionSortMejaByNomorAsc(&tempTab, nMeja)
				i = cariMeja(tempTab, nMeja, nilaiCari, 1, "binary")
				if i != -1 {
					// Cari indeks asli di M
					for j := 0; j < nMeja; j++ {
						if M[j].nomor == tempTab[i].nomor {
							i = j
							break
						}
					}
				}
			} else if cariBerdasarkan == 2 {
				fmt.Print("Kapasitas kursi yang dicari: ")
				fmt.Scan(&nilaiCari)
				var tempTab tabMeja
				for j := 0; j < nMeja; j++ {
					tempTab[j] = M[j]
				}
				selectionSortKapasitasMejaAsc(&tempTab, nMeja)
				i = cariMeja(tempTab, nMeja, nilaiCari, 2, "binary")
				if i != -1 {
					// Cari indeks asli di M
					for j := 0; j < nMeja; j++ {
						if M[j].kapasitasKursi == tempTab[i].kapasitasKursi && M[j].nomor == tempTab[i].nomor {
							i = j
							break
						}
					}
				}
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
			if i != -1 {
				fmt.Println("Nomor Meja     :", M[i].nomor)
				fmt.Println("Kapasitas Kursi:", M[i].kapasitasKursi)
			} else {
				fmt.Println("Meja tidak ditemukan")
			}
		} else if pilihan == 11 {
			var sortPilihan int
			fmt.Println("Selection Sort Kapasitas Meja:")
			fmt.Println("1. Ascending (Kecil ke Besar)")
			fmt.Println("2. Descending (Besar ke Kecil)")
			fmt.Print("Pilihan: ")
			fmt.Scan(&sortPilihan)
			
			if sortPilihan == 1 {
				selectionSortKapasitasMejaAsc(&M, nMeja)
				fmt.Println("\nData meja berhasil diurutkan (Selection Sort - Kapasitas Kecil ke Besar)")
				tampilMeja(M, nMeja)
			} else if sortPilihan == 2 {
				selectionSortKapasitasMejaDesc(&M, nMeja)
				fmt.Println("\nData meja berhasil diurutkan (Selection Sort - Kapasitas Besar ke Kecil)")
				tampilMeja(M, nMeja)
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		} else if pilihan == 12 {
			var sortPilihan int
			fmt.Println("Insertion Sort Kapasitas Meja:")
			fmt.Println("1. Ascending (Kecil ke Besar)")
			fmt.Println("2. Descending (Besar ke Kecil)")
			fmt.Print("Pilihan: ")
			fmt.Scan(&sortPilihan)
			
			if sortPilihan == 1 {
				insertionSortKapasitasMejaAsc(&M, nMeja)
				fmt.Println("\nData meja berhasil diurutkan (Insertion Sort - Kapasitas Kecil ke Besar)")
				tampilMeja(M, nMeja)
			} else if sortPilihan == 2 {
				insertionSortKapasitasMejaDesc(&M, nMeja)
				fmt.Println("\nData meja berhasil diurutkan (Insertion Sort - Kapasitas Besar ke Kecil)")
				tampilMeja(M, nMeja)
			} else {
				fmt.Println("Pilihan tidak valid!")
			}
		} else if pilihan == 13 {
			tambahReservasi(&R, &nReservasi, M, nMeja, P, nPelanggan)
		} else if pilihan == 14 {
			var tanggal string
			fmt.Print("Masukkan tanggal (dd/mm/yyyy): ")
			fmt.Scan(&tanggal)
			statistikJumlah(R, M, nReservasi, nMeja, tanggal)
		} else if pilihan == 15 {
			statusKetersediaan(R, nReservasi, M, nMeja)
		} else if pilihan == 99 {
			fmt.Println("================================================================")
			fmt.Println("         Terima kasih telah menggunakan ReservaResto!")
			fmt.Println("================================================================")
		} else {
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

// ==================== POIN A: CRUD MEJA ====================

func tambahMeja(M *tabMeja, n *int) {
	var i, jumlah int
	fmt.Print("Jumlah meja yang ditambah: ")
	fmt.Scan(&jumlah)

	for i = 0; i < jumlah; i++ {
		fmt.Println("Data meja ke-", i+1)
		fmt.Print("Nomor meja: ")
		fmt.Scan(&M[*n].nomor)
		fmt.Print("Kapasitas kursi: ")
		fmt.Scan(&M[*n].kapasitasKursi)
		M[*n].count = 0
		*n = *n + 1
	}
	fmt.Println("Data meja berhasil ditambahkan!")
}

func tampilMeja(M tabMeja, n int) {
	var i int
	if n == 0 {
		fmt.Println("Belum ada data meja.")
	} else {
		fmt.Println("\n==================== DATA MEJA ====================")
		for i = 0; i < n; i++ {
			fmt.Printf("Meja %-7d | Kapasitas: %-7d kursi\n", M[i].nomor, M[i].kapasitasKursi)
		}
		fmt.Println("===================================================\n")
	}
}

func ubahMeja(M *tabMeja, n, ubah int) {
	var i int
	var kapasitasBaru, nomorBaru int
	var found bool = false

	for i = 0; i < n; i++ {
		if found == false && ubah == (*M)[i].nomor {
			fmt.Print("Masukkan nomor meja baru: ")
			fmt.Scan(&nomorBaru)
			fmt.Print("Masukkan kapasitas baru: ")
			fmt.Scan(&kapasitasBaru)
			(*M)[i].nomor = nomorBaru
			(*M)[i].kapasitasKursi = kapasitasBaru
			found = true
			fmt.Println("Data meja berhasil diubah")
		}
	}
	if found == false {
		fmt.Println("Meja tidak ditemukan")
	}
}

func hapusMeja(M *tabMeja, n *int, nomor int) {
	var i int
	var found bool = false
	i = 0
	for i < *n && found == false {
		if (*M)[i].nomor == nomor {
			found = true
		} else {
			i = i + 1
		}
	}
	if found == true {
		for i < *n-1 {
			(*M)[i] = (*M)[i+1]
			i = i + 1
		}
		*n = *n - 1
		fmt.Println("Data meja berhasil dihapus")
	} else {
		fmt.Println("Meja tidak ditemukan")
	}
}

// ==================== POIN A: CRUD PELANGGAN ====================

func tambahPelanggan(P *tabPelanggan, n *int) {
	var i, jumlah int
	fmt.Print("Jumlah pelanggan yang ditambah: ")
	fmt.Scan(&jumlah)

	for i = 0; i < jumlah; i++ {
		fmt.Println("Data pelanggan ke-", i+1)
		fmt.Print("ID pelanggan: ")
		fmt.Scan(&P[*n].id)
		fmt.Print("Nama: ")
		fmt.Scan(&P[*n].nama)
		fmt.Print("No. Telp: ")
		fmt.Scan(&P[*n].telp)
		*n = *n + 1
	}
	fmt.Println("Data pelanggan berhasil ditambahkan!")
}

func tampilPelanggan(P tabPelanggan, n int) {
	var i int
	if n == 0 {
		fmt.Println("Belum ada data pelanggan.")
	} else {
		fmt.Println("\n==================== DATA PELANGGAN ====================")
		for i = 0; i < n; i++ {
			fmt.Printf("ID: %d \t| Nama: %s \t| Telp: %s\n", P[i].id, P[i].nama, P[i].telp)
		}
		fmt.Println("========================================================\n")
	}
}

func ubahDataPelanggan(P *tabPelanggan, n, ubah int) {
	var i int
	var idBaru int
	var namaPelanggan string
	var telpBaru string
	var found bool = false

	for i = 0; i < n; i++ {
		if found == false && ubah == (*P)[i].id {
			fmt.Print("Masukkan ID pelanggan baru: ")
			fmt.Scan(&idBaru)
			fmt.Print("Masukkan nama pelanggan baru: ")
			fmt.Scan(&namaPelanggan)
			fmt.Print("Masukkan no. telp baru: ")
			fmt.Scan(&telpBaru)
			(*P)[i].id = idBaru
			(*P)[i].nama = namaPelanggan
			(*P)[i].telp = telpBaru
			found = true
			fmt.Println("Data pelanggan berhasil diubah")
		}
	}
	if found == false {
		fmt.Println("Data pelanggan tidak ditemukan")
	}
}

func hapusDataPelanggan(P *tabPelanggan, n *int, ID int) {
	var i int
	var found bool = false
	i = 0
	for i < *n && found == false {
		if (*P)[i].id == ID {
			found = true
		} else {
			i = i + 1
		}
	}
	if found == true {
		for i < *n-1 {
			(*P)[i] = (*P)[i+1]
			i = i + 1
		}
		*n = *n - 1
		fmt.Println("Data pelanggan berhasil dihapus")
	} else {
		fmt.Println("Data pelanggan tidak ditemukan")
	}
}

// ==================== POIN B: STATUS KETERSEDIAAN & JADWAL ====================

func checkStatus(meja Meja, R tabReservasi, nR int, tanggal string, jam int) string {
	var i int
	var status string = "O"
	for i = 0; i < nR; i++ {
		if R[i].nomorMeja == meja.nomor && R[i].tanggal == tanggal && R[i].jam == jam {
			status = "X"
		}
	}
	return status
}

func statusKetersediaan(R tabReservasi, nR int, M tabMeja, nM int) {
	var i, j int
	var tanggal string
	if nM == 0 {
		fmt.Println("Belum ada data meja.")
	} else {
		fmt.Print("\nMasukkan tanggal (dd/mm/yyyy): ")
		fmt.Scan(&tanggal)
		fmt.Println("\n======================= STATUS KETERSEDIAAN MEJA =======================")
		fmt.Printf("                         Tanggal: %s\n", tanggal)
		fmt.Println("------------------------------------------------------------------------")

		fmt.Printf("| %-8s |", "Meja\\Jam")
		for j = 10; j <= 21; j++ {
			fmt.Printf(" %2d |", j)
		}
		fmt.Println()

		fmt.Print("|----------|")
		for j = 10; j <= 21; j++ {
			fmt.Print("----|")
		}
		fmt.Println()

		for i = 0; i < nM; i++ {
			fmt.Printf("| %-8d |", M[i].nomor)
			for j = 10; j <= 21; j++ {
				fmt.Printf(" %2s |", checkStatus(M[i], R, nR, tanggal, j))
			}
			fmt.Println()
		}
		fmt.Println("========================================================================")
		fmt.Println("Keterangan: O = Tersedia, X = Sudah Dipesan")
		fmt.Println()
	}
}

// ==================== POIN C: SEARCHING (SATU FUNGSI) ====================

func cariMeja(M tabMeja, n int, nilaiCari int, berdasarkan int, metode string) int {
	var hasil int = -1

	if metode == "sequential" {
		var i int = 0
		for hasil == -1 && i < n {
			if berdasarkan == 1 {
				if M[i].nomor == nilaiCari {
					hasil = i
				}
			} else if berdasarkan == 2 {
				if M[i].kapasitasKursi == nilaiCari {
					hasil = i
				}
			}
			i = i + 1
		}
	} else if metode == "binary" {
		var left, right, mid int
		left = 0
		right = n - 1
		for left <= right && hasil == -1 {
			mid = (left + right) / 2
			if berdasarkan == 1 {
				if nilaiCari < M[mid].nomor {
					right = mid - 1
				} else if nilaiCari > M[mid].nomor {
					left = mid + 1
				} else {
					hasil = mid
				}
			} else if berdasarkan == 2 {
				if nilaiCari < M[mid].kapasitasKursi {
					right = mid - 1
				} else if nilaiCari > M[mid].kapasitasKursi {
					left = mid + 1
				} else {
					hasil = mid
				}
			}
		}
	}
	return hasil
}

// ==================== POIN D: SORTING ====================

func selectionSortKapasitasMejaAsc(M *tabMeja, n int) {
	var pass, i, idx int
	var temp Meja
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if (*M)[idx].kapasitasKursi > (*M)[i].kapasitasKursi {
				idx = i
			}
		}
		temp = (*M)[idx]
		(*M)[idx] = (*M)[pass-1]
		(*M)[pass-1] = temp
	}
}

func selectionSortKapasitasMejaDesc(M *tabMeja, n int) {
	var pass, i, idx int
	var temp Meja
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if (*M)[idx].kapasitasKursi < (*M)[i].kapasitasKursi {
				idx = i
			}
		}
		temp = (*M)[idx]
		(*M)[idx] = (*M)[pass-1]
		(*M)[pass-1] = temp
	}
}

func insertionSortKapasitasMejaAsc(M *tabMeja, n int) {
	var pass, i int
	var temp Meja
	pass = 1
	for pass <= n-1 {
		temp = (*M)[pass]
		i = pass
		for i > 0 && temp.kapasitasKursi < (*M)[i-1].kapasitasKursi {
			(*M)[i] = (*M)[i-1]
			i = i - 1
		}
		(*M)[i] = temp
		pass = pass + 1
	}
}

func insertionSortKapasitasMejaDesc(M *tabMeja, n int) {
	var pass, i int
	var temp Meja
	pass = 1
	for pass <= n-1 {
		temp = (*M)[pass]
		i = pass
		for i > 0 && temp.kapasitasKursi > (*M)[i-1].kapasitasKursi {
			(*M)[i] = (*M)[i-1]
			i = i - 1
		}
		(*M)[i] = temp
		pass = pass + 1
	}
}

func selectionSortMejaByNomorAsc(M *tabMeja, n int) {
	var pass, i, idx int
	var temp Meja
	for pass = 1; pass < n; pass++ {
		idx = pass - 1
		for i = pass; i < n; i++ {
			if (*M)[idx].nomor > (*M)[i].nomor {
				idx = i
			}
		}
		temp = (*M)[idx]
		(*M)[idx] = (*M)[pass-1]
		(*M)[pass-1] = temp
	}
}

// ==================== POIN E: STATISTIK ====================

func findMaxCountMeja(M tabMeja, m int) int {
	var i, idx int
	if m <= 0 {
		idx = -1
	} else {
		idx = 0
		for i = 1; i < m; i++ {
			if M[i].count > M[idx].count {
				idx = i
			}
		}
	}
	return idx
}

func statistikJumlah(R tabReservasi, M tabMeja, n, m int, tanggal string) {
	var i, j, idx int
	var totalReservasi int = 0

	for j = 0; j < m; j++ {
		M[j].count = 0
	}

	for i = 0; i < n; i++ {
		if R[i].tanggal == tanggal {
			totalReservasi = totalReservasi + 1
			idx = -1
			for j = 0; j < m; j++ {
				if M[j].nomor == R[i].nomorMeja {
					idx = j
					break
				}
			}
			if idx != -1 {
				M[idx].count = M[idx].count + 1
			}
		}
	}

	fmt.Println("\n==================== STATISTIK RESERVASI ====================")
	fmt.Printf("Tanggal %s\n", tanggal)
	fmt.Printf("Total reservasi: %d\n", totalReservasi)

	if totalReservasi == 0 {
		fmt.Println("Belum ada reservasi pada tanggal tersebut.")
	} else {
		if m > 0 {
			idx = findMaxCountMeja(M, m)
			if idx != -1 && M[idx].count > 0 {
				fmt.Printf("\n--- MEJA PALING SERING DIPESAN ---\n")
				fmt.Printf("Meja nomor %d dipesan sebanyak %d kali\n", M[idx].nomor, M[idx].count)
			}
		}
	}
	fmt.Println("=============================================================\n")
}

// ==================== FITUR RESERVASI ====================

func cariPelangganByID(P tabPelanggan, n int, id int) int {
	var i int
	var hasil int = -1
	for i = 0; i < n; i++ {
		if P[i].id == id {
			hasil = i
		}
	}
	return hasil
}

func reservasiBentrok(R tabReservasi, nR int, nomorMeja int, tanggal string, jam int) bool {
	var i int
	var bentrok bool = false
	for i = 0; i < nR; i++ {
		if R[i].nomorMeja == nomorMeja && R[i].tanggal == tanggal && R[i].jam == jam {
			bentrok = true
		}
	}
	return bentrok
}

func tambahReservasi(R *tabReservasi, nR *int, M tabMeja, nM int, P tabPelanggan, nP int) {
	var idPelanggan, nomorMeja, jam int
	var tanggal string

	fmt.Println("\n==================== TAMBAH RESERVASI ====================")
	fmt.Print("ID Reservasi: ")
	fmt.Scan(&(*R)[*nR].idReservasi)
	fmt.Print("ID Pelanggan: ")
	fmt.Scan(&idPelanggan)

	if cariPelangganByID(P, nP, idPelanggan) == -1 {
		fmt.Println("Error: Pelanggan tidak ditemukan!")
	} else {
		fmt.Print("Nomor Meja: ")
		fmt.Scan(&nomorMeja)

		idxMeja := -1
		for i := 0; i < nM; i++ {
			if M[i].nomor == nomorMeja {
				idxMeja = i
			}
		}
		if idxMeja == -1 {
			fmt.Println("Error: Meja tidak ditemukan!")
		} else {
			fmt.Print("Tanggal (dd/mm/yyyy): ")
			fmt.Scan(&tanggal)
			fmt.Print("Jam (10-21): ")
			fmt.Scan(&jam)

			if jam < 10 || jam > 21 {
				fmt.Println("Error: Jam operasional restoran hanya dari jam 10-21!")
			} else if reservasiBentrok(*R, *nR, nomorMeja, tanggal, jam) == true {
				fmt.Println("Error: Reservasi bentrok! Meja sudah dipesan pada tanggal dan jam tersebut.")
			} else {
				(*R)[*nR].idPelanggan = idPelanggan
				(*R)[*nR].nomorMeja = nomorMeja
				(*R)[*nR].tanggal = tanggal
				(*R)[*nR].jam = jam
				*nR = *nR + 1
				fmt.Println("Reservasi berhasil ditambahkan!")
			}
		}
	}
	fmt.Println("==========================================================\n")
}