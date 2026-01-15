target "docker-metadata-action" {}

variable "APP" {
  default = "decypharr"
}

variable "VERSION" {
  // renovate: datasource=github-releases depName=sirrobot01/decypharr
  default = "v1.1.6"
}

variable "SOURCE" {
  default = "https://github.com/sirrobot01/decypharr"
}

group "default" {
  targets = ["image-local"]
}

target "image" {
  inherits = ["docker-metadata-action"]
  args = {
    VERSION = "${VERSION}"
  }
  labels = {
    "org.opencontainers.image.source" = "${SOURCE}"
  }
}

target "image-local" {
  inherits = ["image"]
  output = ["type=docker"]
  tags = ["${APP}:${VERSION}"]
}

target "image-all" {
  inherits = ["image"]
  platforms = [
    "linux/amd64",
    "linux/arm64"
  ]
}
