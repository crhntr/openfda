db.math490_test.find().count() //   1794
db.math490_test.find({sex: '1'}).count() //   955
db.math490_test.find({sex: '2'}).count() //   839

db.math490_test.find({other: true}).count() //  663
db.math490_test.find({other: true, sex: '1'}).count() //  283
db.math490_test.find({other: true, sex: '2'}).count() //  380

db.math490_test.find({congenitalAnomali: true}).count() //  7
db.math490_test.find({congenitalAnomali: true, sex: '1'}).count() //  6
db.math490_test.find({congenitalAnomali: true, sex: '2'}).count() //  1

db.math490_test.find({lifeThreatening: true}).count() //  137
db.math490_test.find({lifeThreatening: true, sex: '1'}).count() //  81
db.math490_test.find({lifeThreatening: true, sex: '2'}).count() //  56

db.math490_test.find({hospitalization: true}).count() //  798
db.math490_test.find({hospitalization: true, sex: '1'}).count() //  520
db.math490_test.find({hospitalization: true, sex: '2'}).count() //  278

db.math490_test.find({disabling: true}).count() //  104
db.math490_test.find({disabling: true, sex: '1'}).count() //  48
db.math490_test.find({disabling: true, sex: '2'}).count() //  56

db.math490_test.find({death: true}).count() //  146
db.math490_test.find({death: true, sex: '1'}).count() //  65
db.math490_test.find({death: true, sex: '2'}).count() //  81
