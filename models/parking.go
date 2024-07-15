package models

import (
    "gorm.io/gorm"
    "time"
)

type ParkingLot struct {
    gorm.Model
    Name            string
    MotorcycleSpots int
    CarSUVSpots     int
    BusTruckSpots   int
    Tariffs         []Tariff
}

type Tariff struct {
    gorm.Model
    ParkingLotID uint
    VehicleType  string
    RatePerHour  float64
    MaxRatePerDay float64
}

type Vehicle struct {
    gorm.Model
    LicensePlate string
    VehicleType  string
    ParkingLotID uint
    EntryTime    time.Time
    ExitTime     time.Time
    TotalCost    float64
}
