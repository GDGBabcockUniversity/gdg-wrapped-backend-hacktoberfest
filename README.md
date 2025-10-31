# gdg-wrapped-backend-hacktoberfest 🎁

[![Stars](https://img.shields.io/github/stars/GDGBabcockUniversity/gdg-wrapped-backend-hacktoberfest?style=social)](https://github.com/GDGBabcockUniversity/gdg-wrapped-backend-hacktoberfest)
[![Forks](https://img.shields.io/github/forks/GDGBabcockUniversity/gdg-wrapped-backend-hacktoberfest?style=social)](https://github.com/GDGBabcockUniversity/gdg-wrapped-backend-hacktoberfest)

## Description 📝

The `gdg-wrapped-backend-hacktoberfest` is a Go-based backend application designed to provide a "wrapped" summary, similar to Spotify Wrapped, for Google Developer Groups (GDG) members. It aggregates and presents data related to member activity, group performance, and individual contributions within the GDG context. The application utilizes the Gin framework for handling HTTP requests and responses, and it retrieves data from a JSON file (`data/new_results.json`).

## Table of Contents 📚

1.  [Features](#features-)
2.  [Tech Stack](#tech-stack-)
3.  [Installation](#installation-)
4.  [Usage](#usage-)
5.  [Project Structure](#project-structure-)
6.  [API Reference](#api-reference-)
7.  [Contributing](#contributing-🤝)
8.  [License](#license-)
9.  [Important links](#important-links-)
10. [Footer](#footer-)

## Features ✨

- **General Wrapped Data**: Provides aggregated statistics for the entire GDG, such as the most active group, most active time, and most active members. 📊
- **Member-Specific Wrapped Data**: Offers personalized summaries for individual members, including total messages sent, link sharing activity, question asking habits, and overall message impact. 👤
- **REST API**: Exposes endpoints for retrieving both general and member-specific data via HTTP requests. 🌐
- **CORS Support**: Includes CORS configuration to allow cross-origin requests.
- **CORS Support**: Includes CORS configuration using an allowlist. Configure allowed origins via a `.env` file to control which origins are permitted for cross-origin requests.

## Tech Stack 💻

- **Go**: The primary programming language. 🚀
- **Gin**: A web framework for building the API. ⚙️
- **JSON**: Used for data serialization and storage (`data/new_results.json`). 🗄️
- **github.com/bytedance/sonic** - A high-performance JSON encoder/decoder.
- **github.com/gin-contrib/cors** - Gin middleware for handling Cross-Origin Resource Sharing (CORS).
- **github.com/stretchr/testify** - Testing framework.

## Installation 🛠️

1.  **Install Go**: Ensure you have Go installed (version 1.21 or later). ⬇️

2.  **Clone the Repository**: Clone the project repository to your local machine.

    ```bash
    git clone https://github.com/GDGBabcockUniversity/gdg-wrapped-backend-hacktoberfest.git
    cd gdg-wrapped-backend-hacktoberfest
    ```

3.  **Download Dependencies**: Use `go mod download` to download the required dependencies.

    ```bash
    go mod download
    ```

4.  **Configure Environment**: Copy `.env.example` to `.env` and configure your settings.

    ```bash
    cp .env.example .env
    # Edit .env to set your ALLOWED_ORIGINS for CORS
    ```

## Usage 🚀

1.  **Run the Application**: Execute the `main.go` file to start the backend server.

    The application will automatically load settings from your `.env` file, including CORS configuration.

    ```bash
    go run main.go
    ```

    **CORS Configuration**: Edit your `.env` file to set `ALLOWED_ORIGINS`:
    
    - Allow specific origins: `ALLOWED_ORIGINS=http://localhost:3000,https://example.com`
    - Allow any origin (use with caution): `ALLOWED_ORIGINS=*`2.  **Access API Endpoints**: Once the server is running, you can access the following endpoints:
    - `GET /ping`: Health check endpoint.
    - `GET /2023/general`: Retrieve general wrapped data.
    - `GET /2023/member/:num`: Retrieve member-specific wrapped data by providing a member number.

Example usage:

```bash
curl http://localhost:8080/2023/general
curl http://localhost:8080/2023/member/2347066550353
```

## Project Structure 📂

```
gdg-wrapped-backend-hacktoberfest/
├── data/
│   └── new_results.json       # JSON data containing wrapped information
├── types/
│   └── types.go               # Go data structures for wrapped data
├── util/
│   └── util.go                # Utility functions for data transformation
├── .env.example               # Environment configuration template
├── go.mod                     # Go module file
├── go.sum                     # Go dependencies checksum file
└── main.go                    # Main application entry point
```

## API Reference 🌐

### `GET /ping`

- Description: Health check endpoint. Returns a simple `pong` message.
- Response:

  ```json
  {
    "message": "pong"
  }
  ```

### `GET /2023/general`

- Description: Retrieves general wrapped data for the GDG.
- Response:

  ```json
  {
    "success": true,
    "data": {},
    "error": null
  }
  ```

### `GET /2023/member/:num`

- Description: Retrieves member-specific wrapped data.
- Path Parameters:
  - `num`: The member identifier.
- Response (Success):

  ```json
  {
    "success": true,
    "data": {},
    "error": null
  }
  ```

- Response (Error):

  ```json
  {
    "success": false,
    "error": "number does not exist",
    "data": null
  }
  ```

## Contributing 🤝

Contributions are welcome! Please follow these steps:

1.  Fork the repository. 🍴
2.  Create a new branch for your feature or bug fix. 🌿
3.  Make your changes and commit them with descriptive commit messages. ✍️
4.  Submit a pull request. 🚀

## License 📜

This project has no specified license.

## Important links 🔗

- Repository Link: [gdg-wrapped-backend-hacktoberfest](https://github.com/GDGBabcockUniversity/gdg-wrapped-backend-hacktoberfest)

## Footer <footer>

- **Repository:** [gdg-wrapped-backend-hacktoberfest](https://github.com/GDGBabcockUniversity/gdg-wrapped-backend-hacktoberfest)
- **Author:** GDGBabcockUniversity
- **Contact:** [GDG Babcock community.](https://gdg.community.dev/gdg-on-campus-babcock-university-ilishan-remo-nigeria/)

If you found this project helpful, consider giving it a star ⭐, forking it 🍴, or reporting any issues 🐛. Your support is greatly appreciated!

---
