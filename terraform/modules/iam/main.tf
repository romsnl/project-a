provider "aws" {
  region = "us-east-1"
}

# Role creation
resource "aws_iam_role" "no_permissions_role" {
  name = var.role_name

  assume_role_policy = aws_iam_policy.assume_role_policy
}

# Role policy creation - this is the policy that allows users and entities to assume the role
resource "aws_iam_policy" "assume_role_policy" {
  name        = var.policy_name
  description = "Policy allowing users and entities to assume the role with no permissions"

  policy = jsonencode({
    Version = "2012-10-17",
    Statement = [
      {
        Effect = "Allow",
        Action = "sts:AssumeRole",
        Principal = {
          Service = "ec2.amazonaws.com" # Adjust this based on your use case
        }
      }
    ]
  })
}

# Group creation
resource "aws_iam_group" "assume_role_group" {
  name = var.group_name
}

# Group policy binding
resource "aws_iam_group_policy_attachment" "assume_role_group_attachment" {
  group      = aws_iam_group.assume_role_group.name
  policy_arn = aws_iam_policy.assume_role_policy.arn
}

# User creation
resource "aws_iam_user" "group_user" {
  name = var.user_name
}

# User group binding
resource "aws_iam_user_group_membership" "group_membership" {
  user   = aws_iam_user.group_user.name
  groups = [aws_iam_group.assume_role_group.name]
}
