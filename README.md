# go-scanln

Continously fetches user input from the console via `fmt.Scanln` and provides it on a channel.

### Example
```go
import (
	"fmt"

	"github.com/franklange/go-scanln"
)

func main() {
	input := scanln.NewScanln()
	defer input.Stop()

	select {
	case s := <-input.C:
		fmt.Println(s)
	}
}
```