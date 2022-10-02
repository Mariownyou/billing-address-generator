# Billing address generator
This package generates valid billing adress for given city


### Run 
You will need mapbox token and provide city in args

```bash
go build -ldflags "-X main.token=<MAPBOX_TOKEN>"
./postal_code -c <CITY>
```
  

```bash

./postal_code -c New-York
# 245 East 11th Street, New York, New York 10003, United States

./postal_code -c Москва
# Russia, Moscow, округ Тверской, 109012, Проезд Воскресенские Ворота 1а

```


### References
https://docs.mapbox.com/playground/  
https://docs.mapbox.com/playground/geocoding/
