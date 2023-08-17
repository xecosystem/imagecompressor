# Image Compression Web Application

This is a simple Go web application that fetches, compresses, and serves images from remote URLs.

## Features

- Accepts a query parameter `url` to fetch and compress images.
- Checks if the image is already cached and serves it if available.
- Compresses images without resizing them.
- Supports caching of images in memory for 24 hours.

## Prerequisites

- Go 1.16 or later
- Docker (optional, for containerization)

## Getting Started

1. Clone this repository to your local machine:

  ```sh
  git clone https://github.com/xecosystem/imagecompressor.git
  ```

2. Navigate to the project directory:

  ```sh
  cd imagecompressor
  ```

3. Run the application:

  ```sh
  go run main.go
  ```
  The application will start a server that listens on port 8080.

4. Access the application in your browser at http://localhost:8080.

## Docker Support

You can also run the application using Docker. Make sure you have Docker installed on your system.

1. Build the Docker image:
  
  ```sh
  docker build -t imagecompressor .
  ```

2. Run the Docker container:

  ```sh
  docker run -p 8080:8080 imagecompressor
  ```

3. Access the application in your browser at http://localhost:8080.

## API Endpoint

To use the image compression feature, make a GET request to the following endpoint:

```
GET /?url=YOUR_IMAGE_URL
```
Replace `YOUR_IMAGE_URL` with the URL of the image you want to compress.

## Caching

The application caches compressed images in memory for 24 hours. If the same image URL is requested within this period, the cached image will be served without recompression.

## Contributing

Contributions are welcome! If you find a bug or have an improvement in mind, please create an issue or submit a pull request.

## License

This project is licensed under the MIT License.
