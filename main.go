package main

import (
  "os"
  "text/template"
  "strings"
  "bytes"
  "github.com/Sirupsen/logrus"
  //"honnef.co/go/augeas"
)

func main() {
  confdir := os.Getenv("CONTUNER_CONFDIR")
  if confdir == "" {
    confdir = "/etc/contuner/conf.d"
  }

  envMap, err := envToMap()
  if err != nil {
    logrus.Errorf("Failed to parse environment: %v", err)
  }

  // Create a new template and parse the letter into it.
  t, err := template.ParseGlob(confdir+"/*.tmpl")
  if err != nil {
    logrus.Errorf("Failed to initialize template: %v", err)
    return
  }

  var augCode bytes.Buffer
  err = t.Execute(&augCode, envMap)
  if err != nil {
    logrus.Errorf("Failed to execute template: %v", err)
  }

  logrus.Infof("Code is:\n%v", augCode.String())

  //aug, _ := augeas.New("/", "", augeas.None)
  // TODO: err

  //aug.Srun
}

func envToMap() (map[string]string, error) {
  envMap := make(map[string]string)
  var err error

  for _, v := range os.Environ() {
    split_v := strings.Split(v, "=")
    envMap[split_v[0]] = split_v[1]
  }

  return envMap, err
}
