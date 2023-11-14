# Traceroute Marker Docs

## Operating System: Ubuntu 22.04.3 LTS (Jammy Jellyfish)  --Linux
`Host:` localhost or 127.0.0.1
`Port:` 80 (http)

`Backend: ` Golang
`Front End: ` Javascript - HTML - CSS

After cloning this repository, an API key must be provided in order to perform the traceroute marking system.

[developer.here.com](https://)

From this web site, the developer needs to register and get an API key by selecting one of the membership packages (Free version is adequate for this project).

API key should be pasted instead of API-KEY on line 113.

In order to manage requests, pages etc. 'chi' is used. It can be installed as following:

`go get -u github.com/go-chi/chi/v5`

Since there is no additional program or database is required for this project yet, and the necessary libraries is already included, the only thing to do is run the program.

If there is any error related to import functionality, the command below must be typed once or more times to fetch the missing packages or libraries:

`go mod tidy`

In go, to run the program, all the necessary files with .go extention must be run as well so it is conventional to use Makefile to reduce the typing time and mistake probability.

The program can be run with the following command:

###### Linux
`sudo make run`

###### MacOS / Windows
`make run`

The port is ':80' and since the ports that are under 1024 is required authority, 'sudo' must be used in this case.

After running the program, simply by typing '127.0.0.1', '127.0.0.1:80', 'localhost', or 'localhost:80' in the browser, the web site can be seen.

To mark the locations, the browser may ask for permission for geolocation access to find and mark the client's location.

When the host address is sent, it may take a little while to fetch the latitude and longitude of the route IP's from the external API.