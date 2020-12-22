package wait

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url, string) error {
	const timeout =1 * time.Minute
	deadline := time.Now().Add(timeout)
	for tries :=0; time.Now().Before(deadline); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil //成功
		}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << unit(tries)) //指数退避策略
	}

	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}
