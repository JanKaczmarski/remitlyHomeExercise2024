package support


type Policy struct {
  PolicyName string `json:"PolicyName"`
  PolicyDocument PolicyDocument `json:"PolicyDocument"`
} 

type PolicyDocument struct {
  Version string `json:"Version"`
  Statement []Element `json:"Statement"`
}

type Element struct {
  Sid string `json:"Sid"`
  Effect string `json:"Effect"`
  Action []string `json:"Action"`
  Resource interface{} `json:"Resource"`
}