package main

import (
    "parking-lot-service/db"
    "parking-lot-service/handlers"
    "parking-lot-service/models"
    "github.com/labstack/echo/v4"
)

func main() {
    db.Init()
    db.DB.AutoMigrate(&models.ParkingLot{}, &models.Tariff{}, &models.Vehicle{})

    e := echo.New()

    e.POST("/parking-lots", handlers.CreateParkingLot)
    e.POST("/vehicles/park", handlers.ParkVehicle)
    e.POST("/vehicles/unpark/:id", handlers.UnparkVehicle)
    e.GET("/parking-lots/:id/free-slots", handlers.GetFreeSlots)

    e.Logger.Fatal(e.Start(":8080"))
}
