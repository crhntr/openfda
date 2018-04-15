var drugData = {
  'METOPROLOL': {
    ids: [],
    names: ['Lopressor', 'Acebutolol'],
    pharmClass: "beta-Adrenergic Agonist [EPC]"
  },
  'ACEBUTOLOL HYDROCHLORIDE': {
    ids: [],
    names: ['Sectral', 'Acebutolol'],
    pharmClass: "beta-Adrenergic Agonist [EPC]"
  },
  'ATENOLOL': {
    ids: [],
    names: ['Tenormin', 'ATENOLOL'],
    pharmClass: "beta-Adrenergic Blocker [EPC]"
  },
  'NADOLOL': {
    ids: [],
    names: ['CORGARD'],
    pharmClass: "beta-Adrenergic Blocker [EPC]"
  },
  'BISOPROLOL': {
    ids: [],
    names: ['Bisoprolol Fumarate'],
    pharmClass: 'beta-Adrenergic Blocker [EPC]'
  }
}

var ids = {}

function getIDsForName(name) {
  regex =  new RegExp(name, 'i')

  var idList = []
  var drugs = db.drug.find({'$or': [{'gn': regex}, {'bn': regex}]}).forEach( function (drug) {
    idList.push(drug._id)
  })
  return idList
}

function getIDs() {
  for (var key in drugData) {
    getIDsForName(key).forEach(function (id) {
      drugData[key].ids.push(id)
      ids[id] = key
    })

    for (var i in drugData[key].names) {
      getIDsForName(drugData[key].names[i]).forEach(function (id) {
        drugData[key].ids.push(id)
        ids[id] = key
      })
    }
  }
}

getIDs()

var idList = []
for (var id in ids) {
  idList.push(id)
}

printjson({
  drugData: drugData,
  ids: ids,
  idList: idList
})

var drugNames = []

for (var key in drugData) {
  drugNames.push(new RegExp(key, 'i'))

  for (var nameIndex in drugData[key]) {
    drugNames.push(new RegExp(drugData[key][nameIndex], 'i'))
  }
}

var match = {
  'sex': {'$in': ['1', '2']},
  "dup" : {"$ne": true},
  "dgs": {"$elemMatch": {
    "char": '1',
    "$or": [
      {"openfda._id": {"$in": idList}},
      {"openfda.gn": {"$in": drugNames}},
      {"openfda.bn": {"$in": drugNames}},
      {"mp": {"$in": drugNames}}
    ],
  }}
}


var project = {
  '_id': 1,
  "sex": 1,
  "seriousness": '$s',
  'weight': '$wt',
  'onsetAgeNumber': '$oage',
  'onsetAgeUnit': '$ount',
  'drugCount': { '$size': '$dgs' },
  'drugs': {
    '$map': {
      'input': '$dgs',
      'as': 'drug',
      'in': {
        '_id': '$$drug.openfda.id',
        'gn': '$$drug.openfda.gn',
        'bn': '$$drug.openfda.bn',
        'char': '$$drug.char',
        'md': '$$drug.mp',
        'epc': '$$drug.openfda.epc'
      }
    }
  }
}


/*

'drug': {
  '$arrayElemAt': [
    '$dgs', {
      '$indexOfArray': ['$dgs', {
        "char": '1',
        "$or": [
          {"openfda._id": {"$in": idList}},
          {"openfda.gn": {"$in": drugNames}},
          {"openfda.bn": {"$in": drugNames}},
          {"mp": {"$in": drugNames}}
        ]
      }]
    }
  ]
}


'drugs': {
  '$map': {
    'input': '$dgs',
    'as': 'drug',
    'in': {
      'name': '$$drug.openfda.gn',
      'char': '$$drug.char',
      'md': '$$drug.mp',
      'epc': '$$drug.openfda.epc'
    }
  }
}

*/

function rip (out) {
  return db.getCollection('drug_event').aggregate([
    { $match: match},
    { $project: project},
    { $limit: 5 } // ,
    // { $out: out }
  ])
}

//
// function drug_event_math490_4 () {
//   return db.getCollection('drug_event').aggregate([
//     {'$match': {
//       "dup" : {"$ne": true},
//       "sex": {"$in": ["1", "2"]},
//       "dgs": {
//
//         "$elemMatch": {
//           "char": '1',
//           "$or": [
//             {"openfda.gn": {"$in": drugRegex}},
//             {"mp": {"$in": drugRegex}}
//           ],
//         }
//       }
//     }},
//     {'$project': {
//       '_id': 1,
//       "sex": 1,
//       "seriousness": '$s',
//       'weight': '$wt',
//       'onsetAgeNumber': '$oage',
//       'onsetAgeUnit': '$ount',
//       'drugs': {
//           '$map': {
//               'input': '$dgs',
//               'as': 'drug',
//               'in': {
//                     'name': '$$drug.openfda.gn',
//                     'characterization': '$$drug.char',
//                     'medicinalProduct': '$$drug.mp',
//                     'epc': '$$drug.openfda.epc'
//               }
//            }
//        }
//     }},
//     { $out: 'drug_event_math490_4'}
// ])
// }
//
//
//
// function stats () {
//   var drugs = ["ACEBUTOLOL", "ATENOLOL", "CORGARD", "NADOLOL", "LOPRESSOR", "METOPROLOL", "BISOPROLOL"]
//
//   var data = {}
//   for (var i in drugStrings) {
//     data[drugStrings[i]] = db.drug_event_math490_1.find({
//       "$or": [
//         {"drugs.name": new RegExp(drugStrings[i], 'i')},
//         {"medicinalProduct": new RegExp(drugStrings[i], 'i')}
//       ]
//     }).count()
//   }
//   return data
// }
