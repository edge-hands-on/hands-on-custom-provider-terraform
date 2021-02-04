terraform {
}

provider "multiverse" {}

resource "multiverse_custom_resource" "custom-resource" {
  executor = "python3"
  script = "my-script.py"
  id_key = "id"
  name = "test"
  data = <<JSON
{
  "name": "test-terraform-test",
  "mlb_id": "lb-124",
  "mlb_deployment_id": "dp-123",
  "mlb_listener_ids": ["ls-124", "ls-456"],
  "test_group_callback_fqdn": "test.fqdn.com",
  "control_group_callback_fqdn": "control.fqdn.com"
}
JSON
}
