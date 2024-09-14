package api

import (
  
"User-Mgt/utils"
"github.com/gofiber/fiber/v2"

  
    "User-Mgt/dao"
    
  

  
  
  
  	"User-Mgt/dto"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
  
)

// @Summary      UpdateUser 
// @Description   This API performs the USERMGT_UPDATE operation on User. It allows you to  User records.
// @Tags          User
// @Accept       json
// @Produce      json
// @Param        
// @Success      200  {array}   dto.User "Status OK"
// @Success      202  {array}   dto.User "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /UpdateUser [USERMGT_UPDATE]

    func UpdateUserApi(c *fiber.Ctx) error {





    
  
inputObj := dto.User{}

if err := c.BodyParser(&inputObj); err != nil {
    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
}
validate := validator.New()
if validationErr := validate.Struct(&inputObj); validationErr != nil {
    return utils.SendErrorResponse(c, fiber.StatusBadRequest, validationErr.Error())
}
err := updateAuth0User(inputObj)
if err != nil {
    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
}
err = dao.DB_UpdateUser(&inputObj)
if err != nil {
    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
}



        return utils.SendSuccessResponse(c)
        
    
}


    func updateAuth0User( user dto.User) error {
	var Auth0Domain string = os.Getenv("AUTH0_DOMAIN")
	var ClientId string = os.Getenv("AUTH0_CLIENTID")
	url := "https://" + Auth0Domain + "/api/v2/users/" + user.UserId
	method := "PATCH"
	metadata := make(map[string]interface{})
	metadata["active"] = false
	metadata["role"] = user.RoleName

	userDataA0 := dto.UserAuth0{
		Email:         user.Email,
		Connection:    "Username-Password-Authentication",
		Name:          user.FirstName + " "+ user.LastName,
		UserMetadata:  metadata,
	}
	payload, err := json.Marshal(userDataA0)
	if err != nil {
		return err
	}
	//payload := strings.NewReader(`{"user_metadata":{"phone_number":"` + user.Phone + `","role":"` + user.Role + `"},"picture":"`+user.ProfileImage+`","name":"` + user.Name + `","client_id":"` + clientId + `","email":"` + user.Email + `"}`)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	token, err := RetrieveAccessToken(Auth0Domain, ClientId)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", "Bearer "+*token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	var responseData map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New(fmt.Sprintf("%s", responseData["message"]))
	}

	defer res.Body.Close()
	return nil
}
