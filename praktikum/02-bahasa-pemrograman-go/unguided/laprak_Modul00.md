# <h1 align="center">Laporan Praktikum Modul [00] - [Dasar Bahasa Pemrograman Go]</h1>
<p align="center">[Rafif Ramadhan] - [109092630001]</p>

## Dasar Teori

### A. [A. Pengenalan Dasar Bahasa Pemrograman Go]
[Go atau Golang adalah bahasa pemrograman prosedural yang dibuat di Google menggunakan bahasa pemrograman C oleh Robert Griesemer, Rob Pike dan Ken Thompson pada tahun 2007 dan dirilis sebagai bahasa pemrograman open source pada tahun 2009.

Golang mulai populer sejak digunakan untuk membuat Docker pada tahun 2011. Saat ini Go-Lang mulai populer untuk pembuatan Backend API pada arsitektur Microservices serta mulai banyak teknologi baru yang dibuat menggunakan bahasa Go-Lang daripada bahasa C, seperti Kubernetes, Prometheus, CockroachDB dan lain-lain. [ Irvan Eksa Mahendra.](2021])

### B. [Package dan Struktur Program di Go atau Golang]

#### 1. [Pengertian Package main dan func main()]
[Setiap file program harus memiliki package. Setiap project harus ada minimal satu file dengan nama package main. File yang ber-package main, akan dieksekusi pertama kali ketika program dijalankan.

Dalam sebuah proyek harus ada file program yang di dalamnya berisi sebuah fungsi bernama main(). Fungsi tersebut harus berada di file yang package-nya bernama main.

Fungsi main() adalah yang dipanggil pertama kali pada saat eksekusi program.
Menurut Noval Agung Prayogo. 2019]

#### 2. [Deklarasi Variabel dan Tipe Data di Go]
[Go adalah bahasa statically typed, artinya tipe data variabel dicek saat kompilasi. Namun, Go memiliki fitur canggih bernama Type Inference.

Deklarasi Variabel

Ada dua cara umum untuk membuat variabel:

Cara 1: Deklarasi Eksplisit

Digunakan ketika Anda ingin menentukan tipe data secara manual atau mendeklarasikan tanpa nilai awal.]

var nama string = "Budi"

var umur int = 25

Cara 2: Short Variable Declaration (:=)

Ini adalah cara paling umum di dalam fungsi. Go akan otomatis menebak tipe datanya.

negara := "Indonesia" // Go tahu ini String

skor := 95.5          // Go tahu ini Float64

Catatan: Di Go, jika Anda mendeklarasikan variabel tetapi tidak menggunakannya, program akan error (gagal kompilasi). Ini memaksa kode tetap bersih.

ditulis oleh Abdullah Fawwaz Qudamah. 2026
## Guided

### 1.Suhu.go


```go
package main
import "fmt"

func main(){
	var suhu float64
 //Membaca input

 fmt.Print("Masukkan suhu dalam Celcius: ")
 fmt.Scan(&suhu)

 //Menampilkan output
  fmt.Println("Suhu dalam Fahrenheit: ", (suhu*9/5)+32)
  fmt.Println("Suhu dalam Kelvin: ", suhu+273.15)
  fmt.Println("Suhu dalam Reannur: ", suhu*0.8)
  
}

```
#### Deskripsi
Program di atas membaca nilai celsius yang kita input lalu dikonvert mennjadi reamur, farenheit, dan kelvin. Nilai yang diinput disimpan dalan variable "celsius" (float64) lalu dikonvert menjadi beberapa jenis suhu

### 2.Lingkaran.go

```go
import "fmt"

func main() {
	var jari float64
	var pi = 3.14
	var luas float64

	fmt.Print("Masukkan jari-jari lingkaran: ")
	fmt.Scan(&jari)

	luas = pi * jari * jari

    fmt.Println("================ OUTPUT ==================")
	fmt.Println("Jari-jari           :", jari)
	fmt.Println("Luas lingkaran      :", luas)
	fmt.Println("==========================================")
}
```
#### Deskripsi 
Program di atas menghitung luas lingkaran dengan cara menghitung jari - jari yang diinputkan dan disimpan di variable "r" (float64) dengan "pi (float64). Cara menghitungnya mengikuti rmus menghitung luas lingkaran yaitu: pi _ r _.


### 3.Skor.go
```go
 package main
 import "fmt"

 func main(){ 
	var nama string
	var skorMatematika, skorBahasaInggris int
	
	//Membaca input
	fmt.Scan(&nama)
	fmt.Scan(&skorMatematika)
	fmt.Scan(&skorBahasaInggris)

	//Menghitung total & rata-rata skor (Pembagian bilangan bulat) 
	total := skorMatematika + skorBahasaInggris
	rataRata := total / 2

	//Menampilkan output
	fmt.Println(nama)
	fmt.Println(total)
	fmt.Println(rataRata)
	
 }
 ```
 #### Deskripsi
 Program di atas membaca input nama siswa dengan variable (string), C. Kemudian program akan menghitung total dari kedua nilai variable (skorMatematika, skorBahasaInggris) untuk menghitung seluruh total nilai, setelah itu total nilai akan masuk ke proses pembagian (skorMatematika, skorBahasaInggris)/2) guna menghitung rata - rata dari nilai atau skor siswa yang telah diinputkan.

### 4.Tukar.go

```go
package main
import "fmt"

func main() {
	var a,b int

	//Membaca input
	fmt.Scan(&a)
	fmt.Scan(&b)

	//Menukar a dan b
	a,b=b,a

	//Menampilkan output
	fmt.Println(a)
	fmt.Println(b)
}

```
#### Deskripsi
Program di atas membaca nilai yang diinput ke dalam variable a (int) dan b (int) lalu menampilkan hasil inputan kita secara tertukar contoh: kita menginputkan a = 50, b = 40 hasilnya Nilai a: 40, Nilai b: 50.




## Unguided

### 1.Kalkulator.go
```go
package main

import "fmt"

func main() {

	var a, b int


	fmt.Scan(&a)
	fmt.Scan(&b)


	fmt.Println("Hasil penjumlahan:", a + b)
	fmt.Println("Hasil pengurangan:", a - b)
	fmt.Println("Hasil perkalian:", a * b)
	fmt.Println("Hasil pembagian:", a / b)
	fmt.Println("Hasil sisa bagi:", a % b)

}
```

##### Output
![Screenshot Output Unguided](https://github.com/rafif-1310/Rafif/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/kalkulator/output.png)


#### Deskripsi
Program di atas adalah program kalkulator, dimana program membaca nilai yang kita inputkan lalu disimpan dalam variabel "a" dan "b" keduanya adalah (int). Variable yang kita inputkan akan dieksekusi dengan metode matematika dasar seperti penjumalahan, pengurangan, perkalian, pembagian, dan sisa bagi.

### 2.Cacahuang.go

```go
package main

import "fmt"

func main() {

var Jumlah int32
fmt.Println("Masukkan jumlah uang: ")
fmt.Print("jumlah uang ")
fmt.Scan(&Jumlah)

uang10000 := Jumlah / 10000
sisa := Jumlah % 10000
uang5000 := sisa / 5000
sisa = sisa % 5000
uang1000 := sisa / 1000
sisa = sisa % 1000

fmt.Println("Jumlah uang:", Jumlah)
fmt.Println("Uang 10000:", uang10000)
fmt.Println("Uang 5000 :", uang5000)
fmt.Println("Uang 1000 :", uang1000)
fmt.Println("Sisa uang  :", sisa)
}
```

##### Output
![Screenshot Output Unguided](https://github.com/rafif-1310/Rafif/blob/main/praktikum/02-bahasa-pemrograman-go/unguided/cacahuang/output.png)

#### Deskripsi
Program di atas bertujuan untuk mencacah uang yang kita inputkan ke dalam variable "Jumlah" (int32) dengan cara Jumlah dibagi 10000 lalu sisa Jumlah dibagi sisa 10000 lalu sisa dibagi 5000 lalu sisanya dibagi hasil 5000 lalu sisa akhir dibagi 1000.



## Kesimpulan
Program ini bertujuan agar kita mengenali dasar-dasar pemrograman go,melalui berbagai latihan diatas

## Referensi
1. Irvan Eksa Mahendra. (2021). medium.com. 
2. Noval Agung Prayogo. (2019). dasarpemrogramangolang.novalagung.com.
3. Abdullah Fawwaz Qudamah. (2026). fawwaz.id