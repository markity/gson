package main

import (
	"encoding/json"
	"fmt"
	"gson/unmarshal"
	"reflect"
)

var com = `
{
    "name":"xxx",
    "num":3,
    "sites": [
        { "name":"Google", "info":[ "x", "xx xxxx", "xxx x" ] },
        { "name":"Runoob", "info":[ "xxxx", "xxxxxx", "xxxx" ] },
        { "name":"Taobao", "info":[ "xxxx", "我测试下utf8哈" ] }
    ]
}
`

var com2 = `
{  
    "resultcode":"200",  
    "reason":"成功的返回",  
    "result":{  
    "company":"顺丰",  
    "com":"sf",  
    "no":"575677355677",  
    "list":[  
        {  
        "datetime":"2013-06-25 10:44:05",  
        "remark":"已收件",  
        "zone":"台州市"  
        },  
        {  
        "datetime":"2013-06-25 11:05:21",  
        "remark":"快件在 台州 ,准备送往下一站 台州集散中心 ",  
        "zone":"台州市"  
        }  
    ],  
    "status":1  
    },  
    "error_code":0  
}  
`

var com3 = `
{"latitude":-35.2809}
`

var com4 = `
{
	"latitude": -35.2809,
	"longitude": 149.13,
	"timezone": "Australia/Sydney",
	"offset": 10,
	"hourly": {
		"summary": "Clear throughout the day.",
		"icon": "clear-day",
		"data": [{
			"time": 1492610400,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 48.24,
			"apparentTemperature": 48.24,
			"dewPoint": 47.73,
			"humidity": 0.98,
			"windSpeed": 2.91,
			"windBearing": 152,
			"visibility": 5.65,
			"cloudCover": 0.03,
			"pressure": 1030
		}, {
			"time": 1492614000,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 48.42,
			"apparentTemperature": 48.42,
			"dewPoint": 48.42,
			"humidity": 1,
			"windSpeed": 0,
			"visibility": 6.2,
			"pressure": 1029.99
		}, {
			"time": 1492617600,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 48.56,
			"apparentTemperature": 47.7,
			"dewPoint": 48.05,
			"humidity": 0.98,
			"windSpeed": 3.38,
			"windBearing": 31,
			"visibility": 6.2,
			"pressure": 1029.74
		}, {
			"time": 1492621200,
			"summary": "Foggy",
			"icon": "fog",
			"precipType": "rain",
			"temperature": 46.49,
			"apparentTemperature": 46.49,
			"dewPoint": 45.93,
			"humidity": 0.98,
			"windSpeed": 0.69,
			"windBearing": 16,
			"visibility": 0.59,
			"cloudCover": 0.14,
			"pressure": 1029.42
		}, {
			"time": 1492624800,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 48.35,
			"apparentTemperature": 46.51,
			"dewPoint": 48.35,
			"humidity": 1,
			"windSpeed": 4.65,
			"windBearing": 167,
			"visibility": 6.2,
			"pressure": 1029.58
		}, {
			"time": 1492628400,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 46.73,
			"apparentTemperature": 46.73,
			"dewPoint": 46.25,
			"humidity": 0.98,
			"windSpeed": 1.05,
			"windBearing": 7,
			"visibility": 6.2,
			"pressure": 1029.61
		}, {
			"time": 1492632000,
			"summary": "Mostly Cloudy",
			"icon": "partly-cloudy-night",
			"precipType": "rain",
			"temperature": 47.05,
			"apparentTemperature": 47.05,
			"dewPoint": 45.96,
			"humidity": 0.96,
			"windSpeed": 2.95,
			"windBearing": 90,
			"visibility": 5.61,
			"cloudCover": 0.75,
			"pressure": 1029.93
		}, {
			"time": 1492635600,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 49.93,
			"apparentTemperature": 49.93,
			"dewPoint": 49.93,
			"humidity": 1,
			"windSpeed": 2.86,
			"windBearing": 90,
			"visibility": 6.2,
			"pressure": 1030.28
		}, {
			"time": 1492639200,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 52.4,
			"apparentTemperature": 52.4,
			"dewPoint": 51.47,
			"humidity": 0.97,
			"windSpeed": 2.32,
			"windBearing": 101,
			"visibility": 6.2,
			"pressure": 1030.84
		}, {
			"time": 1492642800,
			"summary": "Overcast",
			"icon": "cloudy",
			"precipType": "rain",
			"temperature": 56.47,
			"apparentTemperature": 56.47,
			"dewPoint": 52.32,
			"humidity": 0.86,
			"windSpeed": 1.51,
			"windBearing": 68,
			"visibility": 5.25,
			"cloudCover": 1,
			"pressure": 1031.06
		}, {
			"time": 1492646400,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 62.51,
			"apparentTemperature": 62.51,
			"dewPoint": 54.09,
			"humidity": 0.74,
			"windSpeed": 4.85,
			"windBearing": 61,
			"visibility": 6.2,
			"pressure": 1030.51
		}, {
			"time": 1492650000,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 64.46,
			"apparentTemperature": 64.46,
			"dewPoint": 51.78,
			"humidity": 0.63,
			"windSpeed": 1.28,
			"windBearing": 337,
			"visibility": 6.2,
			"pressure": 1029.87
		}, {
			"time": 1492653600,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 67.6,
			"apparentTemperature": 67.6,
			"dewPoint": 49.74,
			"humidity": 0.53,
			"windSpeed": 4.52,
			"windBearing": 3,
			"visibility": 6.2,
			"pressure": 1028.82
		}, {
			"time": 1492657200,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 69.93,
			"apparentTemperature": 69.93,
			"dewPoint": 48.56,
			"humidity": 0.47,
			"windSpeed": 5.43,
			"windBearing": 225,
			"visibility": 6.2,
			"pressure": 1028.01
		}, {
			"time": 1492660800,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 69.82,
			"apparentTemperature": 69.82,
			"dewPoint": 47.8,
			"humidity": 0.46,
			"windSpeed": 3.54,
			"windBearing": 311,
			"visibility": 6.2,
			"pressure": 1026.99
		}, {
			"time": 1492664400,
			"summary": "Partly Cloudy",
			"icon": "partly-cloudy-day",
			"precipType": "rain",
			"temperature": 69.81,
			"apparentTemperature": 69.81,
			"dewPoint": 46.36,
			"humidity": 0.43,
			"windSpeed": 5.82,
			"windBearing": 350,
			"visibility": 6.79,
			"cloudCover": 0.31,
			"pressure": 1026.38
		}, {
			"time": 1492668000,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 69.72,
			"apparentTemperature": 69.72,
			"dewPoint": 48.89,
			"humidity": 0.48,
			"windSpeed": 5.73,
			"windBearing": 265,
			"visibility": 6.2,
			"pressure": 1026.34
		}, {
			"time": 1492671600,
			"summary": "Clear",
			"icon": "clear-day",
			"precipType": "rain",
			"temperature": 64.75,
			"apparentTemperature": 64.75,
			"dewPoint": 47.57,
			"humidity": 0.54,
			"windSpeed": 3.76,
			"windBearing": 307,
			"visibility": 6.2,
			"pressure": 1026.36
		}, {
			"time": 1492675200,
			"summary": "Partly Cloudy",
			"icon": "partly-cloudy-night",
			"precipType": "rain",
			"temperature": 57.57,
			"apparentTemperature": 57.57,
			"dewPoint": 48.9,
			"humidity": 0.73,
			"windSpeed": 4.49,
			"windBearing": 67,
			"visibility": 6.2,
			"cloudCover": 0.31,
			"pressure": 1026.83
		}, {
			"time": 1492678800,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 62.72,
			"apparentTemperature": 62.72,
			"dewPoint": 48.74,
			"humidity": 0.6,
			"windSpeed": 8.12,
			"windBearing": 77,
			"pressure": 1027.15
		}, {
			"time": 1492682400,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 54.89,
			"apparentTemperature": 54.89,
			"dewPoint": 46.81,
			"humidity": 0.74,
			"windSpeed": 2.82,
			"windBearing": 84,
			"pressure": 1027.33
		}, {
			"time": 1492686000,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 54.22,
			"apparentTemperature": 54.22,
			"dewPoint": 47.43,
			"humidity": 0.78,
			"windSpeed": 1.91,
			"windBearing": 63,
			"visibility": 6.2,
			"cloudCover": 0.02,
			"pressure": 1027.75
		}, {
			"time": 1492689600,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 54.21,
			"apparentTemperature": 54.21,
			"dewPoint": 48.32,
			"humidity": 0.8,
			"windSpeed": 1.51,
			"windBearing": 190,
			"visibility": 6.2,
			"cloudCover": 0,
			"pressure": 1027.64
		}, {
			"time": 1492693200,
			"summary": "Clear",
			"icon": "clear-night",
			"precipType": "rain",
			"temperature": 52.26,
			"apparentTemperature": 52.26,
			"dewPoint": 47.73,
			"humidity": 0.85,
			"windSpeed": 2.65,
			"windBearing": 81,
			"visibility": 6.2,
			"cloudCover": 0,
			"pressure": 1027.49
		}]
	},
	"daily": {
		"data": [{
			"time": 1492610400,
			"summary": "Clear th我偷偷地放点utf8字符进去roughout the day.",
			"icon": "clear-day",
			"sunriseTime": 1492633936,
			"sunsetTime": 1492673690,
			"moonPhase": 0.77,
			"precipType": "rain",
			"temperatureMin": 46.49,
			"temperatureMinTime": 1492621200,
			"temperatureMax": 69.93,
			"temperatureMaxTime": 1492657200,
			"apparentTemperatureMin": 46.49,
			"apparentTemperatureMinTime": 1492621200,
			"apparentTemperatureMax": 69.93,
			"apparentTemperatureMaxTime": 1492657200,
			"dewPoint": 48.63,
			"humidity": 0.77,
			"windSpeed": 1.07,
			"windBearing": 53,
			"visibility": 5.88,
			"cloudCover": 0.28,
			"pressure": 1028.66
		}]
	}
}
`

/*

{
    "name":"xxx",
    "num":3,
    "sites": [
        { "name":"Google", "info":[ "x", "xx xxxx", "xxx x" ] },
        { "name":"Runoob", "info":[ "xxxx", "xxxxxx", "xxxx" ] },
        { "name":"Taobao", "info":[ "xxxx", "我测试下utf8哈" ] }
    ]
}*/

type mystruct struct {
	Name string   `gson:"name"`
	Info []string `gson:"info"`
}

type MyStruct struct {
	Name  *string     `gson:"name,omitempty"`
	Num   int         `gson:"num"`
	Sites []*mystruct `gson:"sites"`
}

func main() {
	var m map[string]any
	fmt.Println(unmarshal.Unmarshal(com, &m), m)
	fmt.Println(unmarshal.Unmarshal(com2, &m), m)
	fmt.Println(unmarshal.Unmarshal(com3, &m), m)
	fmt.Println(unmarshal.Unmarshal(com4, &m), m)

	err := unmarshal.Unmarshal(com4, &m)
	if err != nil {
		panic(err)
	}
	m2 := make(map[string]interface{})
	json.Unmarshal([]byte(com4), &m2)
	if fmt.Sprint(m) != fmt.Sprint(m2) {
		fmt.Println("failed")
	} else {
		fmt.Println("ok!")
	}

	a := 1
	t := reflect.TypeOf(&a)
	fmt.Println(t.Kind())

	println("-------------")

	ms := MyStruct{}
	err = unmarshal.Unmarshal(com, &ms)
	if err != nil {
		panic(err)
	}

	fmt.Println(*ms.Name)
	fmt.Println(ms.Num)
}
