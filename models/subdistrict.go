// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 24/3/2564 21:19
// ---------------------------------------------------------------------------
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//model
type Subdistrict struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	SubDistrictTh string             `bson:"sub_district_th" json:"sub_district_th" validate:"required"`
	SubDistrictEn string             `bson:"sub_district_en" json:"sub_district_en" validate:"required"`
	PostCode      string             `bson:"post_code" json:"post_code" validate:"required"`
	DistrictID    string             `bson:"district_id" json:"district_id" validate:"required"`
	ProvinceID    string             `bson:"province_id" json:"province_id" validate:"required"`
}

// New instance Subdistrict
func (u *Subdistrict) New() *Subdistrict {
	return &Subdistrict{
		ID:            primitive.NewObjectID(),
		SubDistrictTh: u.SubDistrictTh,
		SubDistrictEn: u.SubDistrictEn,
		PostCode:      u.PostCode,
		DistrictID:    u.DistrictID,
		ProvinceID:    u.ProvinceID,
	}
}
