# Go1PConnector

This project provides a Go client for interacting with 1Password via the 1Password Connect API. It allows fetching vaults and other secure operations directly from your Go application.

## Prerequisites

Before you can run this project, you will need to have Go installed on your machine. [Download and install Go](https://golang.org/dl/).

## Setting Up

To use this project, you need to set up the necessary environment variables for 1Password Connect. These variables will allow the application to authenticate and interact with the 1Password API.

### Environment Variables

Set the following environment variables before running the application:

- `OP_CONNECT_HOST`: The URL to your 1Password Connect API.
- `OP_CONNECT_TOKEN`: Your 1Password Connect API token.

You can set these variables in your shell like so:

```bash
export OP_CONNECT_HOST='https://your-1password-connect-host.com'
export OP_CONNECT_TOKEN='your-1password-api-token'
```

Make sure to replace `https://your-1password-connect-host.com` and `your-1password-api-token` with your actual 1Password Connect host and API token.

## Running the Application

Navigate to the root directory of the project where the `main.go` file is located inside the `cmd` folder. Run the application using the following command:

```bash
go run cmd/main.go
```

This command will start the application, which will then interact with the 1Password Connect API using the credentials provided via the environment variables.

## Functionality

The application currently supports fetching all available vaults from 1Password and displaying them. Future functionalities can include managing items within the vaults, adding new items, and other API interactions as needed.

