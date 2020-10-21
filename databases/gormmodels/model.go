package gormmodels

// gorm.Model definition
type Model struct {
	ID uint `gorm:"primaryKey" json:"id"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}
