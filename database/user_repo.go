// ---------------------------------------------------------------------------
// LN-ELECTRINIC PROJECT LN-16C10R
// wwww.ln-electronic.com  ProjectManager : @JJOY, @Kimera
// FileData: 28/3/2564 4:46
// ---------------------------------------------------------------------------
package database

//
//// FindByEmail retrieves user row by email address
//func (r *LnDatabase) FindByEmail(ctx context.Context, email string) (*models.User, error) {
//	user := &models.User{}
//
//	query := "SELECT * FROM users WHERE email=$1"
//
//	if err := r.DB.GetContext(ctx, user, query, email); err != nil {
//		log.Printf("Unable to get user with email address: %v. Err: %v\n", email, err)
//		return user, apperrors.NewNotFound("email", email)
//	}
//
//	return user, nil
//}
//

// Signin reaches our to a UserRepository check if the user exists
// and then compares the supplied password with the provided password
// if a valid email/password combo is provided, u will hold all
// available user fields
//func (s *LnDatabase) Signin(ctx context.Context, u *model.User) error {
//	//uFetched, err := s.UserRepository.FindByEmail(ctx, u.Email)
//
//
//	// Will return NotAuthorized to client to omit details of why
//	if err != nil {
//		return apperrors.NewAuthorization("Invalid email and password combination")
//	}
//
//	// verify password - we previously created this method
//	match, err := comparePasswords(uFetched.Password, u.Password)
//
//	if err != nil {
//		return apperrors.NewInternal()
//	}
//
//	if !match {
//		return apperrors.NewAuthorization("Invalid email and password combination")
//	}
//
//	*u = *uFetched
//	return nil
//}
