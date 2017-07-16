# Envoy <small>ˈenˌvoi,ˈänˌvoi</small>

1. a messenger or representative, especially one on a diplomatic mission.
2. short for envoy extraordinary.

---

Envoy is a simple utility to query the Meetup API for groups by a criteria:

```bash
./envoy -location Berlin -text Elixir | jq '.[] | { name, organizer, score }'
```
