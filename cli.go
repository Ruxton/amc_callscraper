package main

import (
	"errors"
	"fmt"
	"github.com/ruxton/amc_callscraper/versions"
	"github.com/ruxton/term"
	flag "github.com/spf13/pflag"
	"os"
)

// <select class="form-control form-control-sm" id="filterRegionId" name="filterRegionId"><option value=""></option>
// <option value="3">Australian Capital Territory</option>
// <option value="4">New South Wales</option>
// <option value="5">Northern Territory</option>
// <option value="6">Queensland</option>
// <option value="7">South Australia</option>
// <option value="8">Tasmania</option>
// <option value="9">Victoria</option>
// <option value="10">Western Australia</option>
// <option value="2">Antarctic</option>
// <option value="11">Australian External Territories</option>
// </select>

var aboutVar bool
var helpVar bool
var callSignType string
var letterTypeCode string
var regionCode string
var filename string

var regionMap = map[string]string{
	"ACT":       "3",
	"NSW":       "4",
	"NT":        "5",
	"QLD":       "6",
	"SA":        "7",
	"TAS":       "8",
	"VIC":       "9",
	"WA":        "10",
	"ANTARCTIC": "2",
	"AET":       "11",
}

var callSignMap = map[string]string{
	"Advanced":   "3",
	"Standard":   "6",
	"Foundation": "7",
	"Beacon":     "4",
	"Repeater":   "5",
}

var letterTypes = []string{
	"One",
	"Two",
	"Three",
}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func CliIntro() {
	setupFlags(&aboutVar, &helpVar, &callSignType, &letterTypeCode, &regionCode, &filename)
	flag.Parse()

	if aboutVar {
		showWelcomeMessage()
		showAboutMessage()
		os.Exit(exitOK)
	}

	if helpVar {
		showWelcomeMessage()
		flag.PrintDefaults()
		os.Exit(exitOK)
	}
}

func CommandSwitches() (string, string, string, string, error) {
	returnCallSignType := ""
	returnLetterType := ""
	returnRegion := ""

	if callSignType != "" {
		if item, ok := callSignMap[callSignType]; ok {
			returnCallSignType = item
		} else {
			return "", "", "", "", errors.New("Callsign type can only be Advanced,Standard,Foundation,Beacon or Repeater")
		}
	}

	if letterTypeCode != "" {
		if isValueInList(letterTypeCode, letterTypes) {
			returnLetterType = letterTypeCode
		} else {
			return "", "", "", "", errors.New("Letters can only be One, Two or Three")
		}
	}

	if regionCode != "" {
		if item, ok := regionMap[regionCode]; ok {
			returnRegion = item
		} else {
			return "", "", "", "", errors.New("Regions can only be one of ACT,NSW,NT,QLD,SA,TAS,VIC,WA,AET,ANTARCTIC")
		}
	}

	return returnCallSignType, returnLetterType, returnRegion, filename, nil
}

func setupFlags(aboutVar *bool, helpVar *bool, callSignType *string, letterTypeCode *string, regionCode *string, filename *string) {
	flag.BoolVarP(aboutVar, "about", "a", false, "Show the about message")
	flag.BoolVarP(helpVar, "help", "h", false, "Show the help menu")
	flag.StringVarP(callSignType, "callsigntype", "c", "", "Advanced,Standard,Foundation,Beacon or Repeater")
	flag.StringVarP(letterTypeCode, "letters", "l", "", "One,Two,Three")
	flag.StringVarP(regionCode, "region", "r", "", "ACT,NSW,NT,QLD,SA,TAS,VIC,WA,AET,ANTARCTIC")
	flag.StringVarP(filename, "outputfile", "o", "output.txt", "The filename to write the results to")
}

func showWelcomeMessage() {
	term.OutputMessage(term.Green + "AMC Publicly Available Callsigns v" + versions.VERSION + term.Reset + "\n\n")
}

func showAboutMessage() {
	term.OutputMessage(fmt.Sprintf("Build Number: %s\n", versions.MINVERSION))
	term.OutputMessage("Created by: Greg Tangey\n")
}
