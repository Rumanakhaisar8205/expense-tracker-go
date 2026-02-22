# Expense Tracker with Bill Splitting API

A production-ready REST API for tracking shared expenses and optimizing bill settlements among groups, similar to Splitwise.

Built using Go, Gin, GORM, and PostgreSQL.

---

## 🎯 Project Overview

This system allows:

- Users to create groups and add shared expenses
- Members to track who paid and who owes
- The system to calculate optimized settlements
- Minimize the number of transactions required

---

## 🏗️ Architecture

### Technology Stack

- Backend: Go 1.21+
- Framework: Gin
- ORM: GORM
- Database: PostgreSQL
- API Testing: Thunder Client (VS Code Extension)

---

### Key Design Decisions

1. PostgreSQL Database: Chosen for reliability and support for decimal money types
2. GORM ORM: Simplifies database interactions
3. Gin Framework: Provides fast HTTP routing
4. Greedy Algorithm: Used for optimized settlement calculation
5. Decimal Money Type: Prevents floating-point errors

---

## 📁 Project Structure

expense-tracker-go/
│
├── main.go
├── go.mod
├── README.md
│
├── config/
│ └── db.go
│
├── models/
│ ├── user.go
│ ├── group.go
│ ├── expense.go
│ └── expense_share.go
│
├── controllers/
│ ├── user_controller.go
│ ├── group_controller.go
│ ├── expense_controller.go
│ └── settlement_controller.go
│
├── algorithm/
│ └── settlement.go
│
├── docs/
│ ├── settlement.md
│ ├── money_handling.md
│ └── examples.md
│
└── prompts/
└── AI_Assistance_Log.md


---

## 🗄️ Database Schema

### Users Table
- id
- name
- email

### Groups Table
- id
- name

### Expenses Table
- id
- group_id
- paid_by
- amount (DECIMAL(10,2))
- note

### Expense Shares Table
- expense_id
- user_id
- share

All monetary values use DECIMAL(10,2) to ensure financial accuracy.

---

## ⚙️ Getting Started

### Prerequisites

- Go 1.21+
- PostgreSQL 13+
- Git

---

### Installation

#### 1. Clone the Repository

```bash
git clone <your-repo-url>
cd expense-tracker-go
```
2. **Install dependencies**
```bash
go mod tidy
```
3.**Configure Database**
```plain text
host=localhost
user=postgres
password=your_password
dbname=expense_db
port=5432
sslmode=disable
```
4. **Create Database**
Login to PostgreSQL and run:
```sql
CREATE DATABASE expense_db;
```
5.**Run Database Migration**

Tables are automatically created using GORM AutoMigrate when the server starts.
4. **Run the server**
```bash
go run main.go
```

Server starts at `http://localhost:8080`

**API Endpoints**
- **Create User**: `http://localhost:8080/user` (POST)
- **Create Group**: `http://localhost:8080/groups` (POST)
- **Get Users**: `http://localhost:8080/user` (GET)
- **Add Expense**: `http://localhost:8080/Expense` (POST)
- **Get Group Expenses**: `http://localhost:8080/groups/:id/expenses` (GET)
- **Get Settelment**: `http://localhost:8080//groups/:id/settle` (GET)

##  API Endpoints

**🧮 Settlement Algorithm**
The Problem

When multiple users pay different amounts, manually calculating settlements leads to many transactions.

Example:
A paid 600
B paid 200
C paid 0
Without optimization → many transfers.

The Solution
The system calculates net balance for each user:
Balance > 0 → Receives money
Balance < 0 → Pays money
A greedy matching algorithm pairs debtors and creditors.
This minimizes total transactions.

**Algorithm Overview**
Calculate balance for each user
Separate debtors and creditors
Match smallest debts with largest credits
Repeat until balances are zero
Time Complexity: O(n)
Details in docs/settlement.md.
## settlement algorithm

**💰 Money Handling Strategy**
All amounts are stored as:
```code
DECIMAL(10,2)
```
This avoids floating-point rounding errors.

Example:
100.25
500.00

Details in docs/money_handling.md.

**🧪 Testing**
**Manual Testing (Thunder Client)**

APIs were tested using Thunder Client in VS Code.
Testing Flow:
Create users
Create groups
Add expenses
Fetch settlement
**Sample Test Flow**
Create User → POST /users
Create Group → POST /groups
Add Expense → POST /expenses
Get Settlement → GET /groups/1/settle
**📚 Example Scenario**
Trip Group

A paid: 600
B paid: 200
C paid: 0

Settlement:
```code
B → A : 200
C → A : 400
```
**📈 Learning Outcomes**

- Implemented greedy optimization algorithms
- Learned financial data handling using decimals
- Designed REST APIs using Gin
- Worked with PostgreSQL and GORM
- Improved system documentation skills

**🚀 Future Enhancements**

- User authentication (JWT)
- Expense categories
- Notification system
- Mobile application
- Cloud deployment
- Analytics dashboard

**👤 Author**

Rumana
Capstone Project – Infosys Springboard
February 2026

**📜 License**

This project is developed for academic purposes.

**🙏 Acknowledgments**

- Infosys Springboard for guidance
- Go Community
- PostgreSQL Documentation