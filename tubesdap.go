/*
	Anggota Kelompok : Muhammad Aga Ananda Said Muharam (1302194061)
					   Ryan Fahreza Maliki (1302190079)
					   
	Judul Tubes : Aplikasi Pendaftaran Les 
	Kelas : SE-43-01

*/

package main

import "fmt"

type data struct {
	nama    string
	nohp    string
	sekolah string
	jenjang string
}

var arrdata [10000]data

var daftar int
var (
	namah, code, nama, school, jjg string
	nomor                          int
)

func menu() {
//	var pilihhh int

	fmt.Println("=========== Pendaftaran Les ==========")
	fmt.Println("=\t\t Menu                =")
	fmt.Println("======================================")
	fmt.Println("1. Input Pendaftar\t4. Lihat Data ")
	fmt.Println("2. Modifikasi Data   	5. Cari Data ")
	fmt.Println("3. Hapus Data 		6. EXIT")
	fmt.Println("======================================")
}


func inputan() {
	var pilih, pilihh, harga int
	var ipaSD, ipsSD, mtkSD, ipaSMP, ipsSMP, mtkSMP, fis, bio, mtk int

	ipaSD = 200000
	ipsSD = 150000
	mtkSD = 250000
	ipaSMP = 300000
	ipsSMP = 275000
	mtkSMP = 350000
	fis = 400000
	bio = 385000
	mtk = 500000

	fmt.Print("Masukan banyak pendaftar :")
	fmt.Scan(&daftar)
	for i := 0; i < daftar; i++ {
		fmt.Print("Nama : ")
		fmt.Scan(&arrdata[i].nama)
		fmt.Print("Nomor Hp :+62")
		fmt.Scan(&arrdata[i].nohp)
		fmt.Print("Asal sekolah: ")
		fmt.Scan(&arrdata[i].sekolah)
		fmt.Print("Jenjang Pendidikan : ")
		fmt.Scan(&arrdata[i].jenjang)
		fmt.Println("1. SD")
		fmt.Println("2. SMP")
		fmt.Println("3. SMA/SMK")
		fmt.Print("Pilih : ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			fmt.Println()
			fmt.Println("Pelajaran")
			fmt.Println("------------")
			fmt.Println("1. IPA")
			fmt.Println("2. IPS")
			fmt.Println("3. Matematika")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilihh)
			 if pilihh == 1 {
				harga = ipaSD
				fmt.Println("Harga Bayar :", harga)
			} else if pilihh == 2 {
				harga = ipsSD
				fmt.Println("Harga Bayar :", harga)
			} else if pilihh == 3 {
				harga = mtkSD
				fmt.Print("Harga Bayar :", harga)
			} else {
				fmt.Println("Tidak Valid")
			}
		} else if pilih == 2 {
			fmt.Println()
			fmt.Println("PIlih Pelajaran")
			fmt.Println("---------------")
			fmt.Println("1. IPA")
			fmt.Println("2. IPS")
			fmt.Println("3. Matematika")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilihh)
			 if pilihh == 1 {
				harga = ipaSMP
				fmt.Println("Harga Bayar :", harga)
			} else if pilihh == 2 {
				harga = ipsSMP
				fmt.Println("Harga Bayar :", harga)
			} else if pilihh == 3 {
				harga = mtkSMP
				fmt.Println("Harga Bayar :", harga)
			} else {
				fmt.Println("Tidak Valid")
			}
		} else if pilih == 3 {
			fmt.Println()
			fmt.Println("Pilih Pelajaran")
			fmt.Println("---------------")
			fmt.Println("1. Fisika ")
			fmt.Println("2. Biologi")
			fmt.Println("3. Matematika Wajib")
			fmt.Print("Pilihan : ")
			fmt.Scan(&pilihh)
			if pilihh == 1 {
				harga = fis
				fmt.Println("Harga Bayar :", harga)
			} else if pilihh == 2 {
				harga = bio
				fmt.Println("Harga Bayar :", harga)
			} else if pilihh == 3 {
				harga = mtk
				fmt.Println("Harga Bayar :", harga)
			} else {
				fmt.Println("Tidak Valid")
			}
		} else {
			fmt.Print()
		}
	}
	//menu()
}

func modifiKasi() {
	var code int
	var nama, nomor, school, jjg string

	fmt.Println("----Modifikasi Data----")
	fmt.Print("Masukan Nama : ")
	fmt.Scan(&namah)
	for i := 0; i < daftar; i++ {
		if namah == arrdata[i].nama {
			fmt.Println("Data yang akan di modifikasi ")
			fmt.Println("1. Nama")
			fmt.Println("2. Nomor Hape")
			fmt.Println("3. Asal Sekolah")
			fmt.Println("4. Jenjang Sekolah")
			fmt.Println("------------------------")
			fmt.Print("Masukan Pilihan    : ")
			fmt.Scan(&code)

			if code == 1 {
				fmt.Print(" Masukan Nama baru :")
				fmt.Scan(&nama)
				arrdata[i].nama = nama
			} else if code == 2 {
				fmt.Print(" Masukan Nomor HP baru : ")
				fmt.Scan(&nomor)
				arrdata[i].nohp = nomor
			} else if code == 3 {
				fmt.Print(" Masukan Asal Sekolah baru : ")
				fmt.Scan(&school)
				arrdata[i].sekolah = school
			} else if code == 4 {
				fmt.Print(" Masukan Jenjang Terbaru : ")
				fmt.Scan(&jjg)
				arrdata[i].jenjang = jjg
			} else {
				fmt.Print("Tidak Valid")

			}
		}
		//menu()
	}
}

func hapus() {
	var hapusDaata string
	var i int

	fmt.Println()
	fmt.Println("----Hapus Data----")
	fmt.Print("Masukkan Data Yang Dihapus (NAMA) : ")
	fmt.Scan(&hapusDaata)
	for i < daftar {
		if arrdata[i].nama == hapusDaata {
			arrdata[i].nama = ""
			arrdata[i].nohp = ""
			arrdata[i].sekolah = ""
			arrdata[i].jenjang = ""
			fmt.Println("BERHASIL")
		}
		i++
	}
	//menu()
}

func cariData() {
	var name string
	var kos int

	if daftar > 0 {
		fmt.Println()
		fmt.Print("Masukan Nama yang di cari : ")
		fmt.Scan(&name)
		for i := 0; i <= daftar; i++ {
			if name == arrdata[i].nama {
				fmt.Println()
				fmt.Println("Data ditemukan")
				fmt.Println("Nama : ", arrdata[i].nama)
				fmt.Println("Nomor hp : ", arrdata[i].nohp)
				fmt.Println("Nama Sekolah : ", arrdata[i].sekolah)
				fmt.Println("Jenjang : ", arrdata[i].jenjang)
				kos = 1
				}
		}
		if kos != 1 {
			fmt.Println()
			fmt.Println("Data tidak ada")
		}

		//menu()
	}
}

func lihatdata() {
	var coba data
	var jumlah int
	var urut string
	for i := 0; i < daftar; i++ {
		fmt.Println()
		fmt.Println("----Lihat Data----")
		fmt.Println("Data Ke-", i+1)
		fmt.Println("Nama\t\t:", arrdata[i].nama)
		fmt.Println("Nomor Hp\t:+62", arrdata[i].nohp)
		fmt.Println("Nama sekolah\t:", arrdata[i].sekolah)
		fmt.Println("jenjang\t\t:", arrdata[i].jenjang)
		fmt.Println("")
	}
	fmt.Print("Ingin Mengurutkan data sesuai nama (y/n) :")
	fmt.Scan(&urut)
		if urut == "y" {	
		urutkan(coba,jumlah)
		for i:=0; i < daftar; i++{
		fmt.Println()
		fmt.Println("----Data Nama Setelah di urutkan----")
		fmt.Println("Data Ke-", i+1)
		fmt.Println("Nama\t\t:", arrdata[i].nama)
		fmt.Println("Nomor Hp\t:+62", arrdata[i].nohp)
		fmt.Println("Nama sekolah\t:", arrdata[i].sekolah)
		fmt.Println("jenjang\t\t:", arrdata[i].jenjang)
	}
		
		} else {
			fmt.Println("Terimakasih")
		}
	//enu()
}

func urutkan(coba data, jumlah int) data{

	for i:=0; i < daftar-1; i ++ {
		for j:= 0; j < daftar-i-1; j++ {
				if arrdata[j].nama > arrdata[j+1].nama {
					temp:= arrdata[j]
					arrdata[j] = arrdata[j+1]
					arrdata[j+1] = temp
				}
		}
	}
	return coba
}
	/*var pass,i int
	var temp data

	i =0 
	pass = i

	for pass < daftar - 1 {
		i = pass + 1
		temp = arrdata[i]
		for i > 0 && temp.nama < arrdata[i-1].nama {
			arrdata[i] = arrdata[i-1]
			i--
		}
	arrdata[i] = temp
	pass++
	}*/
func main() {
	var pilihhh  int

pilihhh = 0
	for pilihhh != 6 {
			menu()
	fmt.Print("Pilih Nomor :")
	fmt.Scan(&pilihhh)
		if pilihhh == 1 {
			inputan()
		} else if pilihhh == 2 {
			modifiKasi()
		} else if pilihhh == 3 {
			hapus()
		} else if pilihhh == 4 {
			lihatdata()
		} else if pilihhh == 5 {
			cariData()
		} 
	}
	//menu()
	//inputan()
	//hapus()
	//lihatdata()
	//baru()
	//mencariData()
}
