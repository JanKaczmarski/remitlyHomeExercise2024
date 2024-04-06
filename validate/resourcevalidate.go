package validate

import (
    "encoding/json"
    "os"
    "io"
    "aiamrpv/support"
    )


func ResourceValidate(fileName string)(bool){
    // open json file
    jsonFile, err := os.Open(fileName)

    // handle error
    support.HandleError(err)

    // close file after all functions finished running
    defer jsonFile.Close()

    // get file content as bytes
    byteValue, err := io.ReadAll(jsonFile)
    support.HandleError(err)

    // policy is a struct that will hold our json
    // in go-readable format
    var policy support.Policy

    // extract json to policy struct
    json.Unmarshal(byteValue, &policy)

    // extract "Statement" from json
    var statement = policy.PolicyDocument.Statement

    // There can be multiple elements in Statement
    // so we'll check all of them
    for i := 0; i< len(statement); i++ {
        if statement[i].Resource == "*"{
            return false
        }
    }
    return true
}


