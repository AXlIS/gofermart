package gophermart

type (
	Order struct {
		Number     int      `json:"order,omitempty" db:"id"`
		Status     string   `json:"status,omitempty" db:"status"`
		Accrual    *float64 `json:"accrual,omitempty" db:"accrual, omitempty"`
		UploadedAt string   `json:"uploaded_at,omitempty" db:"uploaded_at"`
	}

	Balance struct {
		Current   float32 `json:"current,omitempty"`
		Withdrawn float32 `json:"withdrawn,omitempty"`
	}

	Withdrawal struct {
		Order       string  `json:"order"`
		Sum         float32 `json:"sum"`
		ProcessedAt string  `json:"processed_at" db:"processed_at"`
	}
)
