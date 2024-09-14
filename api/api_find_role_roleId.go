package api

import (
  
"User-Mgt/utils"
"github.com/gofiber/fiber/v2"

  
    "User-Mgt/dao"
    
  

  
  
  
)

// @Summary      FindRole 
// @Description   This API performs the GET operation on Role. It allows you to retrieve Role records.
// @Tags          Role
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Role "Status OK"
// @Success      202  {array}   dto.Role "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /FindRole [GET]

    func FindRoleApi(c *fiber.Ctx) error {





    roleId := c.Query("roleId")
        
    
  returnValue, err := dao.DB_FindRolebyRoleId(roleId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }


return c.Status(fiber.StatusAccepted).JSON(&returnValue)
}

