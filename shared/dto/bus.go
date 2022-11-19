package dto

const (
	// Crowded Status
	EMPTY    BusStatus = "EMPTY"
	MODERATE BusStatus = "MODERATE"
	FULL     BusStatus = "FULL"

	// Route
	RED  Route = "RED"
	BLUE Route = "BLUE"
)

type (
	BusStatus string

	Route string

	Bus struct {
		ID       uint      `gorm:"primaryKey;autoIncrement"`
		Number   int       `gorm:"column:number;unique"`
		Plate    string    `gorm:"column:plate;unique"`
		Status   BusStatus `gorm:"column:status;default:EMPTY"`
		Route    Route     `gorm:"column:route"`
		IsActive bool      `gorm:"column:is_active;default:false"`
		Username string    `gorm:"column:username;unique"`
		Password string    `gorm:"password"`
	}

	// CreateBusDto CreateBusDto
	CreateBusDto struct {
		Number   int    `json:"number" validate:"required"`
		Plate    string `json:"plate" validate:"required"`
		Route    Route  `json:"route" validate:"required,oneof=RED BLUE"`
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// CreateBusResponse CreateBusResponse
	CreateBusResponse struct {
		ID       uint      `json:"id"`
		Number   int       `json:"number"`
		Plate    string    `json:"plate"`
		Status   BusStatus `json:"status"`
		Route    Route     `json:"route"`
		IsActive bool      `json:"isActive"`
		Username string    `json:"username"`
	}

	// DriverLoginDto DriverLoginDto
	DriverLoginDto struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	// DriverLoginResponse DriverLoginResponse
	DriverLoginResponse struct {
		ID       uint      `json:"id"`
		Number   int       `json:"number"`
		Plate    string    `json:"plate"`
		Status   BusStatus `json:"status"`
		Route    Route     `json:"route"`
		IsActive bool      `json:"isActive"`
		Token    string    `json:"token"`
	}
)

func (b *Bus) ToCreateBusResponse() CreateBusResponse {
	return CreateBusResponse{
		ID:       b.ID,
		Number:   b.Number,
		Plate:    b.Plate,
		Status:   b.Status,
		Route:    b.Route,
		IsActive: b.IsActive,
		Username: b.Username,
	}
}

func (b *Bus) ToDriverLoginResponse(token string) DriverLoginResponse {
	return DriverLoginResponse{
		ID:       b.ID,
		Number:   b.Number,
		Plate:    b.Plate,
		Status:   b.Status,
		Route:    b.Route,
		IsActive: b.IsActive,
		Token:    token,
	}
}
