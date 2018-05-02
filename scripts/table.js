

var results = []
for (var i = 0; i < 64; i++) {
  let query = {}
  query.death = i&1 ? true : false
  query.other = i&2 ? true : false
  query.congenitalAnomali = i&4 ? true : false
  query.lifeThreatening = i&8 ? true : false
  query.hospitalization = i&16 ? true : false
  query.disabling = i&32 ? true : false

  let n = db.math490_test.find(query).count()
  query.sex = '1'
  let n_m = db.math490_test.find(query).count()
  query.sex = '2'
  let n_f = db.math490_test.find(query).count()

  var srs = []

  for (var key in query) {
    if (key != 'sex' && query[key]) {
      srs.push(key)
    }
  }

  results.push({
    n, n_m, n_f,
    flags: srs
  })
}
results = results.filter(function(a){return a.n > 0})
results.sort(function(a, b){return a.n - b.n})

printjson(results)
