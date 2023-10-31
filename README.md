# MOEX ISS Api client

**Currently in progress....**

## How to use at current stage
**On now working only request builder**
``` golang
import "moexapplication/internal/requests" 

request := requests.New()
request.NewEngine("SOME_ENGINE")
request.NewMarket("SOME_MARKET")
request.NewSecurity("IMOEX")
response, err :=	return request.ExexuteWithType(requests.JSON)
```
See moexapp.go for more

