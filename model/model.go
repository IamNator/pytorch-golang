package model

import (
  "encoding/json"
)


type PytorchModel struct {
  Name string. `json:"name"`
	Version string `json:"version"`
	Model []byte `json:"model"`
	Key []byte `json:"key"`
}

func (p *PytorchModel) SetModel(b any) error {
  byteData, err := json.Marshal(b)
  if err != nil {
    return err
  }
  
  copy(&p.Model, byteData)
}


func (p *PytorchModel) SetKey(b any) error {
  byteData, err := json.Marshal(b)
  if err != nil {
    return err
  }
  
  copy(&p.Key, byteData)
}




type Request struct {
  Input string `json:input`
}


type Response struct {
  Output string `json:output`
}


