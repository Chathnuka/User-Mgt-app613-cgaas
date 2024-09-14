package api

import (
  
"User-Mgt/utils"
"github.com/gofiber/fiber/v2"

  
    "User-Mgt/dao"
    
  

  
  
  "errors"
	"net/http"
	"os"
  
  
)

// @Summary      DeleteUser 
// @Description   This API performs the USERMGT_DELETE operation on User. It allows you to  User records.
// @Tags          User
// @Accept       json
// @Produce      json
// @Param        
// @Success      200  {array}   dto.User "Status OK"
// @Success      202  {array}   dto.User "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /DeleteUser [USERMGT_DELETE]

    func DeleteUserApi(c *fiber.Ctx) error {





    
  
userId := c.Query("userId")
err := deleteAuth0User(userId)
if err != nil {
    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
}
err = dao.DB_DeleteUser(userId)
if err != nil {
    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
}



        return utils.SendSuccessResponse(c)
        
    
}


    func deleteAuth0User(userId string) error {
	var Auth0Domain string = os.Getenv("AUTH0_DOMAIN")
	var ClientId string = os.Getenv("AUTH0_CLIENTID")
	url := "https://" + Auth0Domain + "/api/v2/users/" + userId
	method := "DELETE"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return err
	}
	token, err := RetrieveAccessToken(Auth0Domain, ClientId)
	if err != nil {
		return err
	}
	req.Header.Add("authorization", "Bearer "+*token)

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 204 {
		return errors.New("unable to delete user")
	}
	return nil
}

