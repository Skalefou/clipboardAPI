# Clipboard API

An API that acts as a clipboard, allowing you to copy/paste text and retrieve it from different devices.
It's made in Go, using the Gin, GoDotEnv, Cron and Gorm libraries. You must use a postgres database

## Prerequisites

Prerequisites : Goland (1.20.4), Postgresql

### Dependencies

| Dependency                 | Version | 	Description                                 | 
|----------------------------|---------|----------------------------------------------|
| github.com/gin-gonic/gin   | v1.9.1	 | Web framework for Go                         | 
| github.com/joho/godotenv	  | v1.5.1  | Go library for loading environment variables | 
| gorm.io/driver/postgres	   | v1.5.2  | Postgres driver for GORM, a Go ORM           |
| gorm.io/gorm               | v1.25.2 | 	Go ORM with support for multiple databases  | 

## Installation

> Clone the repo
```shell
git clone https://github.com/Skalefou/clipboardAPI.git
```
> Edit config.env file

| Variable Name                       | Value Type | Description                                                                                                               |
|-------------------------------------|------------|---------------------------------------------------------------------------------------------------------------------------|
| DB_HOST                             | string     | The hostname of the database server.                                                                                      |
| DB_NAME                             | string     | The name of the database to connect to.                                                                                   |
| DB_USERNAME                         | string     | The username to use when connecting to the database.                                                                      |
| DB_PASSWORD                         | string     | The password to use when connecting to the database.                                                                      |
| DB_PORT                             | int        | Database port                                                                                                             |
| ANTI_DDOS_TIME                      | int        | Set in milliseconds the length of time that the same IP address may not make a new request,set to 0 to disable anti DDOS. |
| CLIP_DELETION_TIME_WITHOUT_PASSWORD | int        | Time after which clipboards without password are deleted, set to 0 to disable.                                            |
| CLIP_DELETION_TIME_WITH_PASSWORD    | int        | Time after which clipboards with password are deleted, set to 0 to disable.                                               |
| LOG_DELETION_TIME                   | int        | Time after which logs are deleted.                                                                                        |

Create a database on your postgresql, assign user rights and enter the database information in the "config.env" file. Then execute this SQL file in your database.

```sql
DROP TABLE IF EXISTS clipboard;
DROP TABLE IF EXISTS log;

CREATE TABLE clipboard (
    id INT PRIMARY KEY,
    message TEXT,
    password VARCHAR(72),
    ip_owner VARCHAR(15) NOT NULL,
    creation_date TIMESTAMP NOT NULL,
    last_update TIMESTAMP NOT NULL,
    last_see TIMESTAMP NOT NULL
);

CREATE TABLE log (
    id SERIAL PRIMARY KEY,
    type_request INT NOT NULL,
    ip_user VARCHAR(21),
    clipboard INT NOT NULL,
    date_request TIMESTAMP NOT NULL,
    active BOOLEAN DEFAULT FALSE
);
```

## Usage

>Run the API
```shell
go run main.go
```

To create a clipboard, you need to make a POST request.
>/createPort

You can also send a json file containing your password, as shown below. You need to replace the three dots with a bcrypt hash
```json
{
  "password": "..."
}
```

this will return a JSON clipboard with a maximum ID of 6 digits. You can obtain the contents of the page using this GET request (you must enter the id and password).
>/port

To update clipboard content, you can make this POST request. You'll need to enter at least the id and password if you've set one.
>/port

Finally, to delete a clipboard, you need to specify the ID and password if you have one, and make a DELETE request.
>/port

