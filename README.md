# Case Insensitive Emails

An [xpb](https://github.com/pocketbuilds/xpb) plugin for [Pocketbase](https://pocketbase.io/) that automatically lowercases email addresses on auth records.

Email addresses are automatically lowercased on auth record create and auth record update.

Email addresses are also automatically lowercased when someone attempts to log in with email and password.

## Installation

1. [Install XPB](https://docs.pocketbuilds.com/installing-xpb).
2. [Use the builder](https://docs.pocketbuilds.com/using-the-builder):

```sh
xpb build --with github.com/pocketbuilds/case_insensitive_emails@latest
```
