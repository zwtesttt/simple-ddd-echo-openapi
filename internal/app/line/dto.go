package line

// CreateLanDTO is the data transfer object for creating a LAN.
type CreateLineDTO struct {
	Name string `json:"name"`
	CIDR string `json:"cidr"`
}
