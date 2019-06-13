# Pairot

Pairot is a slack bot api that returns pairs for a provided team

## Setup

Update your docker-compose accordingly so that the all the environment variables match the database configuration files.

After the DB is create do not forget to run `docker-compose up -d --build --force-recreate mongo-seed` in order to seed the database and include the desired username and password.

## Installation

Follow the slackbot slash commands tutorial to add a new `/pair` command.
Add the endpoint where the bot is deployed and voila.

## Usage

```bash
curl -X POST \
  http://localhost:8080/v1/api/pair \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  -H 'Host: localhost:8080' \
  -H 'Postman-Token: 65e2c320-f459-45b7-9c1d-5705e5688129,dcc34069-989a-40aa-994b-237379b18f57' \
  -H 'User-Agent: PostmanRuntime/7.15.0' \
  -H 'accept-encoding: gzip, deflate' \
  -H 'cache-control: no-cache' \
  -H 'content-length: 156' \
  -d 'token=gIkuvaNzQIHg97ATvDxqgjtO&team_id=T0001&team_domain=example&enterprise_id=E0001&channel_name=team-balboa&user_id=U2147483697&user_name=Steve&text=94070'
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
