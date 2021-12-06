# Gobizfly

Gobizfly is a Go client library for accessing BizFly API

You can view the client API docs here: https://pkg.go.dev/github.com/bizflycloud/gobizfly

# Install
```bash
go get github.com/bizflycloud/gobizfly@vX
```

# Usage
```go
import "github.com/bizflycloud/gobizfly"
```
Create a new BizFly Cloud client, then use the exposed services to access different part of the BizFly API 

# Authentication
Currently, token is the only method of authenticating with the API. You can generate token via username/password or application credential

```go
package main
import (
	"github.com/bizflycloud/gobizfly"
	"log"
)

func main() {
	tok, err := client.Token.Create(ctx, &gobizfly.TokenCreateRequest{Username: username, Password: password, AuthMethod: "password"})
	if err != nil {
		log.Fatal(err)
	}
	client.SetKeystoneToken(tok)
}
```

# Example

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/bizflycloud/gobizfly"
)

const (
	host     = "https://manage.bizflycloud.vn"
	username = "foo@bar"
	password = "foobar"
)

func main() {
	client, err := gobizfly.NewClient(
		gobizfly.WithAPIUrl(host),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*10)
	defer cancelFunc()
	tok, err := client.Token.Create(ctx, &gobizfly.TokenCreateRequest{Username: username, Password: password})
	if err != nil {
		log.Fatal(err)
	}
	client.SetKeystoneToken(tok)

	lbs, err := client.LoadBalancer.List(ctx, &gobizfly.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v\n", lbs)
}
```

# Documentation
For details on all the functionality in this library, checkout [Gobizfly documentation](https://pkg.go.dev/github.com/bizflycloud/gobizfly)