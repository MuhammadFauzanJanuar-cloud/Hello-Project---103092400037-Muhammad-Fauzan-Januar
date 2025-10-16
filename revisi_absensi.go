package main

import "fmt"

type RekapKehadiran struct {
	Nama       string
	Hadir      int
	Total      int
	Persentase float64
}

type Absensi struct {
	NamaSiswa string
	Tanggal   string
	Hadir     bool
}

var dataAbsensi [1000]Absensi
var jumlahData int
var rekapKehadiran [100]RekapKehadiran
var jumlahRekap int

func tambahAbsensi(data *[1000]Absensi, jumlah *int) {
	var nama, tanggal, hadir string

	fmt.Print("Masukkan nama siswa: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan tanggal (yyyy-mm-dd): ")
	fmt.Scanln(&tanggal)

	fmt.Print("Apakah hadir? (y/n): ")
	fmt.Scanln(&hadir)

	status := false
	if hadir == "y" || hadir == "Y" {
		status = true
	}

	dataAbsensi[*jumlah] = Absensi{NamaSiswa: nama, Tanggal: tanggal, Hadir: status}
	*jumlah++
	fmt.Println("Data absensi ditambahkan.")
}

func ubahAbsensi(data *[1000]Absensi, jumlah int) {
	var nama, tanggal, hadir string

	fmt.Print("Masukkan nama siswa yang ingin diubah: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan tanggal (yyyy-mm-dd): ")
	fmt.Scanln(&tanggal)

	fmt.Print("Apakah hadir? (y/n): ")
	fmt.Scanln(&hadir)

	ubah := false
	for i := 0; i < jumlahData; i++ {
		if dataAbsensi[i].NamaSiswa == nama && dataAbsensi[i].Tanggal == tanggal {
			if hadir == "y" || hadir == "Y" {
				dataAbsensi[i].Hadir = true
			} else {
				dataAbsensi[i].Hadir = false
			}
			ubah = true
			break
		}
	}

	if ubah {
		fmt.Println("Data absensi diubah.")
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
}

func hapusAbsensi(data *[1000]Absensi, jumlah *int) {
	var nama, tanggal string

	fmt.Print("Masukkan nama siswa yang ingin dihapus: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan tanggal (yyyy-mm-dd): ")
	fmt.Scanln(&tanggal)

	for i := 0; i < *jumlah; i++ {
		if dataAbsensi[i].NamaSiswa == nama && dataAbsensi[i].Tanggal == tanggal {
			for j := i; j < *jumlah-1; j++ {
				dataAbsensi[j] = dataAbsensi[j+1]
			}
			*jumlah--
			fmt.Println("Data absensi dihapus.")
			return
		}
	}

	fmt.Println("Data tidak ditemukan.")
}

func cariAbsensi(data *[1000]Absensi, jumlah int) {
	var nama, tanggal string

	fmt.Print("Masukkan nama siswa: ")
	fmt.Scanln(&nama)

	fmt.Print("Masukkan tanggal (yyyy-mm-dd): ")
	fmt.Scanln(&tanggal)

	ditemukan := false
	for i := 0; i < jumlahData; i++ {
		if dataAbsensi[i].NamaSiswa == nama && dataAbsensi[i].Tanggal == tanggal {
			status := "Tidak Hadir"
			if dataAbsensi[i].Hadir {
				status = "Hadir"
			}
			fmt.Printf("Status pada %s: %s\n", tanggal, status)
			ditemukan = true
			break
		}
	}

	if !ditemukan {
		fmt.Println("Data tidak ditemukan.")
	}
}

func statistikKehadiran(data *[1000]Absensi, jumlahData int, rekap *[100]RekapKehadiran, jumlahRekap *int) {

	*jumlahRekap = 0

	for i := 0; i < jumlahData; i++ {
		nama := dataAbsensi[i].NamaSiswa
		ditemukan := false
		indeksRekap := -1

		for j := 0; j < *jumlahRekap; j++ {
			if rekapKehadiran[j].Nama == nama {
				ditemukan = true
				indeksRekap = j
				break
			}
		}

		if ditemukan {
			rekapKehadiran[indeksRekap].Total++
			if dataAbsensi[i].Hadir {
				rekapKehadiran[indeksRekap].Hadir++
			}
		} else {
			rekap[*jumlahRekap].Nama = nama
			rekap[*jumlahRekap].Total = 1
			if data[i].Hadir {
				rekap[*jumlahRekap].Hadir = 1
			} else {
				rekap[*jumlahRekap].Hadir = 0
			}
			*jumlahRekap++
		}
	}

	fmt.Println("\nStatistik Kehadiran:")
	for i := 0; i < *jumlahRekap; i++ {
		rekapKehadiran[i].Persentase = float64(rekapKehadiran[i].Hadir) / float64(rekapKehadiran[i].Total) * 100
		fmt.Printf("- %s: %d hadir dari %d hari (%.2f%%)\n", rekapKehadiran[i].Nama, rekapKehadiran[i].Hadir, rekapKehadiran[i].Total, rekapKehadiran[i].Persentase)
	}
}

func urutkanKehadiran(data *[1000]Absensi, jumlahData int, rekap *[100]RekapKehadiran, jumlahRekap *int) {

	statistikKehadiran(data, jumlahData, rekap, jumlahRekap)

	for i := 0; i < *jumlahRekap-1; i++ {
		for j := 0; j < *jumlahRekap-i-1; j++ {

			if rekapKehadiran[j].Persentase < rekapKehadiran[j+1].Persentase {
				rekapKehadiran[j], rekapKehadiran[j+1] = rekapKehadiran[j+1], rekapKehadiran[j]
			}
		}
	}

	fmt.Println("\nPengurutan Berdasarkan Persentase Kehadiran:")
	for i := 0; i < *jumlahRekap; i++ {
		fmt.Printf("- %s: %.2f%%\n", rekapKehadiran[i].Nama, rekapKehadiran[i].Persentase)
	}
}

func binarySearchNama(data [100]RekapKehadiran, jumlah int, target string) int {
	low := 0
	high := jumlah - 1

	for low <= high {
		mid := (low + high) / 2
		if data[mid].Nama == target {
			return mid
		} else if data[mid].Nama < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func insertionSortPersentase(data *[100]RekapKehadiran, jumlah int) {
	for i := 1; i < jumlah; i++ {
		temp := data[i]
		j := i - 1
		for j >= 0 && data[j].Persentase < temp.Persentase {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = temp
	}
}

func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Tambah Absensi")
		fmt.Println("2. Ubah Absensi")
		fmt.Println("3. Hapus Absensi")
		fmt.Println("4. Cari Absensi")
		fmt.Println("5. Statistik Kehadiran")
		fmt.Println("6. Urutkan Kehadiran")
		fmt.Println("7. Tampilkan dan Cari Rekap Kehadiran")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		if pilihan == 1 {
			tambahAbsensi(&dataAbsensi, &jumlahData)
		} else if pilihan == 2 {
			ubahAbsensi(&dataAbsensi, jumlahData)
		} else if pilihan == 3 {
			hapusAbsensi(&dataAbsensi, &jumlahData)
		} else if pilihan == 4 {
			cariAbsensi(&dataAbsensi, jumlahData)
		} else if pilihan == 5 {
			statistikKehadiran(&dataAbsensi, jumlahData, &rekapKehadiran, &jumlahRekap)
		} else if pilihan == 6 {
			urutkanKehadiran(&dataAbsensi, jumlahData, &rekapKehadiran, &jumlahRekap)
		} else if pilihan == 7 {
			var rekap [100]RekapKehadiran
			jumlah := 0

			for i := 0; i < jumlahData; i++ {
				nama := dataAbsensi[i].NamaSiswa
				indeks := -1
				for j := 0; j < jumlah; j++ {
					if rekap[j].Nama == nama {
						indeks = j
						return
					}
				}

				if indeks == -1 {
					rekap[jumlah].Nama = nama
					if dataAbsensi[i].Hadir {
						rekap[jumlah].Hadir = 1
					}
					rekap[jumlah].Total = 1
					jumlah++
				} else {
					if dataAbsensi[i].Hadir {
						rekap[indeks].Hadir++
					}
					rekap[indeks].Total++
				}
			}

			for i := 0; i < jumlah; i++ {
				if rekap[i].Total > 0 {
					rekap[i].Persentase = float64(rekap[i].Hadir) / float64(rekap[i].Total) * 100
				}
			}

			insertionSortPersentase(&rekap, jumlah)

			fmt.Println("\nRekap Kehadiran (Diurutkan Descending):")
			for i := 0; i < jumlah; i++ {
				fmt.Printf("- %s: %.2f%%\n", rekap[i].Nama, rekap[i].Persentase)
			}

			for i := 0; i < jumlah-1; i++ {
				for j := 0; j < jumlah-i-1; j++ {
					if rekap[j].Nama > rekap[j+1].Nama {
						rekap[j], rekap[j+1] = rekap[j+1], rekap[j]
					}
				}
			}

			var target string
			fmt.Print("\nMasukkan nama mahasiswa yang ingin dicari: ")
			fmt.Scanln(&target)

			index := binarySearchNama(rekap, jumlah, target)
			if index != -1 {
				fmt.Printf("Ditemukan: %s - Hadir: %d, Total: %d, Persentase: %.2f%%\n",
					rekap[index].Nama, rekap[index].Hadir, rekap[index].Total, rekap[index].Persentase)
			} else {
				fmt.Println("Nama tidak ditemukan.")
			}

		} else if pilihan == 0 {
			fmt.Println("Terima kasih.")
			return
		} else {
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
