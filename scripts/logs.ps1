param(
  [string]$Service = ""
)

if ($Service -eq "") {
  docker compose -f .\infra\docker-compose.yml logs -f --tail=200
} else {
  docker compose -f .\infra\docker-compose.yml logs -f --tail=200 $Service
}