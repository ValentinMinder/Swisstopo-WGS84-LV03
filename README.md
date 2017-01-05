# Swisstopo Scripts GPS WGS84 <-> LV03 (CH1903)

**Programmatic computer scripts to transpose GPS WGS-84 coordinates to/from the swiss militay and civilian coordinate system (LV-03/CH-1903).**

## Description

### Scripts for WGS84<->LV03

**swisstopo created some script examples allowing to calculate the Swiss projection. This allows you to convert WGS84 coordinates to CH1903/LV03 or vice versa.**

In these scripts, the approximate formulas are used for the direct conversion of ellipsoidal WGS84 coordinates to Swiss plane coordinates LV03 (CH1903). 
These formulas are primarily for navigation purposes and may not be used for official or geodetic survey applications!

### Skripts für WGS84 <-> CH1903

**swisstopo erstellt einige Skript-Beispiele, die erlauben die Schweizer Projektion zu berechnen. Damit können Sie WGS84-Koordinaten nach CH1903/LV03 umwandeln oder umgekehrt.**

In diese Skripts werden Näherungsformeln für die direkte Umrechnung von ellipsoidischen WGS84-Koordinaten nach Schweizer Projectionskoordinaten LV03 (CH1903) benutzt. 
Diese Formeln sind vor allem für Navigationszwecke vorgesehen und dürfen nicht für die amtliche Vermessung oder für geodätische Anwendungen verwendet werden.

### Scripts pour WGS84 <-> LV03

**swisstopo a créé quelques exemples de script permettant de calculer la projection suisse. Avec ceux-ci, vous pouvez transformer des coordonnées WGS84 en coordonnées suisses CH1903/MN et réciproquement.**

Les formules approchées pour la transformation directe de coordonnées WGS84 en coordonnées en projection suisse MN03 (CH1903) sont utilisées dans ces scripts. 
Ces formules sont avant tout destinées à des applications de navigation et ne doivent pas être utilisées en mensuration nationale ou pour des travaux géodésiques !

## Details

### Supported reference frames

- WGS84 global geographic coordinates (degrees or degrees / minutes / seconds)
- Swiss national coordinates (CH1903/LV03)
 

### Map projection

- Oblique, conformal cylindrical projection (Mercator projection)
- Bessel ellipsoid 1841
- The projection center is the fundamental point at the old observatory in Bern (Longitude 7 ° 26 '22:50 "/ latitude 46 ° 57' 08.66" -> coordinates 600'000 .000 East / North 200'000 .000)
- Approximation (accuracy on the 1-meter level) 

### REST web geoservices (REFRAME Web API)

*REST geoservices for reference frame changes WGS84-LV95-LV03/Bessel-LHN95-LN02*

swisstopo offers different REST ("REpresentational State Transfer") services which allow you to embed coordinates transformations in your own software products or web services.
[Swisstopo Webservices](https://www.swisstopo.admin.ch/en/maps-data-online/calculation-services/m2m.html)

## About 

- Version: WGS84<->CH1903/LV03 (05.1999)
- Authors: U. Marti swisstopo / H. Dupraz EPFL
- Published: 10.07.2015
- Software system:	Windows / DOS / Unix / Linux / MacOS
- **Programming language:	C# / JAVA / GO / JS / PHP / Python / R / SQL**
- **Price:	Free of charge**
- **License:	Open Source (MIT License)**
- Remarks:	Comments in English, No support

***On 1st March 2016, this content was freely available on swisstopo website.***

- [Original (broken on 1st July 2016)](http://www.swisstopo.admin.ch/internet/swisstopo/en/home/products/software/products/skripts.html)
- [Copy (available on 1st July 2016)](http://www.mont-terri.ch/internet/swisstopo/en/home/products/software/products/skripts.html)
