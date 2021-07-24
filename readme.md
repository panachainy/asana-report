# Asana report

[![Coverage Status](https://coveralls.io/repos/github/panachainy/asana-report/badge.svg?branch=main)](https://coveralls.io/github/panachainy/asana-report?branch=main)

ASAR make for easy to get progress of all your task in each project under workspace, In my use case I need to track all task each project in one command.

## Prerequisites

- Golang

## Usage [Mac]

1. Clone this repository

2. Build

    ```sh
    make build
    ```

3. Install

    ```sh
    make install
    ```

4. Test

    ```sh
    make try
    ```

### Set alias

1. Run this command for set alias to your command

    ```sh
    echo "alias asar='/Users/panachainy/bin/asar'" >> ~/.bash_profile
    ```

    > you can change `~/.bash_profile` follow your command.

2. Restart your command.

3. Try `asar`

### Command

| Command   | Description                                       |
| --------- | ------------------------------------------------- |
| asar ast  | Get task status                                   |
| asar asaa | Assign all task in assignee with your assigneeId. |

### ENV

Must set prefix when you set config with env.

| ENV key           | Description                                    | Example                    | Remark                                                                                       |
| ----------------- | ---------------------------------------------- | -------------------------- | -------------------------------------------------------------------------------------------- |
| ASAR_WORKSPACE_ID | Asana workspace id for scope your task         | 0000000000000000           | [API check workspaceId from your account](https://app.asana.com/api/1.0/workspaces)          |
| ASAR_TOKEN        | Asana token for access your task               | "0/0000000000000000:xxxxx" | [Create your personal token](https://app.asana.com/0/developer-console)                      |
| ASAR_ASSIGNEE_ID  | AssigneeId for assign all your task to that ID | 0000000000000000           | You can get from ASAR_TOKEN in `"0/0000000000000000:xxxxx"` at `0000000000000000` is your id |

#### Example

```sh
export ASAR_WORKSPACE_ID=0000000000000000
export ASAR_TOKEN="0/0000000000000000:xxxxx"
# For feature asaa
export ASAR_ASSIGNEE_ID="0000000000000000"
```

## Getting stared

### Development

Run with development

```sh
go run main.go
```

Build

```sh
make build
```

Install

```sh
make install
```

Try Run

```sh
make try
```

## To do

- Brew install

## REF

### External

- [Blog create golang cli](https://sbstjn.com/blog/create-golang-cli-application-with-cobra-and-goxc/)
- [Hello-Cobra](https://github.com/KEINOS/Hello-Cobra)
- [Issue-188](https://github.com/spf13/viper/issues/188#issuecomment-399884438)
- [Http to go](https://mholt.github.io/curl-to-go/)
- [gjson](https://github.com/tidwall/gjson)
- [go-linq](https://github.com/ahmetb/go-linq)
