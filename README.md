# 🐊 Gator CLI

A simple CLI tool to manage feeds and aggregate posts using PostgreSQL.  

---

## 📦 Requirements
- **PostgreSQL**: `v17.5`
- **Go**: `v1.24.3`

---

## ⚡️ Installation

```bash
go install github.com/mauricebonnesdev/gator@latest
```
After installation, the CLI is available as:
```bash
gator
```
## ⚙️ Configuration
Gator uses a config file located in your home directory:
`~/.gatorconfig.json`
Example:
```json
{
  "db_url": "postgres://<username>:@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```
- db_url -> Connection string to your Postgres instance
- current_user_name -> Set automatically by the `login` command

## 🛠️ Commands
### 🔑 Authentication
- Login
  ```bash
  gator login <username>
  ```
  Set the current user in your config.
- Regiser
  ```bash
  gator register <username>
  ```
  Create a new user in the system.
### 🗄️ Database
- Reset
  ```bash
  gator reset
  ```
  Deletes all data and resets the database.
- Users
  ```bash
  gator users
  ```
  Lists all registered users.
### 📡 Feeds
- Add Feed
  ```bash
  gator addfeed <name> <url>
  ```
  Adds a new feed to the system
- Feeds
  ```bash
  gator feeds
  ```
  Lists all the registered feeds.
### 👤 Following
- Follow
  ```bash
  gator follow <feed_url>
  ```
  Follow a feed by its URL.
- Unfollow
  ```bash
  gator unfollow <feed_url>
  ```
  Stop following a feed.
- Following
  ```bash
  gator following
  ```
  Shows all feeds that the current user follows.
### 🪪 Posts
- Browse
  ```bash
  gator browse [limit]
  ```
  Displays aggregated posts for the logged-in user. Defaults to 2 if `limit` is not provided.
- Aggregator (long-running)
  ```bash
  gator agg <interval>
  ```
  Starts a background process that continuously fetches posts from followed feeds. The `interval` takes the form of `5m` for every 5 minutes, `1h` for every hour etc.

### 🚀 Example Workflow
```bash
# Register and login
gator register alice
gator login alice

# Add a feed and follow it
gator addfeed "Boot.dev" https://blog.boot.dev/index.xml
gator follow https://blog.boot.dev/index.xml

# See feeds and followed feeds
gator feeds
gator following

# Start aggregation
gator agg

# Browse posts
gator browse

```
