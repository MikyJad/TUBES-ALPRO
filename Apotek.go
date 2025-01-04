package main

import (
	"fmt"
)

const NMAX int = 20

var saveObat = 0
var z = 0

type dataobat struct {
	nama, bulan, tahun, kategori, pabrikan, kode string
	harga                                        float64
	jumlah                                       int
}

type data2 struct {
	nama, bulan, tahun, kategori, pabrikan, kode string
	harga                                        float64
	jumlah                                       int 
}

// type data3 struct {
// 	nama, bulan, tahun, kategori, pabrikan, kode string
// 	harga                                        float64
// 	jumlah                                       int 
// }

type tabInt [NMAX]dataobat

type tabData [NMAX]data2

// type tabData2 [NMAX]data3

func main() {
	var obat tabInt
	var obat2 tabData
	var x int

	fmt.Println()
	fmt.Println("=====================================")
	fmt.Println("||  Selamat Datang di FisikaFarma  ||")
	menu()
	fmt.Println("=====================================")
	inputMenu(&obat, &obat2, x)
}

func menu() {
	fmt.Println("||  1. Cari Obat                   ||")
	fmt.Println("||  2. Mode Admin                  ||")
	fmt.Println("||  3. Pemesanan Obat              ||")
	fmt.Println("||  4. Exit                        ||")
}

func inputMenu(obat *tabInt, obat2 *tabData, x int) {
	var namaobat string
	var totalharga float64
	totalharga = 0

	fmt.Print("-> ")
	fmt.Scan(&x)
	fmt.Println()

	for x != 4 {
		for x < 1 || x > 4{
			fmt.Println("MASUKKAN ANGKA YANG TERSEDIA!")
			fmt.Print("-> ")
			fmt.Scan(&x)
			fmt.Scanln()
		}
		if x == 1 {
			cariobat(obat, namaobat)
		} else if x == 2 {
			modeadmin(obat, &saveObat)
		} else if x == 3 {
			pesanobat(obat, obat2, saveObat, &totalharga)
		} 
		fmt.Println("")
		fmt.Println("=====================================")
		fmt.Println("||  Selamat Datang di FisikaFarma  ||")
		fmt.Println("||  1. Cari Obat                   ||")
		fmt.Println("||  2. Mode Admin                  ||")
		fmt.Println("||  3. Pemesanan Obat              ||")
		fmt.Println("||  4. Exit                        ||")
		fmt.Println("=====================================")
		fmt.Print("-> ")
		fmt.Scan(&x)
		fmt.Println()
	}

}

func cariobat(obat *tabInt, namaobat string) {
	var lanjut string
	var found bool = false

	for {
		fmt.Println()
		fmt.Println("Masukkan Nama atau Kode Obat")
		fmt.Print("-> ")
		fmt.Scan(&namaobat)
		fmt.Println()
		fmt.Println()
		for i := 0; i < len(*obat) && !found; i++ {
			if (*obat)[i].nama == namaobat || (*obat)[i].kode == namaobat {
				fmt.Println("Obat ditemukan:")
				fmt.Println("Nama     :", (*obat)[i].nama)
				fmt.Println("Kode     :", (*obat)[i].kode)
				fmt.Println("Kategori :", (*obat)[i].kategori)
				fmt.Println("Pabrikan :", (*obat)[i].pabrikan)
				fmt.Println("Harga    :", (*obat)[i].harga)
				fmt.Println("Jumlah   :", (*obat)[i].jumlah)
				fmt.Println("Bulan    :", (*obat)[i].bulan)
				fmt.Println("Tahun    :", (*obat)[i].tahun)
				found = true
			}
		}
		if !found {
			fmt.Println("Maaf, Obat tidak ditemukan.")
		}
		found = false
		fmt.Println()
		fmt.Println("Ingin mencari obat lagi? (ya/tidak)")
		fmt.Print("-> ")
		fmt.Scan(&lanjut)
		fmt.Println()
		for lanjut != "ya" && lanjut != "Ya" && lanjut != "tidak" && lanjut != "Tidak" {
			fmt.Println("Masukkan tidak valid")
			fmt.Print("-> ")
			fmt.Scan(&lanjut)
			fmt.Println()
		}

		if lanjut != "ya" {
			break
		}
	}
}

func modeadmin(obat *tabInt, i *int) {
	var lagi, lagi2 string
	var pil1, n, edit, edit2 int

	fmt.Println()
	fmt.Println("=================")
	fmt.Println("||  1. Tambah  ||")
	fmt.Println("||  2. Hapus   ||")
	fmt.Println("||  3. Edit    ||")
	fmt.Println("||  4. Back    ||")
	fmt.Println("=================")
	fmt.Print("-> ")
	fmt.Scan(&pil1)
	fmt.Println()


	
	for pil1 < 1 || pil1 > 4 {
		fmt.Println("Masukkan angka yang benar")
		fmt.Print("-> ")
		fmt.Scan(&pil1)
		fmt.Println()
	}

	if pil1 == 1 {
		for {
			fmt.Println()
			fmt.Println("Masukkan Data Obat")
			fmt.Println()

			fmt.Print("Nama Obat         : ")
			fmt.Scan(&(*obat)[*i].nama)

			fmt.Print("Kode Obat         : ")
			fmt.Scan(&(*obat)[*i].kode)

			fmt.Print("Kategori Obat     : ")
			fmt.Scan(&(*obat)[*i].kategori)

			fmt.Print("Pabrikan Obat     : ")
			fmt.Scan(&(*obat)[*i].pabrikan)

			fmt.Print("Harga Obat        : ")
			fmt.Scan(&(*obat)[*i].harga)

			fmt.Print("Jumlah Obat       : ")
			fmt.Scan(&(*obat)[*i].jumlah)
			
			fmt.Print("Bulan Kadaluwarsa : ")
			fmt.Scan(&(*obat)[*i].bulan)
			
			fmt.Print("Tahun Kadaluwarsa : ")
			fmt.Scan(&(*obat)[*i].tahun)

			fmt.Println()
			*i++
			fmt.Println("Masukkan data lagi? (ya/tidak)")
			fmt.Print("-> ")
			fmt.Scan(&lagi)
			fmt.Println()

			for lagi != "ya" && lagi != "Ya" && lagi != "tidak" && lagi != "Tidak" {
				fmt.Println("Masukkan tidak valid")
				fmt.Print("-> ")
				fmt.Scan(&lagi)
				fmt.Println()
			}

			if lagi != "ya" || *i >= NMAX {
				break
			}
		}
	} else if pil1 == 2 {
		if *i == 0 {
			fmt.Println("Tidak ada data yang bisa dihapus.")
			return
		}
		for {
			cetakobat(obat, *i)
			fmt.Print("Masukkan index data yang akan dihapus : ")
			fmt.Scan(&n)
			fmt.Println()
			if n < 1 || n > *i {
				fmt.Println("Nomor data tidak valid.")
				continue
			}
			for k := n - 1; k < *i-1; k++ {
				(*obat)[k] = (*obat)[k+1]
			}
			*i--
			fmt.Println("Data berhasil dihapus.")
			fmt.Println("Hapus data lagi? (ya/tidak)")
			fmt.Print("-> ")
			fmt.Scan(&lagi)
			fmt.Println()

			for lagi != "ya" && lagi != "Ya" && lagi != "tidak" && lagi != "Tidak" {
				fmt.Println("Masukkan tidak valid")
				fmt.Print("-> ")
				fmt.Scan(&lagi)
				fmt.Println()
			}

			if (lagi == "ya" || lagi == "Ya") && *i == 0 {
				fmt.Println("Tidak ada data yang bisa dihapus")
			}

			if lagi != "ya" && lagi != "Ya" || *i == 0 {
				break
			}
		}
	} else if pil1 == 3 {
		var a, b, c, d, g, h string
		var e float64
		var f int

		if *i == 0 {
			fmt.Println("Data tidak ada")
		} else {
			for lagi2 != "tidak" && lagi2 != "Tidak" {
				fmt.Println()
				cetakobat(obat, *i)
				fmt.Print("Masukkan index data yang akan diedit : ")
				fmt.Print("-> ")
				fmt.Scan(&edit)
				fmt.Println()

				for edit < 1 || edit > *i {
					fmt.Println("Data tidak ada")
					fmt.Print("-> ")
					fmt.Scan(&edit)
					fmt.Println()
				}

				edit2 = edit - 1
				fmt.Println()
	
				fmt.Println("Data yang akan diedit :")
				fmt.Println("Nama     :", obat[edit2].nama)
				fmt.Println("Kode     :", obat[edit2].kode)
				fmt.Println("Kategori :", obat[edit2].kategori)
				fmt.Println("Pabrikan :", obat[edit2].pabrikan)
				fmt.Println("Harga    :", obat[edit2].harga)
				fmt.Println("Jumlah   :", obat[edit2].jumlah)
				fmt.Println("Bulan    :", obat[edit2].bulan)
				fmt.Println("Tahun    :", obat[edit2].tahun)
				
				fmt.Print("Nama       : ")
				fmt.Scan(&a)
	
				fmt.Print("Kode       : ")
				fmt.Scan(&b)
	
				fmt.Print("Kategori   : ")
				fmt.Scan(&c)
	
				fmt.Print("Pabrikan   : ")
				fmt.Scan(&d)
	
				fmt.Print("Harga      : ")
				fmt.Scan(&e)
	
				fmt.Print("Jumlah     : ")
				fmt.Scan(&f)
				
				fmt.Print("Bulan      : ")
				fmt.Scan(&g)
				
				fmt.Print("Tahun      : ")
				fmt.Scan(&h)
				
				fmt.Println()
				obat[edit].nama = a
				obat[edit].kode = b
				obat[edit].kategori = c
				obat[edit].pabrikan = d
				obat[edit].harga = e
				obat[edit].jumlah = f
				obat[edit].bulan = g
				obat[edit].tahun = h
	
				fmt.Println("Edit data lagi?")
				fmt.Print("-> ")
				fmt.Scan(&lagi2)
				fmt.Println()

				for lagi2 != "ya" && lagi2 != "Ya" && lagi2 != "tidak" && lagi2 != "Tidak" {
					fmt.Println("Masukkan tidak valid")
					fmt.Print("-> ")
					fmt.Scan(&lagi2)
					fmt.Println()
				}
			}
		}
	} else if pil1 == 4 {
		
	}
}

func cetakobat(obat *tabInt, j int) {
	for i := 0; i < j; i++ {
		fmt.Println(i+1, ".", obat[i].nama, obat[i].kategori, obat[i].pabrikan, obat[i].harga, obat[i].jumlah, obat[i].bulan, obat[i].tahun)
	}
}

func pesanobat(obat *tabInt, obat2 *tabData, j int, totalharga *float64) {
	var x int
	var bayar, memberType string

	if z == 0 {
		*totalharga = 0
	}

	fmt.Println()
	fmt.Println("KATALOG OBAT APOTEK FISIKA FARMA")
	fmt.Println()
	for i := 0; i <= j-1; i++ {
		fmt.Println(i+1, ". ", obat[i].nama, obat[i].kategori, obat[i].pabrikan, obat[i].harga, obat[i].jumlah, obat[i].bulan, obat[i].tahun)
	}
	fmt.Println()
	fmt.Println("Pilih obat dengan memasukkan nomor obat")

	fmt.Scan(&x)
	for x != 0 {
		for x < 1 || x > j {
			fmt.Println("Mohon masukkan angka yang benar")
			fmt.Scan(&x)
		}
		obat[x-1].jumlah = obat[x-1].jumlah - 1

		for l := z; l > 0; l-- {
			if obat2[z] == obat2[l] {
				obat2[z].jumlah = obat2[z].jumlah + 1
				z = z - 1
			}
		}

		obat2[z].nama = obat[x-1].nama
		obat2[z].kategori = obat[x-1].kategori
		obat2[z].pabrikan = obat[x-1].pabrikan
		obat2[z].harga = obat[x-1].harga
		obat2[z].jumlah = obat2[z].jumlah + 1
		obat2[z].bulan = obat[x-1].bulan
		obat2[z].tahun = obat[x-1].tahun
		
		
		z++
		*totalharga = *totalharga + obat[x-1].harga

		fmt.Scan(&x)
	}

	fmt.Println()
	fmt.Println("Apakah Anda member? (ya/tidak)")
	fmt.Print("-> ")
	fmt.Scan(&memberType)
	fmt.Println()

	for memberType != "ya" && memberType != "Ya" && memberType != "tidak" && memberType != "Tidak" {
		fmt.Println("Masukkan tidak valid")
		fmt.Print("-> ")
		fmt.Scan(&memberType)
		fmt.Println()
	}

	if memberType == "ya" || memberType == "Ya" {
		fmt.Print("Masukkan tipe member (gold/silver) : ")
		fmt.Scan(&memberType)
		fmt.Println()

	for memberType != "silver" && memberType != "Silver" && memberType != "gold" && memberType != "Gold" {
		fmt.Println("Masukkan tidak valid")
		fmt.Print("-> ")
		fmt.Scan(&memberType)
		fmt.Println()
	}
	}

	fmt.Println()
	fmt.Println("Lanjutkan pembayaran?")
	fmt.Print("-> ")
	fmt.Scan(&bayar)
	fmt.Println()

	for bayar != "ya" && bayar != "Ya" && bayar != "tidak" && bayar != "Tidak" {
		fmt.Println("Masukkan tidak valid")
		fmt.Print("-> ")
		fmt.Scan(&bayar)
		fmt.Println()
	}

	if bayar == "ya" || bayar == "Ya" {
		checkout(obat, obat2, *totalharga, memberType)
	} else {
		fmt.Println()
		fmt.Println("Pilih obat dengan memasukkan nomor obat")
		fmt.Scan(&x)
		for x != 0 {
			obat[x-1].jumlah = obat[x-1].jumlah - 1

			obat2[z].nama = obat[x-1].nama
			obat2[z].kategori = obat[x-1].kategori
			obat2[z].pabrikan = obat[x-1].pabrikan
			obat2[z].harga = obat[x-1].harga
			obat2[z].jumlah = obat2[z].jumlah + 1
			obat2[z].bulan = obat[x-1].bulan
			obat2[z].tahun = obat[x-1].tahun
			z++

			*totalharga = *totalharga + obat[x-1].harga
			fmt.Scan(&x)
		}
		checkout(obat, obat2, *totalharga, memberType)
	}
}

func checkout(obat *tabInt, obat2 *tabData, totalharga float64, memberType string) {
	var ambil, pay, urut, urut2 int
	var nohp, nobank string
	var ongkir float64

	fmt.Println()
	menuch()
	fmt.Print("-> ")
	fmt.Scan(&ambil)
	fmt.Println()
	if ambil < 1 || ambil > 2 {
		for ambil < 1 || ambil > 2 {
			fmt.Println("Masukkan tidak valid")
			fmt.Print("-> ")
			fmt.Scan(&ambil)
			fmt.Println()
		}
	}
	if ambil == 1 {
		ongkir = 0
		totalharga = totalharga + ongkir
	} else if ambil == 2 {
		ongkir = totalharga * 0.05 
		totalharga = totalharga + ongkir
	}

	fmt.Println()
	discountedPrice := calculateDiscount(memberType, totalharga)
	receipt(obat2, discountedPrice, ongkir)

	fmt.Println()
	fmt.Println("1. Mengurutkan berdasarkan harga")
	fmt.Println("2. Next")
	fmt.Print("-> ")
	fmt.Scan(&urut)
	fmt.Println()
	for urut < 1 || urut > 2 {
		fmt.Println("Masukkan tidak valid")
		fmt.Print("-> ")
		fmt.Scan(&urut)
		fmt.Println()
	}
	if urut == 1 {
		fmt.Println("1. Menaik")
		fmt.Println("2. Menurun")
		fmt.Println()

		fmt.Print("Urut : ")
		fmt.Scan(&urut2)

		for urut2 != 3 {
			for urut2 < 1 || urut2 > 3 {
				fmt.Println("Masukkan angka yang benar!")
				fmt.Print("Urut : ")
				fmt.Scan(&urut2)
			}
			if urut2 == 1 {
				fmt.Println()
				sortasc(obat2)
				fmt.Println()
			} else if urut2 == 2 {
				fmt.Println()
				sortdsc(obat2)
				fmt.Println()
			} else if urut2 == 3 {

			}
			fmt.Println()
			fmt.Print("Urut : ")
			fmt.Scan(&urut2)
		}
		fmt.Println()
	} else if urut == 2 {
		fmt.Println()
	}

	fmt.Println()
	menupay()
	fmt.Print("-> ")
	fmt.Scan(&pay)
	fmt.Println()
	if pay < 1 || pay > 3 {
		for pay < 1 || pay > 3 {
			fmt.Println("Masukkan tidak valid")
			fmt.Print("-> ")
			fmt.Scan(&pay)
			fmt.Println()
		}
	}
	if pay == 1 {
		fmt.Println("Masukkan nomor Handphone yang tersambung dengan E-Wallet : ")
		fmt.Scan(&nohp)
	} else if pay == 2 {
		fmt.Print("Masukkan nomor rekening : ")
		fmt.Scan(&nobank)
	} else if pay == 3 {
		fmt.Println("Tolong siapkan uang yang sesuai dengan total harga")
	}
	fmt.Println()
	fmt.Println("Terimakasih atas pemesanan anda")
	
	z = 0
	obat2[z].jumlah = 0
}

func menuch() {
	fmt.Println("Opsi Pengambilan: ")
	fmt.Println("1. Ambil di Apotek")
	fmt.Println("2. Antar")
}

func menupay() {
	fmt.Println("Metode Pembayaran :")
	fmt.Println("1. E-Wallet")
	fmt.Println("2. Bank")
	fmt.Println("3. Cash")
}

func calculateDiscount(memberType string, totalharga float64) float64 {
	var discount float64
	switch memberType {
	case "gold":
		discount = 0.20
	case "silver":
		discount = 0.10
	default:
		discount = 0
	}
	return totalharga - (totalharga * discount)
}

func receipt(obat2 *tabData, totalharga, ongkir float64) {
	fmt.Printf("%20s %6s %6s\n", "Obat", "Jumlah", "Harga")
	for i := 0; i < z; i++ {
		fmt.Printf("%20s %6d %6.2f\n", obat2[i].nama, obat2[i].jumlah, obat2[i].harga)
	}
	fmt.Println()
	fmt.Printf("%30s %6.2f\n", "Ongkir : ", ongkir)
	fmt.Printf("%30s %6.2f\n", "Total Harga : ", totalharga)
}

func sortdsc(obat2 *tabData) {
	var i, j, idxMax int
	var temp data2
	for i = 0; i < z-1; i++ {
		idxMax = i
		for j = i + 1; j < z; j++ {
			if obat2[j].harga > obat2[idxMax].harga {
				idxMax = j
			}
		}
		temp = obat2[i]
		obat2[i] = obat2[idxMax]
		obat2[idxMax] = temp
	}
	fmt.Printf("%20s %6s %6s\n", "Obat", "Jumlah", "Harga")
	for k := 0; k < z; k++ {
		fmt.Printf("%20s %6d %6.2f\n", obat2[k].nama, obat2[k].jumlah, obat2[k].harga)
	}

}

func sortasc(obat2 *tabData) {
	var i, j, idxMax int
	var temp data2
	for i = 0; i < z-1; i++ {
		idxMax = i
		for j = i + 1; j < z; j++ {
			if obat2[j].harga < obat2[idxMax].harga {
				idxMax = j
			}
		}
		temp = obat2[i]
		obat2[i] = obat2[idxMax]
		obat2[idxMax] = temp
	}
	fmt.Printf("%20s %6s %6s\n", "Obat", "Jumlah", "Harga")
	for k := 0; k < z; k++ {
		fmt.Printf("%20s %6d %6.2f\n", obat2[k].nama, obat2[k].jumlah, obat2[k].harga)
	}

}