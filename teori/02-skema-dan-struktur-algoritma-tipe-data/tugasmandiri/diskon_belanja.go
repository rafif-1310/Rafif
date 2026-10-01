package main

import "fmt"

func main() {
    // variabel
    var harga, persen int
    var potongan, hargaAkhir float64

    fmt.Scan(&harga, &persen)

    // Proses perhitungan
    potongan = float64(harga * persen) / 100.0
    hargaAkhir = float64(harga) - potongan

    fmt.Printf("%.1f\n", hargaAkhir)
}