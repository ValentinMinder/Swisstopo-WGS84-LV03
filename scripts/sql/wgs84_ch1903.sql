create OR REPLACE package convert_GIS_PCKG
as
-- The MIT License (MIT)
-- 
-- Copyright (c) 2014 Federal Office of Topography swisstopo, Wabern, CH and Joerg Schmidt, Rola AG, Zürich, CH
-- 
-- Permission is hereby granted, free of charge, to any person obtaining a copy
--  of this software and associated documentation files (the "Software"), to deal
--  in the Software without restriction, including without limitation the rights
--  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
--  copies of the Software, and to permit persons to whom the Software is
--  furnished to do so, subject to the following conditions:
-- 
-- The above copyright notice and this permission notice shall be included in
--  all copies or substantial portions of the Software.
-- 
-- THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
--  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
--   FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
--  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
--   LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
--  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
--   THE SOFTWARE.
--

-- Source: http://www.swisstopo.admin.ch/internet/swisstopo/en/home/topics/survey/sys/refsys/projections.html (see PDFs under "Documentation")
-- Updated 9 dec 2014
-- translated from python to oracle from Joerg Schmidt (Rola AG)
-- tested on Oracle 11.2.3.0
-- Please validate your results with NAVREF on-line service: http://www.swisstopo.admin.ch/internet/swisstopo/en/home/apps/calc/navref.html (difference ~ 1-2m)
   function WGStoCHy(lat in out float, lng in out float) return float;
   function WGStoCHx(lat in out float, lng in out float) return float;
   function CHtoWGSlat(y float, x float) return float;
   function CHtoWGSlng(y float, x float) return float;
   function DECtoSEX(angle float) return float;
-- Convert WGS lat/long (° dec) to CH y
end convert_GIS_PCKG;


create OR REPLACE package body convert_GIS_PCKG 
is
   function WGStoCHy(lat in out float, lng in out float) return float
   is
   lat_aux float;
   lng_aux float;
   y float;
   begin
-- Converts decimal degrees to sexagecimal seconds
    lat := DECtoSEX(lat);
    lng := DECtoSEX(lng);
--    # Auxiliary values (% Berne)
    lat_aux := (lat - 169028.66)/10000;
    lng_aux := (lng - 26782.5)/10000;

--    # Process Y
    y := (600072.37
         + 211455.93 * lng_aux
         - 10938.51 * lng_aux * lat_aux
         - 0.36 * lng_aux * power( lat_aux, 2 )
         - 44.54 * power( lng_aux, 3 ) );
    return y;
   end WGStoCHy;

--# Convert WGS lat/long (° dec) to CH x
   function WGStoCHx(lat in out float, lng in out float) return float
   is
   lat_aux float;
   lng_aux float;
   x float;
   begin
--    # Converts decimal degrees to sexagecimal seconds
    lat := DECtoSEX(lat);
    lng := DECtoSEX(lng);
--    # Auxiliary values (% Bern)
    lat_aux := (lat - 169028.66)/10000;
    lng_aux := (lng - 26782.5)/10000;
--    # Process X
    x := (200147.07
         + 308807.95 * lat_aux
         + 3745.25 * power(lng_aux, 2 )
         + 76.63 * power( lat_aux, 2 )
         - 194.56 * power( lng_aux, 2) * lat_aux
         + 119.79 * power( lat_aux, 3) );    
    return x;   
   end WGStoCHx;

--# Convert CH y/x to WGS lat
   function CHtoWGSlat(y float, x float) return float
   is
   y_aux float;
   x_aux float;
   lat float;
   begin
--    # Converts military to civil and to unit = 1000km
--    # Auxiliary values (% Bern)
    y_aux := (y - 600000)/1000000;
    x_aux := (x - 200000)/1000000;
--    # Process lat
    lat := (16.9023892
           + 3.238272 * x_aux
           - 0.270978 * power( y_aux, 2 )
           - 0.002528 * power( x_aux, 2 )
           - 0.0447 * power( y_aux, 2 ) * x_aux
           - 0.0140 * power( x_aux, 3 ) );

--    # Unit 10000" to 1 " and converts seconds to degrees (dec)
    lat := lat * 100/36;
    return lat;
  end CHtoWGSlat;

--# Convert CH y/x to WGS long
   function CHtoWGSlng(y float, x float) return float
   is
   y_aux float;
   x_aux float;
   lng float;   
   begin
--    # Converts military to civil and  to unit = 1000km
--    # Auxiliary values (% Bern)
    y_aux := (y - 600000)/1000000;
    x_aux := (x - 200000)/1000000;

--    # Process long
    lng := (2.6779094
           + 4.728982 * y_aux
           + 0.791484 * y_aux * x_aux
           + 0.1306 * y_aux * power( x_aux, 2 )
           - 0.0436 * power( y_aux, 3 ) );

--    # Unit 10000" to 1 " and converts seconds to degrees (dec)
    lng := lng * 100/36;
    return lng;
   end CHtoWGSlng;

--# Convert decimal angle to sexagesimal seconds
   function DECtoSEX(angle float) return float
   is
     deg float;
     mnt float;
     sec float;
   begin
--    # Extract DMS
    deg := angle;
    mnt := (angle-deg)*60;
    sec := (((angle-deg)*60)-mnt)*60;

--    # Result in seconds
    return sec + mnt * 60 + deg * 3600;
   end DECtoSEX;    
end convert_GIS_PCKG;

