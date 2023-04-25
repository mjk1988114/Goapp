package main

import (
      "fmt"
      "os"
      "log"
      "net"
      "net/http"
)
func gohandler(w http.ResponseWriter , r * http.Request){
  name, err := os.Hostanme()
  if err ! =nil {
    fmt.Printf("error: %v\n", err)
    return
  }
  fmt.Fprintln(w, "Hostname: ", name)
  
  addr, err := net.LookupHost(name)
   if err ! =nil {  
    fmt.Printf("error: %v\n", err)
    return
  }
  fmt.Fprintln(w, "IP: ", addr)
}
func main() {
  fmt.Fprintln(os.Stdout, "GO!!! Go Application ......")
  http.HandleFunc("/",gohandler)
  log.Fatal(http.ListenAndServe(":9090", nil))
}
  
