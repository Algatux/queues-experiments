
variable "region" {
  type = string
  description = "AWS region"
  default = "eu-west-1"
}


variable "tfstate_name" {
  type = string
  description = "TFState filename"
  default = "queues.tfstate"
}