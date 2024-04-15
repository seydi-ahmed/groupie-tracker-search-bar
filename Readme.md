# Groupie Trackers Geolocalization

Groupie Trackers consists of receiving a given API and manipulating the data contained in it to create a site that displays the information.

## API Description

The API consists of four parts:

1. **Artists:** Contains information about bands and artists, including their name(s), image, the year they began their activity, the date of their first album, and the members.

2. **Locations:** Contains the last and/or upcoming concert locations for the artists.

3. **Dates:** Contains the last and/or upcoming concert dates for the artists.

4. **Relation:** Links all the other parts (artists, dates, and locations) together.

## Project Objective

To build a user-friendly website that displays bands' information through various data visualizations (e.g., blocks, cards, tables, lists, pages, graphics, etc.). The choice of how to display the information is up to you.

The project also focuses on creating and visualizing events/actions. The main event/action we want you to implement is a client call to the server (client-server). This feature should trigger an action that communicates with the server to receive information ([request-response](https://en.wikipedia.org/wiki/Request%E2%80%93response)). An event consists of a system that responds to some kind of action triggered by the client, time, or any other factor.

## Instructions

- The backend must be written in Go.
- The site and server cannot crash at any time.
- All pages must work correctly, and you must handle any errors gracefully.
- The code must adhere to good programming practices.
- It is recommended to have test files for unit testing.

## Allowed Packages

Only the standard Go packages are allowed.

## Usage

You can see an example of a RESTful API [here](example-api-link).

## Developers

This project is developed by:

- Mouhamed Diouf (Git: mouhameddiouf)
- Atoumane Der (Git: ader)

## Learning Objectives

By working on this project, you will learn about:

- Manipulation and storage of data.
- JSON files and format.
- HTML.
- Event creation and display.
- Client-server communication.

## structure of my project

GROUPIE-TRACKER
|-src
|---artist.go
|---dates.go
|---handle.go
|---locations.go
|---otherApi.go
|---struct.go
|-static
|---calendrier.png
|---image.jpg
|---location.png
|---style_artist.css
|---style_dates.css
|---style.css
|-templates
|---artists.html
|---dates.html
|---error.html
|---index.html
|---locations.html
|-go.mod
|-main.go
|-Readme.md