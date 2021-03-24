// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 24/3/2564 21:17
// ---------------------------------------------------------------------------
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//models  Province
type Province struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	ProvinceTh string             `bson:"province_th" json:"province_th" validate:"required"`
	ProvinceEn string             `bson:"province_en" json:"province_en" validate:"required"`
	CountryID  primitive.ObjectID `bson:"country_id" json:"country_id"`
}

// New instance Province
func (u *Province) New() *Province {
	return &Province{
		ID:         primitive.NewObjectID(),
		ProvinceTh: u.ProvinceTh,
		ProvinceEn: u.ProvinceEn,
		CountryID:  primitive.NewObjectID(),
	}
}
