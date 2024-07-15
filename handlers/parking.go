package handlers

import (
    "net/http"
    "time"
    "github.com/labstack/echo/v4"
    "parking-lot-service/db"
    "parking-lot-service/models"
)

func CreateParkingLot(c echo.Context) error {
    parkingLot := new(models.ParkingLot)
    if err := c.Bind(parkingLot); err != nil {
        return err
    }
    db.DB.Create(parkingLot)
    return c.JSON(http.StatusCreated, parkingLot)
}

func ParkVehicle(c echo.Context) error {
    vehicle := new(models.Vehicle)
    if err := c.Bind(vehicle); err != nil {
        return err
    }
    vehicle.EntryTime = time.Now()
    db.DB.Create(vehicle)
    return c.JSON(http.StatusCreated, vehicle)
}

func UnparkVehicle(c echo.Context) error {
    id := c.Param("id")
    var vehicle models.Vehicle
    if err := db.DB.First(&vehicle, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Vehicle not found")
    }
    vehicle.ExitTime = time.Now()
    hours := vehicle.ExitTime.Sub(vehicle.EntryTime).Hours()
    vehicle.TotalCost = calculateCost(vehicle.VehicleType, vehicle.ParkingLotID, hours)
    db.DB.Save(&vehicle)
    return c.JSON(http.StatusOK, vehicle)
}

func calculateCost(vehicleType string, parkingLotID uint, hours float64) float64 {
    var tariff models.Tariff
    db.DB.Where("parking_lot_id = ? AND vehicle_type = ?", parkingLotID, vehicleType).First(&tariff)
    if hours > 24 {
        return float64(int(hours/24)) * tariff.MaxRatePerDay
    }
    return float64(int(hours+1)) * tariff.RatePerHour
}

func GetFreeSlots(c echo.Context) error {
    id := c.Param("id")
    var parkingLot models.ParkingLot
    if err := db.DB.First(&parkingLot, id).Error; err != nil {
        return c.JSON(http.StatusNotFound, "Parking lot not found")
    }

    var vehicles []models.Vehicle
    db.DB.Where("parking_lot_id = ? AND exit_time IS NULL", id).Find(&vehicles)

    motorcycleSpots := parkingLot.MotorcycleSpots
    carSUVSpots := parkingLot.CarSUVSpots
    busTruckSpots := parkingLot.BusTruckSpots

    for _, vehicle := range vehicles {
        switch vehicle.VehicleType {
        case "Motorcycle", "Scooter":
            motorcycleSpots--
        case "Car", "SUV":
            carSUVSpots--
        case "Bus", "Truck":
            busTruckSpots--
        }
    }

    return c.JSON(http.StatusOK, map[string]int{
        "Motorcycles/Scooters": motorcycleSpots,
        "Cars/SUVs":            carSUVSpots,
        "Buses/Trucks":         busTruckSpots,
    })
}
