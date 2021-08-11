variable "cluster" {
  default = "golang-app"
}
variable "app" {
  type        = string
  description = "Name of application"
  default     = "cycling-blog-app"
}
variable "zone" {
  default = "eu-west-2"
}
variable "docker-image" {
  type        = string
  description = "name of the docker image to deploy"
  default     = "jackstock8904/cycling-blog:1.2.3"
}
variable "namespace" {
  type        = string
  description = "Name of namespace"
  default     = "golang-app"
}