[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]


<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">AMC Publicly Available Callsign Scraper</h3>

  <p align="center">
    Scrape publicly available call signs from the AMC database into a txt file.
    <br />
    <a href="https://github.com/ruxton/call_scraper/issues">amc_callscraperrt Bug</a>
    ·
    <a href="https://github.com/ruxton/call_scraper/issues">Request Feature</a>
  </p>
</p>



<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
  * [Built With](#built-with)
* [Getting Started](#getting-started)
  * [Prerequisites](#prerequisites)
  * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## About The Project

A command line applications that scrapes available call signs from the AMC's publicly available callsign database


### Built With

* [GoLang v1.15](golang-url)
* [Collly](colly-url)



<!-- GETTING STARTED -->
## Getting Started

[Download a release](release-url) and run it.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.
* npm
```sh
npm install npm@latest -g
```

### Installation

1. Unzip or install the Package
2. Run the binary file



<!-- USAGE EXAMPLES -->
## Usage

Be very aware, running the scraper with no options will scape _ALL_ the available callsigns.

```
-a, --about                 Show the about message
-c, --callsigntype string   Advanced,Standard,Foundation,Beacon or Repeater
-h, --help                  Show the help menu
-l, --letters string        One,Two,Three
-o, --outputfile string     The filename to write the results to (default "output.txt")
-r, --region string         ACT,NSW,NT,QLD,SA,TAS,VIC,WA,AET,ANTARCTIC
```


<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request



<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.



<!-- CONTACT -->
## Contact

Your Name - [@ruxton](https://twitter.com/ruxton)

Project Link: [https://github.com/ruxton/call_scraper](https://github.com/ruxton/call_scraper)



<!-- ACKNOWLEDGEMENTS -->
## Acknowledgements

* [AMC Publicly Available Callsigns Database](amcdb-url)
* [Ham College (for teaching me)](hamcollege-url)




<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/ruxton/amc_callscraper.svg?style=flat-square
[contributors-url]: https://github.com/ruxton/amc_callscraper/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/ruxton/amc_callscraper.svg?style=flat-square
[forks-url]: https://github.com/ruxton/amc_callscraper/network/members
[stars-shield]: https://img.shields.io/github/stars/ruxton/amc_callscraper.svg?style=flat-square
[stars-url]: https://github.com/ruxton/amc_callscraper/stargazers
[issues-shield]: https://img.shields.io/github/issues/ruxton/amc_callscraper.svg?style=flat-square
[issues-url]: https://github.com/ruxton/amc_callscraper/issues
[license-shield]: https://img.shields.io/github/license/ruxton/amc_callscraper.svg?style=flat-square
[license-url]: https://github.com/ruxton/amc_callscraper/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=flat-square&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/ruxtonau
[hamcollege-url]: https://www.hamcollege.org.au/
[amcdb-url]: https://www.amc.edu.au/industry/amateur-radio/callsigns/publicly-available-callsigns
[golang-url]: https://golang.org/
[release-url]: https://github.com/ruxton/amc_callscraper/releases/
[colly-url]: https://github.com/gocolly/colly/
