package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var currentServerCount = 0

const (
  SERVER1 = "https://www.google.com"
  SERVER2 = "https://www.facebook.com"
  SERVER3 = "https://www.yahoo.com"
  PORT = "3005"
)

func getProxyAddress() string{

  var servers = []string{SERVER1, SERVER2, SERVER3}

  server := servers[currentServerCount]
  currentServerCount ++

  if currentServerCount >= len(servers){
    currentServerCount = 0
  }

  return server

}

func startReverseProxy(proxyUrl string, req *http.Request, res http.ResponseWriter) {

  url,_ := url.Parse(proxyUrl)

  proxy := httputil.NewSingleHostReverseProxy(url)

  proxy.ServeHTTP(res, req)

}

func logRequestDetails(proxyUrl string) {
  log.Printf("proxy_url: %s\n", proxyUrl)
}

func handleIncomingRequestAndRedirect(res http.ResponseWriter, req *http.Request) {

  url := getProxyAddress()

  logRequestDetails(url)

  startReverseProxy(url, req, res)

}

func main() {

  http.HandleFunc("/", handleIncomingRequestAndRedirect)

  log.Fatal(http.ListenAndServe(":"+PORT, nil))

}