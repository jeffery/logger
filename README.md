# logger

Golang Logger

This project demonstrates proper log level filtering with test cases using the logrus logging library.

## Project Structure
```
logger/
├── examples/            # Example library usage
│   └── loglevel/
│       └── main.go
├── pkg/
│   └── config/          # Logger configuration package
│       └── logger.go
│   └── logger/          # Logger package
│       └── logger.go
├── go.mod
└── README.md
└── run.sh               # Script to run the application
```

## Set up the library for your project

```bash
go get github.com/jeffery/logger@v0.1.0 
```

## Running the example Application

```bash
./run.sh
```
