terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = ">=5.95"
    }
  }
  backend "s3" {
    bucket = "algatux-opentofu-states"
    key    = "states/eks-infrastructure/${var.tfstate_name}"
    region = "eu-west-1"
  }
}

provider "aws" {
  region = var.region

  default_tags {
    tags = {
      Project = "queue-experiments"
      Owner   = "algatux"
    }
  }
}

module "sns" {
  source = "terraform-aws-modules/sns/aws"
  version = ">=v6.1"

  name = "event-collector"

  fifo_topic = true

  topic_policy_statements = {
    sqs = {
      sid = "SQSSubscribe"
      actions = [
        "sns:Subscribe",
        "sns:Receive",
      ]

      principals = [{
        type        = "AWS"
        identifiers = ["*"]
      }]

      conditions = [{
        test     = "StringLike"
        variable = "sns:Endpoint"
        values   = [module.sqs.queue_arn]
      }]
    }
  }

  subscriptions = {
    sqs = {
      protocol = "sqs"
      endpoint = module.sqs.queue_arn
    }
  }
}

module "sqs" {
  source = "terraform-aws-modules/sqs/aws"
  version = ">=v4.3"

  name = "queue-1"

  create_queue_policy = true
  queue_policy_statements = {
    sns = {
      sid     = "SNSPublish"
      actions = ["sqs:SendMessage"]

      principals = [
        {
          type        = "Service"
          identifiers = ["sns.amazonaws.com"]
        }
      ]

      conditions = [{
        test     = "ArnEquals"
        variable = "aws:SourceArn"
        values   = [module.sns.topic_arn]
      }]
    }
  }
}

