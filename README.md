# Parking Lot Service

This HTTP service manages parking lots, allowing different types of vehicles to be parked, generating entry tickets, managing slots, and calculating parking fees.

## Technologies Used

- Framework: Echo
- ORM: GORM
- Database: PostgreSQL

## Project Structure

```plaintext
parking-lot-service/
├── main.go
├── handlers/
│   └── parking.go
├── models/
│   └── parking.go
├── db/
│   └── db.go
└── .env


Features
Vehicle Types & Parking Slots:

Motorcycles/scooters, Cars/SUVs, Buses/Trucks with varying slot sizes per parking lot.
Real-time Slot Tracking:

Ability to check free slots dynamically.
Ticket Generation:

Generate entry tickets with vehicle and parking lot details.
Receipt Generation:

Create receipts upon vehicle unparking, calculating fees based on parking duration and tariff.
Tariff Models
Parking Lot A
Motorcycles/scooters: Rs 5/hr
Cars/SUVs: Rs 20.50/hr
Buses/Trucks: Rs 50/hr up to 1 day, then Rs 500/day
Parking Lot B
Motorcycles/scooters: Rs 10.50/hr
Cars/SUVs: Rs 50 for the 1st hour, then Rs 25/hr
Buses/Trucks: Rs 100/hr
Development Setup
Prerequisites
Go 1.18+
PostgreSQL
Docker (optional

Installation
Clone the repository:

git clone https://github.com/yourusername/parking-lot-service.git
cd parking-lot-service

Set up environment variables:
# Example .env file
DB_HOST=localhost
DB_PORT=5432
DB_USER=myusername
DB_PASS=mypassword
DB_NAME=parking_lot_db

Install dependencies:
go mod tidy

Running the Service
Compile and run the application:
go run main.go

Docker Support
Build and run using Docker
docker-compose up --build

API Endpoints
POST /parking-lots: Create a new parking lot.
POST /vehicles/park: Park a vehicle.
POST /vehicles/unpark/:id: Unpark a vehicle.
GET /parking-lots/:id/free-slots: Get free slots in a parking lot.
Contributing
Contributions are welcome! Please fork the repository and create a pull request.

API Endpoints
POST /parking-lots: Create a new parking lot.
POST /vehicles/park: Park a vehicle.
POST /vehicles/unpark/:id: Unpark a vehicle.
GET /parking-lots/:id/free-slots: Get free slots in a parking lot.
Contributing
Contributions are welcome! Please fork the repository and create a pull request.

This concise README covers all essential aspects of your Parking Lot Service project, including setup instructions, features, development guidelines, and licensing information. Adjust the placeholders (`yourusername`, `myusername`, `mypassword`, `DB_HOST`) according to your project's actual details before publishing.
