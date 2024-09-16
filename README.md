# Go Email Server

## Overview

A simple email server implemented in Go that uses Gmail's SMTP server to send emails. This project is dockerized for easy deployment and includes CORS support.

## Features

- Send emails using Gmail's SMTP server.
- CORS support for cross-origin requests.
- Error handling for various scenarios.
- Dockerized for easy deployment.

## Setup

### Environment Variables

Create a `.env` file in the project root with the following content:

```plaintext
GMAIL_USER=your-email@gmail.com
GMAIL_PASSWORD=your-email-password
```
