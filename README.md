# bitcoin_checker_api

## About The Project:
Project can be use for checking exchange rates form one to another currency.
Evaluable endpoints:
 - `/api/rate` - get rate
 - `/api/subscribe` - subscribe user on emails
 - `/api/sendEmails` - send emails to our subscribers

## How to run:
For run application use make command:

On locale machine
```azure
make run
```
On docker
```azure
make docker
```
## Installation:
In _env/example.toml add information about:

Enter path for internal storage:
```
[storage]
path = "./storage.json"
```

Enter your endpoint for cheking rates:
```
[exchangerate]
url_mask = "URL_WITH_MASK"
in_rate_name = "IN_RATE_NAME"
out_rate_name = "OUT_RATE_NAME"
```
Enter your SendGrid credential
```
[emailservice]
api_key="API_KEY"
from_name="FROM_NAME"
from_address="FROM_ADDRESS"
```

## Roadmap:
- [X] Refactor old project structure and add layer for app
- [ ] Add tests
- [ ] Review code and try implement SOLID and GRASP