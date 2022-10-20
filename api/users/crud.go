package users

type UserApi struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
}

// func CreateUser(c *gin.Context) {
// 	var user UserApi

// 	if err := c.ShouldBindJSON(&user); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	result, err := services.CreateUser(&services.CreateUserDomain{FirstName: user.FirstName, LastName: user.LastName, Email: user.Email})

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	}

// 	c.JSON(http.StatusOK, gin.H{"result": result})
// }
