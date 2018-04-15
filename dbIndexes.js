

db.drug_event.createIndex({
  'sex': 1,
  's': 1,
  'drugs.openfda._id'
}, {'sparse': true})
