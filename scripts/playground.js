load('scripts/metaData.js')










/*

Serious           Seriousness = 0
CongenitalAnomali Seriousness = 2
Death 4
Disabling 8
Hospitalization 16
LifeThreatening 32
Other 64


*/


for (var key in drugData) {
  db.math490_test.find().forEach(function (doc) {
    var set = {}

    set['death'] = (parseInt(doc.s)&4)!==0

    set['disabling'] = (parseInt(doc.s)&8)!==0
    set['hospitalization'] = (parseInt(doc.s)&16)!==0
    set['lifeThreatening'] = (parseInt(doc.s)&32)!==0
    set['other'] = (parseInt(doc.s)&64)!==0
    set['congenitalAnomali'] = (parseInt(doc.s)&2)!==0

    db.math490_test.update({'_id': doc._id}, {'$set': set})
  })
}















// for (var key in drugData) {
//   db.math490_test.updateMany({
//     'dgs': {
//       "$elemMatch": {
//         "char": '1',
//         'openfda._id': {'$in': drugData[key].ids}
//       }
//     }
//   }, { '$set': {'drug': key} })
// }
//
//
// for (var key in drugData) {
//   db.math490_test.find().forEach(function (doc) {
//     db.math490_test.update({'_id': doc._id}, {'$set': {'numDrugs': NumberInt(doc.dgs.length)}})
//   })
// }

// db.math490_test.find().forEach(function (doc) {
//   db.math490_test.update({'_id': doc._id}, {'$unset': {'dgs': ''}})
// })
// var idList = []
// for (var key in drugData) {
//   idList = idList.concat(drugData[key].ids)
// }

// printjson({
//   idList: idList.length,
//   drugData: drugData
// })
//
// function run () {
//   var matchOr = []
//
//   for (var key in drugData) {
//     matchOr.push({
//       'dgs': {
//         "$elemMatch": {
//           "char": '1',
//           'openfda._id': {'$in': drugData[key].ids}
//         }
//       }
//     })
//   }
//
//   return db.drug_event.aggregate([
//     {$match: {
//       "sex": {$in: ['1', '2']},
//       "dup": {$ne: true},
//       $or: matchOr
//     }},
//     {$project: {'dgs.openfda._id': 1, 'sex': 1,'s': 1, 'dgs.char': 1, '_id': 1}},
//     {$out: 'drug_event_math490'}
//   ], {
//     "hint": 'drugs.openfda._id_1_dup_1_sex_1_s_1_drugs.char_1__id_1'
//   })
// }
