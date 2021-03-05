terraform {}

provider "dashdog" {
  host = "api.datadoghq.eu"
  api_key = "fe02fc133c9f3b6ad4a5fd2bffa2c415"
  app_key = "08147e0086776f39cd311457ad4613c26affc31c"
}

resource "dashdog_widget" "test" {
  type = "timeseries"
  title = "other title"
  requests = [
    {
      query = "avg:system.load.1{$host}"
      display_type = "area"
    }
  ]
}

resource "dashdog_dashboard" "dash" {
  title = "dash1"
  description = "super cool dashboard"
  read_only = false
  template_variables = [
    {
      name = "host"
      tag = "host"
      default = "*"
    }
  ]
  widgets = [
    {
      title = "group 1"
      json = [
        "${dashdog_widget.test.json}"
      ]
    }
  ]
}

resource "dashdog_dashboard" "dash2" {
  title = "dash2"
  description = "super cool dashboard"
  read_only = false
  template_variables = [
    {
      name = "host"
      tag = "host"
      default = "*"
    }
  ]
  widgets = [
    {
      title = "group 2"
      json = [
        "${dashdog_widget.test.json}"
      ]
    }
  ]
}
