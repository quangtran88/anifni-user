package services

//func TestUserService_Register(t *testing.T) {
//	// Arrange
//	mAuthService := new(serviceAdapter.MockedAuthenticationService)
//	mTokenService := new(serviceAdapter.MockedTokenService)
//	mUserRepo := new(repositories.MockedUserRepository)
//	mAuthService.On("CheckEmailOTP", "12345", "randy@gmail.com").Return(true, nil)
//	mTokenService.On("GetToken", domain.UserTokenDomain, domain.UserTokenLength).Return("USER_TOKEN", nil)
//	mUserRepo.On("Create", mock.Anything).Return(domain.ID("INSERTED_ID"), nil)
//	userService := NewUserService(mUserRepo, mAuthService, mTokenService)
//	input := domain.RegisterUserInput{
//		OTPCode:   "12345",
//		Email:     "randy@gmail.com",
//		Password:  "secret",
//		LastName:  "Tran",
//		FirstName: "Quang",
//	}
//	// Act
//	user, _ := userService.Register(input)
//	// Assert
//	assert.Equal(t, string(user.Id), "INSERTED_ID")
//	mAuthService.AssertExpectations(t)
//	mTokenService.AssertExpectations(t)
//	userToCreate := mUserRepo.Calls[0].Arguments[0].(domain.User)
//	assert.Equal(t, string(userToCreate.Pid), "USER_TOKEN")
//	assert.Equal(t, userToCreate.Name, "Quang Tran")
//	assert.True(t, utils.CheckPasswordHash("secret", user.Password))
//}
