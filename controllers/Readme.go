# Web Scrapping

Build a script / tool in Node.JS or Go-lang that will scrape the data on this website and 
store it in a JSON file. Share this script. 

## Explicit Features/Tasks

- [X] Scrap mcc and mnc data from website and package into json
- [X] Rest api to query The network name and country by specifying mcc and mnc
- [X] Rest API All the networks in a specific country based on mcc or country name
- [X] Share this code and screenshots of API requests and the response (either on browser/postman etc)
- [X] Create a docker image of this code base and public on docker hub. Share link to this docker image. 
- [ ] BONUS - Use this API to create a visualization on google maps (using the google maps API) that displays the networks for each country in Africa when you hover over a country or click on a pin 


## Features Implemented in the application

- [X] In memory DB storage (Database functionlity)
- [ ] Application testability (Unit tests to achieve coverage of < 90%)
- [X] Application documentation (Feature and code documentation on readme file)
- [X] Code versioning (Git hub)
- [X] Deployment and delivery via Heroku
- [X] Logging and log management.


### Application setup

Get the application source code from github

```bash
git clone https://github.com/ryananyangu/roamtech.git
```

Change directory into the application directory

```bash
cd roamtech
```

Install application dependancies

```bash
go mod tidy
```

Build the application docker image in unix systems
```bash
make build
```

Run the built image in unix systems

```bash
make run
```


#### APIs Implementation and Examples

### Lookup by country

[loookup by country](http://localhost:8080/api/v1/lookup/country/networks)

Sample request via curl

```bash
curl --location --request GET 'http://localhost:8080/api/v1/lookup/country/networks?country=Zimbabwe'
```

or mcc code

```bash
curl --location --request GET 'http://localhost:8080/api/v1/lookup/country/networks?mcc=412'
```

Sample response

```json
[
    {
        "MCC": "648",
        "MNC": "01",
        "ISO": "zw",
        "Country": "Zimbabwe",
        "CountryCode": "263",
        "Network": "Net One "
    },
    {
        "MCC": "648",
        "MNC": "03",
        "ISO": "zw",
        "Country": "Zimbabwe",
        "CountryCode": "263",
        "Network": "Telecel "
    },
    {
        "MCC": "648",
        "MNC": "04",
        "ISO": "zw",
        "Country": "Zimbabwe",
        "CountryCode": "263",
        "Network": "Econet "
    }
]
```

### Scrap data from site

[scrap data](http://localhost:8080/api/v1/mcc-mnc/scrapper)

Sample request

```bash
curl --location --request GET 'http://localhost:8080/api/v1/mcc-mnc/scrapper'
```

Sample response

```json
[
    {
        "MCC": "648",
        "MNC": "01",
        "ISO": "zw",
        "Country": "Zimbabwe",
        "CountryCode": "263",
        "Network": "Net One "
    },
    {
        "MCC": "648",
        "MNC": "03",
        "ISO": "zw",
        "Country": "Zimbabwe",
        "CountryCode": "263",
        "Network": "Telecel "
    },
    {
        "MCC": "648",
        "MNC": "04",
        "ISO": "zw",
        "Country": "Zimbabwe",
        "CountryCode": "263",
        "Network": "Econet "
    }
	...
]
```


### Lookup by mcc and mnc combined

[mcc & mnc lookup](http://localhost:8080/api/v1/lookup/mcc-mnc)

Sample request

```bash
curl --location --request GET 'http://localhost:8080/api/v1/lookup/mcc-mnc?mcc=289&mnc=88'
```

Sample response for network

```json
{
    "MCC": "289",
    "MNC": "88",
    "ISO": "ge",
    "Country": "Abkhazia",
    "CountryCode": "7",
    "Network": "A-Mobile "
}
```