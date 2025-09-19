# 🏢 Go Pegawai - Employee Management System

A comprehensive employee management application built with Go and GORM, demonstrating full CRUD operations with MariaDB database integration, environment-based configuration, and enterprise-grade best practices.

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
- **MariaDB Integration**: Production-ready MariaDB database with GORM ORM
- **Environment Configuration**: Secure configuration management with environment variables
- **Connection Pooling**: Optimized database connections for performance
- **Annual Salary Calculation**: Automatic calculation of yearly salary (12 × monthly salary)
- **Data Persistence**: MariaDB database for reliable, scalable data storage
- **Soft Deletes**: Safe deletion with GORM's built-in soft delete functionality
- **Clean Architecture**: Organized code structure with separate packages
- **Error Handling**: Comprehensive error handling for all operations
- **Security Best Practices**: Environment variables for sensitive data
- **Console Interface**: User-friendly console output with formatted display

## 🛠 Tech Stack

- **Language**: Go 1.21+
- **ORM**: GORM v1.31.0
- **Database**: MariaDB/MySQL
- **Database Driver**: MySQL Driver v1.8.1
- **Configuration**: godotenv v1.5.1 for environment variable management

## 📋 Prerequisites

Before running this application, make sure you have the following installed:

- [Go](https://golang.org/dl/) (version 1.21 or later)
- [MariaDB](https://mariadb.org/download/) or [MySQL](https://dev.mysql.com/downloads/) (version 8.0 or later)
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

3. **Set up MariaDB database**

   ```sql
   -- Connect to MariaDB as root
   mysql -u root -p

   -- Create database
   CREATE DATABASE pegawai_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

   -- Create user (optional, for security)
   CREATE USER 'pegawai_user'@'localhost' IDENTIFIED BY 'your_secure_password';
   GRANT ALL PRIVILEGES ON pegawai_db.* TO 'pegawai_user'@'localhost';
   FLUSH PRIVILEGES;
   ```

4. **Configure environment variables**

   Copy the example environment file and update with your database credentials:

   ```bash
   cp .env.example .env
   ```

   Edit `.env` file with your database configuration:

   ```env
   # Database Configuration
   DB_HOST=localhost
   DB_PORT=3306
   DB_USERNAME=root
   DB_PASSWORD=your_password
   DB_NAME=pegawai_db

   # Application Configuration
   APP_PORT=8080
   APP_ENV=development

   # Database Connection Settings
   DB_MAX_OPEN_CONNS=25
   DB_MAX_IDLE_CONNS=10
   DB_CONN_MAX_LIFETIME=300
   ```

5. **Run the application**

   ```bash
   go run main.go
   ```

## 💻 Usage

The application runs a demonstration of all CRUD operations automatically:

```bash
go run main.go
```

### Application Flow

1. **Configuration Loading**: Loads environment variables from `.env` file
2. **Database Initialization**: Connects to MariaDB and migrates schema
3. **CREATE**: Adds 5 sample employees to the database
4. **READ**: Displays all employees with salary calculations
5. **UPDATE**: Modifies an employee's salary and shows the changes
6. **DELETE**: Removes an employee and displays updated list
7. **SUMMARY**: Shows final employee list with statistics

## 📁 Project Structure

```text
go-pegawai/
├── main.go                 # Main application entry point
├── go.mod                  # Go module definition
├── go.sum                  # Go module checksums
├── .env                    # Environment variables (not in git)
├── .env.example            # Environment variables template
├── .gitignore              # Git ignore file
├── config/                 # Configuration package
│   └── config.go           # Environment variable loading and config management
├── pegawai/                # Employee package
│   └── data.go             # Employee struct and CRUD operations
└── README.md               # Project documentation
```

## 🗄 Database Schema

### Pegawai Table (MariaDB)

| Column       | Type         | Constraints           | Description                    |
|--------------|------------- |-----------------------|--------------------------------|
| id           | BIGINT       | PRIMARY KEY, AUTO_INC | Unique employee identifier     |
| created_at   | DATETIME(3)  | NOT NULL              | Record creation timestamp      |
| updated_at   | DATETIME(3)  | NOT NULL              | Record last update timestamp   |
| deleted_at   | DATETIME(3)  | NULL                  | Soft delete timestamp          |
| nama         | VARCHAR(100) | NOT NULL              | Employee full name             |
| posisi       | VARCHAR(100) | NOT NULL              | Employee position/role         |
| gaji_bulanan | DECIMAL(15,2)| NOT NULL              | Monthly salary amount          |

### Database Configuration

- **Character Set**: utf8mb4
- **Collation**: utf8mb4_unicode_ci
- **Engine**: InnoDB (default)
- **Connection Pool**: Configurable via environment variables

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

#### Configuration Management

```go
// Load configuration from environment variables
func LoadConfig() (*Config, error)

// Get current configuration instance
func GetConfig() *Config

// Configuration structures
type Config struct {
    Database DatabaseConfig
    App      AppConfig
}

type DatabaseConfig struct {
    Host            string
    Port            int
    Username        string
    Password        string
    Name            string
    MaxOpenConns    int
    MaxIdleConns    int
    ConnMaxLifetime int
}

type AppConfig struct {
    Port string
    Env  string
}
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
MariaDB database connected successfully to localhost:3306/pegawai_db
Database schema migrated successfully

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
- Environment variable loading errors
- Record not found scenarios
- Invalid data constraints
- SQL operation errors
- Configuration validation errors

## 🔧 Configuration

### Environment Variables

The application uses environment variables for configuration. Copy `.env.example` to `.env` and configure:

```env
# Database Configuration
DB_HOST=localhost                # Database host
DB_PORT=3306                     # Database port
DB_USERNAME=root                 # Database username
DB_PASSWORD=your_password        # Database password
DB_NAME=pegawai_db              # Database name

# Application Configuration
APP_PORT=8080                    # Application port
APP_ENV=development              # Environment (development/production)

# Database Connection Settings
DB_MAX_OPEN_CONNS=25            # Maximum open connections
DB_MAX_IDLE_CONNS=10            # Maximum idle connections
DB_CONN_MAX_LIFETIME=300        # Connection lifetime in seconds
```

### Security Best Practices

- ✅ Environment variables for sensitive data
- ✅ `.env` file excluded from version control
- ✅ `.env.example` template for setup
- ✅ Database connection pooling
- ✅ Proper error handling without exposing internals

## 🧪 Testing

To run the application and see all operations:

```bash
# Ensure MariaDB is running
sudo systemctl start mariadb  # or: brew services start mariadb

# Create database
mysql -u root -p -e "CREATE DATABASE IF NOT EXISTS pegawai_db;"

# Run the main demo
go run main.go

# Check database tables (optional)
mysql -u root -p pegawai_db -e "SHOW TABLES; DESCRIBE pegawais;"
```

### Development Commands

```bash
# Build the application
go build

# Run with specific environment
APP_ENV=production go run main.go

# Check dependencies
go mod tidy
go mod verify

# Format code
go fmt ./...
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
- [MariaDB](https://mariadb.org/) - Reliable and scalable database
- [godotenv](https://github.com/joho/godotenv) - Environment variable management
- Go community for excellent documentation and support

---

**Note**: This application demonstrates Go programming concepts, GORM usage, MariaDB integration, and enterprise-grade configuration management. The application includes production-ready features like connection pooling, environment-based configuration, and comprehensive error handling.
