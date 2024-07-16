<div align="center">
  <br>
  <h1>AMA</h1>
  <p>🙋 A simple attendance manager CLI app 🙋</p>
  <br>
</div>

## Table of Contents

- [Description](#description)
- [Minimum Requirements](#minimum-requirements)
- [Usage](#usage)
- [Screenshots](#screenshots)

## Description

[`^ back to top ^`](#table-of-contents)

**AMA** is a simple CLI app to manage employee attendance. The app is made with Go. It is created as an exercise in the Go phase of the Backend Development Training.

## Minimum Requirements

[`^ back to top ^`](#table-of-contents)

- Show list of employees and their presence status.
- Add a new employee to the presence list.
- Update presence status (true/ present or false/ absent).
- Delete an employee from the presence list.
- Use map or slice to store employees data.
- Each employee has ID, name, & presence status (boolean).
- Use pointer to access & manipulate data.

## Usage

[`^ back to top ^`](#table-of-contents)

- Start the app.

  ```bash
  go run ./cmd/app
  ```

- Enter a command.

  Choose one of these commands:

  **Get list of employees**

  ```bash
  \l
  ```

  **Add a new employee**

  ```bash
  \a
  ```

  **Update employee presence status**

  ```bash
  \u [employee-id] [status (y/n)]
  ```

  **Delete an employee**

  ```bash
  \d [employee-id]
  ```

  **Show all commands**

  ```bash
  \c
  ```

  **Quit the app**

  ```bash
  \q
  ```

## Screenshots

### Show list of employees & add a new employee

![list & add](./docs/img/list-add.png)

### Update presence status

![update](./docs/img/update.png)

### Delete an employee

![delete](./docs/img/delete.png)
