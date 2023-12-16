package queries

const (
	CarrySaveQuery           = "INSERT INTO carries (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)"
	CarryCIDExistsQuery      = "SELECT cid FROM carries WHERE cid=?"
	CarryLocalityExistsQuery = "SELECT id FROM localities WHERE id=?"
)
