package api

import (
  
"User-Mgt/utils"
"github.com/gofiber/fiber/v2"

  
    "User-Mgt/dao"
    
  

  
  
  
)

// @Summary      UsermgtGetconfigRole 
// @Description   This API performs the USERMGT_GETCONFIG operation on Role. It allows you to  Role records.
// @Tags          Role
// @Accept       json
// @Produce      json
// @Param        
// @Success      200  {array}   dto.Role "Status OK"
// @Success      202  {array}   dto.Role "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UsermgtGetconfigRole [USERMGT_GETCONFIG]

    func UsermgtGetconfigRoleApi(c *fiber.Ctx) error {





    userId := c.Query("userId")
        
    
  
    user, err := dao.DB_FindUserbyUserId(userId)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	config, err := dao.DB_FindRolebyRoleId(user.RoleId)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest,err.Error())
	}



     return c.Status(fiber.StatusCreated).JSON(&config)
     
}

