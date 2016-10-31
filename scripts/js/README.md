# JS lib

## Usage

Import this file with:

```html
<script src="wgs84_ch1903.js">
```

Then call:

```js
var coordsGPS = Swisstopo.CHtoWGS(y, x); // coords = [longitude, latitude]
var CoordsCH = Swisstopo.WGStoCH(lat, lng); // coords = [y, x]
```
