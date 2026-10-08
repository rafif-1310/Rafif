# <h1 align="center">Tugas Pendahuluan Modul [03] - [Variabel dan Operator]</h1>
<p align="center">[Rafif Ramadhan] - [109092630001]</p>

### 1. Sisa Kue

```go
package main

import "fmt"

func main() {

	var x,y int
	
	fmt.Print("Masukkan jumlah kue: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan jumlah anggota keluarga: ")	
	fmt.Scan(&y)

	sisa := x % y

	fmt.Print("Jumlah kue yang tersisa:", sisa )

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/sisa_kue/Output%20sisa_kue.go.png)


#### Deskripsi
Program diatas berfungsi untuk membagi jumlah kue agar sama rata untuk semua anggota kelaurga,dengan memasukkan 2 variabel(int) lalu dibagi menggunakan (&) 

### 2.Boolean

```go
package main

import "fmt"

func main() {
	var nilai bool
	
	fmt.Print("nilai bool(true/false): ")
	fmt.Scan(&nilai)

	fmt.Println(nilai)

}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/boolean/output%20boolean.go.png)


#### Deskripsi
Program diatas bertujuan untuk menukar nilai true/false atau biasa disebut boolean

### Konversi Jarak

```go
package main

import "fmt"

func main(){
	var mil float64

	fmt.Println("masukkan jarak: ")
	fmt.Scan(&mil)

	km := mil * 1.6

	fmt.Print(km)
}
```

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/tp/konversi_jarak/output%20koversi_jarak.go.png)

#### Deskripsi
Program diatas mengkonveri jarak dari mil diubah menjadi km

## Kesimpulan
Kesimpulan dari tugas pendahuluan diatas adalah untuk mengenalkan  dasar-dasar bahasa golang, agar dapat dipahami oleh mahasisa baru