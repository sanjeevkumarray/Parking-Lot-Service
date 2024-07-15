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



## Features

- Supports Motorcycles/scooters, Cars/SUVs, Buses/Trucks with varying slot sizes.
- Real-time slot tracking.
- Generates entry tickets and receipts.
- Different tariff models per parking lot.

## Setup

1. **Clone the repository:**
   ```sh
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
 -POST /parking-lots: Create a new parking lot.
 -POST /vehicles/park: Park a vehicle.
-POST /vehicles/unpark/:id: Unpark a vehicle.
-GET /parking-lots/:id/free-slots: Get free slots in a parking lot.


Contributing
Contributions are welcome! Please fork the repository and create a pull request.


This concise README covers all essential aspects of your Parking Lot Service project, including setup instructions, features, development guidelines, and licensing information. Adjust the placeholders (`yourusername`, `myusername`, `mypassword`, `DB_HOST`) according to your project's actual details before publishing.
