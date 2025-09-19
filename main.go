package main

import (
	"fmt"
	"log"

	"go-pegawai/config"
	"go-pegawai/pegawai"
)

func main() {
	fmt.Println("🏢 APLIKASI MANAJEMEN DATA PEGAWAI 🏢")
	fmt.Println("=====================================")

	// Load configuration from environment variables
	_, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	// Initialize database
	err = pegawai.InitDatabase()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	fmt.Println("\n🔧 DEMO OPERASI CRUD PEGAWAI")
	fmt.Println("============================")

	// Step 1: Create 5 employees (INSERT)
	fmt.Println("\n📝 STEP 1: MENAMBAH 5 DATA PEGAWAI")
	fmt.Println("----------------------------------")

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
		{
			Nama:        "Rina Marlina",
			Posisi:      "Data Analyst",
			GajiBulanan: 9000000.0,
		},
		{
			Nama:        "Doni Prasetyo",
			Posisi:      "DevOps Engineer",
			GajiBulanan: 10500000.0,
		},
	}

	for i, emp := range employees {
		err := pegawai.CreatePegawai(&employees[i])
		if err != nil {
			log.Printf("Error creating employee %s: %v", emp.Nama, err)
		}
	}

	// Step 2: Display all employees (READ)
	fmt.Println("\n📊 STEP 2: MENAMPILKAN SEMUA DATA PEGAWAI")
	fmt.Println("------------------------------------------")

	err = pegawai.DisplayAllPegawai()
	if err != nil {
		log.Printf("Error displaying employees: %v", err)
	}

	// Step 3: Update salary of one employee (UPDATE)
	fmt.Println("\n✏️ STEP 3: UPDATE GAJI PEGAWAI ID 2")
	fmt.Println("-----------------------------------")

	// Get employee before update
	empBefore, err := pegawai.GetPegawaiByID(2)
	if err != nil {
		log.Printf("Error getting employee: %v", err)
	} else {
		fmt.Printf("📋 Data SEBELUM update:\n")
		empBefore.TampilkanInformasi()
	}

	// Update salary
	newSalary := 15000000.0
	err = pegawai.UpdatePegawaiSalary(2, newSalary)
	if err != nil {
		log.Printf("Error updating salary: %v", err)
	}

	// Get employee after update
	empAfter, err := pegawai.GetPegawaiByID(2)
	if err != nil {
		log.Printf("Error getting employee: %v", err)
	} else {
		fmt.Printf("\n📋 Data SETELAH update:\n")
		empAfter.TampilkanInformasi()
	}

	// Step 4: Delete one employee (DELETE)
	fmt.Println("\n🗑️ STEP 4: MENGHAPUS PEGAWAI ID 4")
	fmt.Println("----------------------------------")

	// Get employee before delete
	empToDelete, err := pegawai.GetPegawaiByID(4)
	if err != nil {
		log.Printf("Error getting employee: %v", err)
	} else {
		fmt.Printf("📋 Data pegawai yang akan dihapus:\n")
		empToDelete.TampilkanInformasi()
	}

	// Delete employee
	err = pegawai.DeletePegawai(4)
	if err != nil {
		log.Printf("Error deleting employee: %v", err)
	}

	// Step 5: Display final list (READ)
	fmt.Println("\n📋 STEP 5: DAFTAR PEGAWAI SETELAH SEMUA OPERASI")
	fmt.Println("================================================")

	err = pegawai.DisplayAllPegawai()
	if err != nil {
		log.Printf("Error displaying employees: %v", err)
	}

	fmt.Println("\n🎉 DEMO APLIKASI SELESAI!")
	fmt.Println("Aplikasi berhasil mendemonstrasikan semua operasi CRUD:")
	fmt.Println("✅ CREATE: Menambah 5 data pegawai")
	fmt.Println("✅ READ: Menampilkan semua data pegawai")
	fmt.Println("✅ UPDATE: Mengubah gaji pegawai")
	fmt.Println("✅ DELETE: Menghapus data pegawai")
}
