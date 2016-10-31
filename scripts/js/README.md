# JS lib

## Usage

Import this file with:

```html
<script src="wgs84_ch1903.js">
```

```js
var coords = Swisstopo.WGStoCH(lat, lng); // coords = [y, x]
var coords = Swisstopo.CHtoWGS(y, x); // coords = [longitude, latitude]
```
