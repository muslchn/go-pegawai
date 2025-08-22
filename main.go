package main

import (
	"fmt"

	"github.com/muslchn/go-pegawai/pegawai"
)

func main() {
	fmt.Println("🏢 APLIKASI MANAJEMEN DATA PEGAWAI 🏢")
	fmt.Println("=====================================")

	// Create dummy employee data
	employees := []pegawai.Pegawai{
		{
			Nama:        "Ahmad Wijaya",
			Posisi:      "Software Engineer",
			GajiBulanan: 8500000.0,
		},
		{
			Nama:        "Siti Nurhaliza",
			Posisi:      "Product Manager",
			GajiBulanan: 12000000.0,
		},
		{
			Nama:        "Budi Santoso",
			Posisi:      "UI/UX Designer",
			GajiBulanan: 7500000.0,
		},
	}

	// Display all employees information
	for i, emp := range employees {
		fmt.Printf("\n📊 PEGAWAI %d:\n", i+1)

		// Calculate and display annual salary using method
		gajiTahunan := emp.HitungGajiTahunan()
		fmt.Printf("Gaji Tahunan %s: Rp %.2f\n", emp.Nama, gajiTahunan)

		// Use interface to display employee information
		var info pegawai.InformasiPegawai = emp
		info.TampilkanInformasi()
	}

	// Summary statistics
	fmt.Println("\n📈 RINGKASAN STATISTIK:")
	fmt.Println("=======================")

	var totalGajiBulanan, totalGajiTahunan float64
	for _, emp := range employees {
		totalGajiBulanan += emp.GajiBulanan
		totalGajiTahunan += emp.HitungGajiTahunan()
	}

	fmt.Printf("Total Pegawai: %d\n", len(employees))
	fmt.Printf("Total Gaji Bulanan Semua Pegawai: Rp %.2f\n", totalGajiBulanan)
	fmt.Printf("Total Gaji Tahunan Semua Pegawai: Rp %.2f\n", totalGajiTahunan)
	fmt.Printf("Rata-rata Gaji Bulanan: Rp %.2f\n", totalGajiBulanan/float64(len(employees)))
	fmt.Printf("Rata-rata Gaji Tahunan: Rp %.2f\n", totalGajiTahunan/float64(len(employees)))
}
