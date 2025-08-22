package pegawai

import (
	"fmt"
)

// Pegawai represents an employee with basic information
type Pegawai struct {
	Nama        string  // Employee name
	Posisi      string  // Employee position
	GajiBulanan float64 // Monthly salary
}

// InformasiPegawai defines the interface for displaying employee information
type InformasiPegawai interface {
	TampilkanInformasi()
}

// HitungGajiTahunan calculates the annual salary for an employee
// Formula: Monthly salary * 12 months
func (p Pegawai) HitungGajiTahunan() float64 {
	return p.GajiBulanan * 12
}

// TampilkanInformasi displays employee information to the console
// This method implements the InformasiPegawai interface
func (p Pegawai) TampilkanInformasi() {
	fmt.Println("=== INFORMASI PEGAWAI ===")
	fmt.Printf("Nama: %s\n", p.Nama)
	fmt.Printf("Posisi: %s\n", p.Posisi)
	fmt.Printf("Gaji Bulanan: Rp %.2f\n", p.GajiBulanan)
	fmt.Printf("Gaji Tahunan: Rp %.2f\n", p.HitungGajiTahunan())
	fmt.Println("========================")
}
