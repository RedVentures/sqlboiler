package main

import (
	"github.com/RedVentures/sqlboiler/v4/drivers"
	"github.com/RedVentures/sqlboiler/v4/drivers/sqlboiler-mysql/driver"
)

func main() {
	drivers.DriverMain(&driver.MySQLDriver{})
}
