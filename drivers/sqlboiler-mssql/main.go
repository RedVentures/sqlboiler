package main

import (
	"github.com/RedVentures/sqlboiler/v5/drivers"
	"github.com/RedVentures/sqlboiler/v5/drivers/sqlboiler-mssql/driver"
)

func main() {
	drivers.DriverMain(&driver.MSSQLDriver{})
}
