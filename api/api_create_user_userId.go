package api

import (
  
"User-Mgt/utils"
"github.com/gofiber/fiber/v2"

  
    "User-Mgt/dao"
    
  

  "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "io"
    "math/rand"
    "net/http"
    "os"
    "strings"
    "time"
  	"User-Mgt/dto"
  
  
  
)

// @Summary      CreateUser 
// @Description   This API performs the USERMGT_CREATE operation on User. It allows you to  User records.
// @Tags          User
// @Accept       json
// @Produce      json
// @Param        
// @Success      200  {array}   dto.User "Status OK"
// @Success      202  {array}   dto.User "Status Accepted"
// @Failure      404 "Not Found"
// @Router      /CreateUser [USERMGT_CREATE]

    func CreateUserApi(c *fiber.Ctx) error {





    
  
var Auth0Domain string = os.Getenv("AUTH0_DOMAIN")
var ClientId string = os.Getenv("AUTH0_CLIENTID")
inputObj := dto.User{}
if err := c.BodyParser(&inputObj); err != nil {
    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
}
password := generatePassword()
token, err := RetrieveAccessToken(Auth0Domain, ClientId)
if err != nil {
    return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
}
users, err := GetAuth0User(Auth0Domain ,inputObj.Email,*token)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}
	if len(users) !=0{
		if inputObj.Email == users[0]["email"].(string){
			return utils.SendErrorResponse(c, fiber.StatusBadRequest, errors.New("user already exists").Error())
		}
	}
	if len(users) == 0 {
		userId, err := createAuth0User(Auth0Domain, *token, inputObj, password)
		if err != nil {
			return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
		}
		err = passwordChangeRequest(Auth0Domain, ClientId, inputObj.Email)
		if err != nil {
			return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
		}
		inputObj.UserId = *userId
	}else {
		inputObj.UserId = users[0]["user_id"].(string)
	}
err = dao.DB_CreateUser(&inputObj)
	if err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, err.Error())
	}



        return utils.SendSuccessResponse(c)
        
    
}


    func passwordChangeRequest(domain, clientId, email string) error {
	url := "https://" + domain + "/dbconnections/change_password"

	payload := strings.NewReader("{\"client_id\": \"" + clientId + "\",\"email\": \"" + email + "\",\"connection\": \"Username-Password-Authentication\"}")
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		return err
	}

	req.Header.Add("content-type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()
	return nil
}
func createAuth0User(domain, token string, user dto.User,password string) (*string, error) {
	metadata := make(map[string]interface{})
	metadata["active"] = false
	metadata["role"] = user.RoleName
	metadata["workspaceid"] = "WS231"

	userDataA0 := dto.UserAuth0{
		Email:         user.Email,
		Password:      password,
		Connection:    "Username-Password-Authentication",
		Name:          user.FirstName + " " + user.LastName,
		EmailVerified: true,
		UserMetadata:  metadata,
	}

	payload, err := json.Marshal(userDataA0)
	if err != nil {
		return nil, err
	}

	url := "https://" + domain + "/api/v2/users"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("authorization", "Bearer "+token)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var responseData map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&responseData)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 201 {
		return nil, errors.New(fmt.Sprintf("%s", responseData["message"]))
	}
	userId := responseData["user_id"].(string)
	return &userId, nil
}
func RetrieveAccessToken(domain, clientId string) (*string, error) {

	url := "https://" + domain + "/oauth/token"
	audience := "https://" + domain + "/api/v2/"
	payload := strings.NewReader("{\"client_id\":\"" + clientId + "\",\"client_secret\":\"" + os.Getenv("AUTH0_CLIENT_SECRET") + "\",\"audience\":\"" + audience + "\",\"grant_type\":\"client_credentials\"}")

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
  if err !=nil{
    return nil, errors.New(err.Error())
  }
	var data map[string]interface{}
	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		return nil, err
	}

	accessToken, ok := data["access_token"].(string)
	if !ok {
		return nil, errors.New("access token not found or not a string")
	}

	return &accessToken, nil
}

func generatePassword() string {
	rand.Seed(time.Now().UnixNano())
	caps := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lows := "abcdefghijklmnopqrstuvwxyz"
	num := "0123456789"
	sym := "!@#$%^&*()-_=+[]{}|;:'\\\",.<>?/"
	length := 9
	randomString := make([]byte, length)
	for i := 0; i < length; i++ {
		if i == 3 {
			randomString[i] = caps[rand.Intn(len(caps))]
		} else if i == 5 {
			randomString[i] = sym[rand.Intn(len(sym))]
		} else if i == 1 || i == 2 || i == 8 {
			randomString[i] = lows[rand.Intn(len(lows))]
		} else {
			randomString[i] = num[rand.Intn(len(num))]
		}
	}
	return string(randomString)
}

func GetAuth0User(domain string, email string,token string) ([]map[string]interface{}, error) {
	url := "https://"+domain+"/api/v2/users-by-email?email="+email

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("authorization", "Bearer "+token)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK status: %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	var users []map[string]interface{}
	err = json.Unmarshal(body, &users)
	if err != nil {
		return nil, err

	}
	return users, nil

}

