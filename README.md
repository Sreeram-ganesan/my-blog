## About

API Server application provides very basic functionality to work with Address Book contacts.
You can create new contact with phone numbers, update existing contact, fetch contact by id,
list all contacts, and delete existing contact.


## How to build application

Open the directory with newly created project and run:

```shell 
go build -o my-blog main.go
```
it will result in building executable file "my-blog" (feel free to name it differently).


## How to run application

**IMPORTANT:** depending on your configuration, e.g. if you have added database support
etc, starting of your code may fail because you need to complete configuration settings (e.g.
your database URL and credentials). So in this case keep reading README past this section.

To perform an initial launch an application, run this in shell:

```shell
./my-blog run --deployment=local
```

We launch API server by specifying `run` command. `--deployment=local` tells our code to
perform a local deployment. Local deployment settings can contain options such as URL of
your local database or other local specific environment. `--deployment` flag loads
file `local.yaml` from `config/` directory that resides in the same directory where
your executable file is. You can create a copy of this file and name it,
for example `prod.yaml` where you can add production-specific settings, then running

```shell
./my-blog run --deployment=prod
```

will load this production settings for your API server.


## How to override configuration values

Sometimes editing configuration file to add values is not the best strategy. As an example,
if you have database settings in your `prod.yaml` file, having URL of database specified
there is not a bad idea, but storing a password there - is not good. The better approach
would be to pass sensitive settings via environment variables. And because we use
viper library to load yaml configuration file, it allows us to override values specified
in it with something different. The typical syntax of environment variable is:
`[EnvPrefix]_[YamlConfigKey] = value`. `EnvPrefix` is something you have previously entered when
generated project with GoQuick.

**internal/core/app/configload.go:**
```
const envVarPrefix = "MYBLOG"
```

Let's give it a try. Let's say you want to change a listening port for your API server.
If you open `local.yaml` you can find something like this:

```
server:
  port: 8080
```

`server/port` translates to SERVER_PORT and combined with environment prefix MYBLOG
you can override it as:

```shell
export MYBLOG_SERVER_PORT=9090
```
```shell
./MYBLOG run --deployment=local
```

now API server code will be listing port 9090 instead of 8080.

## Database configuration

This project was generated with PostgreSQL database support. You can find database-specific
configuration in `./config/[deployment].yaml` file, e.g. if you launch application
with `run --deployment=local` flag then configuration file will be `./config/local.yaml`.
Let's take a look inside (for database-specific setting):

```yaml
database:
  host: 127.0.0.1
  port: 5432
  name: myblog
  user: postgres
  password: admin
  sslmode: disable
  # Connect timeout is seconds
  connect_timeout: 10
```

What you see here is typical PostgreSQL configurations. Some notable exceptions are username and
password, because in many cases you don't want to store your db's username and password in file
that you are most likely going to check into some sort of version control system (well, if you
are not planning to do so, feel free to enter your db credentials into this file). Instead, we
decided to set values to some meaningless strings, in our case it is just a single underscore
character. As we already explained previously, you can easily override this value via environment
variables, such as:

```shell
export MYBLOG_DATABASE_USER=postgres
export MYBLOG_DATABASE_PASSWORD=admin
```

or pass it via settings if you run your code in IDE. For instance, in IntelliJ IDEA you can
open **Run** / **Edit Configuration** and for your launch configuration select **Environment** text box and
enter the variables from above separated by semicolumn (without *export* command).

### Tip: How to install PostgreSQL locally

If you have a local PostgreSQL running on your computer - you are all set.
If you don't have PostgreSQL installed, you can do it now by installing
from: https://www.postgresql.org/download/ and set it up with username and
password that can later be used to configure.

Another option is to create a docker image with PostgreSQL database running:

```shell
docker run -d -p 5432:5432 -e POSTGRES_DB=myblog -e POSTGRES_PASSWORD=admin postgres:latest
```

This will create and run a local instance of PostgreSQL database with user name "postgres" and password "postgrespass".

Or use

```shell
docker start localpostgres
```

if localpostgres container was already created.

Make sure to create **addrbook** database in your PostgreSQL instance.

## Cache: REDIS configuration

Code will be generated with Redis support for caching data. By default, `configs/local.yaml` configuration
file contains address and port for local instance. You can test generated app by launching local redis
instance:

```shell
  docker run --name some-redis -p 6379:6379 -d redis
```

Now when you run API server and try Address Book API requests, it will be using redis instance to set,
update, delete and get contacts.

There is another switch that in yaml configuration that allows you to configure cache. By default it is:

```yaml
cache:
  type: redis
```

but you can change type to `inmem` to enable in-memory cache only instead of Redis or you can completely
turn cache off by providing `none` value.

## Access REST API

Generated application uses REST protocol to store and fetch address book records.
Once you have the application launched, you can perform HTTP calls to test REST APIs
exposed by API server.

Please note that each HTTP response contains **X-Request-Id** header with value that
is displayed with application logs (as **requestId** field). It helps you to troubleshoot
application, because logger provided with generated code prints request id with
every log line.

### Examples of REST requests

#### Get service version 

Request:
```shell
curl --location 'http://localhost:8080/api/version'
```
Response (could be slightly different):
```
{
  "service": "rest-net/http",
  "version": "0.1.0",
  "build": "1"
}
```

#### Add new contact

Request:
```shell
curl --location 'http://localhost:8080/api/contacts' \
--data '{
    "first_name": "Joe",
    "last_name": "Doe",
    "phones": [
        {
            "phone_type": "mobile",
            "phone_number": "+1-503-777-0001"
        },
        {
            "phone_type": "home",
            "phone_number": "+1-503-777-9999"
        }
    ]
}
'
```
Response:
```
{
    "first_name": "Joe",
    "last_name": "Doe",
    "phones": [
        {
            "phone_type": "mobile",
            "phone_number": "+1-503-777-0001"
        },
        {
            "phone_type": "home",
            "phone_number": "+1-503-777-9999"
        }
    ]
}
```

#### Get existing contact

Request:
```shell
curl --location 'http://localhost:8080/api/contacts/36'
```
Response (truncated):
```
{
    "id": "36",
    "first_name": "Joe",
    "last_name": "Doe",
    ...
}
```

#### Attempt to get non-existing contact

Request:
```shell
curl --location 'http://localhost:8080/api/contacts/9999'
```
Error response:
```
{
  "status": "Internal server error",
  "error": "contact id=9999 not found"
}
```

#### Get all existing contacts

Request:
```shell
curl --location 'http://localhost:8080/api/contacts'
```
Response (truncated):
```
[
  {
    "id": "36",
    "first_name": "Joe",
    "last_name": "Doe",
    ...
  },
  ...
]
```

#### Delete existing contact

Request:
```shell
curl --location --request DELETE 'http://localhost:8080/api/contacts/36'
```
No response payload is received

### Logging

Each HTTP request returns `X-Request-Id` header as part of response. This `X-Request-Id`
is always unique, unless you specify it explicitly as part of request. What makes it useful
is that each application log line contains `{requestId="...."}` tag, and it matches
`X-Request-Id` value. It makes debugging code much easier because you can filter logs
scoped to specific request.

## Build React/Typescript application with Vite

This generated project uses React with Typescript for front-end application served
by your Go project. Instead of webpack for app packaging it uses Vite (https://vitejs.dev).

### Pre-requisites

To start serving your front-end app, first you need to build it. This code requires typescript
compile so try to run `tsc -v` to see if it is installed already, if not, you might want to install it

```shell
npm install -g typescript
```

We use `yarn` for this tutorial, so make sure to install it first, if you don't
have it yet, e.g.

```shell
npm install --global yarn
```

### Run web packaging

Next step is to go inside a `web` directory in your newly generated project and install
dependencies:

```shell
cd web
yarn install
```

once dependencies are installed, you are ready to go.

Run this in your `web` directory:

```shell
yarn run dev 
```

This command builds your front-end web app and launches file change listener. So every
time when you change something in `web` directory, web app will be automatically repackaged.
Keep this command running during your development.

### Launch API server with Web application

Your Go code is already setup to run API server and to serve SPA web
application (HTML/Typescript content). Actually you can find this code in your
app's `internal/adapters/apiserver/apiserver.go`, there are lines that expose local
web app files to HTTP server.

Now in your browser go to http://localhost:8080/ and it should open your React web
application.


To Run postgres locally on 5432:
```shell
  docker run -d --name my-postgres -e POSTGRES_PASSWORD=admin -p 5432:5432 postgres
```
To Run redis locally on 
To Run this app as a docker image locally exposed in 8080:

```shell
  docker build -t my-blog . && docker run -p 8080:8080 -d my-blog
```