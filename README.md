[![Coverage Status](https://coveralls.io/repos/github/OpenIndustryCloud/fission-go-compute-weather-risk/badge.svg?branch=master)](https://coveralls.io/github/OpenIndustryCloud/fission-go-compute-weather-risk?branch=master)


# Compute Weather Risk API

Compute Weather Risk API uses the weather data as Input and computes risk based on it

## API reference

- [Weather History - Wunderground](https://www.wunderground.com/weather/api/d/docs?d=data/history)



## Authentication

No Authentication needed for this API

## Error hanlding

Empty Payload or malformed JSON would result in error reponse.

- Technical error :  `{"status":400,"message":"EOF"}`
- Malformed JSON `{"status":400,"message":"invalid character 'a' looking for beginning of object key string"}`

## Sample Input/Output

- Request payload

```
{"status":200,"response":{"version":"0.1"},"history":{"dailysummary":[{"fog":"0","rain":"1","maxtempm":"17","mintempm":"12","tornado":"0","maxpressurem":"1014","minpressurem":"1005","maxwspdm":"50","minwspdm":"13"}],"observations":[{"tempm":"17.0","tempi":"62.6","dewptm":"13.0","dewpti":"55.4","hum":"77","wspdm":"29.6","wspdi":"18.4","wgustm":"-9999.0","wgusti":"-9999.0","wdird":"230","wdire":"SW","vism":"10.0","visi":"6.2","pressurem":"1005","pressurei":"29.68","windchillm":"-999","windchilli":"-999","heatindexm":"-9999","heatindexi":"-9999","precipm":"-9999.00","precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR EGCN 012320Z 23016KT 9999 SCT031 17/13 Q1005"},{"tempm":"16.0","tempi":"60.8","dewptm":"12.0","dewpti":"53.6","hum":"77","wspdm":"24.1","wspdi":"15.0","wgustm":"-9999.0","wgusti":"-9999.0","wdird":"240","wdire":"WSW","vism":"10.0","visi":"6.2","pressurem":"1005","pressurei":"29.68","windchillm":"-999","windchilli":"-999","heatindexm":"-9999","heatindexi":"-9999","precipm":"-9999.00","precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR EGCN 012350Z 24013KT 9999 FEW030 16/12 Q1005"},{"tempm":"16.0","tempi":"60.8","dewptm":"11.0","dewpti":"51.8","hum":"72","wspdm":"25.9","wspdi":"16.1","wgustm":"-9999.0","wgusti":"-9999.0","wdird":"250","wdire":"WSW","vism":"10.0","visi":"6.2","pressurem":"1005","pressurei":"29.68","windchillm":"-999","windchilli":"-999","heatindexm":"-9999","heatindexi":"-9999","precipm":"-9999.00","precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR EGCN 020020Z 25014KT 9999 SCT030 SCT035 BKN041 16/11 Q1005"},{"tempm":"15.0","tempi":"59.0","dewptm":"11.0","dewpti":"51.8","hum":"77","wspdm":"20.4","wspdi":"12.7","wgustm":"-9999.0","wgusti":"-9999.0","wdird":"250","wdire":"WSW","vism":"10.0","visi":"6.2","pressurem":"1005","pressurei":"29.68","windchillm":"-999","windchilli":"-999","heatindexm":"-9999","heatindexi":"-9999","precipm":"-9999.00","precipi":"-9999.00","conds":"Partly Cloudy","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR EGCN 020050Z 25011KT 9999 FEW033 15/11 Q1005"},{"tempm":"15.0","tempi":"59.0","dewptm":"11.0","dewpti":"51.8","hum":"77","wspdm":"18.5","wspdi":"11.5","wgustm":"-9999.0","wgusti":"-9999.0","wdird":"230","wdire":"SW","vism":"10.0","visi":"6.2","pressurem":"1006","pressurei":"29.71","windchillm":"-999","windchilli":"-999","heatindexm":"-9999","heatindexi":"-9999","precipm":"-9999.00","precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR EGCN 020120Z 23010KT 9999 FEW029 BKN035 BKN040 15/11 Q1006"},{"tempm":"15.0","tempi":"59.0","dewptm":"11.0","dewpti":"51.8","hum":"77","wspdm":"13.0","wspdi":"8.1","wgustm":"-9999.0","wgusti":"-9999.0","wdird":"240","wdire":"WSW","vism":"10.0","visi":"6.2","pressurem":"1006","pressurei":"29.71","windchillm":"-999","windchilli":"-999","heatindexm":"-9999","heatindexi":"-9999","precipm":"-9999.00","precipi":"-9999.00","conds":"Scattered Clouds","icon":"partlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR EGCN 020150Z 24007KT 9999 SCT038 SCT046 15/11 Q1006"},{"tempm":"15.0","tempi":"59.0","dewptm":"11.0","dewpti":"51.8","hum":"77","wspdm":"18.5","wspdi":"11.5","wgustm":"-9999.0","wgusti":"-9999.0","wdird":"250","wdire":"WSW","vism":"10.0","visi":"6.2","pressurem":"1006","pressurei":"29.71","windchillm":"-999","windchilli":"-999","heatindexm":"-9999","heatindexi":"-9999","precipm":"-9999.00","precipi":"-9999.00","conds":"Mostly Cloudy","icon":"mostlycloudy","fog":"0","rain":"0","snow":"0","hail":"0","thunder":"0","tornado":"0","metar":"METAR EGCN 020220Z 25010KT 9999 BKN028 BKN034 15/11 Q1006"}]}}

```
- Response

```
{
    "status": 200,
    "riskScore": 80,
    "description": "Stormy Weather identified"
}
```


## Example Usage

## 1.  Deploy as Fission Functions

First, set up your fission deployment with the go environment.

```
fission env create --name go-env --image fission/go-env:1.8.1
```

To ensure that you build functions using the same version as the
runtime, fission provides a docker image and helper script for
building functions.


- Download the build helper script

```
$ curl https://raw.githubusercontent.com/fission/fission/master/environments/go/builder/go-function-build > go-function-build
$ chmod +x go-function-build
```

- Build the function as a plugin. Outputs result to 'function.so'

`$ go-function-build compute-weather-risk.go`

- Upload the function to fission

`$ fission function create --name compute-weather-risk --env go-env --package function.so`

- Map /compute-weather-risk to the compute-weather-risk function

`$ fission route create --method POST --url /compute-weather-risk --function compute-weather-risk`

- Run the function

```$ curl -d `sample request` -H "Content-Type: application/json" -X POST http://$FISSION_ROUTER/compute-weather-risk```

## 2. Deploy as AWS Lambda

> to be updated