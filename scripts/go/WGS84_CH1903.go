// The MIT License (MIT)
//
// Copyright (c) 2014 Federal Office of Topography swisstopo, Wabern, CH
// Copyright (c) 2017 Michael Monay https://github.com/micmonay
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
//  of this software and associated documentation files (the "Software"), to deal
//  in the Software without restriction, including without limitation the rights
//  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
//  copies of the Software, and to permit persons to whom the Software is
//  furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
//  all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
//  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
//  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
//  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
//  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
//  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
//  THE SOFTWARE.
//
package swisstopo

import (
	"math"
)

func CHtoWGSheight(y float64, x float64, h float64) float64 {
	// Converts military to civil and to unit = 1000km
	// Auxiliary values (% Bern)
	y_aux := (y - 600000) / 1000000
	x_aux := (x - 200000) / 1000000

	// Process height
	h = (h + 49.55) - (12.60 * y_aux) - (22.64 * x_aux)

	return h
}

// Convert CH y/x to WGS lat
func CHtoWGSlat(y float64, x float64) float64 {
	// Converts military to civil and to unit = 1000km
	// Auxiliary values (% Bern)
	y_aux := (y - 600000) / 1000000
	x_aux := (x - 200000) / 1000000

	// Process lat
	lat := (16.9023892 + (3.238272 * x_aux)) - (0.270978 * math.Pow(y_aux, 2)) - (0.002528 * math.Pow(x_aux, 2)) - (0.0447 * math.Pow(y_aux, 2) * x_aux) - (0.0140 * math.Pow(x_aux, 3))

	// Unit 10000" to 1 " and converts seconds to degrees (dec)
	lat = (lat * 100) / 36

	return lat
}

// Convert CH y/x to WGS long
func CHtoWGSlng(y float64, x float64) float64 {
	// Converts military to civil and to unit = 1000km
	// Auxiliary values (% Bern)
	y_aux := (y - 600000) / 1000000
	x_aux := (x - 200000) / 1000000

	// Process long
	lng := (2.6779094 + (4.728982 * y_aux) + (0.791484 * y_aux * x_aux) + (0.1306 * y_aux * math.Pow(x_aux, 2))) - (0.0436 * math.Pow(y_aux, 3))

	// Unit 10000" to 1 " and converts seconds to degrees (dec)
	lng = (lng * 100) / 36

	return lng
}

// Convert decimal angle (degrees) to sexagesimal angle (seconds)
func DecToSexAngle(dec float64) float64 {
	deg := math.Floor(dec)
	min := math.Floor((dec - deg) * 60)
	sec := (((dec - deg) * 60) - min) * 60

	return sec + min*60.0 + deg*3600.0
}

/**
 * Convert LV03 to WGS84 Return a array of double that contain lat, long,
 * and height
 *
 * @param east
 * @param north
 * @param height
 * @return
 */
func LV03toWGS84(east float64, north float64, height float64) []float64 {

	d := make([]float64, 3)

	d[0] = CHtoWGSlat(east, north)
	d[1] = CHtoWGSlng(east, north)
	d[2] = CHtoWGSheight(east, north, height)
	return d
}

/**
 * Convert WGS84 to LV03 Return an array of double that contain east,
 * north, and height
 *
 * @param latitude
 * @param longitude
 * @param ellHeight
 * @return
 */
func WGS84toLV03(latitude float64, longitude float64,
	ellHeight float64) []float64 {
	// , ref double east, ref double north, ref double height
	d := make([]float64, 3)

	d[0] = WGStoCHy(latitude, longitude)
	d[1] = WGStoCHx(latitude, longitude)
	d[2] = WGStoCHh(latitude, longitude, ellHeight)
	return d
}

// Convert WGS lat/long (° dec) and height to CH h
func WGStoCHh(lat float64, lng float64, h float64) float64 {
	// Converts dec degrees to sex seconds
	lat = DecToSexAngle(lat)
	lng = DecToSexAngle(lng)

	// Auxiliary values (% Bern)
	lat_aux := (lat - 169028.66) / 10000
	lng_aux := (lng - 26782.5) / 10000

	// Process h
	h = (h - 49.55) + (2.73 * lng_aux) + (6.94 * lat_aux)

	return h
}

// Convert WGS lat/long (° dec) to CH x
func WGStoCHx(lat float64, lng float64) float64 {
	// Converts dec degrees to sex seconds
	lat = DecToSexAngle(lat)
	lng = DecToSexAngle(lng)

	// Auxiliary values (% Bern)
	lat_aux := (lat - 169028.66) / 10000
	lng_aux := (lng - 26782.5) / 10000

	// Process X
	x := ((200147.07 + (308807.95 * lat_aux) + (3745.25 * math.Pow(lng_aux, 2)) + (76.63 * math.Pow(lat_aux, 2))) - (194.56 * math.Pow(lng_aux, 2) * lat_aux)) + (119.79 * math.Pow(lat_aux, 3))

	return x
}

// Convert WGS lat/long (° dec) to CH y
func WGStoCHy(lat float64, lng float64) float64 {
	// Converts dec degrees to sex seconds
	lat = DecToSexAngle(lat)
	lng = DecToSexAngle(lng)

	// Auxiliary values (% Bern)
	lat_aux := (lat - 169028.66) / 10000
	lng_aux := (lng - 26782.5) / 10000

	// Process Y
	y := (600072.37 + (211455.93 * lng_aux)) - (10938.51 * lng_aux * lat_aux) - (0.36 * lng_aux * math.Pow(lat_aux, 2)) - (44.54 * math.Pow(lng_aux, 3))

	return y
}
