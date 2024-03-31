# zammad-go

Is a Go client library for interacting with the Zammad REST API.

## Installation

To install `zammad-go`, use `go get`:

```bash
go get github.com/chazari-x/zammad-go
```

## Example

```go
package main

import (
	"fmt"

	"github.com/chazari-x/zammad-go"
)

func main() {
	// Create a client instance
	client, err := zammad.NewClient(&zammad.Client{
		Username: "",
		Password: "",
		Token:    "",
		OAuth:    "",
		Url:      "http://my-zammad-instance.com",
	})
	if err != nil {
		panic(err)
	}

	// Get Users
	users, err := client.UserList()
	if err != nil {
		panic(err)
	}

	fmt.Println("Users:")
	for _, user := range users {
		fmt.Println(user)
	}

	// Get User
	user, err := client.UserShow(1)
	if err != nil {
		panic(err)
	}

	fmt.Println("User:")
	fmt.Println(user)
}
```

## Documentation

For more information on how to use `zammad-go`, please refer to the documentation provided in the [Zammad API documentation](https://docs.zammad.org/en/latest/api/intro.html).

## Contributing

Contributions are welcome! Please feel free to submit issues or pull requests.

## License

This library is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
