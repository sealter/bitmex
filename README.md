# BITMEX-API

### For exmple :
```
package main

import (
	"log"
	"github.com/sunrisedo/bitmex/bitmex"
)

func main() {
    cfg := bitmex.NewConfigurationWithKey("key", "secret")
    userApi := bitmex.NewUserApiWithConfig(cfg)
    wallet, response, err := userApi.UserGetWallet("")
    if err != nil {
    	log.Println("error: ", err)
        return
    }
    log.Println(response.Status)
    log.Println("wallet: ", *wallet)
}

```