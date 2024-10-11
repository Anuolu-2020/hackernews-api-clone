data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./internal/db/migrations",
  ]
}


variable "db_url" {
  type = string
}

env "dev" {
  src = data.external_schema.gorm.url
  dev = var.db_url
  migration {
    dir = "file://internal/db/migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}
