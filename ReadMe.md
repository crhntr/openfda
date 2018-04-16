# openfda
This code is not officially affiliated with the openFDA project.

Right now I am using it for a statistics project and will probably turn it into
a package for studying the various open fda data sets.

Currently the Go structs are not fully compliant with the specification; however,
they have been tested for data/event and work with the majority of records. Not
all special edge cases have been handled.

## Project Navigation
* ./drug - This directory has a go package for manipulating and compressing drug event and label JSON in Go
* ./scripts - has some mongo shell scripts and an R script for data analysis.
* ./www - has a tool for scraping data from the deployed api.fda.org REST api.
* ./ - the root of the package contains standard files (ReadMe.md, .gitignore...) and a compilable cli tool downloading, collecting meta data about, and inserting drug events

## Downloading Data
Note that the data set is huge it is greater than 250 gigs of compressed <filename>.json.zip files for the drug event data alone. Plan accordingly.
I used the `unzip` unix tool to decompress the json files.
The json parsing has been somewhat optimized to deal with the huge data (otherwise it would not run on a macbook air with 8gigs of memory).

![An Example](https://raw.githubusercontent.com/crhntr/openfda/branch/assets/exampleImage.png)
