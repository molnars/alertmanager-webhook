package main

import (
     "net/http"
     "fmt"
     "log"
  }

func webhook(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
  data := template.Data{}
  if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
    asJson(w, http.StatusBadRequest, err.Error())
    return
  }
  fmt.Printf("Alerts: GroupLabels=%v, CommonLabels=%v", data.GroupLabels, data.CommonLabels)
  for _, alert := range data.Alerts {
    fmt.Printf("Alert: status=%s,Labels=%v,Annotations=%v", alert.Status, alert.Labels, alert.Annotations)
  
  severity := alert.Labels["severity"]
  switch strings.ToUpper(severity) {
    case "CRITICAL":
    case "WARNING":
    default:
      log.Printf("no action on severity: %s", severity)
    }
  }
  asJson(w, http.StatusOK, "success")
}
func healthz(w http.ResponseWriter, r *http.Request) {
 fmt.Fprint(w, "Ok!")
}
func main() {
 http.HandleFunc("/healthz", healthz)
 http.HandleFunc("/webhook", webhook)
 listenAddress := ":8080"
 if os.Getenv("PORT") != "" {
  listenAddress = ":" + os.Getenv("PORT")
 }
 fmt.Printf("listening on: %v", listenAddress)
 log.Fatal(http.ListenAndServe(listenAddress, nil))
}
