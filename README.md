# go-testing


--**************************
--**************************
--**************************

primeapp

// mod initialization:
go mod init primeapp

Test files should be located beside the target file.
file name convention: [original_file_name]_test.go

Function name should start from Test_[function_name](t *testing.T){...}

// run test in verbose mode:
go test -v .

// show coverage:
go test -v -cover .

// generate html output with covered code:
go test -coverprofile=coverage.out

// show the covered code: 
go tool cover -html=coverage.out

// combine:
go test -coverprofile=coverage.out && go tool cover -html=coverage.out

--**************************
--**************************
--**************************
