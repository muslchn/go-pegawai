package pegawai

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Pegawai represents an employee with basic information
type Pegawai struct {
	gorm.Model          // Includes ID, CreatedAt, UpdatedAt, DeletedAt
	Nama        string  `gorm:"size:100;not null" json:"nama"`   // Employee name
	Posisi      string  `gorm:"size:100;not null" json:"posisi"` // Employee position
	GajiBulanan float64 `gorm:"not null" json:"gaji_bulanan"`    // Monthly salary
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
	fmt.Printf("ID: %d\n", p.ID)
	fmt.Printf("Nama: %s\n", p.Nama)
	fmt.Printf("Posisi: %s\n", p.Posisi)
	fmt.Printf("Gaji Bulanan: Rp %.2f\n", p.GajiBulanan)
	fmt.Printf("Gaji Tahunan: Rp %.2f\n", p.HitungGajiTahunan())
	fmt.Println("========================")
}

// Database connection variable
var DB *gorm.DB

// InitDatabase initializes the SQLite database connection
func InitDatabase() error {
	var err error

	// Connect to SQLite database
	DB, err = gorm.Open(sqlite.Open("pegawai.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Auto migrate the schema
	err = DB.AutoMigrate(&Pegawai{})
	if err != nil {
		return fmt.Errorf("failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrated successfully")
	return nil
}

// CreatePegawai creates a new employee record in the database
func CreatePegawai(pegawai *Pegawai) error {
	result := DB.Create(pegawai)
	if result.Error != nil {
		return fmt.Errorf("failed to create pegawai: %v", result.Error)
	}

	fmt.Printf("✅ Pegawai berhasil ditambahkan dengan ID: %d\n", pegawai.ID)
	return nil
}

// GetAllPegawai retrieves all employee records from the database
func GetAllPegawai() ([]Pegawai, error) {
	var pegawais []Pegawai
	result := DB.Find(&pegawais)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to get all pegawai: %v", result.Error)
	}

	return pegawais, nil
}

// GetPegawaiByID retrieves a specific employee by ID
func GetPegawaiByID(id uint) (*Pegawai, error) {
	var pegawai Pegawai
	result := DB.First(&pegawai, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("pegawai with ID %d not found", id)
		}
		return nil, fmt.Errorf("failed to get pegawai: %v", result.Error)
	}

	return &pegawai, nil
}

// UpdatePegawaiSalary updates the monthly salary of an employee
func UpdatePegawaiSalary(id uint, newSalary float64) error {
	result := DB.Model(&Pegawai{}).Where("id = ?", id).Update("gaji_bulanan", newSalary)
	if result.Error != nil {
		return fmt.Errorf("failed to update pegawai salary: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("pegawai with ID %d not found", id)
	}

	fmt.Printf("✅ Gaji pegawai ID %d berhasil diupdate menjadi Rp %.2f\n", id, newSalary)
	return nil
}

// DeletePegawai soft deletes an employee record
func DeletePegawai(id uint) error {
	result := DB.Delete(&Pegawai{}, id)
	if result.Error != nil {
		return fmt.Errorf("failed to delete pegawai: %v", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("pegawai with ID %d not found", id)
	}

	fmt.Printf("✅ Pegawai ID %d berhasil dihapus\n", id)
	return nil
}

// DisplayAllPegawai displays all employees with their information and annual salary
func DisplayAllPegawai() error {
	pegawais, err := GetAllPegawai()
	if err != nil {
		return err
	}

	if len(pegawais) == 0 {
		fmt.Println("📋 Tidak ada data pegawai")
		return nil
	}

	fmt.Printf("\n📋 DAFTAR SEMUA PEGAWAI (Total: %d)\n", len(pegawais))
	fmt.Println("=====================================")

	var totalGajiBulanan, totalGajiTahunan float64

	for i, pegawai := range pegawais {
		fmt.Printf("\n👤 PEGAWAI %d:\n", i+1)
		pegawai.TampilkanInformasi()

		totalGajiBulanan += pegawai.GajiBulanan
		totalGajiTahunan += pegawai.HitungGajiTahunan()
	}

	// Display summary
	fmt.Println("\n📊 RINGKASAN:")
	fmt.Printf("Total Gaji Bulanan: Rp %.2f\n", totalGajiBulanan)
	fmt.Printf("Total Gaji Tahunan: Rp %.2f\n", totalGajiTahunan)
	fmt.Printf("Rata-rata Gaji Bulanan: Rp %.2f\n", totalGajiBulanan/float64(len(pegawais)))
	fmt.Printf("Rata-rata Gaji Tahunan: Rp %.2f\n", totalGajiTahunan/float64(len(pegawais)))

	return nil
}
