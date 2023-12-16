package queries

const (
	LocalityGetSellerReportQuery     = "SELECT l.id, l.locality_name, count(c.id) FROM localities l LEFT JOIN sellers c ON c.locality_id = l.id WHERE l.id = ? GROUP BY l.id"
	LocalityGetAllSellerReportsQuery = "SELECT l.id, l.locality_name, count(c.id) FROM localities l LEFT JOIN sellers c ON c.locality_id = l.id GROUP BY l.id"
	LocalityGetAll                   = "SELECT id, locality_name FROM localities;"
	InsertLocality                   = "INSERT INTO localities (id, locality_name) VALUES (?, ?);"
	InsertProvince                   = "INSERT INTO provinces (province_name) VALUES (?);"
	InsertCountry                    = "INSERT INTO countries (country_name) VALUES (?);"
	SelectIdLocality                 = "SELECT id FROM localities WHERE id=?;"
	InsertSeller                     = "INSERT INTO sellers (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?);"
	SelectCidSeller                  = "SELECT cid FROM sellers WHERE cid=?;"
	LocalityGetCarryReportQuery     = "SELECT l.id, l.locality_name, count(c.id) FROM localities l LEFT JOIN carries c ON c.locality_id = l.id WHERE l.id = ? GROUP BY l.id"
	LocalityGetAllCarryReportsQuery = "SELECT l.id, l.locality_name, count(c.id) FROM localities l LEFT JOIN carries c ON c.locality_id = l.id GROUP BY l.id"
)
