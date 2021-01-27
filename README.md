# PoC - Pub/Sub GCP in Golang

You have to create service account JSON file then export as Environment variable where your shell running.

```bash
    make get-credential PATH="[PATH]"
```

### Example

I'm create service account JSON file and name it `POC_SRVACC.json` then save as current directory then comand with

```bash
    make get-credential PATH="~/path-to-repo/POC_SRVACC.json"
```
