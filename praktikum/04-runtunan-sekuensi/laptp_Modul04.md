# <h1 align="center">Tugas Pendahuluan Modul [04] - [Runtutan/Sekuensi]</h1>
<p align="center">[Rafif Ramadhan] - [109092630001]</p>

### 1. Soal 1

```go
package main

import "fmt"

func main() {
	intNum := 5
	intOther := 10
	var sngNum float64 = -3

	fmt.Println(intNum > 5 && intOther > 5 && sngNum > 0)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP_4/soalsatu/output%20soalsatu.png)


#### Deskripsi
Tugas diatas bertujuan untuk menentukan nilai boolean dari beberapa soal yang diberikan

### 2. Soal 2

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP_4/soaldua/output%20soaldua.png)


#### Deskripsi
Tugas tersebut bertujuan untuk mengalisa program tersebut untuk mengecek output dan nlai dari code tersebut

### 3. Soal 3

```go
package main

import "fmt"

func main() {
	var tahun int
	var bulan string

	fmt.Scan(&tahun, &bulan)

	// Cek tahun kabisat
	kabisat := (tahun%4 == 0 && tahun%100 != 0) || tahun%400 == 0

	var hari int

	switch bulan {
	case "Jan", "Mar", "Mei", "Jul", "Agu", "Okt", "Des":
		hari = 31
	case "Apr", "Jun", "Sep", "Nov":
		hari = 30
	case "Feb":
		if kabisat {
			hari = 29
		} else {
			hari = 28
		}
	default:
		// Nama bulan tidak valid: tidak ada keluaran ("-")
		fmt.Println("-")
		return
	}

	fmt.Println(hari)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP_4/soaltiga/output%20soaltiga.png)

#### Deskripsi
Tugas diatas bertujuan untuk menentukan jumlah hari dalam bulan tertentu dengan user memasukkan nama bulan 

### 4. Soal 4

```go
package main

import "fmt"

func main() {
	var hari int
	fmt.Print("Masukkan angka hari (1-7): ")
	fmt.Scan(&hari)

	switch hari {
	case 1:
		fmt.Println("Senin")
	case 2:
		fmt.Println("Selasa")
	case 3:
		fmt.Println("Rabu")
	case 4:
		fmt.Println("Kamis")
	case 5:
		fmt.Println("Jumat")
	case 6:
		fmt.Println("Sabtu")
	case 7:
		fmt.Println("Minggu")
	default:
		fmt.Println("Hari tidak valid")
	}
}
```
##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/TP_4/soalempat/output%20soalempat.png)

#### Deskripsi
Tugas diatas memberikan insrutksi untuk menentukan hari dalam seminggu ke dalam bentuk nomor

## Kesimpulan
Tujuan dari tugas pendahuluan ini adalah untuk mengetahui urutan pemrograman bekerja secara runtutan dan sistematis
