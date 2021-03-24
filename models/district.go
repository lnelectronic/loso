// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 24/3/2564 21:17
// ---------------------------------------------------------------------------
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//modles District Json
type District struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id"`
	DistrictTh string             `bson:"_id" json:"district_th" validate:"required"`
	DistrictEn string             `bson:"_id" json:"district_en" validate:"required"`
	ProvinceID primitive.ObjectID `bson:"_id" json:"province_id"`
}

// New instance District
func (u *District) New() *District {
	return &District{
		ID:         primitive.NewObjectID(),
		DistrictTh: u.DistrictTh,
		DistrictEn: u.DistrictEn,
		ProvinceID: primitive.NewObjectID(),
	}
}
