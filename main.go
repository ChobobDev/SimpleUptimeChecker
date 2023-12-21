package main

import(
  "github.com/Chobobdev/SimpleUptimeChecker/monitoring"
  "time"
)

func main(){
  monitoring.Monitor("https://wrtn.io",1 * time.Minute)
}
