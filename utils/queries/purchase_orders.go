package queries

const (
	PurchaseOrderInsertIntoPO      = "INSERT INTO purchase_orders(order_number,order_date,tracking_code,buyer_id,order_status_id) VALUES (?,?,?,?,?)"
	PurchaseOrderInsertIntoOD      = "INSERT INTO order_details(product_record_id) VALUES (?)"
	PurchaseOrderSelectOrderNumber = "SELECT order_number FROM purchase_orders WHERE order_number=?"
)
