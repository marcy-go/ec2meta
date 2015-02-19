package ec2meta

import (
  "net/http"
  "io/ioutil"
)

func GetMetaData(url string) (string, error) {
  res, err := http.Get("http://169.254.169.254/latest/meta-data" + url)
  defer res.Body.Close()
  if err != nil {
    return "", err
  }
  body, err := ioutil.ReadAll(res.Body)
  if err != nil {
    return "", err
  }
  return string(body), nil
}

func GetInstanceId() (string, error) {
  return getMetaData("/instance-id")
}

func GetRegion() (string, error) {
  meta, err := getMetaData("/placement/availability-zone")
  if err != nil {
    return "", err
  }
  end := len(meta) - 1
  return meta[0:end], nil
}
