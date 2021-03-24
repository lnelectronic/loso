// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 24/3/2564 21:15
// ---------------------------------------------------------------------------
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// models Country
type Country struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	CountryTh string             `bson:"country_th" json:"country_th" validate:"required"`
	CountryEn string             `bson:"country_en" json:"country_en" validate:"required"`
}

// New instance Country
func (c *Country) New() *Country {
	return &Country{
		ID:        primitive.NewObjectID(),
		CountryTh: c.CountryTh,
		CountryEn: c.CountryEn,
	}
}
