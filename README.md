## Installation

### Locally
#### Clone the GitHub repository
```bash
git clone https://github.com/JanKaczmarski/remitlyHomeExercise2024.git
cd remitlyHomeExercise2024
```

#### Build and run
```bash
go build
./aiamrpv -path iamPolicy.json
```
OR
```bash
go run . -path iamPolicy.json
```

You can modify the file that is passed to the script.

### Using Docker
#### Pull and run the container image
```bash
docker pull bigjack213/aiamrpv
docker run -it --entrypoint=/bin/bash bigjack213/aiamrpv -i
```

#### Run the binary
```bash
go run . -path iamPolicy.json
```
OR
```bash
./aiamrpv -path iamPolicy.json
```

Same as above, you can modify the path to the JSON file.

## Directory Structure
```
.
├── README.md
├── go.mod
├── iamPolicy.json
├── main.go
├── support
│   ├── structs.go
│   ├── support.go
│   └── support_test.go
└── validate
    ├── resourcevalidate.go
    ├── resourcevalidate_test.go
    └── test_files
        ├── asteriskResourceAwsIamRolePolicy.json
        ├── multipleResourceAwsIamRolePolicy.json
        └── noResourceAwsIamRolePolicy.json
```

The `support` directory holds structs used to extract JSON into Go-readable data and support functions. The `validate` directory contains packages that perform the validation task along with unit tests. The `validate/test_files` directory contains JSON files used for unit testing the `ResourceValidation` function. `main.go` is used to run the program and accept parameters, and `iamPolicy.json` is a JSON file containing the AWS IAM Role Policy provided in the task description.

