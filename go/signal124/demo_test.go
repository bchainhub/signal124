package signal124_test

import (
	"fmt"

	"github.com/cryptohub-digital/signal124/go/signal124"
)

func Demo() {
	fmt.Println(signal124.EncodeToString([]byte("signal01")))
	fmt.Println(string(signal124.DecodeString("CΩÄ+'6¿£/ù")))
}
