
module "iam" {
  source = "../../modules/iam"
  role_name = "project-a-role"
  policy_name = "project-a-policy"
  group_name = "project-a-group"
  user_name = "project-a"
}
