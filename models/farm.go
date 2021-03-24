// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 24/3/2564 21:16
// ---------------------------------------------------------------------------
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

//models Farm Json
type Farm struct {
	ID         primitive.ObjectID `json:"_id"`
	FarmnameEn string             `bson:"farmname_en" json:"farmname_en" validate:"required"`
	FarmnameTh string             `bson:"farmname_th" json:"farmname_th" validate:"required"`
	AddressID  primitive.ObjectID `bson:"addressid" json:"addressid"`
}

//New Instance Farm
func (f *Farm) New() *Farm {
	return &Farm{
		ID:         primitive.NewObjectID(),
		FarmnameEn: f.FarmnameEn,
		FarmnameTh: f.FarmnameTh,
		AddressID:  primitive.NewObjectID(),
	}
}
