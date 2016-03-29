package main

import (
  "os"
  "text/template"
  "strings"
  "bytes"
  "path/filepath"
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

  tmpls, err := filepath.Glob(confdir+"/*")
  if err != nil {
    logrus.Errorf("Failed to list templates: %v", err)
  }

  for _, tmpl := range tmpls {
    err = executeTemplate(tmpl, envMap)
    if err != nil {
      logrus.Errorf("Failed to execute template %v: %v", tmpl, err)
    }
  }
}

func envToMap() (map[string]string, error) {
  envMap := make(map[string]string)
  var err error

  for _, v := range os.Environ() {
    split_v := strings.Split(v, "=")
    envMap[split_v[0]] = strings.Join(split_v[1:], "=")
  }

  return envMap, err
}

func executeTemplate(tmpl string, envMap map[string]string) (error) {
  var err error


  t, err := template.ParseFiles(tmpl)
  if err != nil {
    return err
  }

  var augCode bytes.Buffer
  err = t.Execute(&augCode, envMap)
  if err != nil {
    return err
  }

  logrus.Infof("Code is:\n%v", augCode.String())

  //aug, _ := augeas.New("/", "", augeas.None)
  // TODO: err

  //aug.Srun

  return err
}
