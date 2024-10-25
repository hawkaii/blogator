
# Blogator
Too fast and simple blog generator.

## Features
- Catered RSS feed for your blogs
- Add, Browse, Edit and Delete blogs
- Follow and Unfollow blogs

## Installation
```bash
go install github.com/hawkaii/blogator
```

## Setup
### Dependencies
- [Go](https://golang.org/)
- [PostgreSQL](https://www.postgresql.org/)

### Database configuration
Create .gatorconfig file in your home directory
```bash
touch ~/.gatorconfig.json
```
Add the following to the file
```json
{
    "db_url": <db_url>,
}
```

## Usage
```bash
gator help
```

## Commands
gator has the following commands:
- `login` - Login to your account
- `register` - Register a new account
- `users` - List all users
- `agg` - Aggregate all blogs



