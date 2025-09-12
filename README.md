# 🏢 Go Pegawai - Employee Management System

A simple yet comprehensive employee management application built with Go and GORM, demonstrating full CRUD operations with SQLite database integration.

## 📋 Table of Contents

- [Features](#-features)
- [Tech Stack](#-tech-stack)
- [Prerequisites](#-prerequisites)
- [Installation](#-installation)
- [Usage](#-usage)
- [Project Structure](#-project-structure)
- [Database Schema](#-database-schema)
- [API Documentation](#-api-documentation)
- [Examples](#-examples)
- [Contributing](#-contributing)
- [License](#-license)

## ✨ Features

- **Complete CRUD Operations**: Create, Read, Update, and Delete employee records
- **Database Integration**: SQLite database with GORM ORM
- **Annual Salary Calculation**: Automatic calculation of yearly salary (12 × monthly salary)
- **Data Persistence**: SQLite database for reliable data storage
- **Soft Deletes**: Safe deletion with GORM's built-in soft delete functionality
- **Clean Architecture**: Organized code structure with separate packages
- **Error Handling**: Comprehensive error handling for all operations
- **Console Interface**: User-friendly console output with formatted display

## 🛠 Tech Stack

- **Language**: Go 1.21+
- **ORM**: GORM v1.31.0
- **Database**: SQLite 3
- **Database Driver**: sqlite3 v1.14.22

## 📋 Prerequisites

Before running this application, make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.21 or later)
- Git (for cloning the repository)

## 🚀 Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/muslchn/go-pegawai.git
   cd go-pegawai
   ```

2. **Install dependencies**

   ```bash
   go mod tidy
   ```

3. **Run the application**

   ```bash
   go run main.go
   ```

## 💻 Usage

The application runs a demonstration of all CRUD operations automatically:

```bash
go run main.go
```

### Application Flow

1. **Database Initialization**: Creates SQLite database and migrates schema
2. **CREATE**: Adds 5 sample employees to the database
3. **READ**: Displays all employees with salary calculations
4. **UPDATE**: Modifies an employee's salary and shows the changes
5. **DELETE**: Removes an employee and displays updated list
6. **SUMMARY**: Shows final employee list with statistics

## 📁 Project Structure

```text
go-pegawai/
├── main.go                 # Main application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── pegawai.db             # SQLite database (auto-generated)
├── pegawai/               # Employee package
│   └── data.go            # Employee struct and CRUD operations
└── README.md              # Project documentation
```

## 🗄 Database Schema

### Pegawai Table

| Column      | Type     | Constraints           | Description                    |
|-------------|----------|-----------------------|--------------------------------|
| id          | INTEGER  | PRIMARY KEY, AUTO_INC | Unique employee identifier     |
| created_at  | DATETIME | NOT NULL              | Record creation timestamp      |
| updated_at  | DATETIME | NOT NULL              | Record last update timestamp   |
| deleted_at  | DATETIME | NULL                  | Soft delete timestamp          |
| nama        | VARCHAR  | NOT NULL, SIZE(100)   | Employee full name             |
| posisi      | VARCHAR  | NOT NULL, SIZE(100)   | Employee position/role         |
| gaji_bulanan| REAL     | NOT NULL              | Monthly salary amount          |

## 📚 API Documentation

### Core Functions

#### Database Operations

```go
// Initialize database connection
func InitDatabase() error

// Create new employee
func CreatePegawai(pegawai *Pegawai) error

// Get all employees
func GetAllPegawai() ([]Pegawai, error)

// Get employee by ID
func GetPegawaiByID(id uint) (*Pegawai, error)

// Update employee salary
func UpdatePegawaiSalary(id uint, newSalary float64) error

// Delete employee (soft delete)
func DeletePegawai(id uint) error

// Display all employees with formatting
func DisplayAllPegawai() error
```

#### Employee Methods

```go
// Calculate annual salary
func (p Pegawai) HitungGajiTahunan() float64

// Display employee information
func (p Pegawai) TampilkanInformasi()
```

### Data Structures

```go
type Pegawai struct {
    gorm.Model           // ID, CreatedAt, UpdatedAt, DeletedAt
    Nama        string   `gorm:"size:100;not null" json:"nama"`
    Posisi      string   `gorm:"size:100;not null" json:"posisi"`
    GajiBulanan float64  `gorm:"not null" json:"gaji_bulanan"`
}
```

## 📖 Examples

### Sample Output

```text
🏢 APLIKASI MANAJEMEN DATA PEGAWAI 🏢
=====================================

🔧 DEMO OPERASI CRUD PEGAWAI
============================

📝 STEP 1: MENAMBAH 5 DATA PEGAWAI
----------------------------------
✅ Pegawai berhasil ditambahkan dengan ID: 1
✅ Pegawai berhasil ditambahkan dengan ID: 2
...

📊 STEP 2: MENAMPILKAN SEMUA DATA PEGAWAI
------------------------------------------
👤 PEGAWAI 1:
=== INFORMASI PEGAWAI ===
ID: 1
Nama: Ahmad Wijaya
Posisi: Software Engineer
Gaji Bulanan: Rp 8,500,000.00
Gaji Tahunan: Rp 102,000,000.00
========================
```

### Sample Employee Data

| ID | Name           | Position          | Monthly Salary  | Annual Salary   |
|----|----------------|-------------------|-----------------|-----------------|
| 1  | Ahmad Wijaya   | Software Engineer | Rp 8,500,000    | Rp 102,000,000  |
| 2  | Siti Nurhaliza | Product Manager   | Rp 12,000,000   | Rp 144,000,000  |
| 3  | Budi Santoso   | UI/UX Designer    | Rp 7,500,000    | Rp 90,000,000   |
| 4  | Rina Marlina   | Data Analyst      | Rp 9,000,000    | Rp 108,000,000  |
| 5  | Doni Prasetyo  | DevOps Engineer   | Rp 10,500,000   | Rp 126,000,000  |

## 🔍 Key Features Demonstrated

### CRUD Operations

- ✅ **Create**: Insert 5 employees with complete information
- ✅ **Read**: Retrieve and display all employee data
- ✅ **Update**: Modify employee salary with before/after comparison
- ✅ **Delete**: Remove employee record with confirmation

### Business Logic

- ✅ **Annual Salary Calculation**: 12 × monthly salary
- ✅ **Data Validation**: GORM constraints and error handling
- ✅ **Soft Deletes**: Safe record removal
- ✅ **Statistics**: Total and average salary calculations

## 🛡 Error Handling

The application includes comprehensive error handling for:

- Database connection failures
- Record not found scenarios
- Invalid data constraints
- SQL operation errors

## 🧪 Testing

To run the application and see all operations:

```bash
# Run the main demo
go run main.go

# Check if database was created
ls -la pegawai.db
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 📞 Contact

- **Author**: [muslchn](https://github.com/muslchn)
- **Repository**: [go-pegawai](https://github.com/muslchn/go-pegawai)

## 🙏 Acknowledgments

- [GORM](https://gorm.io/) - The fantastic ORM library for Go
- [SQLite](https://www.sqlite.org/) - Lightweight database engine
- Go community for excellent documentation and support

---

**Note**: This is a prototype application for educational purposes, demonstrating Go programming concepts, GORM usage, and database operations. The application uses hardcoded data for demonstration purposes.
