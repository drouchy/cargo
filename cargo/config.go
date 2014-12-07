package cargo

import (
  "encoding/json"
  "io/ioutil"
)

type DigitalOcean struct {
  Token string `json="token"`
}

type Config struct {
  DigitalOcean DigitalOcean `json:"digital_ocean"`
}

func LoadConfig(filename string) Config {
  data, _ := ioutil.ReadFile(filename)
  config := Config{}
  json.Unmarshal(data, &config)
  return config
}
