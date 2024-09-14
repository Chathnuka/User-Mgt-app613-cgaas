package integrations

import (
"os"
)

var AUTH0_DOMAIN string
var AUTH0_CLIENTID string
var AUTH0_CLIENT_SECRET string
func SetEnvironmentVariables() {
AUTH0_DOMAIN=os.Getenv("AUTH0_DOMAIN")
    AUTH0_CLIENTID=os.Getenv("AUTH0_CLIENTID")
    AUTH0_CLIENT_SECRET=os.Getenv("AUTH0_CLIENT_SECRET")
    }
