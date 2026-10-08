package main

import (
	"bufio" // Digunakan untuk membaca masukan string yang memiliki spasi
	"fmt"
	"os" // Menyediakan akses ke sistem operasi, dalam hal ini os.Stdin (keyboard)
)

func main() {
	var nama string
	var nilai float64
	var grade string

	// Membuat alat pembaca (scanner) baru yang mengambil masukan dari keyboard
	scanner := bufio.NewScanner(os.Stdin)

	// Proses membaca masukan dari pengguna sampai tombol Enter ditekan
	scanner.Scan()

	// Mengambil teks yang baru saja dibaca dan menyimpannya ke variabel nama
	nama = scanner.Text()

	fmt.Scan(&nilai)

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 && nilai < 90 {
		grade = "B"
	} else if nilai >= 70 && nilai < 80 {
		grade = "C"
	} else if nilai >= 60 && nilai < 70 {
		grade = "D"
	} else {
		grade = "F"
	}

	// %s adalah format penulisan (placeholder) untuk mencetak nilai bertipe string
	// %s pertama akan diisi oleh variabel 'nama', %s kedua diisi oleh 'grade'
	fmt.Printf("%s mendapatkan nilai %s\n", nama, grade)
}
