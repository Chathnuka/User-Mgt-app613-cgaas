package api

import (
  
"User-Mgt/utils"
"github.com/gofiber/fiber/v2"

  
    "User-Mgt/dao"
    
  

  
  
  
)

// @Summary      DeleteRole 
// @Description   This API performs the DELETE operation on Role. It allows you to delete Role records.
// @Tags          Role
// @Accept       json
// @Produce      json
// @Param        objectId query []string false "string collection"  collectionFormat(multi)
// @Success      200  {array}   dto.Role "Status OK"
// @Success      202  {array}   dto.Role "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteRole [DELETE]

    func DeleteRoleApi(c *fiber.Ctx) error {





    roleId := c.Query("roleId")
        
    
  err := dao.DB_DeleteRole(roleId)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

