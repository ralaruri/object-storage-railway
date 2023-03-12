# Object-Storage Railway
Object-Storage-Railway
- Created by Ramzi Al-Aruri
- This application a high level incopration of Object storage using Railway and Google Cloud Storage.


# Application Structure 

- Current App Structure

```
/Object-Storage-Railway
│
├── buckets                        // Database other configuration
|   ├── bucket_operator.go         // Bucket Operations (Write, Read)
│   └── create_bucket.go           // Creating the Bucket, Connection and decoding
|
|── handlers                        // Logic for API (Eventually will include router)
|   └── health.go                   // Health Check on API
|
|── utils                           // constants and tools (i.e.) dotenv
|   └── dotenv.go                   // setting up the .env extract
|
|── object_creator                  // Mocking Json file creator
|   └── cheese_creator.go           // Mock json creation & deletion
|
|── .env_example                    // Example .env
|── go.mod                          // configuration setup by Go
|── go.sum                          // configuration setup by Go
|── *main.go*                       *// Entry point for app *
└── README.md                       // This Readme!
    
```


# Go Setup

### Installing Go using asdf
- This is how I manage my local Go & Python installations 
- You can have a seperate instance of Go versions for each project. 


1. Installing asdf using homebrew
- `brew install asdf`
- `echo -e "\n. $(brew --prefix asdf)/libexec/asdf.sh" >> ~/.zsh_profile`

2. Installing go 
- `asdf plugin add go`
- `asdf install go latest`

3. Setting a Version 
- `asdf global go latest`
if you need to a set a local env.
- `asdf local go latest`


### Running Go Files Locally an Example Project.

1. Create a Directory i.e. `hello`

2. Go into the dircectory `cd hello`

3. get the path of the directory `pwd` 
`>> documents/projects/hello`

OR just name it based on the directory i.e. object-storage.app

4. init the project 
`go mod init object-storage.app`

5. check the go.mod file
`cat go.mod`

6. run `go install`

7. create a main.go file 

add this code snippet
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, SlimeCorp")
}

```

8. finally run the file 
`go run main.go`


# Encoding & Decoding Base64 JSON key.
- Never commit your json service account details from GCP. 
- This is just one method to get service accounts to work in Railway for GCP specfically

1. In your terminal run `cat object-storage-railway.json | base64`
>>> output will be a encoded base64 string copy and use that and .ENV variable or value inside of Railway.

2. Inside the code we have a function that will decode the string at runtime to use the service account permissions. 

```
func CovertStringToJSON(env_details string) []byte {
	decoded_json, err := base64.StdEncoding.DecodeString(env_details)
	if err != nil {
		panic(err)
	}

	return decoded_json

}
```


# Project Packages
- Fiber: https://gofiber.io/
- Google Cloud: https://cloud.google.com


# Railway Setup:
- Railway: railway.app
1. Link Git Repo to Railway
2. Add encoded base64 service account value to the variable of the railway project for this repo. 



