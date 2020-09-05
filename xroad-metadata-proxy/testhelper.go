package proxy

import (
  "testing"

  "github.com/alicebob/miniredis"
  "github.com/go-redis/redis"
  "github.com/mitchellh/mapstructure"
)

func NewRedisMock(t *testing.T) (*redis.Client, *miniredis.Miniredis) {
  t.Helper()

  miniredis, err := miniredis.Run()
  if err != nil {
    t.Fatalf("Error while createing test redis server '%#v'", err)
  }

  client := redis.NewClient(&redis.Options{
    Addr:     miniredis.Addr(),
    Password: "",
    DB:       0,
  })

  return client, miniredis
}

// Make it private function
func HMSetAll(t *testing.T, mapXRoadMembers []map[string]interface{}, client *redis.Client, key string) {
  t.Helper()

  for _, mapXRoadMember := range mapXRoadMembers {
    err := client.HMSet(key + ":" + mapXRoadMember["ID"].(string), mapXRoadMember).Err()
    if err != nil {
      t.Fatalf("Failed to set test data '%#v'", err)
    }
  }
}

func decodeXRoadMembers(t *testing.T, mapXRoadMembers []map[string]interface{}) []*XRoadMember {
  t.Helper()

  var expectedXRoadMembers []*XRoadMember

  for _, mapXRoadMember := range mapXRoadMembers {
    var xRoadMember XRoadMember

    err := mapstructure.Decode(mapXRoadMember, &xRoadMember)
    if err != nil {
      t.Fatalf("Failed to decode '%#v'", err)
    }
    expectedXRoadMembers = append(expectedXRoadMembers, &xRoadMember)
  }

  return expectedXRoadMembers
}
