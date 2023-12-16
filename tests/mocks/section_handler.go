package mocks

func ptrInt(val int) *int {
	return &val
}

type requestTest struct {
	SectionNumber      *int `json:"section_number,omitempty"`
	CurrentTemperature *int `json:"current_temperature,omitempty"`
	MinimumTemperature *int `json:"minimum_temperature,omitempty"`
	CurrentCapacity    *int `json:"current_capacity,omitempty"`
	MinimumCapacity    *int `json:"minimum_capacity,omitempty"`
	MaximumCapacity    *int `json:"maximum_capacity,omitempty"`
	WarehouseID        *int `json:"warehouse_id,omitempty"`
	ProductTypeID      *int `json:"product_type_id,omitempty"`
}

var MockNuevaSectionRequest requestTest = requestTest{
	SectionNumber:      ptrInt(10),
	CurrentTemperature: ptrInt(40),
	MinimumTemperature: ptrInt(30),
	CurrentCapacity:    ptrInt(20),
	MinimumCapacity:    ptrInt(0),
	MaximumCapacity:    ptrInt(100),
	WarehouseID:        ptrInt(5),
	ProductTypeID:      ptrInt(5),
}

var MockNuevaSectionRequestConflict requestTest = requestTest{
	SectionNumber:      nil,
	CurrentTemperature: nil,
	MinimumTemperature: nil,
	CurrentCapacity:    nil,
	MinimumCapacity:    nil,
	MaximumCapacity:    nil,
	WarehouseID:        nil,
	ProductTypeID:      nil,
}

var MockActualizarSectionRequest requestTest = requestTest{
	SectionNumber:      ptrInt(2),
	CurrentTemperature: ptrInt(15),
	MinimumTemperature: ptrInt(10),
	CurrentCapacity:    ptrInt(5),
	MinimumCapacity:    ptrInt(1),
	MaximumCapacity:    ptrInt(100),
	WarehouseID:        ptrInt(2),
	ProductTypeID:      ptrInt(2),
}
