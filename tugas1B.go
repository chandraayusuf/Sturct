package main

import "fmt"

type mahasiswa struct{
    nama string;
    nim int;
    jurusan string;
}

func main(){
	var mhs mahasiswa
	fmt.Println("MEMASUKAN INPUTAN DATA")

	for i := 0; i < 6; i++{
	
	fmt.Println("Masukan Nama :")
	fmt.Scan(&mhs.nama)
	fmt.Println("Masukan Nim :")
	fmt.Scan(&mhs.nim)
	fmt.Println("Masukan Jurusan :")
	fmt.Scan(&mhs.jurusan)
	}

}


