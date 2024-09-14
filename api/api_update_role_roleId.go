package api

import (
  
"User-Mgt/utils"
"github.com/gofiber/fiber/v2"

  "User-Mgt/dto"
    "github.com/go-playground/validator/v10"
    
    "User-Mgt/dao"
    
  

  
  
  
)

// @Summary      UpdateRole 
// @Description   This API performs the PUT operation on Role. It allows you to update Role records.
// @Tags          Role
// @Accept       json
// @Produce      json
// @Param        data body dto.Role false "string collection" 
// @Success      200  {array}   dto.Role "Status OK"
// @Success      202  {array}   dto.Role "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateRole [PUT]

    func UpdateRoleApi(c *fiber.Ctx) error {





    
  
    inputObj := dto.Role{}


    if err := c.BodyParser(&inputObj); err != nil {
    		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }
    

    validate := validator.New()
    if validationErr := validate.Struct(&inputObj); validationErr != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
    }
err := dao.DB_UpdateRole(&inputObj)
    if err != nil {
        return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
    }



        return utils.SendSuccessResponse(c)
        
    
}

