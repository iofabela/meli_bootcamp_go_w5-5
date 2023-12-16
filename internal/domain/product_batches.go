package domain

type ProductBatches struct {
	Id                 int
	BatchNumber        int
	CurrentQuantity    int
	CurrentTemperature int
	DueDate            string
	InitialQuantity    int
	ManufacturingDate  string
	ManufacturingHour  string
	MinumumTemperature int
	ProductId          int
	SectionId          int
}
