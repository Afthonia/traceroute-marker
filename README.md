# Traceroute Marker Docs

## Operating System: Ubuntu 22.04.3 LTS (Jammy Jellyfish)  --Linux
`Host:` localhost or 127.0.0.1

`Port:` 80 (http)

`Backend: ` Golang

`Front End: ` Javascript - HTML - CSS

After cloning this repository, an API key must be provided in order to perform the traceroute marking system.

[developer.here.com](https://)

From this website, the developer needs to register and obtain an API key by selecting one of the membership packages. The free version is adequate for this project.

The API key should be pasted instead of "API-KEY" on line 113 in the index.html file.

To manage requests, pages, etc., the 'chi' library is used. It can be installed as follows, although it is likely already installed:

`go get -u github.com/go-chi/chi/v5`

Since there is no additional program or database is required for this project yet, and the necessary libraries is already included, the only thing to do is run the program.

If there is any error related to import functionality, the command below must be typed once or more times to fetch the missing modules:

`go mod tidy`

In Go, to run a program consisting of multiple files, all the necessary files with the .go extension must be compiled together. It is conventional to use a Makefile to reduce typing time and the probability of mistakes.

The program can be run with the following command:

##### Linux:
`sudo make run`

##### MacOS / Windows:
`make run`


The port is ':80', and since ports below 1024 require authority, 'sudo' must be used in this case for Linux users.

The above command will run the server and that can be seen when 'Server Started!' text is printed out. The traceroute hops and the response json list will be printed out on the server console as the user sends requests and the server can be closed with Ctrl+C shortcut.

#####WARNING: Do not use Ctrl+Z to exit because although it can seem like the server closed, unless the terminal window is closed and a new one is opened, it is not possible to run the server again.

After running the program, simply type '127.0.0.1', '127.0.0.1:80', 'localhost', or 'localhost:80' in the browser to view the website.

To mark the locations, the browser may ask for permission to access geolocation and mark the client's location.

When the host address is sent, it may take a little while to fetch the latitude and longitude of the route IP from the external API.

PS: Most of the code blocks that are wrapped as comments will be used to improve the efficiency and responsiveness of the website in the future and are not completed.

PS2: There are small issues about drawing the routes between markers and they will be fixed or the API for maps will be changed.
