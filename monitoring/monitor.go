package monitoring

import (
	"fmt"
	"net/http"
	"time"
)

func Monitor(siteURL string, interval time.Duration) {
	for {
    // Send a HTTP Get request
		response, err := http.Get(siteURL)
		if err != nil {
			fmt.Printf(" %s is unavailabe : %s \n", siteURL, err.Error())
		} else {
			statusCode := response.StatusCode
			if statusCode != http.StatusOK {
				fmt.Printf("%s has error Status : %d \n", siteURL, statusCode)
			} else {
        fmt.Printf("%s is Up : %d \n", siteURL, statusCode)
			}
			response.Body.Close()
		}

		time.Sleep(interval)
	}
}


