# <h1 align="center">Tugas Pendahuluan Modul [03 Algoritma Pemrograman] - [Variable dan Operator]</h1>
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
![Screenshot Output Unguided](/03-tipe-data-dan-instruksi-dasar/unguided/sisa_kue/Output%20sisa_kue.go.png)

#### Deskripsi
Di program diatas memasukkan variabel berupa x dan y sebagai integer,lalu var x sebagai jumlah kue dan variabel y sebagai jumlah anggota keluarga,lalu x dan y dibagi menggunakan modulus(%) lalu kita bisa mendapatkan sisa dari dari kue tersebut

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
![Screenshot Output Unguided](/03-tipe-data-dan-instruksi-dasar/unguided/boolean/output%20boolean.go.png)


#### Deskripsi
Program tersebut memasukka tipe data boolan yang dimana jika kita memasukkan nilai true maka akan mengeluarkan output true,begitu pula sebaliknya
### 3.Konversi Jarak
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
![Screenshot Output Unguided](/03-tipe-data-dan-instruksi-dasar/unguided/konversi_jarak/output%20koversi_jarak.go.png)

#### Deskripsi
Program diatas masukkan mil menjadi float64 lalu mengubah nya ke satuan km


## Kesimpulan
Tujuan latihan diatas adalah untuk mengetahui dasar-dasar pemrograman
