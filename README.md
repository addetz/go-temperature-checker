# railway-go-demo
This entire application is written in Go! 🥳 
The JS code was generated with GopherJS. 

## Prerequisites
1. [Go 1.21](https://go.dev/doc/install)
1. [GopherJS](https://github.com/gopherjs/gopherjs) 
```
$ go install github.com/gopherjs/gopherjs@v1.18.0-beta3
```

## Environment variables 
Sign up to [OpenWeatherMap](https://openweathermap.org/appid) to get your own API key. 
```
export WEATHER_API_KEY=XXX
export SERVER_PORT=XXX
```

## Run it 
After setting our variables, we run the server like any old server:
```
$ go run server.go
```

Alternatively, you can run the server in Docker: 
```
$ docker build -f Dockerfile -t tempserver .
$ docker run -dt -e WEATHER_API_KEY=$WEATHER_API_KEY -p 8080:8080/tcp tempserver
```