# DevMak Bank CLI

---
- Version: 1.0
- Language: Go
- Architecture: Domain-Driven Design(DDD) & Hexagonal Architecture (Ports & Adapters)
- Interface: Command Line Interface (CLI)
- Storage: JSON Files

![Go](https://img.shields.io/badge/Go-1.25.5-blue)
![Architecture](https://img.shields.io/badge/Architecture-Hexagonal-success)
![Storage](https://img.shields.io/badge/Storage-JSON-orange)
![License](https://img.shields.io/badge/License-MIT-green)

---

## 1. Project Overview
DevMak Bank CLI is a banking management system that simulates the workflow of a real banking institution.
Unlike traditional CLI CRUD applications, this system models actual banking operations such as:
- User registration
- Manager approval workflows
- Role-Based Access Control (RBAC)
- Account creation requests
- Customer onboarding
- Banking transactions
- Notifications
- Session management

The project is designed using Domain-Driven Design & Hexagonal Architecture to ensure that the business logic remains independent of the user interface and storage mechanism.
Future versions can replace the CLI with a REST API or JSON storage with PostgreSQL without changing the core domain logic.

---

## Features

### Authentication
- User registration
- Secure login
- Password hashing
- Session management
- Logout

### User Management
- Guest accounts
- Customer accounts
- Manager accounts
- Admin promotion
- Role-Based Access Control (RBAC)

### Account Requests
- Request a bank account
- Manager approval workflow
- Manager rejection workflow
- Automatic account creation after approval

### Banking
- Create bank accounts
- Deposit funds
- Withdraw funds
- Transfer funds
- Check balances
- View account information

### Transactions
- Deposit transactions
- Withdrawal transactions
- Transfer transactions
- Transaction history
- Account statements

### Architecture
- Hexagonal Architecture
- Domain-Driven Design (DDD)
- Repository Pattern
- Service Layer
- JSON persistence
- UUID identifiers
- Rich domain models

---
## Project Structure
---
```
internal
├── application
├── delivery
│   ├── cli
│   └── http
├── domain
│   ├── account
│   ├── accountrequest
│   ├── transaction
│   └── user
├── infrastructure
└── storage
```

---

# Running the Project
---

Clone the repository

```bash
git clone https://github.com/Mark23Dev/banking_system.git
```

Navigate to the project

```bash
cd banking_system
```

Run

```bash
go run ./internal/cmd/app
```

---

# Example

```
================ DEVMAK BANK =================

guest@devmak-bank > signup

✓ Account created successfully.

guest@devmak-bank > login

✓ Login successful.

alice@devmak-bank > request-account

✓ Account request submitted.

manager@devmak-bank > approve 1

✓ Request approved.

alice@devmak-bank > deposit

Enter amount: 500

✓ Deposit successful.

alice@devmak-bank > transfer

From account: 3479439303
To account: 1234567890
Amount: 100
Description: Rent

✓ Transfer completed.

alice@devmak-bank > statement

══════════════════════════════════════════════════════════
                ACCOUNT STATEMENT
══════════════════════════════════════════════════════════

Account Number : 3479439303
Type           : Checking
Status         : Active
Balance        : KES 400

══════════════════════════════════════════════════════════
```

---

# Learning Objectives

This project was built to deepen understanding of:

- Go fundamentals
- Domain-Driven Design
- Hexagonal Architecture
- Repository Pattern
- Service Layer
- Authentication
- Authorization
- Banking domain modelling
- Backend system design

---

# Future Improvements

- PostgreSQL persistence
- ACID database transactions
- Account numbers generated with checksums
- Audit logging
- Transaction search
- Interest calculations
- Loans
- Fixed deposits
- ATM simulation
- REST API
- Web dashboard
- Docker support
- Unit and integration tests

---

## Author

**Mark Saruni**

Backend Engineer | Go Developer | Systems Programming Enthusiast

- GitHub: https://github.com/Mark23Dev
- LinkedIn: https://www.linkedin.com/in/mark-saruni
- Portfolio: https://www.devmakspace.dev

