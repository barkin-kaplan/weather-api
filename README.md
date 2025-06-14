## Getting started
```bash
git clone https://github.com/barkin-kaplan/weather-api.git
```

### Create .env file
Sample .env is below. Note that postgre db have to be present prior to running this server. A sample postgre connection string is given below
```bash
WEATHER_API_KEY = your_api_key
WEATHER_API_URL = http://api.weatherapi.com/v1/current.json
WEATHER_STACK_KEY = your_api_key
WEATHER_STACK_URL = https://api.weatherstack.com/current
MAX_REQUEST_COUNT = 10
MAX_DELAY_SECONDS = 5
SERVER_PORT = 8080
POSTGRE_CONN_STRING = "host=localhost user=your_username dbname=weather_db port=5432 sslmode=disable"
```

### Start the server
```bash
go run main.go
```
