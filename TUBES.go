package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type akun struct {
	username string
	password string
}

type books struct {
	id_book  int
	judul    string
	kategori string
	penulis  string
	tahun    int
	pinjam   peminjaman
}
type peminjaman struct {
	pinjam         bool
	jumPinjam      int
	tarif          int
	denda          int
	tanggal        time.Time
	tanggalkembali time.Time
	waktu          int
}

type book [1024]books

func main() {
	fmt.Printf("\x1bc")
	var ac akun
	var jum int
	var buku book
	var p peminjaman
	buku[0] = books{50, "Matahari", "Novel", "Tere Liye", 2016, p}
	buku[1] = books{21, "A Study in Scarlet", "Novel", "Arthur Conan Doyle", 1887, p}
	buku[2] = books{26, "Sherlock Holmes", "Novel", "Sherlock Holmes", 1892, p}
	buku[3] = books{10, "Hunter X Hunter", "Komik", "Yoshihiro Togashi", 1998, p}
	buku[4] = books{1, "The Theory of Everything", "Ilmiah", "Stephen Hawking", 2002, p}
	jum = 5
	ac.username = "admin"
	ac.password = "admin"
	var pilihan int
	fmt.Println("Selamat Datang di Perpustakaan")
	fmt.Println("1. Login")
	fmt.Println("2. Ganti Password")
	fmt.Println("3. Exit")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		login(ac, jum, buku)
	case 2:
		register(&ac)
	case 3:
		exit()
	}
}

func login(ac akun, jum int, buku book) {
	var username, password string
	fmt.Printf("\x1bc")
	fmt.Println("Login")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	if ac.username == username && ac.password == password {
		fmt.Printf("\x1bc")
		fmt.Println("Login Berhasil")
		fmt.Scanln()
		menu_utama(jum, buku)
	} else {
		fmt.Printf("\x1bc")
		fmt.Println("Login Gagal")
		fmt.Scanln(jum, buku)
		main()
	}
}

func register(ac *akun) {
	var username, password string
	fmt.Println("Ganti Password")
	fmt.Print("Username: ")
	fmt.Scanln(&username)
	fmt.Print("Password: ")
	fmt.Scanln(&password)
	fmt.Printf("\x1bc")
	ac.username = username
	ac.password = password
	fmt.Println("Pergantian Berhasil")
	fmt.Scanln()
	main()
}

func exit() {
	fmt.Printf("\x1bc")
	fmt.Println("Terima Kasih")
	os.Exit(0)
}
func baca() string {
	reader := bufio.NewReader(os.Stdin)
	nama, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	nama = strings.TrimSpace(nama)
	return nama
}

func menu_utama(jum int, buku book) {
	var pilihan int
	fmt.Println("Menu Utama")
	fmt.Println("1. Daftar Buku")
	fmt.Println("2. Manajemen Buku")
	fmt.Println("3. Manajemen Peminjaman")
	fmt.Println("4. Cari buku")
	fmt.Println("5. Lihat Tarif Peminjaman dan Denda")
	fmt.Println("10. Logout")
	fmt.Println("99. Exit")
	fmt.Scanln(&pilihan)
	fmt.Printf("\x1bc")
	switch pilihan {
	case 1:
		daftar(jum, buku)
	case 2:
		menu_buku(jum, buku)
	case 3:
		pinjam(jum, buku)
	case 4:
		cari(jum, buku)
	case 5:
		tarif(jum, &buku)
	case 10:
		main()
	case 99:
		exit()
	}
}

func menu_buku(jum int, buku book) {
	var pilihan int
	fmt.Println("1. Tambah buku")
	fmt.Println("2. Edit buku")
	fmt.Println("3. Hapus buku")
	fmt.Println("0. Kembali")
	fmt.Scanln(&pilihan)
	fmt.Printf("\x1bc")
	switch pilihan {
	case 1:
		tambah(&jum, &buku)
	case 2:
		edit(jum, &buku)
	case 3:
		hapus(&jum, &buku)
	case 0:
		menu_utama(jum, buku)
	}
}
func daftar(jum int, buku book) {
	var pilihan int
	fmt.Println("Urutkan sesuai:")
	fmt.Println("1. ID buku")
	fmt.Println("2. Buku yang dipinjam")
	fmt.Println("3. Buku favorit")
	fmt.Println("9. Kembali")
	fmt.Scanln(&pilihan)
	fmt.Printf("\x1bc")
	switch pilihan {
	case 1:
		var pil int
		fmt.Println("1.Ascending")
		fmt.Println("2.Descending")
		fmt.Scanln(&pil)
		if pil == 1 {
			urutIDasc(jum, &buku)
		} else if pil == 2 {
			urutIDdsc(jum, &buku)
		} else {
			fmt.Println("PILIHAN TIDAK TERSEDIA")
			menu_utama(jum, buku)
		}
		cetak_daftar(jum, buku)
	case 2:
		cetak_pinjam(jum, &buku)
	case 3:
		cetak_favorit(jum, buku)
	case 9:
		menu_utama(jum, buku)
	default:
		fmt.Println("Pilihan Tidak Tersedia")
		menu_utama(jum, buku)
	}
}
func cetak_pinjam(jum int, buku *book) {
	fmt.Printf("%6s %40s %20s %20s %20s %20s\n", "ID buku", "Judul", "Kategori", "Author", "Tahun", "Tanggal Peminjaman")
	for i := 0; i < jum; i++ {
		if buku[i].pinjam.pinjam == true {
			fmt.Printf("%6d %40s %20s %20s %20d %20s\n", buku[i].id_book, buku[i].judul, buku[i].kategori, buku[i].penulis, buku[i].tahun, buku[i].pinjam.tanggal.Format("2006-01-02"))
		}
	}
	fmt.Scanln()
	menu_utama(jum, *buku)
}
func cetak_favorit(jum int, buku book) {
	urut_pinjam(jum, &buku)
	fmt.Printf("%6s %40s %20s %20s %20s\n", "ID buku", "Judul", "Kategori", "Author", "Tahun")
	for i := 0; i < 5; i++ {
		if buku[i].pinjam.jumPinjam > 0 {
			fmt.Printf("%6d %40s %20s %20s %20d \n", buku[i].id_book, buku[i].judul, buku[i].kategori, buku[i].penulis, buku[i].tahun)
		}
	}
	fmt.Scanln()
	menu_utama(jum, buku)
}
func urut_pinjam(jum int, buku *book) {
	var i, j, min int
	var temp books
	for i < jum-1 {
		min = i
		j = i + 1
		for j < jum {
			if buku[min].pinjam.jumPinjam < buku[j].pinjam.jumPinjam {
				min = j
			}
			j = j + 1
		}
		temp = buku[min]
		buku[min] = buku[i]
		buku[i] = temp
		i = i + 1
	}
}
func urutIDasc(jum int, buku *book) {
	var i, j, min int
	var temp books
	for i < jum-1 {
		min = i
		j = i + 1
		for j < jum {
			if buku[min].id_book > buku[j].id_book {
				min = j
			}
			j = j + 1
		}
		temp = buku[min]
		buku[min] = buku[i]
		buku[i] = temp
		i = i + 1
	}
}
func urutIDdsc(jum int, buku *book) {
	var temp books
	for i := 1; i < jum; i++ {
		var j int
		j = i
		temp = buku[i]
		for j > 0 && temp.id_book > buku[j-1].id_book {
			buku[j] = buku[j-1]
			j--
		}
		buku[j] = temp
	}
}

func cetak_daftar(jum int, buku book) {
	fmt.Printf("%6s %40s %20s %20s %20s\n", "ID buku", "Judul", "Kategori", "Author", "Tahun")
	for i := 0; i < jum; i++ {
		fmt.Printf("%6d %40s %20s %20s %20d \n", buku[i].id_book, buku[i].judul, buku[i].kategori, buku[i].penulis, buku[i].tahun)
	}
	fmt.Scanln()
	menu_utama(jum, buku)
}
func hapus(jum *int, buku *book) {
	var id int
	fmt.Println("Hapus Buku")
	fmt.Print("Masukkan ID buku: ")
	fmt.Scanln(&id)
	if !cek_id_buku(id, *jum, *buku) {
		fmt.Printf("\x1bc")
		for i := 0; i < *jum; i++ {
			if buku[i].id_book == id {
				for j := i; j < *jum; j++ {
					buku[j] = buku[j+1]
				}
				*jum--
			}
		}
	} else {
		fmt.Println("Buku tidak ditemukan")
		fmt.Scanln()
		menu_utama(*jum, *buku)
	}
	fmt.Println("Hapus Berhasil")
	fmt.Scanln()
	menu_utama(*jum, *buku)
}

func tambah(jum *int, buku *book) {
	var n int
	var i int
	fmt.Print("Berapa Buku:")
	fmt.Scanln(&n)
	i = 0
	for i < n {
		tambah_id_buku(*jum, buku)
		*jum++
		i++
		fmt.Printf("\x1bc")
	}
	menu_utama(*jum, *buku)
}

func tambah_id_buku(jum int, buku *book) {
	var id int
	fmt.Print("ID_buku:")
	fmt.Scanln(&id)
	if cek_id_buku(id, jum, *buku) {
		buku[jum].id_book = id
		tambah_buku(jum, buku)
	} else {
		fmt.Println("ID sudah ada")
		tambah_id_buku(jum, buku)
	}
}

func cek_id_buku(id, jum int, buku book) bool {
	var cek bool = true
	var j int
	for j < jum && cek {
		if buku[j].id_book == id {
			cek = false
		}
		j++
	}
	return cek
}
func tambah_buku(jum int, buku *book) {
	fmt.Print("Judul:")
	buku[jum].judul = baca()
	fmt.Print("Kategori:")
	buku[jum].kategori = baca()
	fmt.Print("Penulis:")
	buku[jum].penulis = baca()
	fmt.Print("Tahun:")
	fmt.Scanln(&buku[jum].tahun)
}
func edit(jum int, buku *book) {
	var id int
	var pilihan int
	fmt.Print("ID buku:")
	fmt.Scanln(&id)
	if !cek_id_buku(id, jum, *buku) {
		for i := 0; i < jum; i++ {
			if buku[i].id_book == id {
				fmt.Println("Menu Edit")
				fmt.Println("1. ID buku")
				fmt.Println("2. Judul")
				fmt.Println("3. Kategori")
				fmt.Println("4. Penulis")
				fmt.Println("5. Tahun")
				fmt.Scanln(&pilihan)
				switch pilihan {
				case 1:
					fmt.Printf("\x1bc")
					var idd int
					fmt.Print("ID buku:")
					fmt.Scanln(&idd)
					if cek_id_buku(idd, jum, *buku) {
						buku[i].id_book = idd
					} else {
						fmt.Printf("\x1bc")
						fmt.Println("ID sudah ada")
						fmt.Scanln()
						edit(jum, buku)
					}
				case 2:
					fmt.Print("Judul:")
					buku[i].judul = baca()
				case 3:
					fmt.Print("Kategori:")
					buku[i].kategori = baca()
				case 4:
					fmt.Print("Penulis:")
					buku[i].penulis = baca()
				case 5:
					fmt.Print("Tahun:")
					fmt.Scan(&buku[i].tahun)
				}
			}
		}
	} else {
		fmt.Println("ID tidak ada")
		fmt.Scanln()
		menu_utama(jum, *buku)
	}
	fmt.Println("Edit Berhasil")
	fmt.Scanln()
	menu_utama(jum, *buku)
}

func cari(jum int, buku book) {
	var pilihan int
	fmt.Println("1. Cari ID")
	fmt.Println("2. Cari judul")
	fmt.Println("3. Cari Kategori")
	fmt.Println("4. Cari penulis")
	fmt.Println("5. Cari tahun")
	fmt.Println("9. Kembali")
	fmt.Scanln(&pilihan)
	fmt.Printf("\x1bc")
	switch pilihan {
	case 1:
		cari_id(jum, buku)
	case 2:
		cari_judul(jum, buku)
	case 3:
		cari_kategori(jum, buku)
	case 4:
		cari_penulis(jum, buku)
	case 5:
		cari_tahun(jum, buku)
	case 9:
		menu_utama(jum, buku)
	default:
		fmt.Println("Pilihan Tidak Tersedia")
		menu_utama(jum, buku)
	}
}
func cari_penulis(jum int, buku book) {
	var pilihan string
	var cek int
	fmt.Print("Penulis:")
	pilihan = baca()
	fmt.Printf("%6s %40s %20s %20s %20s\n", "ID buku", "Judul", "Kategori", "Author", "Tahun")
	for i := 0; i < jum; i++ {
		cek = findtext(strings.ToLower(buku[i].penulis), strings.ToLower(pilihan))
		if cek != -1 {
			fmt.Printf("%6d %40s %20s %20s %20d \n", buku[i].id_book, buku[i].judul, buku[i].kategori, buku[i].penulis, buku[i].tahun)
		}
	}
	fmt.Scanln()
	menu_utama(jum, buku)
}
func cari_kategori(jum int, buku book) {
	var pilihan string
	var cek int
	fmt.Print("Kategori:")
	pilihan = baca()
	fmt.Printf("%6s %40s %20s %20s %20s\n", "ID buku", "Judul", "Kategori", "Author", "Tahun")
	for i := 0; i < jum; i++ {
		cek = findtext(strings.ToLower(buku[i].kategori), strings.ToLower(pilihan))
		if cek != -1 {
			fmt.Printf("%6d %40s %20s %20s %20d \n", buku[i].id_book, buku[i].judul, buku[i].kategori, buku[i].penulis, buku[i].tahun)
		}
	}
	fmt.Scanln()
	menu_utama(jum, buku)
}
func cari_judul(jum int, buku book) {
	var pilihan string
	var cek int
	fmt.Print("Judul:")
	pilihan = baca()
	fmt.Printf("%6s %40s %20s %20s %20s\n", "ID buku", "Judul", "Kategori", "Author", "Tahun")
	for i := 0; i < jum; i++ {
		cek = findtext(strings.ToLower(buku[i].judul), strings.ToLower(pilihan))
		if cek != -1 {
			fmt.Printf("%6d %40s %20s %20s %20d \n", buku[i].id_book, buku[i].judul, buku[i].kategori, buku[i].penulis, buku[i].tahun)
		}
	}
	fmt.Scanln()
	menu_utama(jum, buku)
}

func cari_id(jum int, buku book) {
	var id int
	fmt.Print("ID buku:")
	urutIDdsc(jum, &buku)
	fmt.Scanln(&id)
	fmt.Printf("%6s %40s %20s %20s %20s\n", "ID buku", "Judul", "Kategori", "Author", "Tahun")
	var cek bool = false
	var med int
	var kr int = 0
	var kn int = jum - 1
	for kr <= kn && !cek {
		med = (kr + kn) / 2
		if id > buku[med].id_book {
			kn = med - 1
		} else if id < buku[med].id_book {
			kr = med + 1
		} else {
			cek = true
		}
	}
	if cek {
		fmt.Printf("%6d %40s %20s %20s %20d \n", buku[med].id_book, buku[med].judul, buku[med].kategori, buku[med].penulis, buku[med].tahun)
	}
	fmt.Scanln()
	menu_utama(jum, buku)
}

func cari_tahun(jum int, buku book) {
	var pilihan int
	var cek int
	fmt.Print("Tahun:")
	fmt.Scanln(&pilihan)
	fmt.Printf("%6s %40s %20s %20s %20s\n", "ID buku", "Judul", "Kategori", "Author", "Tahun")
	for i := 0; i < jum; i++ {
		if pilihan == buku[i].tahun {
			cek = i
			fmt.Printf("%6d %40s %20s %20s %20d \n", buku[cek].id_book, buku[cek].judul, buku[cek].kategori, buku[cek].penulis, buku[cek].tahun)
		}
	}

	fmt.Scanln()
	menu_utama(jum, buku)
}
func findtext(text, pattern string) int {
	barChar := [256]int{}
	for i := range barChar {
		barChar[i] = len(pattern)
	}
	for i := 0; i < len(pattern); i++ {
		barChar[pattern[i]] = len(pattern) - i - 1
	}
	i := len(pattern) - 1
	for i < len(text) {
		j := len(pattern) - 1
		for j >= 0 && pattern[j] == text[i] {
			j--
			i--
		}
		if j < 0 {
			return i + 1
		}
		i = i + barChar[text[i]]
	}
	return -1
}

func pinjam(jum int, buku book) {
	var pilihan int
	fmt.Println("1.Pinjam")
	fmt.Println("2.Edit pinjaman")
	fmt.Println("3.Kembalikan")
	fmt.Println("9.Kembali")
	fmt.Scanln(&pilihan)
	switch pilihan {
	case 1:
		tambah_pinjam(jum, &buku)
	case 2:
		edit_pinjam(jum, &buku)
	case 3:
		kembalikan(jum, &buku)
	case 9:
		menu_utama(jum, buku)
	default:
		menu_utama(jum, buku)

	}
}
func tambah_pinjam(jum int, buku *book) {
	var id int
	var tanggal string
	fmt.Println("PEMINJAMAN")
	fmt.Print("ID buku yang akan dipinjam:")
	fmt.Scan(&id, jum, buku)
	if !cek_id_buku(id, jum, *buku) {
		for i := 0; i < jum; i++ {
			if id == buku[i].id_book {
				if !buku[i].pinjam.pinjam {
					fmt.Print("Tanggal Pinjam Format(YYYY-MM-DD):")
					fmt.Scan(&tanggal)
					a, _ := time.Parse("2006-01-02", tanggal)
					buku[i].pinjam.tanggal = a
					buku[i].pinjam.pinjam = true
					buku[i].pinjam.jumPinjam++
					buku[i].pinjam.tanggalkembali = buku[i].pinjam.tanggal.Add(time.Hour * 24 * 6)
					fmt.Println("Buku berhasil dipinjam")
				} else {
					fmt.Println("Buku sedang dipinjam")
				}
			}
		}
	} else {
		fmt.Println("Buku tidak ditemukan")
	}
	fmt.Scanln()
	pinjam(jum, *buku)
}
func kembalikan(jum int, buku *book) {
	var id int
	fmt.Println("PENGEMBALIAN")
	fmt.Print("ID buku yang akan dikembalikan:")
	fmt.Scan(&id)
	if !cek_id_buku(id, jum, *buku) {
		for i := 0; i < jum; i++ {
			if id == buku[i].id_book {
				if buku[i].pinjam.pinjam {
					buku[i].pinjam.pinjam = false
					fmt.Println("Terima kasih sudah mengembalikan")
				} else {
					fmt.Println("Buku tidak sedang dipinjam")
				}
			}
		}
	} else {
		fmt.Println("Buku tidak ditemukan")
	}
	fmt.Scanln()
	pinjam(jum, *buku)
}
func edit_pinjam(jum int, buku *book) {
	var id, pilihan int
	var tanggal string
	fmt.Println("PENGEDITAN PINJAMAN")
	fmt.Print("ID buku dipinjam yang akan diedit:")
	fmt.Scan(&id)
	if !cek_id_buku(id, jum, *buku) {
		for i := 0; i < jum; i++ {
			if id == buku[i].id_book {
				if buku[i].pinjam.pinjam {
					fmt.Println("1. Edit Tanggal Pinjam")
					fmt.Println("2. Hapus Data Pinjaman Saat ini")
					fmt.Println("3. HAPUS SEMUA DATA PINJAMAN")
					fmt.Scan(&pilihan)
					if pilihan == 1 {
						fmt.Print("Format(YYYY-MM-DD)")
						fmt.Scan(&tanggal)
						a, _ := time.Parse("2006-01-02", tanggal)
						buku[i].pinjam.tanggal = a
						fmt.Println("Pengeditan telah selesai")
						pinjam(jum, *buku)
					} else if pilihan == 2 {
						buku[i].pinjam.jumPinjam--
						buku[i].pinjam.pinjam = false
						fmt.Println("Pengeditan telah selesai")
						pinjam(jum, *buku)
					} else if pilihan == 3 {
						buku[i].pinjam.jumPinjam = 0
						buku[i].pinjam.pinjam = false
						fmt.Println("Pengeditan telah selesai")
						pinjam(jum, *buku)
					}
				} else {
					fmt.Println("Buku tidak sedang dipinjam")
					pinjam(jum, *buku)
				}
			}
		}
	} else {
		fmt.Println("Buku tidak ditemukan")
		pinjam(jum, *buku)
	}
}

func tarif(jum int, buku *book) {
	var terlambat int
	for i := 0; i < jum; i++ {
		if buku[i].pinjam.pinjam == true {
			buku[i].pinjam.waktu = (time.Now().Year()-buku[i].pinjam.tanggal.Year())*360 + (int(time.Now().Month())-int(buku[i].pinjam.tanggal.Month()))*30 + (time.Now().Day() - buku[i].pinjam.tanggal.Day())
			if buku[i].pinjam.waktu < 7 {
				buku[i].pinjam.tarif = (buku[i].pinjam.waktu + 1) * 1000
				terlambat = 0
			} else if buku[i].pinjam.waktu >= 7 {
				terlambat = buku[i].pinjam.waktu - 6
				buku[i].pinjam.tarif = 7000
				buku[i].pinjam.denda = (buku[i].pinjam.waktu - 6) * 5000
			}
		}
	}
	fmt.Printf("%6s %40s %20s %20s %10s %10s %10s %10s \n", "ID buku", "Judul", "Tanggal Peminjaman", "Tanggal Pengembalian", "Tarif", "Terlambat", "Denda", "Total")
	for i := 0; i < jum; i++ {
		if buku[i].pinjam.pinjam == true {
			fmt.Printf("%6d %40s %20s %20s %10d %10d %10d %10d\n", buku[i].id_book, buku[i].judul, buku[i].pinjam.tanggal.Format("2006-01-02"), buku[i].pinjam.tanggalkembali.Format("2006-01-02"), buku[i].pinjam.tarif, terlambat, buku[i].pinjam.denda, buku[i].pinjam.tarif+buku[i].pinjam.denda)
		}
	}
	fmt.Scanln()
	menu_utama(jum, *buku)

}
