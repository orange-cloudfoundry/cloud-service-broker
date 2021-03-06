variable cores { type = number }
variable instance_name { type = string }
variable db_name { type = string }
variable admin_password { type = string }
variable labels { type = map }
variable storage_gb { type = number }
variable publicly_accessible { type = bool }
variable multi_az { type = bool }
variable instance_class { type = string }
variable engine { type = string }
variable engine_version { type = string }
variable aws_vpc_id { type = string }