
# HTTP Status Checker and Notifier

This GoLang program checks the HTTP status code of a given URL. If the status code is not 200, it sends an email notification to the specified recipient. This can be useful for server administrators to be alerted when their server is experiencing issues.

# Features
- Checks the HTTP status code of a specified URL.
- Sends an email notification if the status code is not 200.
- Easy configuration for sender email, recipient email, and SMTP server.


# Getting Started
## Prerequisites

 - GoLang installed on your system. If not, you can download it from [GoLang official website](https://go.dev/dl/).
## Installation

Clone the repository:

```bash
git clone https://github.com/sabarishOfficial/GoLang-HTTP-Status-Checker-Notifier.git
cd GoLang-HTTP-Status-Checker-Notifier
go run main.go
```

## Usage
Run the program with the target URL as a command-line argument:
```bash
sabarish@Sabarishs-MacBook-Air GoLang-HTTP-Status-Checker-Notifier % go run main.go https://google.com
Http Status Code: 200
```
If the URL you provide is invalid and does not equal 200, an error code will appear and an email will be sent to the administrator or the recipient.

```bash
sabarish@Sabarishs-MacBook-Air GoLang-HTTP-Status-Checker-Notifier % go run main.go http://sabarish.shop
somthing went wrong: 404
```

Replace <url> with the URL you want to check. 



## Configuration

Before running the program, make sure to configure the sender's email address, password, recipient's email address, SMTP server hostname, and port in the main.go file.

```bash
from := "sender@example.com"
password := "sender_password"
toList := "recipient@example.com"
hostname := "smtp.gmail.com"
port := "587"
```

