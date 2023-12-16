package queries

const (
	WarehouseGetAllQuery = "SELECT * FROM warehouses"
	WarehouseGetQuery    = "SELECT * FROM warehouses WHERE id=?;"
	WarehouseExistsQuery = "SELECT warehouse_code FROM warehouses WHERE warehouse_code=?;"
	WarehouseSaveQuery   = "INSERT INTO warehouses (address, telephone, warehouse_code, minimum_capacity, minimum_temperature) VALUES (?, ?, ?, ?, ?)"
	WarehouseUpdateQuery = "UPDATE warehouses SET address=?, telephone=?, warehouse_code=?, minimum_capacity=?, minimum_temperature=? WHERE id=?"
	WarehouseDeleteQuery = "DELETE FROM warehouses WHERE id=?"
)
