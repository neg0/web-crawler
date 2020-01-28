# Cuvva Technical Test
* Golang 1.13+


## Environment File
Please create a `.env` file at the root of project from a copy `.env.dist`, this environment file contains the file
name used to set a default TimeOut for crawling the entire website. If this is not provided application uses the 
default value.

    ~$: cp .env.dist .env


## Architecture
I have utilised Effective Go and Uber Go Style Guide to produce this solution. App is created using Domain-Driven Design
(Rich Domain Model) with implementation of SOLID principles. Dependencies, injected into the Core Domain via Ports and Adapters 
Architecture, please read the `README.md` file inside the `adapter` package inside `pkg`.


## Solution
There are two types of crawling strategies:
* __SiteMap Crawler__
* __Explorer Crawler__

if `sitemap.xml` is present at the root of the website it will choose, _SiteMap Crawler_ and when is not it uses explorer to create 
a sitemap and find new links within newly discovered pages.

## How it Works
* __SiteMap Crawler__: This is the fastest and simplest way, it only requests pages listed inside the `sitemap.xml` and generates 
the outcome of crawl.

* __Explorer Crawler__: This is slower strategy compare to SiteMap Crawler, starts from the home page (root of website) `/` to find all 
links that belongs to host and furthers crawls to those pages and if finds new page that hasn't been added to the collection, will visit 
and so on. This is a recursive method to produce find links until there is no more links to find. 


## Tests
I have used TDD approach for creation of most of components, please check the package folder for the relevant tests.


## Run on Local Host Machine
Please ensure you have a Golang 1.3+ installed on your host machine if not you could use the docker implementation in this
project. First ensure you are at the root of this repository.

To run the application please install all dependencies via:

    ~$: go mod download

and then to compile please run:

    ~$: go build cmd/crawler.go

and now you can run the executable crawler the root of the project:

    ~$: ./crawler https://cuvva.com


## Run using Docker & Makefile
Please ensure use following to build and up the Golang 1.3 container:

    ~$: make up

First we need to compile the application, please use:

    ~$: make build

and then you could run the application inside the Golang container, please ensure you are passing target website 
as a `url` argument; as demonstrated below:

    ~$: make run url=https://cuvva.com

In order to run the tests and lints inside the Golang container please run the following:

    ~$: make test

_I have n't written many tests due to time limit but there are few tests in `pkg`_
rest of the make commands created solely for purpose of development such as: `make ssh`


### License
Creative Commons Attribution-NonCommercial-NoDerivs *CC BY-NC-ND*

<img src='https://licensebuttons.net/l/by-nc-nd/3.0/88x31.png' alt='CC BY-NC-ND'>