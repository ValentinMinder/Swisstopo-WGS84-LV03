# Swisstopo Scripts GPS WGS84 <-> LV03 (CH1903)

**Programmatic computer scripts to transpose GPS WGS-84 coordinates to/from the swiss military and civilian coordinate system (LV-03/CH-1903).**

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

### Original swisstopo files

- Version: WGS84<->CH1903/LV03 (05.1999)
- Authors: U. Marti swisstopo / H. Dupraz EPFL
- Published: 10.07.2015
- Software system:	Windows / DOS / Unix / Linux / MacOS
- **Programming language:	C# / JAVA / PHP / Python / R / SQL**
- **Price:	Free of charge**
- **License:	Open Source (MIT License)**
- Remarks:	Comments in English, No support
- Copyright (c) 2014 Federal Office of Topography swisstopo, Wabern, CH

### Addtionnal languages and contributors

- **Additionnal programming language: GO / JS**
- Additionnal Authors and Contributors:
	- [Valentin Minder] (https://github.com/ValentinMinder)
	- [Sacha Bron](https://github.com/BinaryBrain)
	- [Basile Vu](https://github.com/Flagoul)
	- [Michael Monay](https://github.com/micmonay)

**Feel free to add some piece of code or additionnal language and do a pull request**. To make it easy:

- insert the same MIT licence as swisstopo at the start of your file
- add your name in your file, preferably after the swisstopo copyright in the MIT license
- add your name on the list above (main README.md) in the list of contributors

### Copyright notice

**On 1st March 2016, this content was freely available on swisstopo website. It consists of all the above text in this Readme.md, in three languages English / Deutsch / Français, and the zip archives of languages C# / JAVA / PHP / Python / R / SQL programming languages.**

- [Original on swisstopo.ch](http://www.swisstopo.admin.ch/internet/swisstopo/en/home/products/software/products/skripts.html) (available 1st March 2016, broken on 1st July 2016)
- [Copy on mont-terri.ch](http://www.mont-terri.ch/internet/swisstopo/en/home/products/software/products/skripts.html) (available on 1st July 2016, broken 1st January 2017)
- [Copy on cadastre.ch](https://www.cadastre.ch/internet/swisstopo/en/home/products/software/products/skripts.html) (available on 1st July 2016, broken 1st January 2017)

### Alternative services of swisstopo

- [Reference systems](https://www.swisstopo.admin.ch/en/knowledge-facts/surveying-geodesy/reference-systems.html) of surveying geodesy.

- [NAVREF](https://www.swisstopo.admin.ch/en/maps-data-online/calculation-services/navref.html) allows you to transform Swiss projection coordinates into global WGS84 coordinates (GPS) and vice versa.
*It supports both LV03 and LV95.*

- [REFRAME](https://www.swisstopo.admin.ch/en/maps-data-online/calculation-services/reframe.html) allows you to perform coordinates transformations in planimetry and/or altimetry. 
*Supported file formats: text with separator, LTOP, DXF, ESRI Shape, INTERLIS 1, Adalin OneOne, Topobase .K*

- [GIS tools and applications (FME)](https://shop.swisstopo.admin.ch/en/products/geo_software/GIS_info). swisstopo has developed transformers for the transition of reference frames in position an /or height with FME software. All files which can be read by the FME software (many GIS, CAD and text formats) can be easily processed with this transformer. The CHENyx06 dataset (LV03->LV95 ou LV03->ETRS89) is also available as NTv2 regular grid.

- [REST web geoservices (REFRAME Web API)](https://www.swisstopo.admin.ch/en/maps-data-online/calculation-services/m2m.html). swisstopo offers different REST ("REpresentational State Transfer") services which allow you to embed coordinates transformations in your own software products or web services. *REST geoservices for reference frame changes WGS84-LV95-LV03/Bessel-LHN95-LN02*

- [Developer resources (DLL / JAR)](https://shop.swisstopo.admin.ch/en/products/geo_software/DLL_info). Swisstopo provides shared libraries (DLL or JAR) for software developers which allows transformation of coordinates and heights. <br>These transformations can be easily integrated into existing applications (.NET, C++, Visual Basic, Java), websites (Java, MONO, Silverlight) or macros (e.g. VBA). 

- [GeoSuite (REFRAME/TRANSINT)](https://shop.swisstopo.admin.ch/en/products/geo_software/GeoSuite_info) consists in a suite of calculation, file editing and data visualizing tools, grouped in one modern and efficient application. It is also optimized for the latest computers and operating systems.

