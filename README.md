# Instalation

## Clone github repo
```
git clone https://github.com/JanKaczmarski/remitlyHomeExercise2024.git
cd remitlyHomeExercise2024
```

## Build and run
```
go build
./aiamrpv -path iamPolicy.json
```
#### OR
```
go run . -path iamPolicy.json
```
You can modify file that is passed to the script
`aiamrpv` stands for AWS::IAM:Role Policy Validation
# Workdir overview
```
.
├── README.md
├── go.mod
├── iamPolicy.json
├── main.go
├── support
│   ├── structs.go
│   ├── support.go
│   └── support_test.go
└── validate
    ├── resourcevalidate.go
    ├── resourcevalidate_test.go
    └── test_files
        ├── asteriskResourceAwsIamRolePolicy.json
        ├── multipleResourceAwsIamRolePolicy.json
        └── noResourceAwsIamRolePolicy.json
```
`support` dir holds structs, used to extract json to go readable data and support functions \
`validate` dir holds package that does the task and there are units test in there too\
`validate/test_files` are json files used to unit test ResourceValidation func\
`main.go` used to run program and accept parameters\
`iamPolicy.json` is a json file which holds the AWS::IAM::Role Policy provided in task description



