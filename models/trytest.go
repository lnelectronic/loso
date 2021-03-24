// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 24/3/2564 22:15
// ---------------------------------------------------------------------------
package models

// Auth xxx
type Auth struct {
	Adminname string `form:"adminname" json:"adminname" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
}

// New is an instance
func (a *Auth) New() *Auth {
	return &Auth{
		Adminname: a.Adminname,
		Password:  a.Password,
	}
}
