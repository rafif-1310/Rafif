# <h1 align="center">Tugas Pendahuluan Modul 06 - Perulangan While</h1>
<p align="center">Rafif Ramadhan - 109092630001</p>

### 1. Perbedaan if, if-else, else if, dan switch pada bahasa pemrograman Go

a. If

`if` digunakan untuk menjalankan suatu perintah apabila kondisi yang diberikan bernilai benar (`true`).

Contoh:

```go
if nilai >= 75 {
    fmt.Println("Lulus")
}
```

b. If-else

`if-else` digunakan untuk memilih satu dari dua kemungkinan berdasarkan kondisi. Jika kondisi `if` benar, perintah `if` dijalankan. Jika salah, perintah `else` dijalankan.

Contoh:

```go
if nilai >= 75 {
    fmt.Println("Lulus")
} else {
    fmt.Println("Tidak Lulus")
}
```

c. Else if

`else if` digunakan untuk memeriksa beberapa kondisi secara berurutan. Jika kondisi pertama salah, program akan memeriksa kondisi berikutnya.

Contoh:

```go
if nilai >= 90 {
    fmt.Println("A")
} else if nilai >= 80 {
    fmt.Println("B")
} else if nilai >= 70 {
    fmt.Println("C")
} else {
    fmt.Println("D")
}
```

d. Switch

`switch` digunakan untuk memilih dan menjalankan perintah berdasarkan nilai atau kondisi tertentu. Switch cocok digunakan ketika terdapat banyak pilihan.

Contoh:

```go
switch hari {
case 1:
    fmt.Println("Senin")
case 2:
    fmt.Println("Selasa")
case 3:
    fmt.Println("Rabu")
default:
    fmt.Println("Hari tidak ditemukan")
}
```

### 2. Dua perbedaan switch pada Go dibandingkan dengan Java

1. Tidak memerlukan break secara otomatis

   Pada Go, setiap `case` dalam switch tidak otomatis dilanjutkan ke `case` berikutnya (tidak terjadi fall-through). Sementara itu, pada Java, penggunaan `break` biasanya diperlukan agar eksekusi tidak berlanjut ke case berikutnya pada switch tradisional.
2. Switch pada Go lebih fleksibel

   Pada Go, switch dapat digunakan tanpa ekspresi setelah kata `switch`, sehingga bisa memeriksa beberapa kondisi boolean. Sementara itu, switch tradisional pada Java menggunakan nilai atau ekspresi sebagai dasar pencocokan case.

   Contoh switch tanpa ekspresi pada Go:
   ```
   switch {
   case nilai >= 90:
       fmt.Println("A")
   case nilai >= 80:
       fmt.Println("B")
   default:
       fmt.Println("C")
   }
   ```

### Kesimpulan

`if`, `if-else`, dan `else if` digunakan untuk percabangan berdasarkan kondisi, sedangkan `switch` digunakan untuk memilih satu tindakan dari beberapa kemungkinan. Go memiliki switch yang lebih fleksibel dan tidak memerlukan `break` pada setiap `case` seperti switch tradisional Java.
### 2. bilangan.go

```go
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&n)

	if n == 0 {
		fmt.Println("0 adalah bilangan 0")
	} else {
		if n %2 == 0 {
			fmt.Print("genap, ")
		} else {
			fmt.Print("ganjil, ")
		}
	
		if n > 0 {
			fmt.Print("positif")
		} else {
			fmt.Print("negatif")
		}
	
}
}
```

##### Output

![Screenshot Output TP Soal 1](/praktikum/07-percabangan/TP/soal%202/output.png)
#### Deskripsi
Program diatas menentukan apakah nilai yang kita masukkan bernilai positif/negatif atau genap/ganjil


### 3. hari.go

```go
package main

import "fmt"

func main(){
	var hari int

	
	fmt.Scan(&hari)
	
	switch hari {
	case 1:
		fmt.Print("senin")
	case 2:
		fmt.Print("selasa")
	case 3:
		fmt.Print("rabu")
	case 4:
		fmt.Print("kamis")
	case 5:
		fmt.Print("jumat")
	case 6:
		fmt.Print("sabtu(akhir pekan)")
	case 7:
		fmt.Print("minggu(akhir pekan)")
	default:
		fmt.Print("nomot tidak valid")
	}
		
	

}
```

##### Output

![Screenshot Output TP Soal 2](/praktikum/07-percabangan/TP/soal%203/output.png)

#### Deskripsi
program diatas menggunakan switch untuk menentukan hari,user akan memasukkan hari dengan menggunakan nomor 1-7 dengan nomor 1(senin) dan seterusnya


### 4. terbesar.go

```go
package main

import "fmt"

func main() {
    var a, b, c int

    fmt.Print("Masukkan bilangan: ")
    fmt.Scan(&a, &b, &c)

    var terbesar int

    if a >= b && a >= c {
        terbesar = a
    } else if b >= a && b >= c {
        terbesar = b
    } else {
        terbesar = c
    }

    fmt.Println("Bilangan terbesar:", terbesar)
}

```

##### Output

![Screenshot Output TP Soal 3](/praktikum/07-percabangan/TP/soal%204/output.png)

#### Deskripsi
program diatas digunakan untuk menentukan bilangan terbesar yang dimasukkan dari tiga bilangan


## Kesimpulan

Berdasarkan tugas ini, percabangan dalam bahasa Go dapat digunakan untuk menentukan alur program sesuai kondisi dan masukan pengguna. `if`, `else if`, dan `else` berguna untuk memeriksa kondisi seperti menentukan bilangan positif atau negatif, genap atau ganjil, serta mencari nilai terbesar. Sementara itu, `switch` memudahkan pemilihan tindakan dari beberapa nilai, seperti menentukan nama hari berdasarkan angka. Dengan memahami penggunaan struktur percabangan tersebut, program dapat mengambil keputusan secara tepat dan menyajikan hasil yang sesuai.
