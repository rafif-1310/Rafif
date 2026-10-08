# <h1 align="center">Tugas Pendahuluan Modul [03] - [Variabel dan Operator]</h1>
<p align="center">[Rafif Ramadhan] - [109092630001]</p>

## Dasar Teori

### 1.Variabel Go
Tipe Variabel Go
Dalam bahasa pemrograman Go, terdapat berbagai jenis variabel, misalnya:

-int- menyimpan bilangan bulat (bilangan utuh), seperti 123 atau -123

-float32- menyimpan angka floating point, dengan desimal, seperti 19,99 atau -19,99

-string- Menyimpan teks, seperti "Hello World". Nilai string dikelilingi oleh tanda kutip ganda.

-bool- menyimpan nilai dengan dua keadaan: benar   atau salah

### 2.Operator
Operator digunakan untuk melakukan operasi pada variabel dan nilai.

+ Operator ini menjumlahkan dua nilai, seperti pada contoh di bawah ini:

```go
package main
import "fmt"

func main() {
  var a = 15 + 25
  fmt.Println(a)
}
```

Meskipun +operator ini sering digunakan untuk menjumlahkan dua nilai, operator ini juga dapat digunakan untuk menjumlahkan variabel dan suatu nilai, atau variabel dan variabel lain:

```go
package main
import "fmt"

func main() {
  var (
    sum1 = 100 + 50 // 150 (100 + 50)
    sum2 = sum1 + 250 // 400 (150 + 250)
    sum3 = sum2 + sum2 // 800 (400 + 400)
  )
  fmt.Println(sum3)
}
```
Sumber : https://www.w3schools.com

## Guided

### 1.Suhu

```go
package main

import "fmt"

func main() {
	var cel float64
	fmt.Print("Masukan Suhu: ")
	fmt.Scan(&cel)

	fmt.Print(cel + 273)
}
```
#### Deskripsi
Program diatas mengubah satuan suhu dari kelvin ke celcius dengan menggunakan K = C + 273
sehingga ketika kita memasukkan suhu dalam celcius maka akan diubah otomatis oleh sistem ke dalam reamur

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/konversi/Screenshot%202026-10-05%20152735.png)


### 2.Tukar Nilai

```go
package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	temp := x
	x = z
	z = y
	y = temp
	fmt.Println(x, y, z)
}
```
#### Deskripsi
Program diatas berisikan variabel x,y dan z yang bernilai integer lalu menukarnya dengan menggunakan (=) sehingga nilai nya tetap namun bertukar tempatnya

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/tukar/output.png)



### 3.Kasir

```go
package main

import "fmt"

func main() {
	var uang int
	fmt.Print("Uang = ")
	fmt.Scan(&uang)

	sepuluh := uang / 10000
	sisa := uang % 10000
	lima := sisa / 5000
	sisa = sisa % 5000
	satu := sisa / 1000
	sisa = sisa % 1000
	fmt.Println("10000 = ", sepuluh)
	fmt.Println("5000 = ", lima)
	fmt.Println("1000 = ", satu)
	fmt.Print("sisa = ", sisa)
}
```
### Deskripsi
Program diatas berfungsi untuk mempermudah kita dalam uang dalam beberapa pecahan uang yang lebih sederhana, menggunakan % dan /

<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/guided/kasir/Screenshot%202026-10-05%20151458.png)

## Unguided

### 1.Konversi Suhu

```go
package main

import "fmt"

func main() {

	 var cel float64
	 
	 fmt.Print("Masukan Suhu: ")
	 fmt.Scan(&cel)
	
	 reamur := cel * 4 / 5

	 fmt.Print("Suhu dalam Reamur: ", reamur)
	 
}
```
#### Deskripsi
Program diatas mengubah satuan suhu dari celcius ke reamur dengan menggunakan reamur := cel * 4 / 5
sehingga ketika kita memasukkan suhu dalam celcius maka akan diubah otomatis oleh sistem ke dalam reamur

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi_suhu/output.png)

### 2.Konversi Hari

```go
package main

import "fmt"

func main() {
	var totalHari int
	fmt.Print("Masukkan jumlah hari: ")
	fmt.Scanln(&totalHari)

	tahun := totalHari / 360
	bulan := (totalHari % 360) / 30
	minggu := ((totalHari % 360) % 30) / 7
	hari := ((totalHari % 360) % 30) % 7

	fmt.Printf("%d\n%d\n%d\n%d\n", tahun, bulan, minggu, hari)
}
```

#### Deskripsi
Pemrograman diatas bertujuan untuk memecah hari yang ada dalam satu tahun menjadi bulan,minggu dan hari

#### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/03-tipe-data-dan-instruksi-dasar/unguided/konversi_hari/Screenshot%202026-10-05%20153417.png)


## Kesimpulan
Semua materi diatas bertujuan untuk mengenalkan dasar-dasar dari bahasa golang sehingga kita bisa mengetahui bagaimana algoritma bahasa pemrograman golang