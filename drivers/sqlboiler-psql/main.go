package main

import (
	"github.com/RedVentures/sqlboiler/v5/drivers"
	"github.com/RedVentures/sqlboiler/v5/drivers/sqlboiler-psql/driver"
)

func main() {
	drivers.DriverMain(&driver.PostgresDriver{})
}
