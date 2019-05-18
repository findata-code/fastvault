package token

import (
	"fmt"
	"testing"
)

func TestGetTokens(t *testing.T) {
	data := []byte(`
		[
  {
    "_id": "5cdfa8f699e5e081c854d526",
    "index": 0,
    "guid": "98f87e23-fb18-4eb7-9bb7-f4b49240f6a1",
    "isActive": false,
    "balance": "$2,830.87",
    "picture": "http://placehold.it/32x32",
    "age": 32,
    "eyeColor": "green",
    "name": "Monica Mullins",
    "gender": "female",
    "company": "QUILITY",
    "email": "monicamullins@quility.com",
    "phone": "+1 (935) 519-2625",
    "address": "441 Beverly Road, Echo, Georgia, 676",
    "about": "Proident sint aliqua sint amet do nisi reprehenderit ullamco ad. Nulla irure ullamco ut cillum elit adipisicing incididunt sint non incididunt cupidatat laborum. Aute nostrud tempor cupidatat Lorem ipsum adipisicing dolor non Lorem ex aliquip. Commodo sit officia adipisicing elit excepteur nulla quis consequat ullamco voluptate laborum reprehenderit reprehenderit. Sit sit esse cupidatat sit dolore ad labore tempor. Culpa ea nisi in nisi tempor cillum. Dolor ea dolore esse officia veniam dolore dolor occaecat exercitation non non ut elit exercitation.\r\n",
    "registered": "2014-11-20T01:29:20 -07:00",
    "latitude": 33.875483,
    "longitude": 103.415651,
    "tags": [
      "labore",
      "exercitation",
      "aliquip",
      "enim",
      "mollit",
      "culpa",
      "in"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Tammy Stanton"
      },
      {
        "id": 1,
        "name": "Mcpherson Frederick"
      },
      {
        "id": 2,
        "name": "Compton Keith"
      }
    ],
    "greeting": "Hello, Monica Mullins! You have 7 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5cdfa8f6d38a2f586c322e04",
    "index": 1,
    "guid": "a1b0255d-5423-48b0-8fa8-5d1297fef400",
    "isActive": true,
    "balance": "$3,936.95",
    "picture": "http://placehold.it/32x32",
    "age": 30,
    "eyeColor": "brown",
    "name": "Ava Lewis",
    "gender": "female",
    "company": "GENMEX",
    "email": "avalewis@genmex.com",
    "phone": "+1 (996) 481-2128",
    "address": "771 Just Court, Yorklyn, Missouri, 5095",
    "about": "Non ullamco dolore id deserunt elit tempor fugiat Lorem est. Ea reprehenderit cillum nostrud ipsum reprehenderit reprehenderit quis enim. Magna minim fugiat dolor id quis cillum id aute sint voluptate Lorem excepteur voluptate. Eu sunt et velit sint adipisicing nulla nulla. Consequat esse enim ex sint qui commodo et esse aute ea tempor sint. Ipsum occaecat minim labore magna ad quis occaecat elit velit tempor in. Laboris ullamco aliquip officia in consectetur nulla.\r\n",
    "registered": "2018-06-28T04:35:17 -07:00",
    "latitude": 33.183233,
    "longitude": -148.946182,
    "tags": [
      "enim",
      "eu",
      "dolor",
      "aute",
      "ut",
      "sint",
      "ipsum"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Matilda Booker"
      },
      {
        "id": 1,
        "name": "Guadalupe Cantrell"
      },
      {
        "id": 2,
        "name": "Tiffany Sheppard"
      }
    ],
    "greeting": "Hello, Ava Lewis! You have 8 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5cdfa8f6158914c40dae86ff",
    "index": 2,
    "guid": "08a7624b-69d8-41f6-b7e5-f9da8fe392a0",
    "isActive": false,
    "balance": "$3,140.83",
    "picture": "http://placehold.it/32x32",
    "age": 35,
    "eyeColor": "green",
    "name": "Peggy Sanders",
    "gender": "female",
    "company": "CUIZINE",
    "email": "peggysanders@cuizine.com",
    "phone": "+1 (983) 585-3216",
    "address": "916 Taylor Street, Kidder, Connecticut, 795",
    "about": "Esse incididunt ipsum enim dolor dolor officia aliqua ad. Dolore veniam duis aute voluptate commodo laborum magna amet dolor laborum pariatur ut est. Enim sunt in non velit incididunt pariatur elit. Ex do exercitation consequat adipisicing consequat sunt cillum pariatur ea.\r\n",
    "registered": "2016-01-05T12:12:35 -07:00",
    "latitude": -71.181408,
    "longitude": -114.76042,
    "tags": [
      "minim",
      "enim",
      "proident",
      "nulla",
      "occaecat",
      "enim",
      "ea"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Earlene Williams"
      },
      {
        "id": 1,
        "name": "Noel Hendricks"
      },
      {
        "id": 2,
        "name": "Keri Carter"
      }
    ],
    "greeting": "Hello, Peggy Sanders! You have 9 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5cdfa8f621b688307d8858ed",
    "index": 3,
    "guid": "11797997-f01a-48b5-b182-e15800676a7a",
    "isActive": true,
    "balance": "$3,531.39",
    "picture": "http://placehold.it/32x32",
    "age": 35,
    "eyeColor": "brown",
    "name": "Mable Bonner",
    "gender": "female",
    "company": "NETUR",
    "email": "mablebonner@netur.com",
    "phone": "+1 (925) 536-3979",
    "address": "622 Noel Avenue, Aguila, District Of Columbia, 1022",
    "about": "Sunt excepteur id irure dolor pariatur commodo ex mollit enim. Non qui voluptate excepteur consequat ut tempor voluptate in magna aliquip. Sunt deserunt cupidatat veniam adipisicing aliquip ex.\r\n",
    "registered": "2016-02-03T04:55:19 -07:00",
    "latitude": 19.551664,
    "longitude": 24.888042,
    "tags": [
      "aliqua",
      "est",
      "sunt",
      "nostrud",
      "Lorem",
      "veniam",
      "ut"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Luann Wood"
      },
      {
        "id": 1,
        "name": "Pace Daugherty"
      },
      {
        "id": 2,
        "name": "Delores Morgan"
      }
    ],
    "greeting": "Hello, Mable Bonner! You have 7 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5cdfa8f60c641a94e7162887",
    "index": 4,
    "guid": "088d1be3-712b-4dad-9bcf-8bc4bb3fd940",
    "isActive": true,
    "balance": "$3,922.57",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "brown",
    "name": "Dean Chen",
    "gender": "male",
    "company": "EXERTA",
    "email": "deanchen@exerta.com",
    "phone": "+1 (939) 551-2322",
    "address": "510 Sands Street, Kirk, Marshall Islands, 6745",
    "about": "Cupidatat minim aute sint sint exercitation enim ad eiusmod incididunt elit laboris aliqua mollit. Id Lorem mollit aute irure ipsum pariatur eiusmod ea. Non elit ea in sunt. Ut ipsum culpa irure aliqua. Elit nulla occaecat Lorem incididunt proident est amet quis sint sit quis non ex quis. Ex excepteur occaecat deserunt commodo officia id dolor.\r\n",
    "registered": "2016-10-25T11:02:38 -07:00",
    "latitude": -45.765213,
    "longitude": 43.326394,
    "tags": [
      "exercitation",
      "minim",
      "veniam",
      "ad",
      "exercitation",
      "ut",
      "tempor"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Charlotte Burns"
      },
      {
        "id": 1,
        "name": "Estela Jacobs"
      },
      {
        "id": 2,
        "name": "Carmella Martin"
      }
    ],
    "greeting": "Hello, Dean Chen! You have 3 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5cdfa8f6d11226ef4848af0a",
    "index": 5,
    "guid": "62100213-d3b3-46c8-9f20-6c9d5e1296fc",
    "isActive": true,
    "balance": "$2,209.19",
    "picture": "http://placehold.it/32x32",
    "age": 29,
    "eyeColor": "brown",
    "name": "Berg Miranda",
    "gender": "male",
    "company": "CONCILITY",
    "email": "bergmiranda@concility.com",
    "phone": "+1 (885) 541-2165",
    "address": "281 Gem Street, Norwood, South Dakota, 1439",
    "about": "Eiusmod enim eiusmod ad ullamco mollit est ut veniam amet magna. Exercitation id occaecat eiusmod nisi nulla veniam aliquip ex Lorem aute consequat. Esse ea labore ex exercitation irure exercitation aute ipsum excepteur enim esse. Reprehenderit laboris eiusmod do id sint nulla ipsum nisi tempor ea quis. Officia excepteur adipisicing aliquip adipisicing ad culpa. Et ullamco aliquip quis minim ipsum veniam quis eu. Occaecat ad consequat consectetur consequat aute voluptate et in proident dolor in ipsum.\r\n",
    "registered": "2014-04-26T07:14:51 -07:00",
    "latitude": 64.854341,
    "longitude": 4.91043,
    "tags": [
      "sit",
      "sit",
      "adipisicing",
      "Lorem",
      "veniam",
      "anim",
      "ipsum"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Wooten Peterson"
      },
      {
        "id": 1,
        "name": "Patricia Serrano"
      },
      {
        "id": 2,
        "name": "Courtney Patrick"
      }
    ],
    "greeting": "Hello, Berg Miranda! You have 5 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5cdfa8f6446af13c5c0717fb",
    "index": 6,
    "guid": "dca62282-cb25-4a10-bf7c-e9b589784616",
    "isActive": false,
    "balance": "$3,038.23",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "green",
    "name": "Joni Livingston",
    "gender": "female",
    "company": "MULTIFLEX",
    "email": "jonilivingston@multiflex.com",
    "phone": "+1 (938) 588-2756",
    "address": "832 Court Street, Rockbridge, California, 8435",
    "about": "Sunt consequat duis velit enim dolore quis. Id dolor dolore ad ea ullamco quis amet irure labore anim. Eiusmod elit minim in ullamco eiusmod sint officia dolor eu nisi eiusmod pariatur sit deserunt.\r\n",
    "registered": "2014-01-05T03:02:27 -07:00",
    "latitude": -7.763574,
    "longitude": -42.241013,
    "tags": [
      "esse",
      "sit",
      "labore",
      "esse",
      "sit",
      "nostrud",
      "incididunt"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Black Wallace"
      },
      {
        "id": 1,
        "name": "Carrie Ball"
      },
      {
        "id": 2,
        "name": "Lenore Brennan"
      }
    ],
    "greeting": "Hello, Joni Livingston! You have 5 unread messages.",
    "favoriteFruit": "apple"
  }
]
	`)
	_, err := GetTokens(data)
	if err != nil {
		t.Error(err)
	}
}

func TestGetTokensShouldReturnError(t *testing.T) {
	data := []byte(`
		[
  {
    "_id": "5cdfa8f699e5e081c854d526",
    "index": 0,
    "guid": "98f87e23-fb18-4eb7-9bb7-f4b49240f6a1",
    "isActive": false,
    "balance": "$2,830.87",
    "picture": "http://placehold.it/32x32",
    "age": 32,
    "eyeColor": "green",
    "name": "Monica Mullins",
    "gender": "female",
    "company": "QUILITY",
    "email": "monicamullins@quility.com",
    "phone": "+1 (935) 519-2625",
    "address": "441 Beverly Road, Echo, Georgia, 676",
    "about": "Proident sint aliqua sint amet do nisi reprehenderit ullamco ad. Nulla irure ullamco ut cillum elit adipisicing incididunt sint non incididunt cupidatat laborum. Aute nostrud tempor cupidatat Lorem ipsum adipisicing dolor non Lorem ex aliquip. Commodo sit officia adipisicing elit excepteur nulla quis consequat ullamco voluptate laborum reprehenderit reprehenderit. Sit sit esse cupidatat sit dolore ad labore tempor. Culpa ea nisi in nisi tempor cillum. Dolor ea dolore esse officia veniam dolore dolor occaecat exercitation non non ut elit exercitation.\r\n",
    "registered": "2014-11-20T01:29:20 -07:00",
    "latitude": 33.875483,
    "longitude": 103.415651,
    "tags": [
      "labore",
      "exercitation",
      "aliquip",
      "enim",
      "mollit",
      "culpa",
      "in"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Tammy Stanton"
      },
      {
        "id": 1,
        "name": "Mcpherson Frederick"
      },
      {
        "id": 2,
        "name": "Compton Keith"
      }
    ],
    "greeting": "Hello, Monica Mullins! You have 7 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5cdfa8f6d38a2f586c322e04",
    "index": 1,
    "guid": "a1b0255d-5423-48b0-8fa8-5d1297fef400",
    "isActive": true,
    "balance": "$3,936.95",
    "picture": "http://placehold.it/32x32",
    "age": 30,
    "eyeColor": "brown",
    "name": "Ava Lewis",
    "gender": "female",
    "company": "GENMEX",
    "email": "avalewis@genmex.com",
    "phone": "+1 (996) 481-2128",
    "address": "771 Just Court, Yorklyn, Missouri, 5095",
    "about": "Non ullamco dolore id deserunt elit tempor fugiat Lorem est. Ea reprehenderit cillum nostrud ipsum reprehenderit reprehenderit quis enim. Magna minim fugiat dolor id quis cillum id aute sint voluptate Lorem excepteur voluptate. Eu sunt et velit sint adipisicing nulla nulla. Consequat esse enim ex sint qui commodo et esse aute ea tempor sint. Ipsum occaecat minim labore magna ad quis occaecat elit velit tempor in. Laboris ullamco aliquip officia in consectetur nulla.\r\n",
    "registered": "2018-06-28T04:35:17 -07:00",
    "latitude": 33.183233,
    "longitude": -148.946182,
    "tags": [
      "enim",
      "eu",
      "dolor",
      "aute",
      "ut",
      "sint",
      "ipsum"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Matilda Booker"
      },
      {
        "id": 1,
        "name": "Guadalupe Cantrell"
      },
      {
        "id": 2,
        "name": "Tiffany Sheppard"
      }
    ],
    "greeting": "Hello, Ava Lewis! You have 8 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5cdfa8f6158914c40dae86ff",
    "index": 2,
    "guid": "08a7624b-69d8-41f6-b7e5-f9da8fe392a0",
    "isActive": false,
    "balance": "$3,140.83",
    "picture": "http://placehold.it/32x32",
    "age": 35,
    "eyeColor": "green",
    "name": "Peggy Sanders",
    "gender": "female",
    "company": "CUIZINE",
    "email": "peggysanders@cuizine.com",
    "phone": "+1 (983) 585-3216",
    "address": "916 Taylor Street, Kidder, Connecticut, 795",
    "about": "Esse incididunt ipsum enim dolor dolor officia aliqua ad. Dolore veniam duis aute voluptate commodo laborum magna amet dolor laborum pariatur ut est. Enim sunt in non velit incididunt pariatur elit. Ex do exercitation consequat adipisicing consequat sunt cillum pariatur ea.\r\n",
    "registered": "2016-01-05T12:12:35 -07:00",
    "latitude": -71.181408,
    "longitude": -114.76042,
    "tags": [
      "minim",
      "enim",
      "proident",
      "nulla",
      "occaecat",
      "enim",
      "ea"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Earlene Williams"
      },
      {
        "id": 1,
        "name": "Noel Hendricks"
      },
      {
        "id": 2,
        "name": "Keri Carter"
      }
    ],
    "greeting": "Hello, Peggy Sanders! You have 9 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5cdfa8f621b688307d8858ed",
    "index": 3,
    "guid": "11797997-f01a-48b5-b182-e15800676a7a",
    "isActive": true,
    "balance": "$3,531.39",
    "picture": "http://placehold.it/32x32",
    "age": 35,
    "eyeColor": "brown",
    "name": "Mable Bonner",
    "gender": "female",
    "company": "NETUR",
    "email": "mablebonner@netur.com",
    "phone": "+1 (925) 536-3979",
    "address": "622 Noel Avenue, Aguila, District Of Columbia, 1022",
    "about": "Sunt excepteur id irure dolor pariatur commodo ex mollit enim. Non qui voluptate excepteur consequat ut tempor voluptate in magna aliquip. Sunt deserunt cupidatat veniam adipisicing aliquip ex.\r\n",
    "registered": "2016-02-03T04:55:19 -07:00",
    "latitude": 19.551664,
    "longitude": 24.888042,
    "tags": [
      "aliqua",
      "est",
      "sunt",
      "nostrud",
      "Lorem",
      "veniam",
      "ut"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Luann Wood"
      },
      {
        "id": 1,
        "name": "Pace Daugherty"
      },
      {
        "id": 2,
        "name": "Delores Morgan"
      }
    ],
    "greeting": "Hello, Mable Bonner! You have 7 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5cdfa8f60c641a94e7162887",
    "index": 4,
    "guid": "088d1be3-712b-4dad-9bcf-8bc4bb3fd940",
    "isActive": true,
    "balance": "$3,922.57",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "brown",
    "name": "Dean Chen",
    "gender": "male",
    "company": "EXERTA",
    "email": "deanchen@exerta.com",
    "phone": "+1 (939) 551-2322",
    "address": "510 Sands Street, Kirk, Marshall Islands, 6745",
    "about": "Cupidatat minim aute sint sint exercitation enim ad eiusmod incididunt elit laboris aliqua mollit. Id Lorem mollit aute irure ipsum pariatur eiusmod ea. Non elit ea in sunt. Ut ipsum culpa irure aliqua. Elit nulla occaecat Lorem incididunt proident est amet quis sint sit quis non ex quis. Ex excepteur occaecat deserunt commodo officia id dolor.\r\n",
    "registered": "2016-10-25T11:02:38 -07:00",
    "latitude": -45.765213,
    "longitude": 43.326394,
    "tags": [
      "exercitation",
      "minim",
      "veniam",
      "ad",
      "exercitation",
      "ut",
      "tempor"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Charlotte Burns"
      },
      {
        "id": 1,
        "name": "Estela Jacobs"
      },
      {
        "id": 2,
        "name": "Carmella Martin"
      }
    ],
    "greeting": "Hello, Dean Chen! You have 3 unread messages.",
    "favoriteFruit": "banana"
  },
  {
    "_id": "5cdfa8f6d11226ef4848af0a",
    "index": 5,
    "guid": "62100213-d3b3-46c8-9f20-6c9d5e1296fc",
    "isActive": true,
    "balance": "$2,209.19",
    "picture": "http://placehold.it/32x32",
    "age": 29,
    "eyeColor": "brown",
    "name": "Berg Miranda",
    "gender": "male",
    "company": "CONCILITY",
    "email": "bergmiranda@concility.com",
    "phone": "+1 (885) 541-2165",
    "address": "281 Gem Street, Norwood, South Dakota, 1439",
    "about": "Eiusmod enim eiusmod ad ullamco mollit est ut veniam amet magna. Exercitation id occaecat eiusmod nisi nulla veniam aliquip ex Lorem aute consequat. Esse ea labore ex exercitation irure exercitation aute ipsum excepteur enim esse. Reprehenderit laboris eiusmod do id sint nulla ipsum nisi tempor ea quis. Officia excepteur adipisicing aliquip adipisicing ad culpa. Et ullamco aliquip quis minim ipsum veniam quis eu. Occaecat ad consequat consectetur consequat aute voluptate et in proident dolor in ipsum.\r\n",
    "registered": "2014-04-26T07:14:51 -07:00",
    "latitude": 64.854341,
    "longitude": 4.91043,
    "tags": [
      "sit",
      "sit",
      "adipisicing",
      "Lorem",
      "veniam",
      "anim",
      "ipsum"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Wooten Peterson"
      },
      {
        "id": 1,
        "name": "Patricia Serrano"
      },
      {
        "id": 2,
        "name": "Courtney Patrick"
      }
    ],
    "greeting": "Hello, Berg Miranda! You have 5 unread messages.",
    "favoriteFruit": "strawberry"
  },
  {
    "_id": "5cdfa8f6446af13c5c0717fb",
    "index": 6,
    "guid": "dca62282-cb25-4a10-bf7c-e9b589784616",
    "isActive": false,
    "balance": "$3,038.23",
    "picture": "http://placehold.it/32x32",
    "age": 24,
    "eyeColor": "green",
    "name": "Joni Livingston",
    "gender": "female",
    "company": "MULTIFLEX",
    "email": "jonilivingston@multiflex.com",
    "phone": "+1 (938) 588-2756",
    "address": "832 Court Street, Rockbridge, California, 8435",
    "about": "Sunt consequat duis velit enim dolore quis. Id dolor dolore ad ea ullamco quis amet irure labore anim. Eiusmod elit minim in ullamco eiusmod sint officia dolor eu nisi eiusmod pariatur sit deserunt.\r\n",
    "registered": "2014-01-05T03:02:27 -07:00",
    "latitude": -7.763574,
    "longitude": -42.241013,
    "tags": [
      "esse",
      "sit",
      "labore",
      "esse",
      "sit",
      "nostrud",
      "incididunt"
    ],
    "friends": [
      {
        "id": 0,
        "name": "Black Wallace"
      },
      {
        "id": 1,
        "name": "Carrie Ball"
      },
      {
        "id": 2,
        "name": "Lenore Brennan"
      }
    ],
    "greeting": "Hello, Joni Livingston! You have 5 unread messages.",
    "favoriteFruit": "apple"
  }
]
	`)
	_, err := GetTokens(data)
	if err == nil {
		t.Error("should return error due to key length not equals 16")
	}
}

func TestGenerateKey(t *testing.T) {
	key := GenerateKey()
	if len(key) != 16 {
		t.Error("key should return 16 count of number but got", fmt.Sprintf("%s", key), len(key))
	}

	fmt.Sprintf("%s", key)
}
