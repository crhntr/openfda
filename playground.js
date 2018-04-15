var drugStrings = ["METOPROLOL", "METOPROLOL TARTRATE", "LOPRESSOR", "ATENOLOL", "ACEBUTOLOL", "CORGARD", "NADOLOL", "LOPRESSOR", "BISOPROLOL"]

var drugRegex = []

for (var i in drugStrings) {
  drugRegex.push(new RegExp(drugStrings[i], 'i'))
}

function drug_event_math490_4 () {
  return db.getCollection("drug_event").aggregate([
    {'$match': {
      "dup" : {"$ne": true},
      "sex": {"$in": ["1", "2"]},
      "dgs": {
        "$elemMatch": {
          "char": '1',
          "$or": [
            {"openfda.gn": {"$in": drugRegex}},
            {"mp": {"$in": drugRegex}},
            {'openfda.epc': 'beta-Adrenergic Blocker [EPC]'}
          ],
        }
      }
    }},
    {'$project': {
      '_id': 1,
      "sex": 1,
      "seriousness": '$s',
      'weight': '$wt',
      'onsetAgeNumber': '$oage',
      'onsetAgeUnit': '$ount',
      'drugs': {
          '$map': {
              'input': '$dgs',
              'as': 'drug',
              'in': {
                    'name': '$$drug.openfda.gn',
                    'characterization': '$$drug.char',
                    'medicinalProduct': '$$drug.mp',
                    'epc': '$$drug.openfda.epc'
              }
           }
       }
    }},
    { $out: 'drug_event_math490_4'}
])
}



function stats () {
  var drugs = ["ACEBUTOLOL", "ATENOLOL", "CORGARD", "NADOLOL", "LOPRESSOR", "METOPROLOL", "BISOPROLOL"]

  var data = {}
  for (var i in drugStrings) {
    data[drugStrings[i]] = db.drug_event_math490_1.find({
      "$or": [
        {"drugs.name": new RegExp(drugStrings[i], 'i')},
        {"medicinalProduct": new RegExp(drugStrings[i], 'i')}
      ]
    }).count()
  }
  return data
}
