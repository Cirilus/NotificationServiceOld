env "local"{
  url = "postgresql://postgres:postgres@5.42.75.12:5432/postgres?sslmode=disable"
  dev = "postgresql://postgres:postgres@5.42.75.12:5432/dev?sslmode=disable"
  dir="file://migrations"
  migration {
    dir="file://migrations"
  }

}