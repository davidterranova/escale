package file

import (
	"bytes"
	"testing"

	"github.com/matryer/is"
)

func TestJSONPortsLoaderRead(t *testing.T) {
	is := is.New(t)

	reader := NewJSONPortLoader(bytes.NewBufferString(rawJSON))
	n := 0
	for p := range reader.Read() {
		n++
		t.Logf("%+v\n", *p)
	}
	is.Equal(n, 10)
}

const rawJSON = `{
  "AEAJM": {
    "name": "Ajman",
    "city": "Ajman",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.5136433,
      25.4052165
    ],
    "province": "Ajman",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAJM"
    ],
    "code": "52000"
  },
  "AEAUH": {
    "name": "Abu Dhabi",
    "coordinates": [
      54.37,
      24.47
    ],
    "city": "Abu Dhabi",
    "province": "Abu Z¸aby [Abu Dhabi]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEAUH"
    ],
    "code": "52001"
  },
  "AEDXB": {
    "name": "Dubai",
    "coordinates": [
      55.27,
      25.25
    ],
    "city": "Dubai",
    "province": "Dubayy [Dubai]",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEDXB"
    ],
    "code": "52005"
  },
  "AEFJR": {
    "name": "Al Fujayrah",
    "coordinates": [
      56.33,
      25.12
    ],
    "city": "Al Fujayrah",
    "province": "Al Fujayrah",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEFJR"
    ]
  },
  "AEJEA": {
    "name": "Jebel Ali",
    "city": "Jebel Ali",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.0272904,
      24.9857145
    ],
    "province": "Dubai",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEJEA"
    ],
    "code": "52051"
  },
  "AEJED": {
    "name": "Jebel Dhanna",
    "city": "Jebel Dhanna",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      52.6126027,
      24.1915137
    ],
    "province": "Abu Dhabi",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEJED"
    ],
    "code": "52050"
  },
  "AEKLF": {
    "name": "Khor al Fakkan",
    "coordinates": [
      56.35,
      25.33
    ],
    "city": "Khor al Fakkan",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEKLF"
    ]
  },
  "AEPRA": {
    "name": "Port Rashid",
    "city": "Port Rashid",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "coordinates": [
      55.2756505,
      25.284755
    ],
    "province": "Dubai",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEPRA"
    ],
    "code": "52005"
  },
  "AEQIW": {
    "name": "Umm al Qaiwain",
    "coordinates": [
      55.55,
      25.57
    ],
    "city": "Umm al Qaiwain",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "province": "Umm Al Quwain",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AEQIW"
    ]
  },
  "AERKT": {
    "name": "Ras al Khaimah",
    "coordinates": [
      55.95,
      25.78
    ],
    "city": "Ras al Khaimah",
    "country": "United Arab Emirates",
    "alias": [],
    "regions": [],
    "province": "Ras al Khaimah",
    "timezone": "Asia/Dubai",
    "unlocs": [
      "AERKT"
    ],
    "code": "52000"
  }
}`
