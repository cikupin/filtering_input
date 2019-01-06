# Simple Go Climate
Only a really simple Go endpoint with go-ozzo input validation (https://github.com/go-ozzo/ozzo-validation) & minimalist docker image using scratch.


## Quick Start
### 1. Go and Dep Installation
##### A. Go
1. Download Go & follow the installation instruction found here: https://golang.org/doc/install

##### B. Go Dep
1. Go Dep is the official *"experiment"* Go package manager. Please follow the instruction here https://github.com/golang/dep

### 2. Build
To build :

    $ [ -d $GOPATH/src/github.com/cikupin ] || mkdir -p $GOPATH/src/github.com/cikupin
    $ cd $GOPATH/src/github.com/cikupin
    $ git clone git@github.com:cikupin/filtering_input.git
    $ cd filtering_input
    $
    $ dep ensure
    $ go build -race

### 3. Run
To run :

    $ ./filtering_input

### 2. Running Go Apps

Run this app by accessing `http://localhost:8080/` and using **POST** method.  

Use body like the example below:

```json  
{
    "id": 2,
    "first_name": "foo",
    "last_name": "bar",
    "email": "123@123.com",
    "condition": true
}
```

### 3. Minimalist docker image
##### A. Building docker images

To build :

    $ docker build -t filtering-input-img .

where `filtering-input-img` is docker images name.

##### B. Running docker images

Use thid command to run this image:

    $ docker run --rm -p 8080:8080 filtering-input-img
