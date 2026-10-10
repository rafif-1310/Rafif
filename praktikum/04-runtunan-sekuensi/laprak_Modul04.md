# <h1 align="center">Tugas Pendahuluan Modul [04] - [Runtunan/Sekuensi]</h1>
<p align="center">[Rafif Ramadhan] - [109092630001]</p>

## Dasar Teori

### A. Alur Sekuensial dan Percabangan (Branching)
Dalam pemrograman, program pada umumnya dieksekusi secara berurutan baris demi baris dari atas ke bawah (runtunan/sequential flow). Namun, untuk menyelesaikan masalah yang lebih kompleks, program sering kali memerlukan pengambilan keputusan (selection/branching). Struktur percabangan memungkinkan alur eksekusi kode dialihkan ke blok instruksi tertentu berdasarkan evaluasi suatu kondisi ekspresi boolean yang menghasilkan nilai true atau false.

### B. Operator Relasional dan Logika
1. *Operator Relasional (Perbandingan)*: Digunakan untuk membandingkan dua buah nilai atau variabel. Operator yang didukung oleh Go meliputi:
   - == (sama dengan)
   - != (tidak sama dengan)
   - < (kurang dari) dan > (lebih dari)
   - <= (kurang dari atau sama dengan) dan >= (lebih dari atau sama dengan)
2. *Operator Logika*: Digunakan untuk mengombinasikan satu atau lebih ekspresi boolean:
   - && (AND): Menghasilkan true hanya jika kedua kondisi yang dihubungkan bernilai true.
   - || (OR): Menghasilkan true jika salah satu atau kedua kondisi bernilai true.
   - ! (NOT): Membalikkan nilai kebenaran suatu ekspresi (!true menjadi false, dan sebaliknya).
3. *Short-circuit Evaluation*: Go menerapkan evaluasi sirkuit pendek pada operator logika. Pada operasi &&, jika kondisi pertama bernilai false, maka kondisi kedua tidak akan dievaluasi lagi karena hasilnya sudah pasti false. Begitu pula pada operasi ||, jika kondisi pertama bernilai true, kondisi kedua tidak akan diperiksa.

### C. Struktur Percabangan if, else if, dan else
Struktur if mengeksekusi blok kode di dalamnya apabila kondisi pengujian terpenuhi (true). Jika kondisi tidak terpenuhi, program dapat melanjutkan pengecekan ke blok else if berikutnya. Blok else bersifat opsional dan bertindak sebagai penampung akhir yang dieksekusi ketika semua kondisi di atasnya bernilai false.
Selain itu, Go juga mendukung *percabangan bersarang (*nested if)**, yaitu penempatan pernyataan if di dalam blok if atau else lain untuk menangani alur keputusan bertingkat (seperti sistem menu dan validasi lanjutan).

### D. Struktur Pemilihan switch - case
Pernyataan switch menyediakan cara yang lebih rapi dan terstruktur untuk menulis percabangan banyak cabang (multiple branching) dibandingkan rantai if - else if yang panjang.
1. *Switch Berbasis Ekspresi*: Mencocokkan nilai variabel secara langsung dengan klausa case.
2. *Tagless Switch*: Pernyataan switch yang ditulis tanpa parameter variabel di samping kata kunci switch. Pada jenis ini, setiap klausa case mengevaluasi ekspresi boolean tersendiri.
3. Di Go, setelah suatu case terpenuhi dan blok kodenya dieksekusi, aliran program langsung keluar dari blok switch secara otomatis tanpa perlu menambahkan instruksi break manual. Blok default digunakan untuk menangani kondisi yang tidak tercakup oleh seluruh case yang ada.

## Guided

### 1. grade.go

go
package main

import "fmt"

func main() {
	var nama1 string
	var nilai int

	fmt.Print("Masukkan nama siswa: ")
	fmt.Scan(&nama1)
	fmt.Print("Masukkan nilai siswa: ")
	fmt.Scan(&nilai)

	var grade string

	if nilai >= 90 && nilai <= 100 {
		grade = "A"
	} else if nilai >= 80 {
		grade = "B"
	} else if nilai >= 70 {
		grade = "C"
	} else if nilai >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Println(nama1, "mendapatkan nilai", grade)

#### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/guided/grade/output.png)


#### Deskripsi
Program ini bertujuan untuk menentukan predikat atau grade kelulusan siswa berdasarkan nilai angka yang dimasukkan. Pengguna diminta memasukkan nama siswa bertipe string dan nilai siswa bertipe int. Program mengevaluasi nilai tersebut menggunakan struktur percabangan bertingkat if - else if - else:
- Jika nilai berada pada rentang 90 hingga 100 (nilai >= 90 && nilai <= 100), maka grade bernilai "A".
- Jika tidak, tetapi nilai $\ge$ 80, maka grade bernilai "B".
- Jika tidak, tetapi nilai $\ge$ 70, maka grade bernilai "C".
- Jika tidak, tetapi nilai $\ge$ 60, maka grade bernilai "D".
- Jika nilai kurang dari 60, maka alur masuk ke blok else dan grade bernilai "F".

Di akhir eksekusi, program menampilkan nama siswa beserta grade yang didapatkan ke layar.

### 2. penilaian.go

go
package main

import "fmt"

func main() {
	var opsi int

	fmt.Println("============== Menu ==============")
	fmt.Println("1. Sistem penilaian")
	fmt.Println("0. Keluar")
	fmt.Print("Pilih opsi: ")
	fmt.Scan(&opsi)

	if opsi == 1 {
		var nama1 string
		var nilai int

		fmt.Print("Masukkan nama siswa: ")
		fmt.Scan(&nama1)
		fmt.Print("Masukkan nilai siswa: ")
		fmt.Scan(&nilai)

		var grade string

		if nilai >= 90 && nilai <= 100 {
			grade = "A"
		} else if nilai >= 80 {
			grade = "B"
		} else if nilai >= 70 {
			grade = "C"
		} else if nilai >= 60 {
			grade = "D"
		} else {
			grade = "F"
		}

		fmt.Println(nama1, "mendapatkan nilai", grade)
	} else if opsi == 0 {
		fmt.Println("Keluar dari program.")
	} else {
		fmt.Println("Pilihan tidak valid.")
	}

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/guided/penilaian/output.png)


#### Deskripsi
Program ini mengimplementasikan antarmuka menu interaktif dengan memanfaatkan struktur percabangan bersarang (nested if). Alur program adalah sebagai berikut:
1. Program menampilkan menu pilihan awal (opsi 1 untuk sistem penilaian dan opsi 0 untuk keluar), kemudian membaca input opsi bertipe int.
2. Struktur percabangan tingkat pertama (if - else if - else) mengevaluasi pilihan menu:
   - Jika pengguna memilih opsi == 1, program masuk ke blok nested if untuk meminta input nama siswa dan nilai angka, lalu mengklasifikasikan nilai ke dalam grade A, B, C, D, atau F sama seperti logika pada program grade.go.
   - Jika pengguna memilih opsi == 0, program mencetak pesan "Keluar dari program.".
   - Jika opsi yang dimasukkan selain 1 atau 0, blok else akan dijalankan dan menampilkan pesan "Pilihan tidak valid.".

### 3. klasifikasi.go

go
package main

import "fmt"

func main() {
	var usia int
	var gaji int

	fmt.Print("Masukkan usia: ")
	fmt.Scan(&usia)
	fmt.Print("Masukkan gaji tahunan (juta): ")
	fmt.Scan(&gaji)

	switch {
	case usia < 18:
		fmt.Println("Masih sekolah")
	case usia >= 18 && usia <= 25 && gaji >= 50:
		fmt.Println("Muda sukses")
	case usia >= 18 && usia <= 25 && gaji < 50:
		fmt.Println("Masih belajar hidup")
	case usia >= 26 && usia <= 40 && gaji >= 100:
		fmt.Println("Pekerja mapan")
	case usia >= 26 && usia <= 40 && gaji < 100:
		fmt.Println("Perlu perbaikan karier")
	case usia > 40 && gaji >= 150:
		fmt.Println("Profesional berpengalaman")
	case usia > 40 && gaji < 150:
		fmt.Println("Perlu evaluasi finansial")
	}

##### Output
<!-- Path relatif dari readme.md ke folder unguided/[nama_soal]/output.png, contoh: -->
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/guided/klasifikasi/output.png)

#### Deskripsi
Program ini melakukan klasifikasi kategori tahapan kehidupan, karier, dan finansial seseorang berdasarkan masukan usia (tipe int) dan gaji tahunan dalam satuan juta (tipe int). Alur penentuan kategori menggunakan struktur tagless switch (switch tanpa ekspresi variabel di awal) dengan kombinasi operator relasional dan logika &&:
- usia < 18: Menampilkan "Masih sekolah".
- usia >= 18 && usia <= 25: Jika gaji >= 50 juta dicetak "Muda sukses", sedangkan jika gaji < 50 juta dicetak "Masih belajar hidup".
- usia >= 26 && usia <= 40: Jika gaji >= 100 juta dicetak "Pekerja mapan", sedangkan jika gaji < 100 juta dicetak "Perlu perbaikan karier".
- usia > 40: Jika gaji >= 150 juta dicetak "Profesional berpengalaman", sedangkan jika gaji < 150 juta dicetak "Perlu evaluasi finansial".

Penggunaan tagless switch di sini membuat kode evaluasi multi-kondisi menjadi lebih rapi, terstruktur, dan mudah dibaca dibandingkan rantai percabangan if - else if yang panjang.

## Unguided

### 1. grade.go

go
package main

import "fmt"

func main() {
	var nama1, nama2 string
	var nilai int

	fmt.Print("Masukkan nama siswa: ")
	fmt.Scan(&nama1, &nama2)
	fmt.Print("Masukkan nilai siswa: ")
	fmt.Scan(&nilai)

	var grade string

	switch {
	case nilai >= 90 && nilai <= 100:
		grade = "A"
	case nilai >= 80:
		grade = "B"
	case nilai >= 70:
		grade = "C"
	case nilai >= 60:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Println(nama1, nama2, "mendapatkan nilai", grade)
}


##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/soal1/output.png)

#### Deskripsi
Program ini merupakan variasi dari sistem penentuan predikat nilai siswa yang mengimplementasikan struktur pemilihan tagless switch:
1. Program membaca nama siswa yang terdiri dari dua suku kata menggunakan dua variabel string (nama1 dan nama2) serta sebuah variabel nilai bertipe integer.
2. Penentuan huruf mutu (grade) dilakukan menggunakan switch tanpa ekspresi. Evaluasi kondisi dilakukan secara terurut dari atas ke bawah:
   - case nilai >= 90 && nilai <= 100: Menetapkan grade = "A".
   - case nilai >= 80: Menetapkan grade = "B".
   - case nilai >= 70: Menetapkan grade = "C".
   - case nilai >= 60: Menetapkan grade = "D".
   - default: Menangani nilai di bawah 60 atau nilai lain di luar rentang di atas dengan menetapkan grade = "F".
3. Setelah proses evaluasi selesai, program mencetak nama lengkap beserta grade siswa yang bersangkutan.

### 2. pajak.go

go
package main

import "fmt"

func main() {
	var penghasilan float64

	fmt.Print("Masukkan penghasilan (juta): ")
	fmt.Scan(&penghasilan)

	var pajak float64

	if penghasilan <= 50 {
		pajak = 0.05 * penghasilan
	} else if penghasilan <= 100 {
		pajak = (0.05 * 50) + (0.10 * (penghasilan - 50))
	} else if penghasilan <= 200 {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * (penghasilan - 100))
	} else {
		pajak = (0.05 * 50) + (0.10 * 50) + (0.15 * 100) + (0.20 * (penghasilan - 200))
	}

	fmt.Println(pajak)
}


##### Output
![Screenshot Output Unguided](/praktikum/04-runtunan-sekuensi/unguided/soal2/output.png)

#### Deskripsi
Program ini menghitung total beban pajak tahunan menggunakan skema tarif pajak progresif berjenjang (bracketed tax rate) berdasarkan masukan penghasilan tahunan bertipe float64 (dalam satuan juta rupiah). Logika perhitungan menggunakan struktur percabangan if - else if - else:
- *Lapisan 1 ($\le$ 50 juta):* Dikenakan tarif 5%, sehingga perhitungannya 0.05 * penghasilan.
- *Lapisan 2 (> 50 s.d. 100 juta):* Dikenakan tarif 5% untuk 50 juta pertama, ditambah 10% untuk kelebihan penghasilan di atas 50 juta: (0.05 * 50) + (0.10 * (penghasilan - 50)).
- *Lapisan 3 (> 100 s.d. 200 juta):* Dikenakan tarif 5% untuk 50 juta pertama, 10% untuk 50 juta kedua, dan 15% untuk kelebihan di atas 100 juta: (0.05 * 50) + (0.10 * 50) + (0.15 * (penghasilan - 100)).
- *Lapisan 4 (> 200 juta):* Dikenakan akumulasi tarif penuh dari lapisan 1 hingga 3 ditambah 20% untuk kelebihan penghasilan di atas 200 juta: (0.05 * 50) + (0.10 * 50) + (0.15 * 100) + (0.20 * (penghasilan - 200)).

Hasil perhitungan akhir yang tersimpan pada variabel pajak kemudian dicetak langsung ke layar.

## Kesimpulan
Dari praktikum Modul 04 mengenai runtunan sekuensial dan struktur percabangan dalam bahasa pemrograman Go ini, dapat disimpulkan bahwa:
1. Alur eksekusi sekuensial (berurutan) dapat dikendalikan secara dinamis menggunakan struktur kontrol percabangan (selection) berdasarkan evaluasi kondisi boolean (true atau false).
2. Pemahaman mengenai operator perbandingan (==, !=, <, >, <=, >=) dan operator logika (&&, ||, !), termasuk sifat short-circuit evaluation, sangat krusial dalam menyusun logika kondisi yang akurat pada program.
3. Struktur percabangan if - else if - else dan nested if (percabangan bersarang) efektif digunakan untuk memecahkan permasalahan yang melibatkan alur bertingkat, seperti pembuatan sistem menu interaktif dan perhitungan bertingkat (seperti skema pajak progresif).
4. Struktur pemilihan switch-case di bahasa Go memiliki fleksibilitas tinggi, khususnya fitur tagless switch yang mampu mengevaluasi kondisi majemuk secara elegan dan bersih (clean code). Fitur ini mengeliminasi kebutuhan kata kunci break manual serta menyediakan blok default untuk menangani kondisi di luar ekspektasi input.
