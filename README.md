# Tipeee Api Client

Api Client for Tipeee's API ([api.tipeee.com/api/doc/partner](https://api.tipeee.com/api/doc/partner)), written in Go. 

```go
import (
	tipeee "github.com/aduermael/tipeee-api-client"
)

func main() {
	client := tipeee.ClientWithToken(token)
}
```