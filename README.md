# Object-Storage Railway
Object-Storage-Railway
- Created by Ramzi Al-Aruri
- This application a high level incopration of Object storage using Railway and Google Cloud Storage.


# Application Structure 

- Current App Structure

```
/high-alch-go
│
├── buckets                        // Database other configuration
|   ├── bucket_operator.go           // Getting the object values
│   └── create_bucket.go           // Getting the object values
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
|── *main.go*                       *// Entry point for app and connecting to Database *
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

4. init the project 
`go mod init documents/project/hello`

5. check the go.mod file
`cat go.mod`

6. run `go install`

7. create a hello.go file 

add this code snippet
```
package main

import "fmt"

func main() {
    fmt.Println("Hello, SlimeCorp")
}

```

8. finally run the file 
`go run hello.go`


# Project Packages
- Fiber: https://gofiber.io/
- Google Cloud: 


# Railway Setup 


