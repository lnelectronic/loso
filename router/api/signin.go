// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 28/3/2564 1:56
// ---------------------------------------------------------------------------
package api

// signinReq is not exported
type signinReq struct {
	Username string `bson:"username" json:"username" binding:"required"`
	Passwd   string `bson:"passwd" json:"passwd" binding:"required,gte=6,lte=30"`
}
