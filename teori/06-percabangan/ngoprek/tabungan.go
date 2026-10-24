package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func masukan(msg string) string {
	fmt.Print(msg)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func daftarMenu() {
	fmt.Println("1. Tampilkan daftar tabungan")
	fmt.Println("2. Tambahkan tabungan baru")
	fmt.Println("3. Isi tabungan")
	fmt.Println("4. Ambil isi tabungan")
	fmt.Println("5. Edit tabungan")
	fmt.Println("6. Hapus tabungan")
	fmt.Println("0. Keluar")
}

type tabungan struct {
	nama      string
	target    float64
	terkumpul float64
	status    string
}

func updateStatus(t *tabungan) {
	if t.terkumpul >= t.target && t.target > 0 {
		t.status = "tercapai"
	} else {
		t.status = "proses"
	}
}

func buatTabunganBaru(daftar []tabungan, namaTabungan string, target float64) []tabungan {
	if strings.TrimSpace(namaTabungan) == "" {
		fmt.Println("\nNama tabungan tidak boleh kosong!\n")
		return daftar
	}
	if target <= 0 {
		fmt.Println("\nTarget tabungan harus lebih dari 0!\n")
		return daftar
	}

	tabunganBaru := tabungan{
		nama:      namaTabungan,
		target:    target,
		terkumpul: 0.0,
		status:    "proses",
	}
	fmt.Printf("\nTabungan \"%s\" berhasil dibuat!\n\n", namaTabungan)
	return append(daftar, tabunganBaru)
}

func tampilkanTabungan(daftar []tabungan) {
	if len(daftar) == 0 {
		fmt.Println("\nTabungan belum pernah dibuat.\n")
		return
	}

	fmt.Println()
	fmt.Printf("%-4s %-25s %-18s %-18s %-18s %-10s\n", "No", "Nama", "Target", "Terkumpul", "Kekurangan", "Status")
	fmt.Println(strings.Repeat("=", 95))
	for i, v := range daftar {
		kekurangan := v.target - v.terkumpul
		if kekurangan < 0 {
			kekurangan = 0
		}
		fmt.Printf("%-4d %-25s Rp %-15.2f Rp %-15.2f Rp %-15.2f %-10s\n",
			i+1, v.nama, v.target, v.terkumpul, kekurangan, v.status)
	}
	fmt.Println(strings.Repeat("=", 95))
	fmt.Println()
}

func isiTabungan(daftar []tabungan) []tabungan {
	if len(daftar) == 0 {
		fmt.Println("\nTabungan belum pernah dibuat.\n")
		return daftar
	}

	tampilkanTabungan(daftar)
	noStr := masukan("Pilih nomor tabungan yang ingin diisi (0 untuk batal): ")
	no, err := strconv.Atoi(noStr)
	if err != nil || no < 0 || no > len(daftar) {
		fmt.Println("\nNomor tabungan tidak valid!\n")
		return daftar
	}
	if no == 0 {
		return daftar
	}

	idx := no - 1
	nominalStr := masukan("Masukan nominal yang ingin disetor : Rp ")
	nominal, err := strconv.ParseFloat(nominalStr, 64)
	if err != nil || nominal <= 0 {
		fmt.Println("\nNominal harus berupa angka lebih dari 0!\n")
		return daftar
	}

	daftar[idx].terkumpul += nominal
	updateStatus(&daftar[idx])

	fmt.Printf("\nBerhasil menyetor Rp %.2f ke \"%s\"!\n", nominal, daftar[idx].nama)
	fmt.Printf("Total terkumpul saat ini : Rp %.2f (Target: Rp %.2f, Status: %s)\n\n",
		daftar[idx].terkumpul, daftar[idx].target, daftar[idx].status)
	return daftar
}

func ambilTabungan(daftar []tabungan) []tabungan {
	if len(daftar) == 0 {
		fmt.Println("\nTabungan belum pernah dibuat.\n")
		return daftar
	}

	tampilkanTabungan(daftar)
	noStr := masukan("Pilih nomor tabungan yang ingin diambil (0 untuk batal): ")
	no, err := strconv.Atoi(noStr)
	if err != nil || no < 0 || no > len(daftar) {
		fmt.Println("\nNomor tabungan tidak valid!\n")
		return daftar
	}
	if no == 0 {
		return daftar
	}

	idx := no - 1
	if daftar[idx].terkumpul <= 0 {
		fmt.Printf("\nTabungan \"%s\" belum memiliki saldo untuk ditarik!\n\n", daftar[idx].nama)
		return daftar
	}

	nominalStr := masukan("Masukan nominal yang ingin ditarik : Rp ")
	nominal, err := strconv.ParseFloat(nominalStr, 64)
	if err != nil || nominal <= 0 {
		fmt.Println("\nNominal harus berupa angka lebih dari 0!\n")
		return daftar
	}

	if nominal > daftar[idx].terkumpul {
		fmt.Printf("\nSaldo tidak mencukupi! Saldo saat ini: Rp %.2f\n\n", daftar[idx].terkumpul)
		return daftar
	}

	daftar[idx].terkumpul -= nominal
	updateStatus(&daftar[idx])

	fmt.Printf("\nBerhasil menarik Rp %.2f dari \"%s\"!\n", nominal, daftar[idx].nama)
	fmt.Printf("Sisa saldo saat ini : Rp %.2f (Status: %s)\n\n", daftar[idx].terkumpul, daftar[idx].status)
	return daftar
}

func editTabungan(daftar []tabungan) []tabungan {
	if len(daftar) == 0 {
		fmt.Println("\nTabungan belum pernah dibuat.\n")
		return daftar
	}

	tampilkanTabungan(daftar)
	noStr := masukan("Pilih nomor tabungan yang ingin diedit (0 untuk batal): ")
	no, err := strconv.Atoi(noStr)
	if err != nil || no < 0 || no > len(daftar) {
		fmt.Println("\nNomor tabungan tidak valid!\n")
		return daftar
	}
	if no == 0 {
		return daftar
	}

	idx := no - 1
	fmt.Printf("\n--- Mengedit Tabungan \"%s\" ---\n", daftar[idx].nama)
	fmt.Println("(Tekan Enter / kosongkan jika tidak ingin mengubah nilai)")

	namaBaru := masukan(fmt.Sprintf("Nama baru [%s]: ", daftar[idx].nama))
	if strings.TrimSpace(namaBaru) != "" {
		daftar[idx].nama = strings.TrimSpace(namaBaru)
	}

	targetBaruStr := masukan(fmt.Sprintf("Target baru [Rp %.2f]: Rp ", daftar[idx].target))
	if strings.TrimSpace(targetBaruStr) != "" {
		targetBaru, err := strconv.ParseFloat(targetBaruStr, 64)
		if err != nil || targetBaru <= 0 {
			fmt.Println("Format target salah atau <= 0, target tidak diubah.")
		} else {
			daftar[idx].target = targetBaru
		}
	}

	updateStatus(&daftar[idx])
	fmt.Printf("\nTabungan berhasil diperbarui!\n\n")
	return daftar
}

func hapusTabungan(daftar []tabungan) []tabungan {
	if len(daftar) == 0 {
		fmt.Println("\nTabungan belum pernah dibuat.\n")
		return daftar
	}

	tampilkanTabungan(daftar)
	noStr := masukan("Pilih nomor tabungan yang ingin dihapus (0 untuk batal): ")
	no, err := strconv.Atoi(noStr)
	if err != nil || no < 0 || no > len(daftar) {
		fmt.Println("\nNomor tabungan tidak valid!\n")
		return daftar
	}
	if no == 0 {
		return daftar
	}

	idx := no - 1
	konfirmasi := masukan(fmt.Sprintf("Apakah Anda yakin ingin menghapus tabungan \"%s\"? (y/n): ", daftar[idx].nama))
	if strings.ToLower(konfirmasi) == "y" {
		namaDihapus := daftar[idx].nama
		daftar = append(daftar[:idx], daftar[idx+1:]...)
		fmt.Printf("\nTabungan \"%s\" berhasil dihapus!\n\n", namaDihapus)
	} else {
		fmt.Println("\nPenghapusan dibatalkan.\n")
	}

	return daftar
}

func main() {
	daftarTabungan := []tabungan{}

	fmt.Println("=== Aplikasi Tabungan ===")
	fmt.Println()

	for {
		daftarMenu()

		pilihan, err := strconv.Atoi(masukan("Masukan nomor perintah : "))
		if err != nil {
			fmt.Println("Pilihan harus berupa angka!")
			fmt.Println()
			continue
		}

		if pilihan == 0 {
			fmt.Println("Terima kasih telah menggunakan aplikasi!")
			break
		}

		switch pilihan {
		case 1:
			tampilkanTabungan(daftarTabungan)
		case 2:
			nama := masukan("Masukan nama Tabungan baru : ")
			targetStr := masukan("Masukan target tabungan : Rp ")
			target, err := strconv.ParseFloat(targetStr, 64)

			if err != nil {
				fmt.Println("\nTarget tabungan harus berupa angka!\n")
			} else {
				daftarTabungan = buatTabunganBaru(daftarTabungan, nama, target)
			}
		case 3:
			daftarTabungan = isiTabungan(daftarTabungan)
		case 4:
			daftarTabungan = ambilTabungan(daftarTabungan)
		case 5:
			daftarTabungan = editTabungan(daftarTabungan)
		case 6:
			daftarTabungan = hapusTabungan(daftarTabungan)
		default:
			fmt.Println("\nPilihan belum tersedia.\n")
		}
	}
}