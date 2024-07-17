# Weather CLI

Weather CLI is a command-line interface written in Golang that provides current weather by passing a city and using WeatherAPI from RapidAPI

## Installation

1. Clone the repository

```
git clone https://www.github.com/leoalipazaga/weathercli
```

2. Create and `.env` file with your RapidAPI key and host. Get your key and host from [RapidAPI](asddasd)

```
RAPI_API_URL=https://weatherapi-com.p.rapidapi.com/forecast.json
RAPID_API_KEY=
RAPID_API_HOST=
```

3. Build the app

```
make build
```

4. Move the executable to your path

```
mv ./bin/weather-cli /usr/local/bin
```

5. Run the app. City by default is **London**

```
weather-cli {city}
```

## Commands

All commands are run from the root of the project:
| Command | Action |
|--|--|
| make build | Generate executable into ./bin/ |
| make run | Run the app |

## Dependencies

- [Color](https://pkg.go.dev/github.com/fatih/color) colorized your outputs

- [DotEnv](https://pkg.go.dev/github.com/joho/godotenv) load env vars from `.env` file
